package handler

import (
	"context"
	"log/slog"
	"server/internal/repository"
	"time"

	"github.com/coder/websocket"
	"github.com/jackc/pgx/v5/pgxpool"
)

type wsRawMessageDTO struct {
	typ  websocket.MessageType
	data []byte
}

type WsClient struct {
	ctx      context.Context
	conn     *websocket.Conn
	messages chan wsRawMessageDTO
}

func NewWsClient(ctx context.Context, conn *websocket.Conn) *WsClient {
	return &WsClient{
		ctx:      ctx,
		conn:     conn,
		messages: make(chan wsRawMessageDTO, 64),
	}
}

func (w *WsClient) Listen() {
	for {
		select {
		case msg := <-w.messages:
			{
				ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
				if err := w.conn.Write(ctx, msg.typ, msg.data); err != nil {
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

	broadcast chan wsRawMessageDTO
}

func NewHub() *Hub {
	return &Hub{
		wsClients:   make(map[*WsClient]struct{}),
		subscribe:   make(chan *WsClient),
		unsubscribe: make(chan *WsClient),
		broadcast:   make(chan wsRawMessageDTO),
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
