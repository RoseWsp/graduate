// +build wireinject

package main

import (
	"graduate/app/album/task/internal/biz"
	"graduate/app/album/task/internal/conf"
	"graduate/app/album/task/internal/data"
	"graduate/app/album/task/internal/server"
	"graduate/app/album/task/internal/service"

	"github.com/go-kratos/kratos/v2/log"

	"github.com/go-kratos/kratos/v2"
	"github.com/google/wire"
)

func initApp(*conf.Data, *conf.Server, log.Logger) (*kratos.App, func(), error) {
	//	panic(wire.Build(data.ProviderSet, biz.ProviderSet, service.ProviderSet, server.ProviderSet, newApp))
	panic(wire.Build(server.ProviderSet, data.ProviderSet, biz.ProviderSet, service.ProviderSet, newApp))

}
