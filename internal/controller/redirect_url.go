package controller

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

func (s *ToDoService) RedirectUrl(resp http.ResponseWriter, req *http.Request) {
	//Redirect url with short code
	param := mux.Vars(req)
	shortCode := param["id"]
	url, err := s.UseCase.RedirectUrl(req.Context(), shortCode)
	resp.Header().Set("Content-Type", "application/json")
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(resp).Encode(map[string]string{
			"error": err.Error(),
		})
		return
	}
	http.Redirect(resp, req, url.FullUrl, 302)
}
