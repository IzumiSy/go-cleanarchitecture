package drivers

import (
	"go-cleanarchitecture/adapters"
)

type HttpDriver struct{}

func (driver HttpDriver) Run() {
	adapters.RunHTTPServer()
}
