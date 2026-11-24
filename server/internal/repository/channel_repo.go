package repository

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)

type ChannelRepository struct {
	pool *pgxpool.Pool
}
type ChannelType string

const (
	TextChannelType  ChannelType = "text"
	VoiceChannelType ChannelType = "voice"
)

type ChannelRecord struct {
	UID  string
	Type ChannelType
}

func NewChannelRepository(pool *pgxpool.Pool) *ChannelRepository {
	return &ChannelRepository{pool: pool}
}

func (c ChannelRepository) GetAll(ctx context.Context) ([]ChannelRecord, error) {
	var records []ChannelRecord
	rows, err := c.pool.Query(ctx,
		`SELECT type, uid FROM channels`,
	)
	if err != nil {
		return nil, fmt.Errorf("get all channels: sql query %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var record ChannelRecord
		if err := rows.Scan(&record.Type, &record.UID); err != nil {
			return nil, fmt.Errorf("get all channels: row parsing %w ", err)
		}
		records = append(records, record)
	}
	err = rows.Err()
	if err != nil {
		return nil, fmt.Errorf("get all channels: after iterate rows %w")
	}

	return records, nil
}

type CreateChannelParams struct {
	UID  string
	Type ChannelType
}

func (c ChannelRepository) Create(ctx context.Context, channel CreateChannelParams) error {
	_, err := c.pool.Exec(ctx,
		`INSERT INTO channels (uid, type) VALUES ($1, $2)`,
		channel.UID, channel.Type,
	)
	if err != nil {
		return fmt.Errorf("Create channel: %w", err)
	}

	return nil
}

func (c ChannelRepository) Delete(ctx context.Context, channelUID string) error {
	_, err := c.pool.Exec(ctx,
		`DELETE FROM channels WHERE uid = $1`,
		channelUID,
	)
	if err != nil {
		return err
	}

	return nil
}
