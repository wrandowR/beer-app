package repository

import (
	"ZachIgarz/test-beer/domain/model"
	"ZachIgarz/test-beer/infrastructure/datastore"
	"ZachIgarz/test-beer/usecase/repository"
	"database/sql"
	"errors"

	"github.com/ansel1/merry"
	"github.com/doug-martin/goqu/v9"
	"github.com/doug-martin/goqu/v9/exp"
)

var beerTable = goqu.T("beer")

type beerRepository struct {
	db        *goqu.Database
	beerTable exp.IdentifierExpression
}

var BotDBRepository repository.BeerRepository = &beerRepository{
	db:        &datastore.SQLDBGoqu,
	beerTable: beerTable,
}

func (b *beerRepository) Beers() ([]*model.Beer, error) {

	var beers = []*model.Beer{}

	err := b.db.From(b.beerTable).Select(
		b.beerTable.Col("id"),
		b.beerTable.Col("name"),
		b.beerTable.Col("brewery"),
		b.beerTable.Col("country"),
		b.beerTable.Col("price"),
		b.beerTable.Col("currency"),
	).ScanStructs(&beers)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, nil
	}

	if err != nil {
		return nil, merry.Wrap(err)
	}

	return beers, nil
}

func (b *beerRepository) Beer(ID string) (*model.Beer, error) {
	var beer model.Beer

	ok, err := b.db.From(b.beerTable).Select(
		b.beerTable.Col("id"),
		b.beerTable.Col("name"),
		b.beerTable.Col("brewery"),
		b.beerTable.Col("country"),
		b.beerTable.Col("price"),
		b.beerTable.Col("currency"),
	).Where(
		b.beerTable.Col("id").Eq(ID)).ScanStruct(&beer)

	if err != nil {
		return nil, merry.Wrap(err)
	}

	if !ok {
		return nil, nil
	}

	return &beer, nil
}
