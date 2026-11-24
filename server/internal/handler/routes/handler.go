package routes

import (
	"context"
	"encoding/json"
	"log/slog"
	"server/internal/dto"
	"server/internal/repository"
	"time"

	"github.com/coder/websocket"
	"github.com/jackc/pgx/v5/pgxpool"
)

type WsClient struct {
	ctx      context.Context
	conn     *websocket.Conn
	messages chan dto.WSMessageDTO
}

func NewWsClient(ctx context.Context, conn *websocket.Conn) *WsClient {
	return &WsClient{
		ctx:      ctx,
		conn:     conn,
		messages: make(chan dto.WSMessageDTO, 64),
	}
}

func (w *WsClient) Listen() {
	for {
		select {
		case msg, ok := <-w.messages:
			{
				if !ok {
					return
				}
				ctx, cancel := context.WithTimeout(w.ctx, 5*time.Second)

				data, err := json.Marshal(msg)
				if err != nil {
					slog.Error("Error while marshalling for client")
					continue
				}

				if err := w.conn.Write(ctx, websocket.MessageText, data); err != nil {
					slog.Warn("Message to client is not transferred")
				}
				cancel()
			}
		case <-w.ctx.Done():
			return
		}
	}
}

type Hub struct {
	wsClients map[*WsClient]struct{}

	subscribe   chan *WsClient
	unsubscribe chan *WsClient

	broadcast chan dto.WSMessageDTO
}

func NewHub() *Hub {
	return &Hub{
		wsClients:   make(map[*WsClient]struct{}),
		subscribe:   make(chan *WsClient),
		unsubscribe: make(chan *WsClient),
		broadcast:   make(chan dto.WSMessageDTO),
	}
}

func (h *Hub) Run() {
	for {
		select {
		case client := <-h.subscribe:
			h.wsClients[client] = struct{}{}
			go client.Listen()
		case client := <-h.unsubscribe:
			{
				delete(h.wsClients, client)
			}
		case msg := <-h.broadcast:
			for client := range h.wsClients {
				slog.Debug("Message transferred to client")
				select {
				case client.messages <- msg:
				default:
					slog.Error("Client message buffer is full")
				}
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
