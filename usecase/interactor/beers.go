package interactor

import (
	"ZachIgarz/test-beer/domain/model"
	interfaceRepository "ZachIgarz/test-beer/interface/repository"
	"ZachIgarz/test-beer/usecase/repository"
	"net/http"

	"github.com/ansel1/merry"
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type BeerRequest struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Brewery  string `json:"brewery"`
	Country  string `json:"country"`
	Price    int64  `json:"Price"`
	Currency string `json:"currency"`
}

func (request *BeerRequest) Validate() error {
	return validation.ValidateStruct(request,
		validation.Field(&request.Name, validation.Required),
		validation.Field(&request.Brewery, validation.Required),
		validation.Field(&request.Country, validation.Required),
		validation.Field(&request.Price, validation.Required),
		validation.Field(&request.Currency, validation.Required),
	)
}

type BeerInteractorInterface interface {
	CreateBeer(request *BeerRequest) (*model.Beer, error)
	BeerList() ([]*model.Beer, error)
	BeerByID(ID string) (*model.Beer, error)
}

type beer struct {
	BeerRespository repository.BeerRepository
}

var BeerInteractor BeerInteractorInterface = &beer{
	BeerRespository: interfaceRepository.BeerRepository,
}

func (b *beer) CreateBeer(request *BeerRequest) (*model.Beer, error) {

	if err := request.Validate(); err != nil {
		merry.Wrap(err).WithHTTPCode(http.StatusBadRequest)
	}

	auxModel := model.Beer{
		ID:       request.ID,
		Name:     request.Name,
		Brewery:  request.Brewery,
		Country:  request.Country,
		Price:    request.Price,
		Currency: request.Currency,
	}

	createdBeer, err := b.BeerRespository.CreateBeer(&auxModel)
	if err != nil {
		return nil, merry.Wrap(err)
	}

	return createdBeer, nil
}

func (b *beer) BeerList() ([]*model.Beer, error) {
	beers, err := b.BeerRespository.Beers()
	if err != nil {
		return nil, merry.Wrap(err)
	}

	return beers, nil
}

func (b *beer) BeerByID(ID string) (*model.Beer, error) {
	beer, err := b.BeerRespository.Beer(ID)
	if err != nil {
		return nil, merry.Wrap(err)
	}

	return beer, nil
}
