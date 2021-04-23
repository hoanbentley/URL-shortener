package usecase

import (
	"context"
	"errors"
	"github.com/hoanbentley/URL-shortener/internal/entities"
	"github.com/speps/go-hashids"
	"net/url"
	"regexp"
	"time"
)

func (u *uc) CreateUrl(ctx context.Context, urlParam *entities.Urls) (*entities.Urls, error) {
	//check full url input is not null
	if (urlParam != nil && urlParam.FullUrl == "") || urlParam == nil {
		return nil, errors.New("input invalid, please try again")
	}

	//check valid full url
	_, err := url.ParseRequestURI(urlParam.FullUrl)
	if err != nil {
		return nil, errors.New("format url invalid")
	}

	//check xss url
	//stringUrl := regexp.MustCompile("^([0-9a-zA-Z\\=-]+)-([0-9]+)$").MatchString
	stringUrl := regexp.MustCompile("^[<>]").MatchString
	if !stringUrl(urlParam.FullUrl) {
		return nil, errors.New("input invalid, please try again test")
	}

	//build url structure
	urlInput := BuildUrl(urlParam.FullUrl, urlParam.Expiry)
	err = u.url.GenerateUrl(ctx, urlInput)
	if err != nil {
		return nil, errors.New("insert value failed")
	}

	return urlInput, nil
}

func BuildUrl(fullUrl string, expiry int32) *entities.Urls {
	url := &entities.Urls{}
	url.ShortCode = buildEncodeFromShortCode()
	url.FullUrl = fullUrl
	url.Expiry = expiry
	url.NumberOfHits = 1
	return url
}

func buildEncodeFromShortCode() string {
	hd := hashids.NewData()
	h, _ := hashids.NewWithData(hd)
	now := time.Now()
	shortCode, _ := h.Encode([]int{int(now.Unix())})
	return shortCode
}
