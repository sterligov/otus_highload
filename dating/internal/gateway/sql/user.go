package sql

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/go-sql-driver/mysql"

	"go.uber.org/zap"

	"github.com/sterligov/otus_highload/dating/internal/domain"

	"github.com/jmoiron/sqlx"
)

const (
	MALE              = 'M'
	FEMALE            = 'F'
	mysqlUniqueErrNum = 1062
)

type (
	User struct {
		ID        int64         `db:"id"`
		FirstName string        `db:"first_name"`
		LastName  string        `db:"last_name"`
		Password  string        `db:"password"`
		Birthday  time.Time     `db:"birthday"`
		Email     string        `db:"email"`
		Interests string        `db:"interests"`
		Sex       string        `db:"sex"`
		CityID    int64         `db:"city_id"`
		IsFriend  sql.NullInt32 `db:"is_friend"`
		City      `db:"city"`
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

func (ug *UserGateway) FindByID(ctx context.Context, curUserID, id int64) (*domain.User, error) {
	const query = `
SELECT u.*,
       с.id "city.id",
	   с.name "city.name",
       (
       	SELECT 1
    	FROM friends
       	WHERE user_id = ? AND friend_id = ? 
       	OR friend_id = ? AND user_id = ?
       ) is_friend
FROM user u
JOIN city с ON с.id = u.city_id
WHERE u.id = ?`

	u := new(User)

	err := ug.db.
		QueryRowxContext(
			ctx,
			query,
			curUserID,
			id,
			curUserID,
			id,
			id,
		).StructScan(u)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, domain.ErrNotFound
		}

		return nil, fmt.Errorf("query row: %w", err)
	}

	return toDomainUser(u), nil
}

func (ug *UserGateway) FindByEmail(ctx context.Context, email string) (*domain.User, error) {
	const query = `
SELECT u.*,
       c.id "city.id",
       c.name "city.name"
FROM user u
JOIN city c on u.city_id = c.id
WHERE email = ?`

	u := new(User)

	err := ug.db.
		QueryRowxContext(ctx, query, email).
		StructScan(u)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, domain.ErrNotFound
		}

		return nil, err
	}

	return toDomainUser(u), nil
}

func (ug *UserGateway) FindByFirstNameAndLastName(
	ctx context.Context,
	firstName, lastName string,
) ([]*domain.User, error) {
	const query = `
SELECT *
FROM user
WHERE first_name LIKE ? AND last_name LIKE ?`

	var users []*domain.User

	rows, err := ug.db.QueryxContext(ctx, query, firstName+"%", lastName+"%")
	if err != nil {
		return nil, fmt.Errorf("user find by first and last name query: %w", err)
	}
	defer func() {
		if err := rows.Close(); err != nil {
			ug.logger.Warn("rows close failed", zap.Error(err))
		}
	}()

	for rows.Next() {
		u := new(User)
		if err := rows.StructScan(&u); err != nil {
			return nil, fmt.Errorf("user struct scan: %w", err)
		}
		users = append(users, toDomainUser(u))
	}

	if rows.Err() != nil {
		return nil, fmt.Errorf("user rows: %w", rows.Err())
	}

	return users, nil
}

func (ug *UserGateway) FindAll(ctx context.Context) ([]*domain.User, error) {
	const query = `
SELECT u.*,
       ci.id as "city.id",
       ci.name as "city.name"
FROM user u
JOIN city ci ON ci.id = u.city_id`

	var users []*domain.User

	rows, err := ug.db.QueryxContext(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("user find all query: %w", err)
	}
	defer func() {
		if err := rows.Close(); err != nil {
			ug.logger.Warn("rows close failed", zap.Error(err))
		}
	}()

	for rows.Next() {
		u := new(User)
		if err := rows.StructScan(&u); err != nil {
			return nil, fmt.Errorf("user struct scan: %w", err)
		}
		users = append(users, toDomainUser(u))
	}

	if rows.Err() != nil {
		return nil, fmt.Errorf("user rows: %w", rows.Err())
	}

	return users, nil
}

