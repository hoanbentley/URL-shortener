package controller

import (
	"encoding/json"
	"github.com/hoanbentley/URL-shortener/internal/entities"
	"net/http"
)

func (s *ToDoService) CreateUrl(resp http.ResponseWriter, req *http.Request) {
	resp.Header().Set("Content-Type", "application/json")

	//Execute request
	urlInput := &entities.Urls{}
	err := json.NewDecoder(req.Body).Decode(urlInput)
	defer req.Body.Close()
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(resp).Encode(map[string]string{
			"error": err.Error(),
		})
		return
	}

	//Call create Url via use case layer
	url, err := s.UseCase.CreateUrl(req.Context(), urlInput)
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(resp).Encode(map[string]string{
			"error": err.Error(),
		})
		return
	}

	json.NewEncoder(resp).Encode(map[string]*entities.Urls{
		"data": url,
	})
}
