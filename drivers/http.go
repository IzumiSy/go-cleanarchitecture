package drivers

import (
	"context"
	"go-cleanarchitecture/adapters"
)

type HttpDriver struct{}

func (driver HttpDriver) Run(ctx context.Context) {
	adapters.RunHTTPServer(ctx)
}
