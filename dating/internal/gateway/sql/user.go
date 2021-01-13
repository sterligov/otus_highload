package sql

import (
	"context"
	"time"

	"go.uber.org/zap"

	"github.com/sterligov/otus_highload/dating/internal/domain"

	"github.com/jmoiron/sqlx"
)

const (
	MALE   = 'M'
	FEMALE = 'F'
)

type (
	User struct {
		ID        int64     `db:"id"`
		FirstName string    `db:"first_name"`
		LastName  string    `db:"last_name"`
		Birthday  time.Time `db:"birthday"`
		Email     string    `db:"email"`
		Interests string    `db:"interests"`
		Sex       byte      `db:"sex"`
		CityID    int64     `db:"city_id"`
	}
)

type UserGateway struct {
	logger *zap.Logger
	db     *sqlx.DB
}

func NewUserGateway(db *sqlx.DB) *UserGateway {
	return &UserGateway{
		logger: zap.L().Named("user gateway"),
		db:     db,
	}
}

func (ug *UserGateway) FindByID(ctx context.Context, id int64) (*domain.User, error) {
	const query = `
SELECT u.*,
       ci.id as city_id,
	   ci.name as city_name,
       co.id as country_id,
	   co.name as country_name
FROM user u
JOIN city ci ON ci.id = u.city_id
JOIN country co ON ci.id = ci.country_id
WHERE id = :id`

	return nil, nil
}

func (ug *UserGateway) FindByEmailAndPassword(ctx context.Context, email, password string) (*domain.User, error) {
	const query = `
SELECT *
FROM user
WHERE email = :email AND password = :password`

	u := &domain.User{}

	row := ug.db.QueryRowxContext(ctx, query)
	if err := row.StructScan(&u); err != nil {
		return nil, domain.ErrNotFound
	}

	return nil, nil
}

func (ug *UserGateway) FindAfterID(ctx context.Context, id, limit int64) ([]*domain.User, error) {
	const query = `
SELECT u.*,
       ci.name as city_name,
       co.name as country_name
FROM user u
JOIN city ci ON ci.id = u.city_id
JOIN country co on ci.country_id = co.id
WHERE u.id > :id
LIMIT :limit`

	//users := make([]*domain.User, limit)
	//
	//rows, err := ug.db.QueryxContext(ctx, query, id, limit)
	//if err != nil {
	//	return nil, domain.ErrNotFound
	//}
	//
	//for rows.Next() {
	//
	//}

	return nil, nil
}
