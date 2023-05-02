package repo

import (
	"context"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"tgsender/internal/question/model"
	"tgsender/pkg/logging"
	"tgsender/pkg/utils"
	"time"
)

type Client interface {
	Exec(ctx context.Context, sql string, arguments ...any) (pgconn.CommandTag, error)
	Query(ctx context.Context, sql string, args ...any) (pgx.Rows, error)
	QueryRow(ctx context.Context, sql string, args ...any) pgx.Row
	Begin(ctx context.Context) (pgx.Tx, error)
}

type Repo struct {
	client Client
	logger *logging.Logger
}

func NewRepo(client Client, logger *logging.Logger) Repo {
	return Repo{
		client: client,
		logger: logger,
	}
}

func (re *Repo) RListArticlesByDate(ctx context.Context, t *time.Time) (*[]model.Question, error) {
	return nil, nil
}

func (re *Repo) RGetQuestion(ctx context.Context, id int) (*model.Question, error) {
	q := `select * from questions where $1 = id;`
	utils.FormatQuery(q)
	res := &model.Question{}
	err := re.client.QueryRow(ctx, q, id).Scan(&res.ID, &res.Email, &res.Header, &res.Message, &res.Answer, &res.Date, &res.Status)
	if err != nil {
		re.logger.Println(err)
		return nil, err
	}
	return res, nil
}

func (re *Repo) RCreateQuestion(ctx context.Context, qu *model.Question) error {
	tx, err := re.client.Begin(ctx)
	if err != nil {
		re.logger.Println(err)
		return err
	}
	q := `insert into questions (email, header, message)
			values ($1, $2, $3)
		  returning questions.id;`
	utils.FormatQuery(q)
	var res int64
	err = tx.QueryRow(ctx, q, qu.Email, qu.Header, qu.Message).Scan(&res)
	if err != nil {
		if err := tx.Rollback(ctx); err != nil {
			re.logger.Println(err)
			return err
		}
		re.logger.Println(err)
		return err
	}

	err = tx.Commit(ctx)
	if err != nil {
		re.logger.Println(err)
		return err
	}
	return nil
}

func (re *Repo) RDeleteQuestion(ctx context.Context, id int) error {
	return nil
}
