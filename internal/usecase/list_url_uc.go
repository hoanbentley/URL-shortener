package usecase

import (
	"context"
	"github.com/hoanbentley/URL-shortener/internal/entities"
)

func (u *uc) ListUrl(ctx context.Context) ([]*entities.Urls, error) {
	return u.url.ListUrl(ctx)
}
