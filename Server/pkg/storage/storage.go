package storage

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/spf13/viper"
	"log"
	"tgsender/pkg/utils"
	"time"
)

func NewStorage(ctx context.Context) (pool *pgxpool.Pool, err error) {
	dns := fmt.Sprintf("postgres://%s:%s@%s:%s/%s",
		viper.Get("Storage.username"),
		viper.Get("Storage.password"),
		viper.Get("Storage.host"),
		viper.Get("Storage.port"),
		viper.Get("Storage.database"))
	attempt := viper.GetInt("Storage.maxAttempt")
	err = utils.DoWhitTries(func() error {
		ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
		defer cancel()

		pool, err = pgxpool.New(ctx, dns)
		if err != nil {
			return err
		}
		err = pool.Ping(ctx)
		if err != nil {
			return err
		}

		return nil
	}, attempt, 5*time.Second)
	if err != nil {
		log.Fatal(err, "(error do with tries postgresql)")
	}
	return pool, nil
}
