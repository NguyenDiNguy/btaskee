package test

import (
	bookingClient "btaskee/services/booking/client"
	identityClient "btaskee/services/identity/client"
	"btaskee/services/identity/proto"
	identityProto "btaskee/services/identity/proto"
	"fmt"
	"sync"
	"testing"
)

type User struct {
	Jwt  string
	Uuid string
}

var (
	JwtList map[string]*User = make(map[string]*User)
	Pass    string           = "2080"
)

func createUser(cli proto.IdentityClient) error {
	for i := 0; i <= 1000; i++ {
		err := SignUp(cli, fmt.Sprintf("kenny%v@gmail.com", i), Pass)
		if err != nil {
			return err
		}
	}
	return nil
}

func getJWT(cli identityProto.IdentityClient) error {
	for i := 0; i <= 1000; i++ {
		email := fmt.Sprintf("kenny%v@gmail.com", i)
		rs, err := SignIn(cli, email, Pass)
		if err != nil {
			return err
		}
		JwtList[email] = &User{
			Jwt:  rs.Jwt,
			Uuid: rs.User.Uuid,
		}
	}
	return nil
}

func TestCreateUser(t *testing.T) {
	identityCli, err := identityClient.NewIdentityClient("127.0.0.1:80")
	if err != nil {
		panic(err)
	}
	err = createUser(identityCli)
	if err != nil {
		t.Error(err)
	}
}

func TestAcceptedLate(t *testing.T) {
	identityCli, err := identityClient.NewIdentityClient("127.0.0.1:80")
	if err != nil {
		t.Error(err)
	}
	err = getJWT(identityCli)
	if err != nil {
		t.Error(err)
	}

	// booking
	bookingCli, err := bookingClient.NewBookingClient("127.0.0.1:80")

	asker := JwtList["kenny1@gmail.com"]
	task, err := CreateTask(bookingCli, "task1", asker.Jwt)
	if err != nil {
		t.Error(err)
	}

	tasker := JwtList["kenny2@gmail.com"]
	_, err = AcceptTask(bookingCli, task.Task.Uuid, tasker.Jwt)
	if err != nil {
		t.Error(err)
	}

	_, err = ConfirmTasker(bookingCli, task.Task.Uuid, tasker.Uuid, asker.Jwt)
	if err != nil {
		t.Error(err)
	}

	laterTasker := JwtList["kenny3@gmail.com"]
	_, err = AcceptTask(bookingCli, task.Task.Uuid, laterTasker.Jwt)
	if err == nil {
		t.Error("Expected error when late accepter tries to accept the task, but got none")
	} else {
		// Kiểm tra loại lỗi nếu có thể
		expectedErrMsg := "rpc error: code = Unknown desc = Task status wrong" // Thay thế bằng thông báo lỗi thực tế
		if err.Error() != expectedErrMsg {
			t.Errorf("Expected error message '%s', but got '%s'", expectedErrMsg, err.Error())
		}
	}
}

func TestBasicFlow(t *testing.T) {
	identityCli, err := identityClient.NewIdentityClient("127.0.0.1:80")
	if err != nil {
		t.Error(err)
	}
	err = getJWT(identityCli)
	if err != nil {
		t.Error(err)
	}

	// booking
	bookingCli, err := bookingClient.NewBookingClient("127.0.0.1:80")

	asker := JwtList["kenny1@gmail.com"]
	task, err := CreateTask(bookingCli, "task1", asker.Jwt)
	if err != nil {
		t.Error(err)
	}

	tasker := JwtList["kenny2@gmail.com"]
	_, err = AcceptTask(bookingCli, task.Task.Uuid, tasker.Jwt)
	if err != nil {
		t.Error(err)
	}

	_, err = ConfirmTasker(bookingCli, task.Task.Uuid, tasker.Uuid, asker.Jwt)
	if err != nil {
		t.Error(err)
	}
}

func TestMainFlow(t *testing.T) {
	identityCli, err := identityClient.NewIdentityClient("127.0.0.1:80")
	if err != nil {
		t.Error(err)
	}
	err = getJWT(identityCli)
	if err != nil {
		t.Error(err)
	}

	// booking
	bookingCli, err := bookingClient.NewBookingClient("127.0.0.1:80")
	asker := JwtList["kenny1@gmail.com"]
	task, err := CreateTask(bookingCli, "task1", asker.Jwt)
	if err != nil {
		t.Error(err)
	}

	for i := 2; i <= 5; i++ {
		tasker := JwtList[fmt.Sprintf("kenny%v@gmail.com", i)]
		_, err = AcceptTask(bookingCli, task.Task.Uuid, tasker.Jwt)
		if err != nil {
			t.Error(err)
		}
	}

	tasker := JwtList["kenny3@gmail.com"]
	_, err = ConfirmTasker(bookingCli, task.Task.Uuid, tasker.Uuid, asker.Jwt)
	if err != nil {
		t.Error(err)
	}
}

func TestQuality(t *testing.T) {
	identityCli, err := identityClient.NewIdentityClient("127.0.0.1:80")
	if err != nil {
		t.Error(err)
	}
	err = getJWT(identityCli)
	if err != nil {
		t.Error(err)
	}

	// booking
	bookingCli, err := bookingClient.NewBookingClient("127.0.0.1:80")
	wg := &sync.WaitGroup{}

	for i := 1; i <= 1000; i += 5 {
		t.Log("Test", i)
		wg.Add(1)
		go func(begin int, wg *sync.WaitGroup) {
			defer wg.Done()
			asker := JwtList[fmt.Sprintf("kenny%v@gmail.com", begin)]
			taskDetail := fmt.Sprintf("task%v", begin)
			task, err := CreateTask(bookingCli, taskDetail, asker.Jwt)
			if err != nil {
				t.Error(err)
			}

			for j := begin + 1; j <= begin+4; j++ {
				tasker := JwtList[fmt.Sprintf("kenny%v@gmail.com", j)]
				_, err = AcceptTask(bookingCli, task.Task.Uuid, tasker.Jwt)
				if err != nil {
					t.Error(err)
				}
			}

			tasker := JwtList[fmt.Sprintf("kenny%v@gmail.com", begin+3)]
			_, err = ConfirmTasker(bookingCli, task.Task.Uuid, tasker.Uuid, asker.Jwt)
			if err != nil {
				t.Error(err)
			}
		}(i, wg)
	}
	wg.Wait()
}
