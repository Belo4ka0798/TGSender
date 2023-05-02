package question

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"html/template"
	"net/http"
	"tgsender/internal/handlers"
	"tgsender/pkg/logging"
)

type handler struct {
	log *logging.Logger
	uc  UseCase
}

func NewHandler(log *logging.Logger, uc UseCase) handlers.Handler {
	return &handler{log: log, uc: uc}
}

func (h *handler) MainRoutesHandlers(r chi.Router) {
	r.Route("/question", func(r chi.Router) {
		r.Post("/", h.createQuestion)
		r.Get("/", h.showForm)
		r.Get("/{day}-{month}-{year}", h.listArticlesByDate)
		r.Route("/{questionID}", func(r chi.Router) {
			r.Get("/", h.getQuestion)
			r.Delete("/", h.deleteQuestion)
		})
	})
}

func (h *handler) deleteQuestion(w http.ResponseWriter, r *http.Request) {
	err := h.uc.DeleteQuestion(r)
	if err != nil {
		h.log.Println(err)
		w.WriteHeader(400)
		return
	}
	fmt.Fprintf(w, "Question about id delete!\n")
	w.WriteHeader(200)
}

func (h *handler) getQuestion(w http.ResponseWriter, r *http.Request) {
	q, err := h.uc.GetQuestion(r)
	if err != nil {
		h.log.Println(err)
		w.WriteHeader(400)
		return
	}
	fmt.Fprintf(w, "Question about id:\n%v", q)
	w.WriteHeader(200)
}

func (h *handler) listArticlesByDate(w http.ResponseWriter, r *http.Request) {
	q, err := h.uc.ListArticlesByDate(r)
	if err != nil {
		h.log.Println(err)
		w.WriteHeader(400)
		return
	}
	fmt.Fprintf(w, "Questions for the specified date:\n%v", q)
	w.WriteHeader(200)
}

func (h *handler) showForm(w http.ResponseWriter, r *http.Request) {
	err := h.uc.ShowForm()
	t, err := template.ParseFiles("./static/form.html")
	t.Execute(w, nil)
	if err != nil {
		h.log.Println(err)
		w.WriteHeader(400)
	}
	if err != nil {
		h.log.Println(err)
		w.WriteHeader(400)
		return
	}
	fmt.Fprintf(w, "Feedback Form:\n")
	w.WriteHeader(200)
}

func (h *handler) createQuestion(w http.ResponseWriter, r *http.Request) {
	err := h.uc.CreateQuestion(r)
	if err != nil {
		h.log.Println(err)
		w.WriteHeader(400)
		return
	}
	fmt.Fprintf(w, "Question accepted!\nWe will send a response to the specified mail within 7 days.")
	w.WriteHeader(200)
}
