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
		Sex       byte
		CityID    int64
	}

	UserGateway interface {
		FindByID(ctx context.Context, id int64) (*User, error)
		FindByEmailAndPassword(ctx context.Context, email, password string) (*User, error)
		FindAfterID(ctx context.Context, id, limit int64) ([]*User, error)
		Filter(ctx context.Context, user *User) ([]*User, error)
		Create(ctx context.Context, u *User) (int64, error)
		//Delete(ctx context.Context, id int64) (int64, error)
		//Update(ctx context.Context, u *User) (int64, error)
	}
)
