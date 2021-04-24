package storages

import (
	"context"
	"database/sql"
	"errors"
	"reflect"
	"testing"

	"github.com/hoanbentley/URL-shortener/internal/entities"
	"github.com/hoanbentley/URL-shortener/internal/storages/mocks"
	_ "github.com/mattn/go-sqlite3"
)

var m = new(mocks.Url)
var client = Url(m)

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
		{
			name: "case 2",
			args: args{
				ctx: context.Background(),
				url: &entities.Urls{
					ShortCode:    "nYwnyRl",
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

func Test_liteDB_ListUrl(t *testing.T) {
	type fields struct {
		db *sql.DB
	}
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		rs      []*entities.Urls
		want    []*entities.Urls
		err     error
		wantErr bool
	}{
		{
			name: "case 1",
			args: args{
				ctx: context.Background(),
			},
			rs: []*entities.Urls{
				&entities.Urls{
					ShortCode:    "nYwnyRl",
					FullUrl:      "https://www.thepolyglotdeveloper.com/2016/12/create-a-url-shortener-with-golang-and-couchbase-nosql/",
					Expiry:       60,
					NumberOfHits: 6,
				},
			},
			want: []*entities.Urls{
				&entities.Urls{
					ShortCode:    "nYwnyRl",
					FullUrl:      "https://www.thepolyglotdeveloper.com/2016/12/create-a-url-shortener-with-golang-and-couchbase-nosql/",
					Expiry:       60,
					NumberOfHits: 6,
				},
			},
			err:     nil,
			wantErr: false,
		},
		{
			name: "case 2",
			args: args{
				ctx: context.Background(),
			},
			rs:      []*entities.Urls{},
			want:    []*entities.Urls{},
			err:     nil,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m.On("ListUrl", tt.args.ctx).Return(tt.rs, tt.err).Once()
			got, err := client.ListUrl(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("liteDB.ListUrl() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("liteDB.ListUrl() = %v, want %v", got, tt.want)
			}
			m.AssertExpectations(t)
		})
	}
}

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
		{
			name: "case 2",
			args: args{
				ctx:       context.Background(),
				shortCode: "",
			},
			rs:      nil,
			want:    nil,
			err:     nil,
			wantErr: false,
		},
		{
			name: "case 3",
			args: args{
				ctx:       context.Background(),
				shortCode: "nYwnyRl12",
			},
			rs:      nil,
			want:    nil,
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

func Test_liteDB_UpdateUrl(t *testing.T) {
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
		wantErr bool
		err     error
	}{
		{
			name: "case 1",
			args: args{
				ctx: context.Background(),
				url: &entities.Urls{
					ShortCode:    "nYwnyRl",
					FullUrl:      "https://www.thepolyglotdeveloper.com/2016/12/create-a-url-shortener-with-golang-and-couchbase-nosql/",
					Expiry:       60,
					NumberOfHits: 6,
				},
			},
			wantErr: false,
			err:     nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m.On("UpdateUrl", tt.args.ctx, tt.args.url).Return(tt.err).Once()
			if err := client.UpdateUrl(tt.args.ctx, tt.args.url); (err != nil) != tt.wantErr {
				t.Errorf("liteDB.UpdateUrl() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_liteDB_DeleteUrl(t *testing.T) {
	type fields struct {
		db *sql.DB
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
				shortCode: "JY3VgQD",
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
				shortCode: "JY3VgQD123",
			},
			wantErr: true,
			err:     errors.New("sql error"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m.On("DeleteUrl", tt.args.ctx, tt.args.shortCode).Return(tt.err).Once()
			if err := client.DeleteUrl(tt.args.ctx, tt.args.shortCode); (err != nil) != tt.wantErr {
				t.Errorf("liteDB.DeleteUrl() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_liteDB_SearchUrl(t *testing.T) {
	type fields struct {
		db *sql.DB
	}
	type args struct {
		ctx       context.Context
		shortCode string
		fullUrl   string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []*entities.Urls
		rs      []*entities.Urls
		wantErr bool
		err     error
	}{
		{
			name: "case 1",
			args: args{
				ctx:       context.Background(),
				shortCode: "",
				fullUrl:   "",
			},
			rs: []*entities.Urls{
				&entities.Urls{
					ShortCode:    "nYwnyRl",
					FullUrl:      "https://www.thepolyglotdeveloper.com/2016/12/create-a-url-shortener-with-golang-and-couchbase-nosql/",
					Expiry:       60,
					NumberOfHits: 6,
				},
			},
			want: []*entities.Urls{
				&entities.Urls{
					ShortCode:    "nYwnyRl",
					FullUrl:      "https://www.thepolyglotdeveloper.com/2016/12/create-a-url-shortener-with-golang-and-couchbase-nosql/",
					Expiry:       60,
					NumberOfHits: 6,
				},
			},
			err:     nil,
			wantErr: false,
		},
		{
			name: "case 2",
			args: args{
				ctx:       context.Background(),
				shortCode: "nYwnyRl",
				fullUrl:   "",
			},
			rs: []*entities.Urls{
				&entities.Urls{
					ShortCode:    "nYwnyRl",
					FullUrl:      "https://www.thepolyglotdeveloper.com/2016/12/create-a-url-shortener-with-golang-and-couchbase-nosql/",
					Expiry:       60,
					NumberOfHits: 6,
				},
			},
			want: []*entities.Urls{
				&entities.Urls{
					ShortCode:    "nYwnyRl",
					FullUrl:      "https://www.thepolyglotdeveloper.com/2016/12/create-a-url-shortener-with-golang-and-couchbase-nosql/",
					Expiry:       60,
					NumberOfHits: 6,
				},
			},
			err:     nil,
			wantErr: false,
		},
		{
			name: "case 3",
			args: args{
				ctx:       context.Background(),
				shortCode: "",
				fullUrl:   "https://www.thepolyglotdeveloper.com/2016/12/create-a-url-shortener-with-golang-and-couchbase-nosql/",
			},
			rs: []*entities.Urls{
				&entities.Urls{
					ShortCode:    "nYwnyRl",
					FullUrl:      "https://www.thepolyglotdeveloper.com/2016/12/create-a-url-shortener-with-golang-and-couchbase-nosql/",
					Expiry:       60,
					NumberOfHits: 6,
				},
			},
			want: []*entities.Urls{
				&entities.Urls{
					ShortCode:    "nYwnyRl",
					FullUrl:      "https://www.thepolyglotdeveloper.com/2016/12/create-a-url-shortener-with-golang-and-couchbase-nosql/",
					Expiry:       60,
					NumberOfHits: 6,
				},
			},
			err:     nil,
			wantErr: false,
		},
		{
			name: "case 4",
			args: args{
				ctx:       context.Background(),
				shortCode: "nYwnyRl",
				fullUrl:   "https://www.thepolyglotdeveloper.com/2016/12/create-a-url-shortener-with-golang-and-couchbase-nosql/",
			},
			rs: []*entities.Urls{
				&entities.Urls{
					ShortCode:    "nYwnyRl",
					FullUrl:      "https://www.thepolyglotdeveloper.com/2016/12/create-a-url-shortener-with-golang-and-couchbase-nosql/",
					Expiry:       60,
					NumberOfHits: 6,
				},
			},
			want: []*entities.Urls{
				&entities.Urls{
					ShortCode:    "nYwnyRl",
					FullUrl:      "https://www.thepolyglotdeveloper.com/2016/12/create-a-url-shortener-with-golang-and-couchbase-nosql/",
					Expiry:       60,
					NumberOfHits: 6,
				},
			},
			err:     nil,
			wantErr: false,
		},
		{
			name: "case 5",
			args: args{
				ctx:       context.Background(),
				shortCode: "nYwny",
				fullUrl:   "https://www.thepolyglotdeveloper.com/2016/12/create-a-url-shortener-with-golang-and-couchbase-nosql/",
			},
			rs: []*entities.Urls{
				&entities.Urls{
					ShortCode:    "nYwnyRl",
					FullUrl:      "https://www.thepolyglotdeveloper.com/2016/12/create-a-url-shortener-with-golang-and-couchbase-nosql/",
					Expiry:       60,
					NumberOfHits: 6,
				},
			},
			want: []*entities.Urls{
				&entities.Urls{
					ShortCode:    "nYwnyRl",
					FullUrl:      "https://www.thepolyglotdeveloper.com/2016/12/create-a-url-shortener-with-golang-and-couchbase-nosql/",
					Expiry:       60,
					NumberOfHits: 6,
				},
			},
			err:     nil,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m.On("SearchUrl", tt.args.ctx, tt.args.shortCode, tt.args.fullUrl).Return(tt.rs, tt.err).Once()
			got, err := client.SearchUrl(tt.args.ctx, tt.args.shortCode, tt.args.fullUrl)
			if (err != nil) != tt.wantErr {
				t.Errorf("liteDB.SearchUrl() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("liteDB.SearchUrl() = %v, want %v", got, tt.want)
			}
			m.AssertExpectations(t)
		})
	}
}
