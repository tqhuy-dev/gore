//go:build wireinject
// +build wireinject

package api

import (
	"github.com/google/wire"
	"github.com/tqhuy-dev/gore/templates/applications/internal/services"
)

func NewApiApp() (BaseApiApp, func(), error) {
	wire.Build(NewBaseApiApp, services.NewSampleService,
		wire.Bind(new(ISampleService), new(*services.SampleService)))
	return BaseApiApp{}, func() {

	}, nil
}
