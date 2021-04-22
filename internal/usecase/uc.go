package usecase

import (
	"context"
	"database/sql"
	"github.com/hoanbentley/URL-shortener/internal/entities"
	"github.com/hoanbentley/URL-shortener/internal/storages"
)

type uc struct {
	url storages.Url
}

func NewUc(db string) UseCase {
	return &uc{
		url: storages.NewLiteDB(),
	}
}

type UseCase interface {
	CreateUrl(ctx context.Context, url *entities.Urls) error
	ListUrl(ctx context.Context) ([]*entities.Urls, error)
	RedirectUrl(ctx context.Context, shortCode string) (*entities.Urls, error)
	Validate(ctx context.Context, user, password sql.NullString) bool
	CreateToken(id, jwtKey string) (string, error)
}
