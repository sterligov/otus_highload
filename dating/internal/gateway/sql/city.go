package sql

import (
	"context"
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/sterligov/otus_highload/dating/internal/domain"
	"go.uber.org/zap"
)

type City struct {
	ID        int64  `db:"id"`
	Name      string `db:"name"`
	CountryID int64  `db:"country_id"`
	Country   `db:"country"`
}

type CityGateway struct {
	db     *sqlx.DB
	logger *zap.Logger
}

func NewCityGateway(db *sqlx.DB) *CityGateway {
	return &CityGateway{
		db:     db,
		logger: zap.L().Named("city gateway"),
	}
}

func (cg *CityGateway) FindAll(ctx context.Context) ([]*domain.City, error) {
	query := `
SELECT c.*,
       co.id "country.id",
       co.name "country.name"
FROM city c
JOIN country co ON co.id = c.country_id`

	rows, err := cg.db.QueryxContext(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("find all: %w", err)
	}
	defer func() {
		if err := rows.Close(); err != nil {
			cg.logger.Warn("error close failed", zap.Error(err))
		}
	}()

	var cities []*domain.City

	for rows.Next() {
		city := new(City)
		if err := rows.StructScan(&city); err != nil {
			return nil, fmt.Errorf("rows scan: %w", err)
		}

		cities = append(cities, toDomainCity(city))
	}

	return cities, nil
}

func toDomainCity(c *City) *domain.City {
	return &domain.City{
		ID:   c.ID,
		Name: c.Name,
		Country: &domain.Country{
			ID:   c.Country.ID,
			Name: c.Country.Name,
		},
	}
}
