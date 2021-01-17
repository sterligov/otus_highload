package auth

import (
	"context"
	"fmt"

	"github.com/sterligov/otus_highload/dating/internal/domain"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
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
	u, err := uc.userGateway.FindByEmail(ctx, email)
	if err != nil {
		return nil, fmt.Errorf("FindByEmail: %w", err)
	}

	if err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password)); err != nil {
		return nil, fmt.Errorf("CompareHashAndPassword: %w", err)
	}

	return u, nil
}
