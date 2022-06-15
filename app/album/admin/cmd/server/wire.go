// +build wireinject

package main

import (
	"graduate/app/album/service/internal/biz"
	"graduate/app/album/service/internal/conf"
	"graduate/app/album/service/internal/data"
	"graduate/app/album/service/internal/server"
	"graduate/app/album/service/internal/service"

	"github.com/go-kratos/kratos/v2/log"

	"github.com/go-kratos/kratos/v2"
	"github.com/google/wire"
)

func initApp(*conf.Data, *conf.Server, log.Logger) (*kratos.App, func(), error) {
	//	panic(wire.Build(data.ProviderSet, biz.ProviderSet, service.ProviderSet, server.ProviderSet, newApp))
	panic(wire.Build(server.ProviderSet, data.ProviderSet, biz.ProviderSet, service.ProviderSet, newApp))

}
