package cmd

import (
	"Crud/internal/model/entity"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Migrate() *gorm.DB {
	dsn := "host=localhost user=postgres password=Azkoien1033702995 dbname=Crud port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	// List of entities to migrate
	entities := []interface{}{
		&entity.Department{},
		&entity.Employee{},
	}

	// AutoMigrate the entities
	if err := db.AutoMigrate(entities...); err != nil {
		panic(err)
	}
	return db
}
