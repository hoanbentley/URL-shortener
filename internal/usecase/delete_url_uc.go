package usecase

import (
	"context"
	"errors"
)

// DeleteUrl
// Get information url from short code
// If not exists then return error
// Execute delete url
func (u *uc) DeleteUrl(ctx context.Context, shortCode string) error {
	url, err := u.url.GetUrl(ctx, shortCode)
	if err != nil {
		return errors.New("get url failed")
	}

	if url == nil {
		return errors.New("the information of url is not exists")
	}

	return u.url.DeleteUrl(ctx, shortCode)
}
