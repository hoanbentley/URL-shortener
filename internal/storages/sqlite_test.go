package storages

import (
	"context"
	"database/sql"
	"github.com/hoanbentley/URL-shortener/internal/entities"
	"github.com/hoanbentley/URL-shortener/internal/storages/mocks"
	"reflect"
	"testing"
)

var m = new(mocks.Url)
var client = Url(m)

func Test_liteDB_GetUrl(t *testing.T) {
	type fields struct {
		db *sql.DB
	}
	type args struct {
		ctx       context.Context
		shortCode string
	}
	tests := []struct {
		name    string
		args    args
		rs      *entities.Urls
		want    *entities.Urls
		err     error
		wantErr bool
	}{
		{
			name: "case 1",
			args: args{
				ctx:       context.Background(),
				shortCode: "nYwnyRl",
			},
			rs: &entities.Urls{
				ShortCode:    "nYwnyRl",
				FullUrl:      "https://www.thepolyglotdeveloper.com/2016/12/create-a-url-shortener-with-golang-and-couchbase-nosql/",
				Expiry:       60,
				NumberOfHits: 6,
			},
			want: &entities.Urls{
				ShortCode:    "nYwnyRl",
				FullUrl:      "https://www.thepolyglotdeveloper.com/2016/12/create-a-url-shortener-with-golang-and-couchbase-nosql/",
				Expiry:       60,
				NumberOfHits: 6,
			},
			err:     nil,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m.On("GetUrl", tt.args.ctx, tt.args.shortCode).Return(tt.rs, tt.err).Once()
			got, err := client.GetUrl(tt.args.ctx, tt.args.shortCode)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetUrl() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetUrl() got = %v, want %v", got, tt.want)
			}
			m.AssertExpectations(t)
		})
	}
}

func Test_liteDB_GenerateUrl(t *testing.T) {
	type fields struct {
		db *sql.DB
	}
	type args struct {
		ctx context.Context
		url *entities.Urls
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		err     error
		wantErr bool
	}{
		{
			name: "case 1",
			args: args{
				ctx: context.Background(),
				url: &entities.Urls{
					ShortCode:    "",
					FullUrl:      "https://www.thepolyglotdeveloper.com/2016/12/create-a-url-shortener-with-golang-and-couchbase-nosql/",
					Expiry:       60,
					NumberOfHits: 1,
				},
			},
			err:     nil,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m.On("GenerateUrl", tt.args.ctx, tt.args.url).Return(tt.err).Once()
			if err := client.GenerateUrl(tt.args.ctx, tt.args.url); (err != nil) != tt.wantErr {
				t.Errorf("GenerateUrl() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
