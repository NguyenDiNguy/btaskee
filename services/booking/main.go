package main

import (
	"btaskee/libs/logger"
	"btaskee/libs/mongodb"
	"btaskee/libs/redis"
	"btaskee/libs/viper"
	"btaskee/services/booking/controller"
	"btaskee/services/booking/proto"
	"btaskee/services/booking/repository"
	"btaskee/services/booking/service"
	"context"
	"net"
	"net/http"
	"strings"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
)

func main() {
	finish := make(chan struct{})
	logger.Init()

	go startGRPC()
	go startAPI()

	<-finish
}

func startGRPC() {
	urlgRPC := ":" + viper.GlobalConfig.PortGRPC
	log.Info().Any("Address", urlgRPC).Msg("Start GRPC Server")

	lis, err := net.Listen("tcp", urlgRPC)
	if err != nil {
		log.Panic().Err(err).Msgf("Failed to listen %v", urlgRPC)
		return
	}

	db := mongodb.Initialize(viper.GlobalConfig.DBuri, viper.GlobalConfig.DBName)
	redis := redis.Initialize(viper.GlobalConfig.RedisUrl, "")
	repo := repository.NewRepository(db)
	svc := service.NewService(repo, redis)
	controller := controller.NewController(svc)

	s := grpc.NewServer()
	proto.RegisterBookingServer(s, controller)
	err = s.Serve(lis)
	if err != nil {
		log.Error().Any("Error", err).Msg("Failed to serv")
		return
	}
}

func startAPI() {
	urlHTTP := ":" + viper.GlobalConfig.PortHTTP

	log.Info().Any("Address", urlHTTP).Msg("Start HTTP Server")
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	rmux := runtime.NewServeMux(
		runtime.WithMetadata(func(c context.Context, req *http.Request) metadata.MD {
			datas := make([]string, 0)
			for k, v := range req.Header {
				datas = append(datas, k, strings.Join(v, ","))
			}
			return metadata.Pairs(datas[:]...)
		}),
	)

	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	err := proto.RegisterBookingHandlerFromEndpoint(ctx, rmux, ":9080", opts)
	if err != nil {
		log.Error().Any("Error", err).Msg("Failed to RegisterAdminHandlerFromEndpoint")
		return
	}

	mux := http.NewServeMux()
	mux.Handle("/", rmux)
	http.ListenAndServe(urlHTTP, mux)
}
