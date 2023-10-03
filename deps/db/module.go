package db

import (
	"alfath_lms/instructor/domain/entity"

	"flamingo.me/dingo"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Module struct{}

func (*Module) Configure(injector *dingo.Injector) {
	injector.Bind(new(gorm.DB)).ToProvider(func() *gorm.DB {
		dsn := "ahmdisa:Sarah072724!@tcp(127.0.0.1:3306)/alfath_lms?charset=utf8mb4&parseTime=True&loc=Local"
		db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{TranslateError: true})
		if err != nil {
			panic(err)
		}

		db.AutoMigrate(&entity.Instructor{})

		return db
	}).In(dingo.Singleton)
}
