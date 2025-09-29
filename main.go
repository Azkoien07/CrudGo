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
	_ = db

	// Seeders
	cmd.RunAllSeeders(ctx)

	// Servidor
	cmd.Main.Run(ctx)
}
