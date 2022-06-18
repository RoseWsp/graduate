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

type ListAlbum struct {
	Count  int64
	Albums []*Album
}

type AlbumRepo interface {
	UpdateAlbum(ctx context.Context, userId int64, album *Album) error
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

func (uc *AlbumUseCase) UpdateAlbum(ctx context.Context, userId int64, album *Album) error {
	return uc.repo.UpdateAlbum(ctx, userId, album)
}
