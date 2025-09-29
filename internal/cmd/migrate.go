package cmd

import (
	"Crud/internal/model/entity"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Migrate() *gorm.DB {
	dsn := "host=localhost user=fabrica password=fabrica2024* dbname=Crud port=5432 sslmode=disable TimeZone=UTC"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
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
