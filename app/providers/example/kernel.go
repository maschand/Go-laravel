//go:build wireinject
// +build wireinject

package provider_test

import (
	"github.com/google/wire"
)

//provider set
var fooSet = wire.NewSet(NewFooRepository, NewFooService)
var barSet = wire.NewSet(NewBarRepository, NewBarService)

func InitializedFooBarService() *FooBarService {
	wire.Build(
		fooSet,
		barSet,
		NewFooBarService,
	)
	return nil
}

//binding interface
var helloService = wire.NewSet(
	NewSayHelloImpl,
	wire.Bind(new(SayHello), new(*SayHelloImpl)),
)

func InitializedHelloService() *SayHelloService {
	wire.Build(
		helloService,
		NewSayHelloService,
	)

	return nil
}

//Struct Provider
var FooBarSet = wire.NewSet(
	NewFoo,
	NewBar,
)

func InitializedFooBar() *FooBar {
	wire.Build(
		FooBarSet,
		wire.Struct(new(FooBar), "*"),
	)
	return nil
}

//with param injector
func InitializedService(isError bool) (*SimpleService, error) {
	wire.Build(
		NewSimpleRepository, NewSimpleService,
	)
	return nil, nil
}

//Multiple Binding
func InitializedDatabaseRepository() *DatabaseRepository {
	wire.Build(
		NewDatabasePostgreeSQL,
		NewDatabaseMongoDB,
		NewDatabaseRepository,
	)
	return nil
}
