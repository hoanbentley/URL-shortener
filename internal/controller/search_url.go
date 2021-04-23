package controller

import (
	"encoding/json"
	"github.com/hoanbentley/URL-shortener/internal/entities"
	"net/http"
)

func (s *ToDoService) SearchUrl(resp http.ResponseWriter, req *http.Request) {

	url, err := s.UseCase.ListUrl(req.Context())
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(resp).Encode(map[string]string{
			"error": err.Error(),
		})
		return
	}
	resp.Header().Set("Content-Type", "application/json")
	json.NewEncoder(resp).Encode(map[string][]*entities.Urls{
		"data": url,
	})
}
