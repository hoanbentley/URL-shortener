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

func NewUc() UseCase {
	return &uc{
		url: storages.NewLiteDB(),
	}
}

type UseCase interface {
	CreateUrl(ctx context.Context, url *entities.Urls) (*entities.Urls, error)
	ListUrl(ctx context.Context) ([]*entities.Urls, error)
	RedirectUrl(ctx context.Context, shortCode string) (*entities.Urls, error)
	DeleteUrl(ctx context.Context, shortCode string) error
	SearchUrl(ctx context.Context, shortCode, fullUrl string) ([]*entities.Urls, error)
	Validate(ctx context.Context, user, password sql.NullString) bool
	CreateToken(id, jwtKey string) (string, error)
	ValidToken(token, JWTKey string) (string, bool)
}
