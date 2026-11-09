package repository

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type User struct {
	UserUID    string
	Username   string
	AvatarPath string
}

type UserRepository struct {
	pool *pgxpool.Pool
}

func NewUserRepository(pool *pgxpool.Pool) *UserRepository {
	return &UserRepository{pool: pool}
}

func (u UserRepository) Create(ctx context.Context, user User) error {
	_, err := u.pool.Exec(ctx,
		`INSERT INTO users (uid, username, avatar_path) VALUES ($1, $2)`,
		user.Username, user.AvatarPath,
	)
	if err != nil {
		return err
	}

	return nil
}

func (u UserRepository) Delete(ctx context.Context, userUID string) error {
	_, err := u.pool.Exec(ctx,
		`DELETE FROM users WHERE uid = $1`,
		userUID,
	)
	if err != nil {
		return err
	}

	return nil
}
