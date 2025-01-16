package api

import (
	"github.com/litmus-zhang/momentum-backend/util"
	"go.uber.org/fx"
)

var Module = fx.Module("api",
	fx.Provide(
		NewServer,
		util.NewPasetoMaker,
	),
)
