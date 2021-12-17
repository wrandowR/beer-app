package repository

import "ZachIgarz/test-beer/domain/model"

type BeerRepository interface {
	Beers() ([]*model.Beer, error)
	Beer(ID string) (*model.Beer, error)
	//BeersBoxPrice(ID string) []*model.Beer
}
