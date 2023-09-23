package db

import (
	"flamingo.me/dingo"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Module struct{}

func (*Module) Configure(injector *dingo.Injector){
	injector.Bind(new(gorm.DB)).ToProvider( func() *gorm.DB {
		dsn:= "ahmdisa:sarah072724@tcp(127.0.0.1:3306)/alfath.lms?charset=utf8mb4&parseTime=True&loc=Local"
		db, err := gorm.Open(mysql.Open(dsn))
		if err != nil {
			panic(err)
		}

		return db
	}).In(dingo.Singleton)
}