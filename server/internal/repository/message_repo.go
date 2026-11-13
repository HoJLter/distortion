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

type MessageAuthor struct {
	UID        string  `json:"uid"`
	Username   string  `json:"username"`
	AvatarPath *string `json:"avatar_path"`
}

type GetMessageResponse struct {
	MessageID int           `json:"message_id"`
	Author    MessageAuthor `json:"author"`
	Content   string        `json:"content"`
}

type PostMessageParams struct {
	ChannelUID string
	UserUID    string
	Content    string
}

type MessageRepository struct {
	pool *pgxpool.Pool
}

func NewMessageRepository(pool *pgxpool.Pool) *MessageRepository {
	return &MessageRepository{pool: pool}
}

func (m MessageRepository) Post(ctx context.Context, message PostMessageParams) (int, error) {
	row := m.pool.QueryRow(ctx,
		`INSERT INTO messages (channel_uid, user_uid, content) VALUES ($1, $2, $3) RETURNING id`,
		message.ChannelUID, message.UserUID, message.Content,
	)
	var id int
	err := row.Scan(&id)

	if err != nil {
		return 0, err
	}

	return id, nil
}

func (m MessageRepository) Delete(
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

func (m MessageRepository) GetMessagesForChannel(
	ctx context.Context,
	channelUID string,
	offset int,
	limit int,
) ([]GetMessageResponse, error) {
	rows, err := m.pool.Query(ctx,
		`SELECT 
    			messages.id,
    			users.uid,
				users.username, 
				users.avatar_path, 
				messages.content 
			from messages 
			inner join users 
			on messages.user_uid = users.uid
			where messages.channel_uid = $1
			order by id
			limit $2 offset $3;`,
		channelUID, limit, offset,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var values []GetMessageResponse
	for rows.Next() {
		var message GetMessageResponse
		err = rows.Scan(
			&message.MessageID,
			&message.Author.UID,
			&message.Author.Username,
			&message.Author.AvatarPath,
			&message.Content)
		if err != nil {
			return nil, err
		}
		values = append(values, message)
	}
	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return values, nil
}
