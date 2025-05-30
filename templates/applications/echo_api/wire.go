//go:build wireinject
// +build wireinject

package main

import "github.com/google/wire"

func NewApiApp() (BaseApiApp, func(), error) {
	wire.Build(NewBaseApiApp)
	return BaseApiApp{}, func() {

	}, nil
}
