package usecase

import (
	"context"
	"github.com/hoanbentley/URL-shortener/internal/entities"
)

func (t *uc) CreateUrl(ctx context.Context, url *entities.Urls) error {
	return t.url.GenerateUrl(ctx, url)
}

func (t *uc) ListUrl(ctx context.Context) ([]*entities.Urls, error) {
	return t.url.ListUrl(ctx)
}

func (t *uc) GetUrl(ctx context.Context, shortCode string) (*entities.Urls, error) {
	return t.url.GetUrl(ctx, shortCode)
}
