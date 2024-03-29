package ticket

import (
	"context"

	"github.com/google/uuid"

	"github.com/pkg/errors"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

type sqliteRepository struct {
	database *sqlx.DB
}

func NewMysqlRepository() (*sqliteRepository, error) {
	var err error
	database, err := sqlx.Connect("sqlite3", "challenge_db")
	if err != nil {
		return nil, errors.Wrap(err, "couldn't initialize mysql db")
	}

	// TODO: This setup could be done directly into the db or using a migration lib
	database.MustExec(schema)

	return &sqliteRepository{database: database}, nil
}

func (m sqliteRepository) Save(ctx context.Context, ticket Ticket) error {
	if _, err := m.database.Exec("INSERT into tickets (id, country, event) values ($1, $2, $3)", uuid.New().String(), ticket.Country, ticket.Event); err != nil {
		return errors.Wrap(err, "an error occurred while saving a ticket sale")
	}
	return nil
}

func (m sqliteRepository) GetAll(ctx context.Context) (Tickets, error) {
	rows, err := m.database.Queryx("SELECT * FROM tickets")
	if err != nil {
		return nil, errors.Wrap(err, "error fetching all tickets")
	}
	var result Tickets
	for rows.Next() {
		var ticket Ticket
		err = rows.StructScan(&ticket)
		if err != nil {
			return nil, errors.Wrap(err, "failed to scan ticket sale row into struct")
		}
		result = append(result, ticket)
	}
	return result, nil
}

var _ Repository = (*sqliteRepository)(nil)

var schema = `		
	create table if not exists tickets
	(
		id text not null
			constraint tickets_pk
				primary key,
		country text null,
		event text null
	);
	
	create unique index if not exists tickets_id_uindex
		on tickets (id);
	`
