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

// RedirectUrl
// get url from short code
// update number of hits + 1
// return url struct with short url
func (t *uc) RedirectUrl(ctx context.Context, shortCode string) (*entities.Urls, error) {
	//get url from short code
	url, err := t.url.GetUrl(ctx, shortCode)
	if err != nil {
		return nil, err
	}

	//update number of hits + 1
	url.NumberOfHits = url.NumberOfHits + 1
	err = t.url.UpdateUrl(ctx, url)
	if err != nil {
		return nil, err
	}

	//return url struct with short url
	return t.url.GetUrl(ctx, shortCode)
}
