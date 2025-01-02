package db

import "go.uber.org/fx"

var Module = fx.Module("db",
	
	fx.Provide(
		fx.Annotate(
			NewStore,
			fx.As(new(Store)),
		),
	))
