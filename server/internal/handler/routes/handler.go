package routes

import (
	"server/internal/repository"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Handler struct {
	channels *repository.ChannelRepository
	users    *repository.UserRepository
	messages *repository.MessageRepository
	hub      *Hub
}

func NewHandler(pool *pgxpool.Pool) *Handler {
	return &Handler{
		channels: repository.NewChannelRepository(pool),
		users:    repository.NewUserRepository(pool),
		messages: repository.NewMessageRepository(pool),
		hub:      NewHub(),
	}
}
