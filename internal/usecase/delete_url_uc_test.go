package usecase

import (
	"context"
	"errors"
	"github.com/hoanbentley/URL-shortener/internal/storages"
	"github.com/hoanbentley/URL-shortener/internal/usecase/mocks"
	"testing"
)

var m = new(mocks.UseCase)
var client = UseCase(m)

func Test_uc_DeleteUrl(t *testing.T) {
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
		wantErr bool
		err     error
	}{
		{
			name: "case 1",
			args: args{
				ctx:       context.Background(),
				shortCode: "ovw2ZV3",
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
			wantErr: true,
			err:     errors.New("sql error"),
		},
		{
			name: "case 3",
			args: args{
				ctx:       context.Background(),
				shortCode: "ovw2ZV3465",
			},
			wantErr: true,
			err:     errors.New("sql error"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m.On("DeleteUrl", tt.args.ctx, tt.args.shortCode).Return(tt.err).Once()
			if err := client.DeleteUrl(tt.args.ctx, tt.args.shortCode); (err != nil) != tt.wantErr {
				t.Errorf("uc.DeleteUrl() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
