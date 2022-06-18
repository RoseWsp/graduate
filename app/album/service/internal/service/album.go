package service

import (
	"context"
	"time"

	v1 "graduate/api/album/service/v1"
	"graduate/app/album/service/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
)

type AlbumService struct {
	v1.UnimplementedAlbumServer

	ac  *biz.AlbumUseCase
	log *log.Helper
}

func NewAlbumService(ac *biz.AlbumUseCase, logger log.Logger) *AlbumService {
	return &AlbumService{
		ac:  ac,
		log: log.NewHelper(log.With(logger, "module", "service/album")),
	}

}

func (s *AlbumService) ListAlbum(ctx context.Context, req *v1.ListAlbumReq) (*v1.ListAlbumReply, error) {
	rv, err := s.ac.ListAlbum(ctx, req.PageNum, req.PageSize)
	if err != nil {
		return nil, err
	}
	albums := make([]*v1.ListAlbumReply_Album, 0)
	for _, r := range rv.Albums {
		albums = append(albums, &v1.ListAlbumReply_Album{
			Id:       r.Id,
			Title:    r.Title,
			Artist:   r.Artist,
			Price:    r.Price,
			CreateAt: time.Unix(r.CreateAt.Unix(), 0).Format("2006-01-02 15:04:05"),
		})
	}
	return &v1.ListAlbumReply{
		Count:  rv.Count,
		Albums: albums,
	}, nil
}

func (s *AlbumService) GetAlbumById(ctx context.Context, req *v1.GetAlbumByIdReq) (*v1.GetAlbumByIdReply, error) {
	rv, err := s.ac.GetAlbumById(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	album := &v1.GetAlbumByIdReply{
		Album: &v1.GetAlbumByIdReply_Album{
			Id:       rv.Id,
			Title:    rv.Title,
			Artist:   rv.Artist,
			Price:    rv.Price,
			CreateAt: time.Unix(rv.CreateAt.Unix(), 0).Format("2006-01-02 15:04:05"),
		},
	}
	return album, nil
}

func (s *AlbumService) CreateOrders(ctx context.Context, req *v1.CreateOrdersReq) (*v1.CreateOrdersReply, error) {
	order := biz.Orders{
		UserId:   req.Orders.UserId,
		AlbumId:  req.Orders.AlbumId,
		Price:    req.Orders.Price,
		Receiver: req.Orders.Receiver,
		Address:  req.Orders.Address,
		Mobile:   req.Orders.Mobile,
	}
	return &v1.CreateOrdersReply{}, s.ac.CreateOrders(ctx, &order)
}
