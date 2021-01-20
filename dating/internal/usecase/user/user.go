package user

import (
	"context"
	"fmt"

	"golang.org/x/crypto/bcrypt"

	"go.uber.org/zap"

	"github.com/sterligov/otus_highload/dating/internal/domain"
)

type UseCase struct {
	userGateway domain.UserGateway
	//friendGateway domain.Fri
	logger *zap.Logger
}

func NewUserUseCase(gateway domain.UserGateway) *UseCase {
	return &UseCase{
		logger:      zap.L().Named("user use case"),
		userGateway: gateway,
	}
}

func (uc *UseCase) FindByID(ctx context.Context, curUserID, id int64) (*domain.User, error) {
	return uc.userGateway.FindByID(ctx, curUserID, id)
}

func (uc *UseCase) Subscribe(ctx context.Context, userID, friendID int64) (int64, error) {
	return uc.userGateway.AddFriend(ctx, userID, friendID)
}

func (uc *UseCase) Unsubscribe(ctx context.Context, userID, friendID int64) (int64, error) {
	return uc.userGateway.DeleteFriend(ctx, userID, friendID)
}

func (uc *UseCase) CreateUser(ctx context.Context, user *domain.User) (int64, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return 0, fmt.Errorf("hash password: %w", err)
	}
	user.Password = string(hash)

	return uc.userGateway.Create(ctx, user)
}

func (uc *UseCase) FindAll(ctx context.Context) ([]*domain.User, error) {
	return uc.userGateway.FindAll(ctx)
}

func (uc *UseCase) FindFriends(ctx context.Context, userID int64) ([]*domain.User, error) {
	return uc.userGateway.FindFriends(ctx, userID)
}
