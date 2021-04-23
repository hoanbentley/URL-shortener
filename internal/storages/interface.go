package storages

import (
	"context"
	"database/sql"
	"github.com/hoanbentley/URL-shortener/internal/entities"
)

type Url interface {
	GenerateUrl(ctx context.Context, url *entities.Urls) error
	ListUrl(ctx context.Context) ([]*entities.Urls, error)
	GetUrl(ctx context.Context, shortCode string) (*entities.Urls, error)
	UpdateUrl(ctx context.Context, urls *entities.Urls) error
	DeleteUrl(ctx context.Context, shortCode string) error
	SearchUrl(ctx context.Context, shortCode, fullUrl string) ([]*entities.Urls, error)
	ValidateUser(ctx context.Context, userID, pwd sql.NullString) bool
}
