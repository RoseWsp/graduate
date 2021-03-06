package data

import (
	"context"
	"encoding/json"
	"graduate/app/album/admin/internal/biz"
	"time"

	"github.com/go-kratos/kratos/v2/log"
	"gorm.io/gorm"
)

var _ biz.AlbumRepo = (*albumRepo)(nil)

// 实现 biz 定义的 AlbumRepo 接口
type albumRepo struct {
	data *Data
	log  *log.Helper
}

//PO
type Album struct {
	gorm.Model
	Id       int64
	Title    string
	Artist   string
	Price    float64
	CreateAt time.Time
}

func NewAlbumRepo(data *Data, logger log.Logger) biz.AlbumRepo {
	return &albumRepo{
		data: data,
		log:  log.NewHelper(log.With(logger, "module", "data/album")),
	}
}

func (r *albumRepo) UpdateAlbum(ctx context.Context, userId int64, b *biz.Album) error {
	o := Album{}
	result := r.data.db.WithContext(ctx).First(&o, b.Id)
	if result.Error != nil {
		return result.Error
	}

	o.Title = b.Title
	o.Artist = b.Artist
	o.UpdatedAt = time.Now()

	result = r.data.db.WithContext(ctx).Save(&o)
	originObj, _ := json.Marshal(o)
	newObj, _ := json.Marshal(b)
	err := r.data.db.Transaction(func(tx *gorm.DB) error {
		innerErr := tx.WithContext(ctx).Save(&o)
		if innerErr.Error != nil {
			return innerErr.Error
		}
		innerErr = insertIntegrating(ctx,
			string(originObj),
			string(newObj),
			userId, tx,
		)
		return innerErr.Error
	})
	return err
}
