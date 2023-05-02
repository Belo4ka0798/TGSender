package question

import (
	"net/http"
	"tgsender/internal/question/model"
)

type UseCase interface {
	ShowForm() error
	ListArticlesByDate(r *http.Request) (*[]model.Question, error)
	GetQuestion(r *http.Request) (*model.Question, error)
	CreateQuestion(r *http.Request) error
	DeleteQuestion(r *http.Request) error
}
