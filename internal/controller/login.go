package controller

import (
	"encoding/json"
	"net/http"
)

func (s *ToDoService) GetAuthToken(resp http.ResponseWriter, req *http.Request) {
	user := s.Trans.GetValue(req, "user_id")
	pass := s.Trans.GetValue(req, "password")

	if !s.UseCase.Validate(req.Context(), user, pass) {
		resp.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(resp).Encode(map[string]string{
			"error": "incorrect user_id/pwd",
		})
		return
	}

	token, err := s.UseCase.CreateToken(user.String, s.JWTKey)
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(resp).Encode(map[string]string{
			"error": err.Error(),
		})
		return
	}

	resp.Header().Set("Content-Type", "application/json")
	json.NewEncoder(resp).Encode(map[string]string{
		"data": token,
	})
}
