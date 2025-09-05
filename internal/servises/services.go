package servises

import (
	"errors"
	"log/slog"
	"redirect_url/internal/interfaces"
)

type Services struct {
	storage interfaces.Storeges
	log     *slog.Logger
}

func New(storege interfaces.Storeges, log *slog.Logger) interfaces.Servises {
	return &Services{
		storage: storege,
		log:     log,
	}
}

func (s Services) GetByID(ID int) error {
	return errors.New("Not implemented")
}

func (s Services) Get(ID int) error {
	return errors.New("Not implemented")
}
