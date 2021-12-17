package controller

import (
	"ZachIgarz/test-beer/domain/model"
	"ZachIgarz/test-beer/testutil"
	"ZachIgarz/test-beer/usecase/interactor"
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
)

func TestHTTPRequestGetBeerList(t *testing.T) {
	testutil.ConfigDbTest(t)

	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/beers", nil)

	rec := httptest.NewRecorder()

	c := e.NewContext(req, rec)

	BeerController = &beersController{
		BeerInteractor: interactor.BeerInteractor,
	}

	err := BeerController.Beers(c)
	assert.NoError(t, err)

	assert.Equal(t, http.StatusOK, rec.Code)

	var expectedResponse []model.Beer

	err = json.Unmarshal(rec.Body.Bytes(), &expectedResponse)
	assert.NoError(t, err)

	//3 beers inserted on the migration
	assert.Equal(t, 3, len(expectedResponse))
}

func TestHTTPRequestCreateBeer(t *testing.T) {
	testutil.ConfigDbTest(t)

	body := new(bytes.Buffer)

	body.Write([]byte(`{
	"id":"1264e273-d888-45dc-a5df-3700fc474845",
    "name":"testBeer",
    "brewery":"polars",
    "country":"Colombia",
    "price":2,
    "currency":"USD"
	}`))

	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/beers", body)
	req.Header.Set(echo.HeaderContentType, "application/json")

	rec := httptest.NewRecorder()

	c := e.NewContext(req, rec)

	BeerController = &beersController{
		BeerInteractor: interactor.BeerInteractor,
	}

	err := BeerController.CreateBeer(c)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusCreated, rec.Code)

	expectedResponse := model.Beer{
		ID:       "1264e273-d888-45dc-a5df-3700fc474845",
		Name:     "testBeer",
		Brewery:  "polars",
		Country:  "Colombia",
		Price:    2,
		Currency: "USD",
	}

	var response model.Beer
	err = json.Unmarshal(rec.Body.Bytes(), &response)
	assert.NoError(t, err)

	assert.Equal(t, expectedResponse, response)
}

func TestHTTPRequestGetBeerByID(t *testing.T) {
	testutil.ConfigDbTest(t)

	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/beers/:beerID", nil)

	rec := httptest.NewRecorder()

	beerID := "5f21ed03-513a-4731-9323-59d46c1d739b"

	c := e.NewContext(req, rec)
	c.SetParamNames("beerID")
	c.SetParamValues(beerID)

	BeerController = &beersController{
		BeerInteractor: interactor.BeerInteractor,
	}

	err := BeerController.Beer(c)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)

	expectedResponse := model.Beer{
		ID:       "5f21ed03-513a-4731-9323-59d46c1d739b",
		Name:     "Vusenwaiser",
		Brewery:  "GermanyCO",
		Country:  "Germany",
		Price:    3.2,
		Currency: "USD",
	}

	var response model.Beer
	err = json.Unmarshal(rec.Body.Bytes(), &response)
	assert.NoError(t, err)

	assert.Equal(t, expectedResponse, response)
}
