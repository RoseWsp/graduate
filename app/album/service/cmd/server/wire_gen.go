// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//+build !wireinject

package main

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"graduate/app/album/service/internal/biz"
	"graduate/app/album/service/internal/conf"
	"graduate/app/album/service/internal/data"
	"graduate/app/album/service/internal/server"
	"graduate/app/album/service/internal/service"
)

// Injectors from wire.go:

func initApp(confData *conf.Data, confServer *conf.Server, logger log.Logger) (*kratos.App, func(), error) {
	db := data.NewDB(confData, logger)
	client := data.NewRedis(confData)
	dataData, cleanup, err := data.NewData(db, client, logger)
	if err != nil {
		return nil, nil, err
	}
	albumRepo := data.NewAlbumRepo(dataData, logger)
	albumUseCase := biz.NewAlbumUseCase(albumRepo, logger)
	albumService := service.NewAlbumService(albumUseCase, logger)
	grpcServer := server.NewGRPCServer(confServer, logger, albumService)
	app := newApp(logger, grpcServer)
	return app, func() {
		cleanup()
	}, nil
}
