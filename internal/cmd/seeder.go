package cmd

import (
	"Crud/utility/seeders"
	"context"

	"github.com/gogf/gf/v2/frame/g"
	"gorm.io/gorm"
)

func RunAllSeeders(ctx context.Context, db *gorm.DB) {
	if err := seeders.SeedEmployees(ctx, db); err != nil {
		g.Log().Error(ctx, "Employee seeder failed:", err)
	} else {
		g.Log().Info(ctx, "Employees seeded successfully")
	}
}
