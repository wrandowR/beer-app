package controller

import (
	"ZachIgarz/test-beer/usecase/interactor"
	"net/http"

	"github.com/ansel1/merry"
)

type BeersControllerInterface interface {
	Beers(c Context) error
	Beer(c Context) error
	CreateBeer(c Context) error
	BeerBoxPrice(c Context) error
}

type beersController struct {
	BeerInteractor interactor.BeerInteractorInterface
}

var BeerController BeersControllerInterface = &beersController{
	BeerInteractor: interactor.BeerInteractor,
}

//beers return all beers in a list
func (b *beersController) Beers(c Context) error {
	beerList, err := b.BeerInteractor.BeerList()
	if err != nil {
		return merry.Wrap(err)
	}

	return c.JSON(http.StatusOK, beerList)
}

//Beer get a specific beer by id
func (b *beersController) Beer(c Context) error {
	beerID := c.Param("beerID")
	if beerID == "" {
		return merry.New("Param {beerID} is required").WithHTTPCode(http.StatusBadRequest)
	}

	beer, err := b.BeerInteractor.BeerByID(beerID)
	if err != nil {
		return merry.Wrap(err)
	}

	if beer == nil {
		return c.NoContent(http.StatusNotFound)
	}

	return c.JSON(http.StatusOK, beer)
}

//Beer get a specific beer by id
func (b *beersController) CreateBeer(c Context) error {
	var request interactor.BeerRequest
	err := c.Bind(&request)
	if err != nil {
		return merry.Wrap(err)
	}

	beerList, err := b.BeerInteractor.CreateBeer(&request)
	if err != nil {
		return merry.Wrap(err)
	}

	return c.JSON(http.StatusCreated, beerList)
}

func (b *beersController) BeerBoxPrice(c Context) error {

	beerID := c.Param("beerID")
	if beerID == "" {
		return merry.New("Param {beerID} is required").WithHTTPCode(http.StatusBadRequest)
	}

	var boxpriceRequest interactor.BoxPriceRequest
	err := c.Bind(&boxpriceRequest)
	if err != nil {
		return merry.Wrap(err)
	}

	price, err := b.BeerInteractor.BeerBoxPrice(beerID, &boxpriceRequest)
	if err != nil {
		return merry.Wrap(err)
	}

	return c.JSON(http.StatusOK, price)
}
