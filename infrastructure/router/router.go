package router

import (
	"ZachIgarz/test-beer/interface/controller"

	"github.com/labstack/echo"
)

//NewRouter new router
func NewRouter(e *echo.Echo, c controller.AppController) *echo.Echo {

	e.GET("/beers", func(context echo.Context) error {
		return controller.BeerController.Beers(context)
	})

	e.GET("/beers/:beerID", func(context echo.Context) error {
		return controller.BeerController.Beer(context)
	})

	e.POST("/beers", func(context echo.Context) error {
		return controller.BeerController.CreateBeer(context)
	})

	e.GET("/beers/:beerID/boxprice", func(context echo.Context) error {
		return controller.BeerController.BeerBoxPrice(context)
	})

	return e
}
