package controller

import (
	"net/http"

	"github.com/ansel1/merry"
)

//AppController all controllers
type AppController struct {
	Beers BeersController
}

type BeersController interface {
}

type beersController struct {
}

//
func (b *beersController) Beers(c Context) error {

	project, err := proController.projectInteractor.GetAllProjects(orgID)
	if err != nil {
		return merry.Wrap(err)
	}

	return c.JSON(http.StatusCreated, project)
}
