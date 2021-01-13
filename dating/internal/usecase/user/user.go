package user

import (
	"context"

	"go.uber.org/zap"

	"github.com/sterligov/otus_highload/dating/internal/domain"
)

type UseCase struct {
	userGateway domain.UserGateway
	logger      *zap.Logger
}

func NewUserUseCase(gateway domain.UserGateway) *UseCase {
	return &UseCase{
		logger:      zap.L().Named("user use case"),
		userGateway: gateway,
	}
}

func (uc *UseCase) FindByID(ctx context.Context, id int64) (*domain.User, error) {
	u, err := uc.userGateway.FindByID(ctx, id)
	if err != nil {

	}

	return u, nil
}

func (uc *UseCase) FriendRequest() {

}

func (uc *UseCase) AnswerFriendRequest() {

}

func (uc *UseCase) CreateUser(ctx context.Context, user *domain.User) (int64, error) {
	id, err := uc.userGateway.Create(ctx, user)
	if err != nil {
		uc.logger.Error("create user failed", zap.Error(err))

		return 0, err
	}

	return id, nil
}

func (uc *UseCase) Filter(ctx context.Context, user *domain.User) (int64, error) {
	return 0, nil
	//id, err := uc.userGateway.Create(ctx, user)
	//if err != nil {
	//	uc.logger.Error("create user failed", zap.Error(err))
	//
	//	return 0, err
	//}
	//
	//return id, nil
}

func (uc *UseCase) FindAfterID(ctx context.Context, id int64) ([]*domain.User, error) {
	users, err := uc.userGateway.FindAfterID(ctx, id, 10)
	if err != nil {
		uc.logger.Error("FindAfterID failed", zap.Error(err))

		return nil, err
	}

	return users, nil
}
