package service

import (
	"btaskee/libs/redis"
	"btaskee/model/task"
	"btaskee/services/booking/proto"
	"btaskee/services/booking/repository"
	"context"
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
)

type IService interface {
	CreateTask(ctx context.Context, userId string, req *proto.CreateTaskRequest) (*proto.CreateTaskResponse, error)
	GetTask(ctx context.Context, req *proto.GetTaskRequest) (*proto.GetTaskResponse, error)
	AcceptTask(ctx context.Context, userId string, req *proto.AcceptTaskRequest) (*proto.AcceptTaskResponse, error)
	ConfirmTasker(ctx context.Context, userId string, req *proto.ConfirmTaskerRequest) (*proto.ConfirmTaskerResponse, error)
}

type service struct {
	repo  repository.IRepository
	redis redis.IRedis
}

func NewService(repo repository.IRepository, redis redis.IRedis) IService {
	return &service{
		repo:  repo,
		redis: redis,
	}
}

// CreateTask ...
func (s *service) CreateTask(ctx context.Context, userId string, req *proto.CreateTaskRequest) (*proto.CreateTaskResponse, error) {
	taskId, err := s.repo.InsertTask(ctx, req.Detail, userId)
	if err != nil {
		return nil, err
	}

	ttask, err := s.repo.GetTaskById(ctx, *taskId)
	if err != nil {
		return nil, err
	}

	taskStatusKey := "task:" + ttask.Uuid + ":status"
	err = s.redis.Set(ctx, taskStatusKey, task.TaskStatus_PENDING.String())

	status, err := s.redis.Get(ctx, taskStatusKey)
	if err != nil {
		log.Err(err).Send()
		return nil, err
	}
	log.Debug().Msg(status)

	return &proto.CreateTaskResponse{
		Task: ttask,
	}, nil
}

// GetTask ...
func (s *service) GetTask(ctx context.Context, req *proto.GetTaskRequest) (*proto.GetTaskResponse, error) {
	task, err := s.repo.GetTaskByUuid(ctx, req.TaskId)
	if err != nil {
		return nil, err
	}

	return &proto.GetTaskResponse{
		Task: task,
	}, nil
}

// AcceptTask ...
func (s *service) AcceptTask(ctx context.Context, userId string, req *proto.AcceptTaskRequest) (*proto.AcceptTaskResponse, error) {
	// Lock task
	taskLockKey := "task:" + req.TaskId + ":lock"
	lockValue := uuid.New().String()
	if ok, err := s.redis.AcquireLock(ctx, taskLockKey, lockValue, 10*time.Second); !ok {
		return nil, err
	}
	defer s.redis.ReleaseLock(ctx, taskLockKey, lockValue)

	// Check task status
	taskStatusKey := "task:" + req.TaskId + ":status"
	status, err := s.redis.Get(ctx, taskStatusKey)
	if err != nil {
		log.Err(err).Send()
		return nil, err
	}
	if status == task.TaskStatus_CONFIRMED.String() {
		return nil, errors.New("Task confirmed")
	}

	// Validate task
	ttask, err := s.repo.GetTaskByUuid(ctx, req.TaskId)
	if err != nil {
		return nil, err
	}
	if ttask.Status.String() != task.TaskStatus_PENDING.String() {
		return nil, errors.New("Task status wrong")
	}
	if ttask.AskerId == userId {
		return nil, errors.New("You are asker")
	}

	// update task
	if ttask.AcceptedTaskers == nil {
		ttask.AcceptedTaskers = make([]string, 0)
	}
	ttask.UpdatedAt = time.Now().Unix()
	ttask.AcceptedTaskers = append(ttask.AcceptedTaskers, userId)

	err = s.repo.UpdateTask(ctx, ttask)
	if err != nil {
		return nil, err
	}

	// Add tasker vào redis
	acceptedTaskersKey := "task:" + req.TaskId + ":accepted_taskers"
	log.Debug().Msgf("%v - %v", acceptedTaskersKey, userId)
	s.redis.Add(ctx, acceptedTaskersKey, userId)

	return &proto.AcceptTaskResponse{
		IsSuccess: true,
	}, nil
}

// ConfirmTasker ...
func (s *service) ConfirmTasker(ctx context.Context, userId string, req *proto.ConfirmTaskerRequest) (*proto.ConfirmTaskerResponse, error) {
	// Lock task
	taskLockKey := "task:" + req.TaskId + ":lock"
	lockValue := uuid.New().String()
	if ok, err := s.redis.AcquireLock(ctx, taskLockKey, lockValue, 10*time.Second); !ok {
		return nil, err
	}
	defer s.redis.ReleaseLock(ctx, taskLockKey, lockValue)

	// Check task status
	taskStatusKey := "task:" + req.TaskId + ":status"
	status, err := s.redis.Get(ctx, taskStatusKey)
	if err != nil {
		return nil, err
	}
	if status == task.TaskStatus_CONFIRMED.String() {
		return nil, errors.New("Task confirmed")
	}

	// Check taskers
	acceptedTaskersKey := "task:" + req.TaskId + ":accepted_taskers"
	log.Debug().Msgf("%v - %v", acceptedTaskersKey, req.TaskerId)
	isAccepted, err := s.redis.CheckList(ctx, acceptedTaskersKey, req.TaskerId)
	if err != nil {
		return nil, err
	}

	if !isAccepted {
		return nil, errors.New("Tasker has not accepted this task")
	}

	// Validate task
	ttask, err := s.repo.GetTaskByUuid(ctx, req.TaskId)
	if err != nil {
		return nil, err
	}
	if ttask.Status != task.TaskStatus_PENDING {
		return nil, errors.New("Task status wrong")
	}
	if ttask.AskerId != userId {
		return nil, errors.New("You are not owner")
	}

	isExited := false
	for _, tasker := range ttask.AcceptedTaskers {
		if tasker == req.TaskerId {
			isExited = true
			break
		}
	}

	if !isExited {
		return nil, errors.New("Tasker has not accepted this task")
	}

	ttask.UpdatedAt = time.Now().Unix()
	ttask.ConfirmedTasker = req.TaskerId
	ttask.Status = task.TaskStatus_CONFIRMED
	// update task
	s.repo.UpdateTask(ctx, ttask)

	// Update redis status
	err = s.redis.Set(ctx, taskStatusKey, task.TaskStatus_CONFIRMED)

	return &proto.ConfirmTaskerResponse{
		IsSuccess: true,
	}, nil
}
