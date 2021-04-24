package usecase

import (
	"context"
	"errors"
	"github.com/hoanbentley/URL-shortener/internal/storages"
	"reflect"
	"testing"

	"github.com/hoanbentley/URL-shortener/internal/entities"
)

func Test_uc_RedirectUrl(t *testing.T) {
	type fields struct {
		url storages.Url
	}
	type args struct {
		ctx       context.Context
		shortCode string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *entities.Urls
		rs      *entities.Urls
		wantErr bool
		err     error
	}{
		{
			name: "case 1",
			args: args{
				ctx:       context.Background(),
				shortCode: "ovw2ZV3",
			},
			rs: &entities.Urls{
				ShortCode:    "ovw2ZV3",
				FullUrl:      "https://microservices.io/patterns/microservices.html",
				Expiry:       60,
				NumberOfHits: 6,
			},
			want: &entities.Urls{
				ShortCode:    "ovw2ZV3",
				FullUrl:      "https://microservices.io/patterns/microservices.html",
				Expiry:       60,
				NumberOfHits: 6,
			},
			wantErr: false,
			err:     nil,
		},
		{
			name: "case 2",
			args: args{
				ctx:       context.Background(),
				shortCode: "",
			},
			rs:      nil,
			want:    nil,
			wantErr: true,
			err:     errors.New("get url not found"),
		},
		{
			name: "case 3",
			args: args{
				ctx:       context.Background(),
				shortCode: "ovw2ZV31",
			},
			rs:      nil,
			want:    nil,
			wantErr: true,
			err:     errors.New("get url not found"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m.On("RedirectUrl", tt.args.ctx, tt.args.shortCode).Return(tt.rs, tt.err).Once()
			got, err := client.RedirectUrl(tt.args.ctx, tt.args.shortCode)
			if (err != nil) != tt.wantErr {
				t.Errorf("uc.RedirectUrl() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("uc.RedirectUrl() = %v, want %v", got, tt.want)
			}
		})
	}
}
