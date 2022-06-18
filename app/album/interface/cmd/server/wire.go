// +build wireinject

package main

import (
	"graduate/app/album/interface/internal/biz"
	"graduate/app/album/interface/internal/conf"
	"graduate/app/album/interface/internal/data"
	"graduate/app/album/interface/internal/server"
	"graduate/app/album/interface/internal/service"

	"github.com/go-kratos/kratos/v2/log"

	"github.com/go-kratos/kratos/v2"
	"github.com/google/wire"
)

func initApp(*conf.Grpc, *conf.Server, log.Logger) (*kratos.App, func(), error) {
	//	panic(wire.Build(data.ProviderSet, biz.ProviderSet, service.ProviderSet, server.ProviderSet, newApp))
	panic(wire.Build(server.ProviderSet, data.ProviderSet, biz.ProviderSet, service.ProviderSet, newApp))

}
