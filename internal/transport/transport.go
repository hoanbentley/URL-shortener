package transport

import (
	"database/sql"
	"encoding/json"
	"github.com/hoanbentley/URL-shortener/internal/entities"
	"github.com/speps/go-hashids"
	"net/http"
	"time"
)

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

func (t *Transport) BuildEncodeFromShortCode() string {
	hd := hashids.NewData()
	h, _ := hashids.NewWithData(hd)
	now := time.Now()
	shortCode, _ := h.Encode([]int{int(now.Unix())})
	return shortCode
}

func (t *Transport) BuildUrl(req *http.Request) (*entities.Urls, error) {
	url := &entities.Urls{}
	err := json.NewDecoder(req.Body).Decode(url)
	defer req.Body.Close()
	if err != nil {
		return nil, err
	}

	//check full_url is null
	if url.FullUrl == "" {
		return nil, err
	}

	//set again in url structure
	url.ShortCode = t.BuildEncodeFromShortCode()
	url.NumberOfHits = 1
	return url, nil
}
