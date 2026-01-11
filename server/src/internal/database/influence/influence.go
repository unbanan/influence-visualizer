package influence

import (
	"fmt"

	"contest-influence/server/internal/config"
	"contest-influence/server/internal/simulation_types"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/samber/lo"
	"google.golang.org/protobuf/proto"
)

type InfluenceDBRepo interface {
	Close() error
	Register(int64, string) error
	GetSimulationPlayers(uuid.UUID) ([]User, error)
	GetSimulation(uuid.UUID) (*simulation_types.Simulation, error)
}

type InfluenceDBRepoImpl struct {
	db *sqlx.DB
}

func NewInfluenceDBRepo(config config.PgConfig) (*InfluenceDBRepoImpl, error) {
	db, err := sqlx.Connect("postgres", config.DSN())
	if err != nil {
		return nil, fmt.Errorf("Failed to open database: %w", err)
	}

	db.SetMaxOpenConns(config.MaxOpenConns)
	db.SetMaxIdleConns(config.MaxIdleConns)
	db.SetConnMaxLifetime(config.ConnMaxLifetime)

	return &InfluenceDBRepoImpl{
		db: db,
	}, nil
}

func (r *InfluenceDBRepoImpl) Close() error {
	return r.db.Close()
}

func (r *InfluenceDBRepoImpl) Register(id int64, name string) error {
	_, err := r.db.Exec("INSERT INTO influence.users (id, name) VALUES ($1, $2)", id, name)
	return err
}

type User struct {
	ID   int64  `db:"id"`
	Name string `db:"name"`
}

func (r *InfluenceDBRepoImpl) GetSimulationPlayers(id uuid.UUID) ([]User, error) {
	users := make([]User, 0)
	err := r.db.Select(
		users,
		`
		SELECT
			u.*
		FROM
				(
					SELECT
						* 
					FROM
						influence.users_simulations
					WHERE sid = $1
				) AS i
			JOIN 
				influence.users AS u
			ON
				i.uid = u.id
		ORDER BY
			i.order
		`,
		id,
	)
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (r *InfluenceDBRepoImpl) GetSimulation(id uuid.UUID) (*simulation_types.Simulation, error) {
	var err error
	users, err := r.GetSimulationPlayers(id)

	if err != nil {
		return nil, err
	}

	simulation := simulation_types.NewSimulation()
	simulation.Players = lo.Map(users, func(u User, _ int) string {
		return u.Name
	})

	simdata := make([]byte, 0)
	mapdata := make([]byte, 0)

	err = r.db.QueryRowx(
		`
		SELECT
			s.data,
			s.queued_at,
			s.started_at,
			s.finished_at,
			s.state,
			m.data AS map_data,
			m.name AS map_name
		FROM
				(
					SELECT
						*
					FROM
						influence.simulations
					WHERE
						id = $1
				) AS s
			JOIN
				influence.maps AS m
			ON
				s.map_id = m.id
		WHERE
			id = $1
		`,
		id,
	).Scan(
		simdata,
		&simulation.QueuedAt,
		&simulation.StartedAt,
		&simulation.FinishedAt,
		&simulation.State,
		&mapdata,
		&simulation.Map.Name,
	)

	if err != nil {
		return nil, err
	}

	if err := proto.Unmarshal(simdata, &simulation.Data); err != nil {
		return nil, err
	}

	if err := proto.Unmarshal(mapdata, &simulation.Map.Data); err != nil {
		return nil, err
	}

	return simulation, nil
}
