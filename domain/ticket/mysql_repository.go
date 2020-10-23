package ticket

import (
	"context"

	"github.com/google/uuid"

	"github.com/pkg/errors"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

type mysqlRepository struct {
	database *sqlx.DB
}

func NewMysqlRepository() (*mysqlRepository, error) {
	var err error
	database, err := sqlx.Connect("sqlite3", "challenge_db")
	if err != nil {
		return nil, errors.Wrap(err, "couldn't initialize mysql db")
	}

	// TODO: This setup could be done directly into the db or using a migration lib
	database.MustExec(schema)

	return &mysqlRepository{database: database}, nil
}

func (m mysqlRepository) Save(ctx context.Context, ticket Ticket) error {
	if _, err := m.database.Exec("insert into tickets (id, country, event) values ($1, $2, $3)", uuid.New().String(), ticket.Country, ticket.Event); err != nil {
		return errors.Wrap(err, "an error occurred while saving a ticket sale")
	}
	return nil
}

func (m mysqlRepository) GetAll(ctx context.Context) (Tickets, error) {
	rows, err := m.database.Queryx("SELECT * FROM tickets")
	if err != nil {
		return nil, errors.Wrap(err, "error fetching all tickets")
	}
	var result Tickets
	for rows.Next() {
		var ticketDb Ticket
		err = rows.StructScan(&ticketDb)
		if err != nil {
			return nil, nil
		}
		result = append(result, ticketDb)
	}
	return result, nil
}

var _ Repository = (*mysqlRepository)(nil)

var schema = `		
	create table if not exists tickets
	(
		id text not null
			constraint tickets_pk
				primary key,
		country varchar(255) null,
		event varchar(255) null
	);
	
	create unique index if not exists tickets_id_uindex
		on tickets (id);
	`
