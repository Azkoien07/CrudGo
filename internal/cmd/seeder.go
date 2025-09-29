package cmd

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"

	"seeders/employeeSeeder"
)

func RunAllSeeders(ctx context.Context) {
	if err := employeeSeeder.SeedEmployees(ctx); err != nil {
		g.Log().Error(ctx, "Employee seeder failed:", err)
	} else {
		g.Log().Info(ctx, "Employees seeded successfully")
	}
}
