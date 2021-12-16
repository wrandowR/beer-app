package repository

import (
	"ZachIgarz/test-beer/infrastructure/datastore"
	"database/sql"
	"errors"

	"github.com/ansel1/merry"
	"github.com/doug-martin/goqu/v9"
	"github.com/doug-martin/goqu/v9/exp"
)

var botTableGoqu = goqu.T("bot")

type botRepository struct {
	db           *goqu.Database
	botTableGoqu exp.IdentifierExpression
}

var BotDBRepository repository.BotRepository = &botRepository{
	db:           &datastore.SQLDBGoqu,
	botTableGoqu: botTableGoqu,
}

func (repository *botRepository) ListBots(projectID string) ([]*model.Bot, error) {

	var bots = []*model.Bot{}

	err := repository.db.From(repository.botTableGoqu).Select(
		repository.botTableGoqu.Col("id"),
		repository.botTableGoqu.Col("project_id"),
		repository.botTableGoqu.Col("bot_whatsapp_line_id"),
		repository.botTableGoqu.Col("title"),
		repository.botTableGoqu.Col("initial_greet"),
		repository.botTableGoqu.Col("questions"),
		repository.botTableGoqu.Col("simultaneous_meetings"),
		repository.botTableGoqu.Col("meeting_duration"),
		repository.botTableGoqu.Col("schedule_monday"),
		repository.botTableGoqu.Col("schedule_tuesday"),
		repository.botTableGoqu.Col("schedule_wednesday"),
		repository.botTableGoqu.Col("schedule_thursday"),
		repository.botTableGoqu.Col("schedule_friday"),
		repository.botTableGoqu.Col("schedule_saturday"),
		repository.botTableGoqu.Col("schedule_sunday"),
		repository.botTableGoqu.Col("timeoffset"),
		repository.botTableGoqu.Col("calendar"),
		repository.botTableGoqu.Col("created"),
	).Where(
		repository.botTableGoqu.Col("project_id").Eq(projectID),
	).ScanStructs(&bots)
	if errors.Is(err, sql.ErrNoRows) {
		return []*model.Bot{}, nil
	}
	if err != nil {
		return nil, merry.Wrap(err).WithValue("projectId", projectID)
	}

	return bots, nil

}
