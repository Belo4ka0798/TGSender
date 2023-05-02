package question

import (
	"context"
	"tgsender/internal/question/model"
	"time"
)

type Repo interface {
	RListArticlesByDate(ctx context.Context, t *time.Time) (*[]model.Question, error)
	RGetQuestion(ctx context.Context, id int) (*model.Question, error)
	RCreateQuestion(ctx context.Context, qu *model.Question) error
	RDeleteQuestion(ctx context.Context, id int) error
}
