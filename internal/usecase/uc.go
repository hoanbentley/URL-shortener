package usecase

import (
	"context"
	"github.com/hoanbentley/URL-shortener/internal/entities"
	"github.com/hoanbentley/URL-shortener/internal/storages"
)

type uc struct {
	url storages.Url
}

func NewUc(db string) UseCase {
	return &uc{
		url: storages.NewLiteDB(),
	}
}

type UseCase interface {
	CreateUrl(ctx context.Context, url *entities.Urls) error
	ListUrl(ctx context.Context) ([]*entities.Urls, error)
	GetUrl(ctx context.Context, shortCode string) (*entities.Urls, error)
}
