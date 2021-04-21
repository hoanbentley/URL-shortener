package storages

import (
	"context"
	"github.com/hoanbentley/URL-shortener/internal/entities"
)

type Url interface {
	GenerateUrl(ctx context.Context, t *entities.Urls) error
	ListUrl(ctx context.Context) ([]*entities.Urls, error)
	GetUrl(ctx context.Context, shortCode string) (*entities.Urls, error)
}
