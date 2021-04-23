package controller

import (
	"context"
	"encoding/json"
	"github.com/hoanbentley/URL-shortener/internal/entities"
	"github.com/hoanbentley/URL-shortener/internal/transport"
	"net/http"
)

func (s *ToDoService) ListUrl(resp http.ResponseWriter, req *http.Request) {
	resp.Header().Set("Content-Type", "application/json")
	//Get Authorization
	token := s.Trans.GetToken(req)
	id, ok := s.UseCase.ValidToken(token, s.JWTKey)
	req = req.WithContext(context.WithValue(req.Context(), transport.UserAuthKey(0), id))
	if !ok {
		resp.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(resp).Encode(map[string]string{
			"error": "Please get token to authorization",
		})
		return
	}

	//Get list url
	url, err := s.UseCase.ListUrl(req.Context())
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(resp).Encode(map[string]string{
			"error": err.Error(),
		})
		return
	}

	json.NewEncoder(resp).Encode(map[string][]*entities.Urls{
		"data": url,
	})
}
