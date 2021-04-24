package usecase

import (
	"context"
	"github.com/hoanbentley/URL-shortener/internal/entities"
	"github.com/hoanbentley/URL-shortener/internal/storages"
	"reflect"
	"testing"
)

func Test_uc_CreateUrl(t *testing.T) {
	type fields struct {
		url storages.Url
	}
	type args struct {
		ctx      context.Context
		urlParam *entities.Urls
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *entities.Urls
		wantErr bool
	}{
		{
			name: "case 1",
			args: args{
				ctx:      nil,
				urlParam: nil,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "case 2",
			args: args{
				ctx: nil,
				urlParam: &entities.Urls{
					ShortCode:    "",
					FullUrl:      "",
					Expiry:       0,
					NumberOfHits: 0,
				},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "case 3",
			args: args{
				ctx: nil,
				urlParam: &entities.Urls{
					ShortCode:    "",
					FullUrl:      "www.abc.com",
					Expiry:       0,
					NumberOfHits: 0,
				},
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &uc{
				url: tt.fields.url,
			}
			got, err := u.CreateUrl(tt.args.ctx, tt.args.urlParam)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateUrl() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateUrl() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBuildUrl(t *testing.T) {
	type args struct {
		fullUrl string
		expiry  int32
	}
	tests := []struct {
		name string
		args args
		want *entities.Urls
	}{
		// TODO: Add test cases.
		{
			name: "case",
			args: args{
				fullUrl: "",
				expiry:  60,
			},
			want: &entities.Urls{
				ShortCode:    buildEncodeFromShortCode(),
				FullUrl:      "",
				Expiry:       60,
				NumberOfHits: 1,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BuildUrl(tt.args.fullUrl, tt.args.expiry); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BuildUrl() = %v, want %v", got, tt.want)
			}
		})
	}
}
