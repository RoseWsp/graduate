package data

import (
	"context"
	"encoding/json"
	"graduate/app/album/service/internal/biz"
	"graduate/pkg/util/pagination"
	"strconv"
	"time"

	"github.com/Shopify/sarama"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-redis/redis/v8"

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

func (r *albumRepo) GetAlbumById(ctx context.Context, id int64) (*biz.Album, error) {
	var album *biz.Album
	rAlbum, err := r.getAlbumFromRedis(ctx, id)
	if err != nil {
		return nil, err
	}
	album = rAlbum
	if album != nil {
		return album, nil
	}
	album, err = r.getAlbumFromDB(ctx, id)
	if err != nil {
		return nil, err
	}
	err = r.setAlbumToRedis(ctx, album)
	if err != nil {
		return nil, err
	}
	return album, nil
}

func (r *albumRepo) getAlbumFromDB(ctx context.Context, id int64) (*biz.Album, error) {
	a := Album{}
	result := r.data.db.WithContext(ctx).First(&a, id)
	return &biz.Album{
		Id:       a.Id,
		Title:    a.Title,
		Artist:   a.Artist,
		Price:    a.Price,
		CreateAt: a.CreateAt,
	}, result.Error
}

func (r *albumRepo) getAlbumFromRedis(ctx context.Context, id int64) (*biz.Album, error) {
	key := "getAlbumFromRedis" + strconv.Itoa((int)(id))
	val, err := r.data.rdb.Get(ctx, key).Result()
	if err == redis.Nil {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	var rst biz.Album
	json.Unmarshal([]byte(val), &rst)
	return &rst, nil
}

func (r *albumRepo) setAlbumToRedis(ctx context.Context, album *biz.Album) error {
	if val, err := json.Marshal(album); err != nil {
		return err
	} else {
		key := "getAlbumFromRedis" + strconv.Itoa((int)(album.Id))
		err = r.data.rdb.Set(ctx, key, string(val), 2*time.Hour).Err()
		return err
	}
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
	var cnt int64
	r.data.db.WithContext(ctx).Model(&Album{}).Count(&cnt)
	return biz.ListAlbum{
		Albums: rv,
		Count:  cnt,
	}, nil
}

func (r *albumRepo) CreateOrders(ctx context.Context, orders *biz.Orders) error {
	err := createOrders(ctx, orders, r.data.db)
	if err != nil {
		return err
	}

	b, err := json.Marshal(orders)
	if err != nil {
		return err
	}
	r.data.kp.Input() <- &sarama.ProducerMessage{
		Topic: "integrating", // 发送订单 ，供 积分job 消费
		Value: sarama.ByteEncoder(b),
	}
	return nil
}
