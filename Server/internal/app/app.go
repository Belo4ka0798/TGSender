package app

import (
	"github.com/go-chi/chi/v5"
	"tgsender/internal/question"
	"tgsender/internal/question/repo"
	"tgsender/internal/question/usecase"
	"tgsender/internal/server"
	"tgsender/pkg/logging"
)

type App struct {
	server *server.Server
	repo   repo.Repo
	log    *logging.Logger
}

func NewApp(logger *logging.Logger, srv *server.Server, rp repo.Repo) *App {
	return &App{
		server: srv,
		repo:   rp,
		log:    logger,
	}
}

func (a *App) Init() error {
	router := chi.NewRouter()

	qUseCase := usecase.NewUseCase(a.log, &a.repo)
	qHandler := question.NewHandler(a.log, qUseCase)
	qHandler.MainRoutesHandlers(router)

	a.server.Init(router)
	return a.server.Run()
}
