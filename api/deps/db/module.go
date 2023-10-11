package db

import (
	e_inst "alfath_lms/api/instructor/domain/entity"
	e_stud "alfath_lms/api/student/domain/entity"

	"fmt"
	"os"

	"flamingo.me/dingo"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Module struct{}

func (*Module) Configure(injector *dingo.Injector) {
	injector.Bind(new(gorm.DB)).ToProvider(func() *gorm.DB {
		dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", os.Getenv("DB_USERNAME"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_NAME"))
		fmt.Println(dsn)
		db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{TranslateError: true})
		if err != nil {
			panic(err)
		}

		db.AutoMigrate(&e_inst.Instructor{}, &e_stud.Student{})

		return db
	}).In(dingo.Singleton)
}
