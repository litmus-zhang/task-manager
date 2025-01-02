package api

import "go.uber.org/fx"

var Module = fx.Module("api",
	fx.Provide(
		fx.Annotate(NewServer,
			fx.As(new(Server)),
		),
	))
