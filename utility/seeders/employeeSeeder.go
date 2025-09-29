package seeders

import (
	"Crud/internal/model/entity"
	"context"
	"fmt"
	"gorm.io/gorm"
)

func SeedEmployees(ctx context.Context, db *gorm.DB) error {
	employees := []entity.Employee{
		{Name: "Alice", Age: 30, Email: "alice@example.com", Phone: "123456"},
		{Name: "Bob", Age: 25, Email: "bob@example.com", Phone: "654321"},
		{Name: "Carlos", Age: 44, Email: "carlos@example.com", Phone: "5435345"},
	}

	for _, e := range employees {
		// Verifica si existe por email
		var count int64
		if err := db.WithContext(ctx).Model(&entity.Employee{}).
			Where("email = ?", e.Email).Count(&count).Error; err != nil {
			return fmt.Errorf("error checking employee %s: %w", e.Name, err)
		}

		if count == 0 {
			if err := db.WithContext(ctx).Create(&e).Error; err != nil {
				return fmt.Errorf("error inserting employee %s: %w", e.Name, err)
			}
		}
	}

	return nil
}