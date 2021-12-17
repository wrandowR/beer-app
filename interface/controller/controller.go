package controller

import (
	interfaceRepository "ZachIgarz/test-beer/interface/repository"
	"ZachIgarz/test-beer/usecase/repository"
	"net/http"

	"github.com/ansel1/merry"
)

type BeersController interface {
}

type beersController struct {
	BeerRepository repository.BeerRepository
}

var BeerInteractor BeersController = beersController{
	BeerRepository: interfaceRepository.BotDBRepository,
}

//
func (b *beersController) Beers(c Context) error {

	beerList, err := b.BeerRepository.Beers()
	if err != nil {
		return merry.Wrap(err)
	}

	return c.JSON(http.StatusCreated, beerList)
}
