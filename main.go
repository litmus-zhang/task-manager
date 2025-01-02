package main

import (
	"github.com/litmus-zhang/task-manager/lib/api"
	"github.com/litmus-zhang/task-manager/lib/config"
	"github.com/litmus-zhang/task-manager/lib/db"
	"go.uber.org/fx"
)

func main() {
	app := fx.New(
		config.Module,
		db.Module,
		api.Module,
	)
	app.Run()
}
