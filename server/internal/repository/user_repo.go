package repository

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)

type UserRecord struct {
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

func (u UserRepository) GetAll(ctx context.Context) ([]UserRecord, error) {
	var records []UserRecord
	rows, err := u.pool.Query(ctx,
		`SELECT uid, username, avatar_path FROM users`,
	)
	defer rows.Close()
	if err != nil {
		return nil, fmt.Errorf("getAllUsers: sql query %w", err)
	}

	for rows.Next() {
		var record UserRecord
		err = rows.Scan(&record.UserUID, &record.Username, &record.AvatarPath)
		if err != nil {
			return nil, fmt.Errorf("getAllUsers: row parsing %w", err)
		}
		records = append(records, record)
	}
	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("getAllUsers: after iterate %w", err)
	}

	return records, nil
}

func (u UserRepository) Create(ctx context.Context, user UserRecord) error {
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
