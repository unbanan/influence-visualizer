package repos

import (
	"fmt"

	"contest-influence/server/internal/config"

	"github.com/jmoiron/sqlx"
)

type InfluenceDBRepo interface {
	Close() error
	Register(id int64, name string) error
}

type InfluenceDBRepoImpl struct {
	db *sqlx.DB
}

func NewInfluenceDBRepo(config config.PgConfig) (*InfluenceDBRepoImpl, error) {
	db, err := sqlx.Open("postgres", config.DSN())
	if err != nil {
		return nil, fmt.Errorf("Failed to open database: %w", err)
	}

	return &InfluenceDBRepoImpl{
		db: db,
	}, nil
}

func (r *InfluenceDBRepoImpl) Close() error {
	return r.db.Close()
}

func (r *InfluenceDBRepoImpl) Register(id int64, name string) error {
	_, err := r.db.Exec("INSERT INTO users (id, name) VALUES ($1, $2)", id, name)
	return err
}
