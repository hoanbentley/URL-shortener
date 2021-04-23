package controller

import (
	"github.com/hoanbentley/URL-shortener/internal/transport"
	"github.com/hoanbentley/URL-shortener/internal/usecase"
)

type ToDoService struct {
	JWTKey  string
	UseCase usecase.UseCase
	Trans   *transport.Transport
}

func NewToDoService() *ToDoService {
	todo := &ToDoService{
		JWTKey:  "wqGyEBBfPK9w3Lxw",
		UseCase: usecase.NewUc(),
		Trans:   transport.NewTransport(),
	}
	return todo
}
