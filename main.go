package main

import (
	_ "Crud/internal/packed"

	"github.com/gogf/gf/v2/os/gctx"

	"Crud/internal/cmd"
)

func main() {
	db := cmd.Migrate()
	_ = db
	cmd.Main.Run(gctx.GetInitCtx())
}
