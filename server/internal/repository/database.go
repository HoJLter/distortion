package repository

import (
	"context"
	"log/slog"
	"server/internal/config"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Channel struct {
	UID  string
	Type string
}

type User struct {
	UID        string
	Username   string
	AvatarPath string
}

type Message struct {
	TargetChannel Channel
	Author        User
	Content       string
}

func CreateChannel(
	ctx context.Context,
	db *pgxpool.Pool,
	channel Channel) error {
	_, err := db.Exec(ctx,
		`INSERT INTO channels (uid, type) VALUES ($1, $2)`,
		channel.UID, channel.Type,
	)
	if err != nil {
		return err
	}

	return nil
}

func DeleteChannel(
	ctx context.Context,
	db *pgxpool.Pool,
	channelUID string) error {
	_, err := db.Exec(ctx,
		`DELETE FROM channels WHERE uid = $1`,
		channelUID,
	)
	if err != nil {
		return err
	}

	return nil
}

func CreateMessage(
	ctx context.Context,
	db *pgxpool.Pool,
	message Message) error {
	_, err := db.Exec(ctx,
		`INSERT INTO messages (channel_uid, user_uid, content) VALUES ($1, $2, $3)`,
		message.TargetChannel.UID, message.Author.UID, message.Content,
	)
	if err != nil {
		return err
	}

	return nil
}

func DeleteMessage(
	ctx context.Context,
	db *pgxpool.Pool,
	messageID int) error {
	_, err := db.Exec(ctx,
		`DELETE FROM messages WHERE id = $1`,
		messageID,
	)
	if err != nil {
		return err
	}

	return nil
}

func CreateUser(
	ctx context.Context,
	db *pgxpool.Pool,
	user User) error {
	_, err := db.Exec(ctx,
		`INSERT INTO users (uid, username, avatar_path) VALUES ($1, $2, $3)`,
		user.UID, user.Username, user.AvatarPath,
	)
	if err != nil {
		return err
	}

	return nil
}

func DeleteUser(
	ctx context.Context,
	db *pgxpool.Pool,
	userUID string) error {
	_, err := db.Exec(ctx,
		`DELETE FROM users WHERE uid = $1`,
		userUID,
	)
	if err != nil {
		return err
	}

	return nil
}

func SetupDatabase(cfg config.DatabaseConfig) error {
	ctx := context.Background()

	db, err := pgxpool.New(ctx, cfg.Url)
	if err != nil {
		slog.Error("Error while db init", "err", err)
		return err
	}
	defer db.Close()

	if err := db.Ping(ctx); err != nil {
		slog.Error("Error while db ping", "err", err)
		return err
	}

	return nil
}
