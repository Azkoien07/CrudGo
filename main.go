package main

import (
	_ "Crud/internal/packed"

	"github.com/gogf/gf/v2/os/gctx"

	"Crud/internal/cmd"
)

func main() {
	// Contexto global
	ctx := gctx.New()

	// Migraciones
	db := cmd.Migrate()

	// Seeders
	cmd.RunAllSeeders(ctx, db)

	// Servidor
	cmd.Main.Run(ctx)
}
