package repository

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Message struct {
	TargetChannel Channel
	Author        User
	Content       string
}

type MessageRepository struct {
	pool *pgxpool.Pool
}

func NewMessageRepository(pool *pgxpool.Pool) *MessageRepository {
	return &MessageRepository{pool: pool}
}

func (m MessageRepository) CreateMessage(ctx context.Context, message Message) error {
	_, err := m.pool.Exec(ctx,
		`INSERT INTO messages (channel_uid, user_uid, content) VALUES ($1, $2, $3)`,
		message.TargetChannel.UID, message.Author.UID, message.Content,
	)
	if err != nil {
		return err
	}

	return nil
}

func (m MessageRepository) DeleteMessage(
	ctx context.Context,
	messageID int) error {
	_, err := m.pool.Exec(ctx,
		`DELETE FROM messages WHERE id = $1`,
		messageID,
	)
	if err != nil {
		return err
	}

	return nil
}
