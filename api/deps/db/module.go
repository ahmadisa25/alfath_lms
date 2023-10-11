package db

import (
	e_inst "alfath_lms/api/instructor/domain/entity"
	e_stud "alfath_lms/api/student/domain/entity"

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

		db.AutoMigrate(&e_inst.Instructor{}, &e_stud.Student{})

		return db
	}).In(dingo.Singleton)
}
