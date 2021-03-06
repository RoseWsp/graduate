package biz

import (
	"context"
	"time"

	"github.com/go-kratos/kratos/v2/log"
)

//Domain Object
type Album struct {
	Id       int64
	Title    string
	Artist   string
	Price    float64
	CreateAt time.Time
}

// Domain Object
type Orders struct {
	Id       int64
	UserId   int64
	AlbumId  int64
	Price    float64
	Receiver string
	Address  string
	Mobile   string
}

type ListAlbum struct {
	Count  int64
	Albums []*Album
}

type AlbumRepo interface {
	ListAlbum(ctx context.Context, pageNum, pageSiz int64) (ListAlbum, error)
	GetAlbumById(ctx context.Context, id int64) (*Album, error)
	CreateOrders(ctx context.Context, orders *Orders) error
}

type AlbumUseCase struct {
	repo AlbumRepo
	log  *log.Helper
}

func NewAlbumUseCase(repo AlbumRepo, logger log.Logger) *AlbumUseCase {
	return &AlbumUseCase{
		repo: repo,
		log:  log.NewHelper(log.With(logger, "module", "usecase/album")),
	}
}

func (uc *AlbumUseCase) ListAlbum(ctx context.Context, pageNum, pageSize int64) (ListAlbum, error) {
	return uc.repo.ListAlbum(ctx, pageNum, pageSize)
}

func (uc *AlbumUseCase) GetAlbumById(ctx context.Context, id int64) (*Album, error) {
	return uc.repo.GetAlbumById(ctx, id)
}

func (uc *AlbumUseCase) CreateOrders(ctx context.Context, orders *Orders) error {
	return uc.repo.CreateOrders(ctx, orders)
}