func (ug *UserGateway) FindFriends(ctx context.Context, userID int64) ([]*domain.User, error) {
	const query = `
SELECT u.*,
       c.id "city.id",
       c.name "city.name"
FROM (
	SELECT *
	FROM user
	WHERE id IN (
		SELECT friend_id id
		FROM friends
		WHERE user_id = :id
		UNION
		SELECT user_id id
		FROM friends
		WHERE friend_id = :id
	)
) u JOIN city c ON u.city_id = c.id`

	rows, err := ug.db.NamedQueryContext(ctx, query, map[string]interface{}{
		"id": userID,
	})
	if err != nil {
		return nil, fmt.Errorf("named query context: %w", err)
	}
	defer func() {
		if err := rows.Close(); err != nil {
			ug.logger.Warn("rows close failed", zap.Error(err))
		}
	}()

	var users []*domain.User

	for rows.Next() {
		u := new(User)
		if err := rows.StructScan(u); err != nil {
			return nil, fmt.Errorf("struct scan: %w", err)
		}

		users = append(users, toDomainUser(u))
	}

	if rows.Err() != nil {
		return nil, fmt.Errorf("fetch rows: %w", rows.Err())
	}

	return users, nil
}

func (ug *UserGateway) Create(ctx context.Context, u *domain.User) (int64, error) {
	const query = `
INSERT INTO user(
	first_name,
	last_name,
	email,
	password,
	interests,
	sex,
	birthday,
	city_id
) VALUES (?, ?, ?, ?, ?, ?, ?, ?)`
	res, err := ug.db.ExecContext(
		ctx,
		query,
		u.FirstName,
		u.LastName,
		u.Email,
		u.Password,
		u.Interests,
		u.Sex,
		u.Birthday,
		u.City.ID,
	)
	if err != nil {
		if me, ok := err.(*mysql.MySQLError); ok && me.Number == mysqlUniqueErrNum {
			return 0, domain.ErrEmailAlreadyExist
		}

		return 0, fmt.Errorf("user create entity: %w", err)
	}

	id, err := res.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("user get last inserted id: %w", err)
	}

	return id, nil
}

func (ug *UserGateway) AddFriend(ctx context.Context, userID, friendID int64) (int64, error) {
	const query = `
INSERT INTO friends(user_id, friend_id)
VALUES (?, ?)`

	res, err := ug.db.ExecContext(ctx, query, userID, friendID)
	if err != nil {
		return 0, fmt.Errorf("subscribe exec context: %w", err)
	}

	insertedID, err := res.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("subscribe last inserted id: %w", err)
	}

	return insertedID, nil
}

func (ug *UserGateway) DeleteFriend(ctx context.Context, userID, friendID int64) (int64, error) {
	const query = `
DELETE FROM friends
WHERE user_id = :user_id AND friend_id = :friend_id 
	OR friend_id = :user_id AND user_id = :friend_id`

	res, err := ug.db.NamedExecContext(ctx, query, map[string]interface{}{
		"user_id":   userID,
		"friend_id": friendID,
	})
	if err != nil {
		return 0, fmt.Errorf("delete friend exec context: %w", err)
	}

	affected, err := res.RowsAffected()
	if err != nil {
		return 0, fmt.Errorf("delete friend affected: %w", err)
	}

	return affected, nil
}

func toDomainUser(u *User) *domain.User {
	return &domain.User{
		ID:        u.ID,
		FirstName: u.FirstName,
		LastName:  u.LastName,
		Birthday:  u.Birthday,
		Password:  u.Password,
		Email:     u.Email,
		Interests: u.Interests,
		Sex:       u.Sex,
		IsFriend:  int(u.IsFriend.Int32),
		City: &domain.City{
			ID:   u.CityID,
			Name: u.City.Name,
		},
	}
}
