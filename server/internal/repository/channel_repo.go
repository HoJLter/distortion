package repository

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Channel struct {
	UID  string
	Type string
}

type ChannelRepository struct {
	pool *pgxpool.Pool
}

func NewChannelRepository(pool *pgxpool.Pool) *ChannelRepository {
	return &ChannelRepository{pool: pool}
}

func (c ChannelRepository) CreateChannel(ctx context.Context, channel Channel) error {
	_, err := c.pool.Exec(ctx,
		`INSERT INTO channels (uid, type) VALUES ($1, $2)`,
		channel.UID, channel.Type,
	)
	if err != nil {
		return err
	}

	return nil
}

func (c ChannelRepository) DeleteChannel(ctx context.Context, channelUID string) error {
	_, err := c.pool.Exec(ctx,
		`DELETE FROM channels WHERE uid = $1`,
		channelUID,
	)
	if err != nil {
		return err
	}

	return nil
}
