package domain

import (
	"context"
	"time"
)

type (
	User struct {
		ID        int64
		FirstName string
		LastName  string
		Birthday  time.Time
		Email     string
		Password  string
		Interests string
		Sex       string
		City      *City
	}

	UserGateway interface {
		FindByID(ctx context.Context, id int64) (*User, error)
		FindByEmail(ctx context.Context, email string) (*User, error)
		FindAll(ctx context.Context) ([]*User, error)
		FindFriends(ctx context.Context, userID int64) ([]*User, error)
		Create(ctx context.Context, u *User) (int64, error)
		AddFriend(ctx context.Context, userID, friendID int64) (int64, error)
		DeleteFriend(ctx context.Context, userID, friendID int64) (int64, error)
	}
)
