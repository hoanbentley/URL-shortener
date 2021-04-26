package usecase

import (
	"context"
	"github.com/hoanbentley/URL-shortener/internal/entities"
	"html"
)

func (u *uc) SearchUrl(ctx context.Context, shortCode, fullUrl string) ([]*entities.Urls, error) {
	//escape url string xss url
	escapeShortCode := html.EscapeString(shortCode)
	escapeFullUrl := html.EscapeString(fullUrl)

	return u.url.SearchUrl(ctx, escapeShortCode, escapeFullUrl)
}
