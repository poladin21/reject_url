package url

import (
	"database/sql"
	"redirect_url/internal/interfaces"
)

type Storage struct {
	db *sql.DB
}

func New(db *sql.DB) interfaces.Storeges {
	return &Storage{
		db: db,
	}
}

func (s Storage) GetByID(ID int) error {

	return nil
}
