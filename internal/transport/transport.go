package transport

import (
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

func (t *Transport) BuildUrl(req *http.Request) (*entities.Urls, error) {
	url := &entities.Urls{}
	err := json.NewDecoder(req.Body).Decode(url)
	defer req.Body.Close()
	if err != nil {
		return nil, err
	}

	//build short code with encrypt data
	hd := hashids.NewData()
	h, _ := hashids.NewWithData(hd)
	now := time.Now()
	shortCode, _ := h.Encode([]int{int(now.Unix())})

	//set again in url structure
	url.ShortCode = shortCode
	url.Expiry = "60"
	url.NumberOfHits = "1"
	return url, nil
}
