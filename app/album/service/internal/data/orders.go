package data

import (
	"context"
	"graduate/app/album/service/internal/biz"
	"time"

	"gorm.io/gorm"
)

//PO
type Orders struct {
	gorm.Model
	Id       int64
	UserId   int64
	AlbumId  int64
	Price    float64
	Receiver string
	Address  string
	Mobile   string
	CreateAt time.Time
}

func createOrders(ctx context.Context, b *biz.Orders, db *gorm.DB) error {
	orders := Orders{
		UserId:   b.UserId,
		AlbumId:  b.AlbumId,
		Price:    b.Price,
		Receiver: b.Receiver,
		Address:  b.Address,
		Mobile:   b.Mobile,
		CreateAt: time.Now(),
	}

	rst := db.WithContext(ctx).Create(&orders)
	return rst.Error
}
