package cmd

import (
	"Crud/internal/model/entity"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Migrate() *gorm.DB {
	// DSN correcto para MySQL
	dsn := "root:@tcp(127.0.0.1:3306)/Crud?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	// Lista de entidades a migrar
	entities := []interface{}{
		&entity.Department{},
		&entity.Employee{},
	}

	if err := db.AutoMigrate(entities...); err != nil {
		panic(err)
	}

	return db
}
