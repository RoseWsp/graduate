// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//+build !wireinject

package main

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"graduate/app/album/job/internal/biz"
	"graduate/app/album/job/internal/conf"
	"graduate/app/album/job/internal/data"
	"graduate/app/album/job/internal/server"
	"graduate/app/album/job/internal/service"
)

// Injectors from wire.go:

func initApp(confData *conf.Data, confServer *conf.Server, logger log.Logger) (*kratos.App, func(), error) {
	db := data.NewDB(confData, logger)
	consumer := data.NewKafkaConsumer(confData)
	dataData, cleanup, err := data.NewData(db, consumer, logger)
	if err != nil {
		return nil, nil, err
	}
	jobRepo := data.NewJobRepo(dataData, logger)
	jobUseCase := biz.NewJobUseCase(jobRepo, logger)
	jobService := service.NewJobService(jobUseCase, logger)
	grpcServer := server.NewGRPCServer(confServer, logger, jobService)
	app := newApp(logger, grpcServer)
	return app, func() {
		cleanup()
	}, nil
}
