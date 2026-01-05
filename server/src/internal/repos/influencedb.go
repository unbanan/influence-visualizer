package repos

import (
	"fmt"

	"contest-influence/server/internal/config"

	"github.com/jmoiron/sqlx"
)

type InfluenceDBRepo struct {
	db *sqlx.DB
}

func NewInfluenceDBRepo(config config.PgConfig) *InfluenceDBRepo {
	db, err := sqlx.Open("postgres", config.DSN())
	if err != nil {
		panic(fmt.Errorf("failed to open database: %w", err))
	}
	return &InfluenceDBRepo{
		db: db,
	}
}

func (r *InfluenceDBRepo) Close() error {
	return r.db.Close()
}

func (r *InfluenceDBRepo) Register(id int64, name string) {
	_, err := r.db.Exec("INSERT INTO users (id, name) VALUES ($1, $2)", id, name)
	if err != nil {
		panic(err)
	}
}
