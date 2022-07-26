//go:build wireinject
// +build wireinject

package providers

import (
	"github.com/chand19-af/digitels-template/app/controllers"
	"github.com/chand19-af/digitels-template/app/repositories"
	"github.com/chand19-af/digitels-template/app/services"
	"github.com/chand19-af/digitels-template/platform/database"
	"github.com/google/wire"
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
