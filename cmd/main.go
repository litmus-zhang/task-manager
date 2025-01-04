package main

import (
	"github.com/litmus-zhang/task-manager/internal/api"
	"github.com/litmus-zhang/task-manager/internal/config"
	"github.com/litmus-zhang/task-manager/internal/db"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

const (
	VAULT_ADDR  = "http://127.0.0.1:8200"
	VAULT_TOKEN = "root"
	VAULT_PATH  = "secret/data/task-manager"
)

func main() {
	app := fx.New(
		config.Module,
		db.Module,
		api.Module,
		fx.Provide(zap.NewDevelopment),
		fx.Invoke(func(lc fx.Lifecycle, cfg *config.Config, server *api.Server) error {
			return server.Start()
		}),
	)
	app.Run()
}
