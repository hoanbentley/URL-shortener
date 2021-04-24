package transport

import (
	"database/sql"
	"net/http"
)

type UserAuthKey int8
type Transport struct {
}

func NewTransport() *Transport {
	return &Transport{}
}

func (t *Transport) GetToken(req *http.Request) string {
	return req.Header.Get("Authorization")
}

func (t *Transport) GetValue(req *http.Request, param string) sql.NullString {
	return sql.NullString{
		String: req.FormValue(param),
		Valid:  true,
	}
}
