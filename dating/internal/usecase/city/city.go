package city

import (
	"context"

	"github.com/sterligov/otus_highload/dating/internal/domain"

	"go.uber.org/zap"
)

type UseCase struct {
	logger      *zap.Logger
	cityGateway domain.CityGateway
}

func NewCityUseCase(cg domain.CityGateway) *UseCase {
	return &UseCase{
		logger:      zap.L().Named("city use case"),
		cityGateway: cg,
	}
}

func (uc *UseCase) FindAll(ctx context.Context) ([]*domain.City, error) {
	cities, err := uc.cityGateway.FindAll(ctx)
	if err != nil {
		return nil, err
	}

	return cities, nil
}
