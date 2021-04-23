package usecase

import (
	"context"
	"github.com/hoanbentley/URL-shortener/internal/entities"
)

func (u *uc) SearchUrl(ctx context.Context, shortCode string) ([]*entities.Urls, error) {
	return u.url.SearchUrl(ctx, shortCode, "")
}
