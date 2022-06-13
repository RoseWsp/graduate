package data

import (
	"context"
	"graduate/app/album/service/internal/biz"
	"graduate/pkg/util/pagination"
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

func (r *albumRepo) ListAlbum(ctx context.Context, pageNum, pageSize int64) (biz.ListAlbum, error) {
	var as []Album
	result := r.data.db.WithContext(ctx).
		Limit(int(pageSize)).
		Offset(int(pagination.GetPageOffset(pageNum, pageSize))).
		Find(&as)
	if result.Error != nil {
		return biz.ListAlbum{}, result.Error
	}
	rv := make([]*biz.Album, 0)
	for _, o := range as {
		rv = append(rv, &biz.Album{
			Id:       o.Id,
			Title:    o.Title,
			Artist:   o.Artist,
			Price:    o.Price,
			CreateAt: o.CreateAt,
		})
	}

	cnt := r.data.db.WithContext(ctx).Count(0)
	return biz.ListAlbum{
		Albums: rv,
		Count:  cnt,
	}, nil
}
