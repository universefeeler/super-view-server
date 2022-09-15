package dao

import (
	"github.com/scorpiotzh/toolib"
	"gorm.io/gorm"
	"super-view-server/config"
)

type DbDao struct {
	db       *gorm.DB
	parserDb *gorm.DB
}

var (
	GormDb   *gorm.DB
	AllDbDao *DbDao
)

func not_init() {
	dbMysql := config.Cfg.DB.Mysql
	parserMysql := config.Cfg.DB.ParserMysql

	db, err := toolib.NewGormDB(dbMysql.Addr, dbMysql.User, dbMysql.Password, dbMysql.DbName, dbMysql.MaxOpenConn, dbMysql.MaxIdleConn)
	if err != nil {
		panic(err)
	}
	GormDb = db

	parserDb, err := toolib.NewGormDB(parserMysql.Addr, parserMysql.User, parserMysql.Password, parserMysql.DbName, parserMysql.MaxOpenConn, parserMysql.MaxIdleConn)
	if err != nil {
		panic(err)
	}
	AllDbDao = &DbDao{db: db, parserDb: parserDb}
}

type RecordTotal struct {
	Total int `json:"total" gorm:"column:total"`
}
