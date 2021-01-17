package domain

import (
	"context"
)

type (
	City struct {
		ID      int64
		Name    string
		Country *Country
	}

	CityGateway interface {
		FindAll(ctx context.Context) ([]*City, error)
	}
)
