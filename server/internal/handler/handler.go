package handler

import (
	"context"
	"log/slog"
	"server/internal/repository"

	"github.com/coder/websocket"
	"github.com/jackc/pgx/v5/pgxpool"
)

type wsRawMessageDTO struct {
	ctx  context.Context
	typ  websocket.MessageType
	data []byte
}
type Hub struct {
	connections map[*websocket.Conn]struct{}

	subscribe   chan *websocket.Conn
	unsubscribe chan *websocket.Conn

	broadcast chan wsRawMessageDTO
}

func NewHub() *Hub {
	return &Hub{
		connections: make(map[*websocket.Conn]struct{}, 0),
		subscribe:   make(chan *websocket.Conn),
		unsubscribe: make(chan *websocket.Conn),
		broadcast:   make(chan wsRawMessageDTO),
	}
}

func sendMessageToConn(conn *websocket.Conn, msg *wsRawMessageDTO) {
	if err := conn.Write(msg.ctx, msg.typ, msg.data); err != nil {
		slog.Warn("Message to client is not transferred")
	}
}

func (h *Hub) Run() {
	for {
		select {
		case conn := <-h.subscribe:
			h.connections[conn] = struct{}{}
		case conn := <-h.unsubscribe:
			{
				delete(h.connections, conn)
			}
		case msg := <-h.broadcast:
			for conn, _ := range h.connections {
				go sendMessageToConn(conn, &msg)
			}
		}
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
		hub:      NewHub(),
	}
}
