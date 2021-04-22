package controller

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/hoanbentley/URL-shortener/internal/entities"
	"github.com/hoanbentley/URL-shortener/internal/transport"
	"github.com/hoanbentley/URL-shortener/internal/usecase"
	"log"
	"net/http"
)

type ToDoService struct {
	JWTKey  string
	UseCase usecase.UseCase
	Trans   *transport.Transport
}

func NewToDoService(db string) *ToDoService {
	todo := &ToDoService{
		JWTKey:  "wqGyEBBfPK9w3Lxw",
		UseCase: usecase.NewUc(db),
		Trans:   transport.NewTransport(),
	}
	return todo
}

func (s *ToDoService) getAuthToken(resp http.ResponseWriter, req *http.Request) {
	user := s.Trans.GetValue(req, "user_id")
	pass := s.Trans.GetValue(req, "password")
	if !s.UseCase.Validate(req.Context(), user, pass) {
		resp.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(resp).Encode(map[string]string{
			"error": "incorrect user_id/pwd",
		})
		return
	}
	resp.Header().Set("Content-Type", "application/json")

	token, err := s.UseCase.CreateToken(user.String, s.JWTKey)
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(resp).Encode(map[string]string{
			"error": err.Error(),
		})
		return
	}

	json.NewEncoder(resp).Encode(map[string]string{
		"data": token,
	})
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
		"data": url,
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
	resp.Header().Set("Content-Type", "application/json")
	json.NewEncoder(resp).Encode(map[string][]*entities.Urls{
		"data": url,
	})
}

func (s *ToDoService) RedirectUrl(resp http.ResponseWriter, req *http.Request) {
	param := mux.Vars(req)
	shortCode := param["id"]
	log.Println("Short_code:", shortCode)
	url, err := s.UseCase.RedirectUrl(req.Context(), shortCode)
	resp.Header().Set("Content-Type", "application/json")
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(resp).Encode(map[string]string{
			"error": err.Error(),
		})
		return
	}
	/*json.NewEncoder(resp).Encode(map[string]*entities.Urls{
		"data": url,
	})*/
	log.Println("Full URL:", url.FullUrl)
	http.Redirect(resp, req, url.FullUrl, 301)
}

func (s *ToDoService) DeleteUrl(resp http.ResponseWriter, req *http.Request) {
	param := mux.Vars(req)
	shortCode := param["id"]
	log.Println("shortCode:", shortCode)
	resp.Header().Set("Content-Type", "application/json")
}
