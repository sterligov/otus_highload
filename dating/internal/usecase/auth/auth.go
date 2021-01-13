package auth

import (
	"context"
	"errors"

	"github.com/sterligov/otus_highload/dating/internal/domain"
	"github.com/sterligov/otus_highload/dating/internal/util/crypt"
	"go.uber.org/zap"
)

type UseCase struct {
	userGateway domain.UserGateway
	logger      *zap.Logger
}

func NewAuthUseCase(gateway domain.UserGateway) *UseCase {
	return &UseCase{
		logger:      zap.L().Named("auth use case"),
		userGateway: gateway,
	}
}

func (uc *UseCase) Login(ctx context.Context, email, password string) (*domain.User, error) {
	ep, err := crypt.Encrypt(password)
	if err != nil {
		uc.logger.Error("encrypt failed", zap.Error(err))

		return nil, domain.ErrUnexpected
	}

	u, err := uc.userGateway.FindByEmailAndPassword(ctx, email, ep)
	if err != nil && !errors.Is(err, domain.ErrNotFound) {
		uc.logger.Error("find user by email and password failed", zap.Error(err))

		return nil, domain.ErrUnexpected
	}

	return u, err
}

func (uc *UseCase) Register(ctx context.Context, user *domain.User) error {
	_, err := uc.userGateway.Create(ctx, user)
	if err != nil && !errors.Is(err, domain.ErrNotFound) {
		uc.logger.Error("create user failed", zap.Error(err))

		return domain.ErrUnexpected
	}

	return err
}
