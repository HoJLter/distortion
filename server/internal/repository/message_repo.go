package repository

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type MessageRepository struct {
	pool *pgxpool.Pool
}

func NewMessageRepository(pool *pgxpool.Pool) *MessageRepository {
	return &MessageRepository{pool: pool}
}

type CreateMessageParams struct {
	ChannelUID string
	UserUID    string
	Content    string
}

func (m MessageRepository) Create(ctx context.Context, params CreateMessageParams) (MessageRecord, error) {
	var record MessageRecord

	err := m.pool.QueryRow(ctx,
		`WITH inserted AS (INSERT INTO messages (channel_uid, user_uid, content)
    VALUES ($1, $2, $3)
    RETURNING id, channel_uid, user_uid, content
)
SELECT
    inserted.id,
	inserted.channel_uid,
    users.uid,
    users.username,
    users.avatar_path,
    inserted.content
FROM inserted
JOIN users ON users.uid = inserted.user_uid;`,
		params.ChannelUID, params.UserUID, params.Content,
	).Scan(
		&record.MessageID,
		&record.ChannelUID,
		&record.Author.UID,
		&record.Author.Username,
		&record.Author.AvatarPath,
		&record.Content,
	)

	if err != nil {
		return MessageRecord{}, err
	}

	return record, nil
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

type AuthorRecord struct {
	UID        string
	Username   string
	AvatarPath *string
}
type MessageRecord struct {
	ChannelUID string
	MessageID  int
	Author     AuthorRecord
	Content    string
}

func (m MessageRepository) GetForChannel(
	ctx context.Context,
	channelUID string,
	offset int,
	limit int,
) ([]MessageRecord, error) {
	rows, err := m.pool.Query(ctx,
		`SELECT 
    			channel_uid,
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

	var values []MessageRecord
	for rows.Next() {
		var message MessageRecord
		err = rows.Scan(
			&message.ChannelUID,
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
