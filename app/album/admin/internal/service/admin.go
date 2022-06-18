package service

import (
	"context"
	"time"

	v1 "graduate/api/album/admin/v1"
	"graduate/app/album/admin/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
)

type AdminService struct {
	v1.UnimplementedAdminServer

	auc *biz.AlbumUseCase
	ouc *biz.OrdersUseCase
	log *log.Helper
}

func NewAdminService(auc *biz.AlbumUseCase, ouc *biz.OrdersUseCase, logger log.Logger) *AdminService {
	return &AdminService{
		auc: auc,
		ouc: ouc,
		log: log.NewHelper(log.With(logger, "module", "service/admin")),
	}
}

func (s *AdminService) UpdateAlbum(ctx context.Context, req *v1.UpdateAlbumReq) (*v1.UpdateAlbumReply, error) {
	crtAt, err := time.ParseInLocation("2006-01-02 15:04:05", req.Album.CreateAt, time.Local)
	if err != nil {
		return &v1.UpdateAlbumReply{}, err
	}
	album := biz.Album{
		Id:       req.Album.Id,
		Title:    req.Album.Title,
		Artist:   req.Album.Artist,
		Price:    req.Album.Price,
		CreateAt: crtAt,
	}
	return &v1.UpdateAlbumReply{}, s.auc.UpdateAlbum(ctx, req.UserId, &album)
}

func (s *AdminService) UpdateOrders(ctx context.Context, req *v1.UpdateOrdersReq) (*v1.UpdateOrdersReply, error) {
	crtAt, err := time.ParseInLocation("2006-01-02 15:04:05", req.Orders.CreateAt, time.Local)
	if err != nil {
		return &v1.UpdateOrdersReply{}, err
	}
	order := biz.Orders{
		Id:       req.Orders.Id,
		UserId:   req.Orders.UserId,
		AlbumId:  req.Orders.AlbumId,
		Price:    req.Orders.Price,
		Receiver: req.Orders.Receiver,
		Address:  req.Orders.Address,
		Mobile:   req.Orders.Mobile,
		CreateAt: crtAt,
	}
	return &v1.UpdateOrdersReply{}, s.ouc.UpdateOrders(ctx, req.Orders.UserId, &order)
}
