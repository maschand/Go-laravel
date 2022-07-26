//go:build wireinject
// +build wireinject

package providers

import (
	"github.com/google/wire"
	"gitlab.com/d6825/golang_template/app/controllers"
	"gitlab.com/d6825/golang_template/app/repositories"
	"gitlab.com/d6825/golang_template/app/services"
	"gitlab.com/d6825/golang_template/platform/database"
)

var hotelSet = wire.NewSet(
	repositories.NewHotelRepository,
	wire.Bind(new(repositories.HotelRepository), new(*repositories.HotelRepositoryImpl)),
	services.NewHotelService,
	wire.Bind(new(services.HotelService), new(*services.HotelServiceImpl)),
	controllers.NewHotelController,
	wire.Bind(new(controllers.HotelControlller), new(*controllers.HotelControllerImp)),
)

func InitializedServer() *controllers.HotelControllerImp {
	wire.Build(
		database.GormMysqlConnection,
		hotelSet,
	)

	return nil
}
