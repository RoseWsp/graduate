package data

import (
	"graduate/app/album/service/conf"

	"github.com/go-kratos/kratos/v2/log"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Data struct {
	db  *gorm.DB
	log *log.Helper
}

func newDB(conf *conf.Data, logger log.Logger) *gorm.BD {
	log := log.NewHelper(log.With(logger, "module", "album-service/data/gorm"))
	db, err := gorm.Open(mysql.Open(conf.Database.Source), &gorm.Config{})
	if err != nil {
		log.Fatal("failed opening connection to mysql: %v", err)
	}
	if err := db.AutoMigrate(&Album{}); err != nil {
		log.Fatal(err)
	}
	return db
}

// NewData .
func NewData(db *gorm.DB, logger log.Logger) (*Data, func(), error) {
	log := log.NewHelper(log.With(logger, "module", "order-service/data"))

	d := &Data{
		db:  db,
		log: log,
	}
	return d, func() {

	}, nil
}
