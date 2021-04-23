package usecase

import (
	"context"
	"github.com/hoanbentley/URL-shortener/internal/entities"
)

// RedirectUrl
// get url from short code
// update number of hits + 1
// return url struct with short url
func (u *uc) RedirectUrl(ctx context.Context, shortCode string) (*entities.Urls, error) {
	//get url from short code
	url, err := u.url.GetUrl(ctx, shortCode)
	if err != nil {
		return nil, err
	}

	//update number of hits + 1
	url.NumberOfHits = url.NumberOfHits + 1
	err = u.url.UpdateUrl(ctx, url)
	if err != nil {
		return nil, err
	}

	//return url struct with short url
	return u.url.GetUrl(ctx, shortCode)
}
