package url

import "database/sql"

type Storage struct {
	db *sql.DB
}

func New() Storage {
	return Storage{}
}
