package usecase

import (
	"context"
	"database/sql"
	"errors"
	"testing"

	"github.com/hoanbentley/URL-shortener/internal/storages"
)

func Test_uc_CreateToken(t *testing.T) {
	type fields struct {
		url storages.Url
	}
	type args struct {
		id     string
		jwtKey string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
		rs      string
		err     error
	}{
		{
			name: "case 1",
			args: args{
				id:     "admin",
				jwtKey: "wqGyEBBfPK9w3Lxw",
			},
			want:    "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2MTkyNzM5OTEsInVzZXJfaWQiOiJhZG1pbiJ9.ezow_d9HfsmYjZIxiiZPmvDgVou07AKc44S7TA2k1Tk",
			wantErr: false,
			rs:      "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2MTkyNzM5OTEsInVzZXJfaWQiOiJhZG1pbiJ9.ezow_d9HfsmYjZIxiiZPmvDgVou07AKc44S7TA2k1Tk",
			err:     nil,
		},
		{
			name: "case 2",
			args: args{
				id:     "",
				jwtKey: "wqGyEBBfPK9w3Lxw",
			},
			want:    "",
			wantErr: true,
			rs:      "",
			err:     errors.New(""),
		},
		{
			name: "case 3",
			args: args{
				id:     "admin",
				jwtKey: "",
			},
			want:    "",
			wantErr: true,
			rs:      "",
			err:     errors.New(""),
		},
		{
			name: "case 4",
			args: args{
				id:     "",
				jwtKey: "",
			},
			want:    "",
			wantErr: true,
			rs:      "",
			err:     errors.New(""),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m.On("CreateToken", tt.args.id, tt.args.jwtKey).Return(tt.rs, tt.err).Once()
			got, err := client.CreateToken(tt.args.id, tt.args.jwtKey)
			if (err != nil) != tt.wantErr {
				t.Errorf("uc.CreateToken() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("uc.CreateToken() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_uc_Validate(t *testing.T) {
	type fields struct {
		url storages.Url
	}
	type args struct {
		ctx      context.Context
		user     sql.NullString
		password sql.NullString
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
		rs     bool
	}{
		{
			name: "case 1",
			args: args{
				ctx: context.Background(),
				user: sql.NullString{
					String: "admin",
					Valid:  true,
				},
				password: sql.NullString{
					String: "admin",
					Valid:  true,
				},
			},
			want: true,
			rs:   true,
		},
		{
			name: "case 2",
			args: args{
				ctx: context.Background(),
				user: sql.NullString{
					String: "user",
					Valid:  true,
				},
				password: sql.NullString{
					String: "admin",
					Valid:  true,
				},
			},
			want: false,
			rs:   false,
		},
		{
			name: "case 3",
			args: args{
				ctx: context.Background(),
				user: sql.NullString{
					String: "",
					Valid:  true,
				},
				password: sql.NullString{
					String: "",
					Valid:  true,
				},
			},
			want: false,
			rs:   false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m.On("Validate", tt.args.ctx, tt.args.user, tt.args.password).Return(tt.rs).Once()
			if got := client.Validate(tt.args.ctx, tt.args.user, tt.args.password); got != tt.want {
				t.Errorf("uc.Validate() = %v, want %v", got, tt.want)
			}
		})
	}
}
