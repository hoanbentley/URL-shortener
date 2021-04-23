package controller

import (
	"context"
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/hoanbentley/URL-shortener/internal/transport"
	"net/http"
)

func (s *ToDoService) DeleteUrl(resp http.ResponseWriter, req *http.Request) {
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

	//Delete url by short code
	param := mux.Vars(req)
	shortCode := param["id"]
	err := s.UseCase.DeleteUrl(req.Context(), shortCode)
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(resp).Encode(map[string]string{
			"error": err.Error(),
		})
		return
	}
	json.NewEncoder(resp).Encode(map[string]string{
		"mess": "Success",
	})
}
