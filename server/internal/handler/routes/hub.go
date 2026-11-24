package routes

import (
	"context"
	"encoding/json"
	"log/slog"
	"server/internal/dto"
	"time"

	"github.com/coder/websocket"
)

type WsClient struct {
	ctx        context.Context
	conn       *websocket.Conn
	messages   chan dto.WSMessageDTO
	channelUID string
}

func NewWsClient(ctx context.Context, conn *websocket.Conn, channelUID string) *WsClient {
	return &WsClient{
		ctx:        ctx,
		conn:       conn,
		messages:   make(chan dto.WSMessageDTO, 64),
		channelUID: channelUID,
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

				var msgPayload dto.MessageDTO
				if err := json.Unmarshal(msg.Payload, &msgPayload); err != nil {
					slog.Error("Cannot unmarshal broadcast message")
					continue
				}
				slog.Debug("Successfully unmarshalled WS internal", "msg", msgPayload, "clientChannel", client)

				if client.channelUID != msgPayload.ChannelUID {
					continue
				}
				select {
				case client.messages <- msg:
				default:
					slog.Error("Client message buffer is full")
				}
			}
		}
	}
}
