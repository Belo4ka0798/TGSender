package usecase

import (
	"github.com/go-chi/chi/v5"
	"github.com/gorilla/schema"
	"net/http"
	"strconv"
	question "tgsender/internal/question"
	"tgsender/internal/question/model"
	"tgsender/pkg/logging"
	"time"
)

type useCase struct {
	repo question.Repo
	log  *logging.Logger
}

var decoder = schema.NewDecoder()

func NewUseCase(log *logging.Logger, repo question.Repo) question.UseCase {
	return &useCase{log: log, repo: repo}
}

func (u *useCase) ShowForm() error {
	return nil
}

func (u *useCase) ListArticlesByDate(r *http.Request) (*[]model.Question, error) {
	day, err := strconv.Atoi(chi.URLParam(r, "day"))
	month, err := strconv.Atoi(chi.URLParam(r, "month"))
	year, err := strconv.Atoi(chi.URLParam(r, "year"))
	if err != nil {
		return nil, err
	}
	t := &time.Time{}
	t.AddDate(year, month, day)
	resQ, err := u.repo.RListArticlesByDate(r.Context(), t)
	if err != nil {
		return nil, err
	}
	return resQ, nil
}

func (u *useCase) GetQuestion(r *http.Request) (*model.Question, error) {
	q := chi.URLParam(r, "questionID")
	qi, err := strconv.Atoi(q)

	if err != nil {
		return nil, err
	}
	resQues, err := u.repo.RGetQuestion(r.Context(), qi)
	if err != nil {
		return nil, err
	}
	return resQues, nil
}

func (u *useCase) CreateQuestion(r *http.Request) error {
	q, err := u.parsForm(r)
	if err != nil {
		u.log.Println(err)
		return err
	}

	err = u.repo.RCreateQuestion(r.Context(), q)
	if err != nil {
		u.log.Println(err)
		return err
	}
	return nil

	//TODO implement me
	panic("implement me")
}

func (u *useCase) DeleteQuestion(r *http.Request) error {
	q := chi.URLParam(r, "questionID")
	qi, err := strconv.Atoi(q)
	if err != nil {
		u.log.Println(err)
		return err
	}

	err = u.repo.RDeleteQuestion(r.Context(), qi)
	if err != nil {
		u.log.Println(err)
		return err
	}

	return nil
}

func (u *useCase) parsForm(r *http.Request) (*model.Question, error) {
	questionData := new(model.Question)
	if err := r.ParseForm(); err != nil {
		u.log.Println(err)
		return nil, err
	}
	err := decoder.Decode(questionData, r.PostForm)
	if err != nil {
		u.log.Println(err)
		return nil, err
	}
	return questionData, nil
}
