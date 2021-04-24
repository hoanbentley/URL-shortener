package usecase

import (
	"context"
	"errors"
	"github.com/hoanbentley/URL-shortener/internal/entities"
	"time"
)

// RedirectUrl
// Get url from short code
// Check expiry url link (created_date + expiry) second
// Update number of hits + 1
// Return url struct with short url
func (u *uc) RedirectUrl(ctx context.Context, shortCode string) (*entities.Urls, error) {
	//get url from short code
	url, err := u.url.GetUrl(ctx, shortCode)
	if err != nil {
		return nil, err
	}

	//check expiry url link
	expiryDate := url.CreatedDate + int64(url.Expiry)
	currentDate := time.Now().Unix()
	if expiryDate < currentDate {
		return nil, errors.New("the url link expired")
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
