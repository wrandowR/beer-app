package interactor

import (
	"ZachIgarz/test-beer/domain/model"
	interfaceRepository "ZachIgarz/test-beer/interface/repository"
	"ZachIgarz/test-beer/testutil"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type CurrencyLayerMock struct {
	mock.Mock
}

func (c *CurrencyLayerMock) Convert(fromCurrency string, toCurrency string) (float32, error) {
	args := c.Called(fromCurrency, toCurrency)

	return args.Get(0).(float32), args.Error(1)
}

func TestGetBeerList(t *testing.T) {
	testutil.ConfigDbTest(t)
	BeerInteractor = &beer{
		BeerRespository: interfaceRepository.BeerRepository,
	}

	requestCreateBeer := BeerRequest{
		ID:       uuid.New().String(),
		Name:     "Cool Beer",
		Brewery:  "Cool Beer House",
		Country:  "Colombia",
		Price:    3.2,
		Currency: "USD",
	}

	_, err := BeerInteractor.CreateBeer(&requestCreateBeer)
	assert.NoError(t, err)

	requestCreateBeer2 := BeerRequest{
		ID:       uuid.New().String(),
		Name:     "Cool Beer",
		Brewery:  "Cool Beer House",
		Country:  "Colombia",
		Price:    3.2,
		Currency: "USD",
	}

	_, err = BeerInteractor.CreateBeer(&requestCreateBeer2)
	assert.NoError(t, err)

	beers, err := BeerInteractor.BeerList()
	assert.NoError(t, err)

	assert.Equal(t, 2, len(beers))
}

func TestSuccessfullyCreateBeers(t *testing.T) {
	testutil.ConfigDbTest(t)
	BeerInteractor = &beer{
		BeerRespository: interfaceRepository.BeerRepository,
	}

	requestCreateBeer := BeerRequest{
		ID:       uuid.New().String(),
		Name:     "Cool Beer",
		Brewery:  "Cool Beer House",
		Country:  "Colombia",
		Price:    3.2,
		Currency: "USD",
	}

	_, err := BeerInteractor.CreateBeer(&requestCreateBeer)
	assert.NoError(t, err)

	requestCreateBeer2 := BeerRequest{
		ID:       uuid.New().String(),
		Name:     "Cool Beer",
		Brewery:  "Cool Beer House",
		Country:  "Colombia",
		Price:    3.2,
		Currency: "USD",
	}

	_, err = BeerInteractor.CreateBeer(&requestCreateBeer2)
	assert.NoError(t, err)

	resultFirstBeer, err := BeerInteractor.BeerByID(requestCreateBeer.ID)
	assert.NoError(t, err)

	expecteFirstBeerResult := model.Beer{
		ID:       requestCreateBeer.ID,
		Name:     requestCreateBeer.Name,
		Brewery:  requestCreateBeer.Brewery,
		Country:  requestCreateBeer.Country,
		Price:    requestCreateBeer.Price,
		Currency: requestCreateBeer.Currency,
	}

	assert.Equal(t, expecteFirstBeerResult, *resultFirstBeer)

	resultSecondBeer, err := BeerInteractor.BeerByID(requestCreateBeer2.ID)
	assert.NoError(t, err)

	expecteSecondBeerResult := model.Beer{
		ID:       requestCreateBeer2.ID,
		Name:     requestCreateBeer2.Name,
		Brewery:  requestCreateBeer2.Brewery,
		Country:  requestCreateBeer2.Country,
		Price:    requestCreateBeer2.Price,
		Currency: requestCreateBeer2.Currency,
	}

	assert.Equal(t, expecteSecondBeerResult, *resultSecondBeer)
}

func TestGetBeerBoxPrice(t *testing.T) {
	testutil.ConfigDbTest(t)

	currencyLayerMock := new(CurrencyLayerMock)

	BeerInteractor = &beer{
		BeerRespository: interfaceRepository.BeerRepository,
		CurrencyLayer:   currencyLayerMock,
	}

	requestCreateBeer := BeerRequest{
		ID:       uuid.New().String(),
		Name:     "Cool Beer",
		Brewery:  "Cool Beer House",
		Country:  "Colombia",
		Price:    3.2,
		Currency: "USD",
	}

	_, err := BeerInteractor.CreateBeer(&requestCreateBeer)
	assert.NoError(t, err)

	var plnValue float32 = 4.09445

	request := BoxPriceRequest{
		Quantity: 12,
		Currency: "PLN",
	}

	expectedPrice := requestCreateBeer.Price * plnValue
	expectedTotalPriceResult := expectedPrice * float32(request.Quantity)

	currencyLayerMock.On("Convert", requestCreateBeer.Currency, request.Currency).Return(plnValue, nil)

	totalPrice, err := BeerInteractor.BeerBoxPrice(requestCreateBeer.ID, &request)
	assert.NoError(t, err)

	currencyLayerMock.AssertExpectations(t)

	assert.Equal(t, expectedTotalPriceResult, totalPrice)
}
