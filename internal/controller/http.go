package controller

import (
	"encoding/json"
	"github.com/hoanbentley/URL-shortener/internal/entities"
	"github.com/hoanbentley/URL-shortener/internal/transport"
	"github.com/hoanbentley/URL-shortener/internal/usecase"
	"log"
	"net/http"
)

type ToDoService struct {
	UseCase usecase.UseCase
	Trans   *transport.Transport
}

func NewToDoService(db string) *ToDoService {
	todo := &ToDoService{
		UseCase: usecase.NewUc(db),
		Trans:   transport.NewTransport(),
	}
	return todo
}

func (s *ToDoService) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	log.Println(req.Method, req.URL.Path)
	switch req.URL.Path {
	case "/url":
		switch req.Method {
		case http.MethodPost:
			s.CreateUrl(resp, req)
		case http.MethodGet:
			s.ListUrl(resp, req)
		}
	}
}

func (s *ToDoService) CreateUrl(resp http.ResponseWriter, req *http.Request) {
	//Build Url structure
	url, err := s.Trans.BuildUrl(req)
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(resp).Encode(map[string]string{
			"error": err.Error(),
		})
		return
	}

	//Call create Url via use case layer
	err = s.UseCase.CreateUrl(req.Context(), url)
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(resp).Encode(map[string]string{
			"error": err.Error(),
		})
		return
	}
	json.NewEncoder(resp).Encode(map[string]*entities.Urls{
		"data": nil,
	})
}

func (s *ToDoService) ListUrl(resp http.ResponseWriter, req *http.Request) {
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
