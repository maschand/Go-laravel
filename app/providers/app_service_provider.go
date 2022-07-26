//go:build wireinject
// +build wireinject

package providers

import (
	"github.com/chand19-af/digitels-template/app/repositories"
	"github.com/chand19-af/digitels-template/app/services"
	"github.com/google/wire"
)

var hotelService = wire.NewSet(
	services.NewHotelServiceImpl,
	wire.Bind(new(services.Hotel), new(*services.HotelServiceImpl)),
)

var hotelRepository = wire.NewSet(
	repositories.NewHotelRepositoryImpl,
	wire.Bind(new(repositories.Hotel), new(*repositories.HotelRepositoryImpl)),
)

func InitializedHotelRepository() *repositories.HotelRepository {
	wire.Build(
		hotelRepository,
		repositories.NewHotelRepository,
	)

	return nil
}

func InitializedHotelService() *services.HotelService {
	wire.Build(
		hotelService,
		hotelRepository,
		repositories.NewHotelRepository,
		services.NewHotelService,
	)

	return nil
}
