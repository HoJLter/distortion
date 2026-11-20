package handler

import (
	"context"
	"log/slog"
	"server/internal/repository"
	"slices"

	"github.com/coder/websocket"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Hub struct {
	connections []*websocket.Conn
}

func (h *Hub) subscribe(conn *websocket.Conn) {
	h.connections = append(h.connections, conn)
}

func (h *Hub) unsubscribe(connToUnsub *websocket.Conn) {
	for i, conn := range h.connections {
		if conn == connToUnsub {
			h.connections = slices.Delete(h.connections, i, i+1)
			return
		}
	}
}

func (h *Hub) broadcast(ctx context.Context, messageType websocket.MessageType, message []byte) {
	for _, conn := range h.connections {

		if err := conn.Write(ctx, messageType, message); err != nil {
			slog.Error("Message cannot be transferred to client", "conn", conn)
		}
		slog.Debug("Message has been sent to the client")
	}
}

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
		hub: &Hub{
			connections: make([]*websocket.Conn, 0),
		},
	}
}
