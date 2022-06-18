package data

import (
	serviceV1 "graduate/api/album/service/v1"
	taskV1 "graduate/api/album/task/v1"
	"graduate/app/album/interface/internal/conf"

	"github.com/go-kratos/kratos/v2/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/google/wire"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewGrpcClient, NewInterfaceRepo)

type Data struct {
	client *grpcClient
	log    *log.Helper
}

type grpcClient struct {
	albumClient serviceV1.AlbumClient
	taskClient  taskV1.TaskClient
}

func NewGrpcClient(conf *conf.Grpc) (*grpcClient, error) {
	serviceConn, err := grpc.Dial(conf.Service.Addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	albumClient := serviceV1.NewAlbumClient(serviceConn)

	taskConn, err := grpc.Dial(conf.Task.Addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	taskClient := taskV1.NewTaskClient(taskConn)
	return &grpcClient{
		albumClient, taskClient,
	}, nil
}

// NewData .
func NewData(client *grpcClient, logger log.Logger) (*Data, func(), error) {
	log := log.NewHelper(log.With(logger, "module", "album-task/data"))

	d := &Data{
		client: client,
		log:    log,
	}
	return d, func() {

	}, nil
}
