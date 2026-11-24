package routes

import (
	"context"
	"encoding/json"
	"log/slog"
	"net/http"
	"server/internal/dto"
	"server/internal/repository"

	"github.com/coder/websocket"
	"github.com/go-chi/chi"
)

const (
	postMessageType      = "postMessage"
	broadcastMessageType = "broadcastMessage"
)

func (h *Handler) handleWs(w http.ResponseWriter, r *http.Request) {

	conn, err := websocket.Accept(w, r, &websocket.AcceptOptions{
		InsecureSkipVerify: true,
	})
	if err != nil {
		slog.Error("Websocket upgrade failed", "err", err)
		return
	}
	defer conn.CloseNow()
	slog.Debug("Websocket upgrade completed")

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	client := NewWsClient(ctx, conn)
	h.hub.subscribe <- client
	defer func() { h.hub.unsubscribe <- client }()

	for {
		_, data, err := conn.Read(ctx)
		if err != nil {
			status := websocket.CloseStatus(err)
			if status != websocket.StatusNormalClosure &&
				status != websocket.StatusGoingAway {
				slog.Error("ws read error", "err", err)
			}
			return
		}

		var acqWsMessage dto.WSMessageDTO
		err = json.Unmarshal(data, &acqWsMessage)
		if err != nil {
			slog.Error("Unmarshalling error", "err", err)
			return
		}
		slog.Debug("WS Message unmarshalled", "msg", acqWsMessage)

		if acqWsMessage.Type != postMessageType {
			slog.Error("Invalid type: ", "type", acqWsMessage.Type)
			return
		}

		var acqPostMessage dto.WSPostMessageDTO

		err = json.Unmarshal(acqWsMessage.Payload, &acqPostMessage)
		if err != nil {
			slog.Error("Unmarshalling error", "err", err)
			return
		}
		slog.Debug("WS Post Message unmarshalled", "msg", acqPostMessage)

		messageRecord, err := h.messages.Create(ctx, repository.CreateMessageParams{
			ChannelUID: acqPostMessage.ChannelUID,
			UserUID:    acqPostMessage.UserUID,
			Content:    acqPostMessage.Content,
		})
		if err != nil {
			slog.Error("Database transfer failed", "err", err)
			return
		}
		slog.Debug("Message transfer to database completed")

		message := dto.RecordToMessageDTO(messageRecord)
		payload, err := json.Marshal(message)
		if err != nil {
			slog.Error("Marshalling after mapping is failed")
			return
		}

		h.hub.broadcast <- dto.WSMessageDTO{
			Type:    broadcastMessageType,
			Payload: payload,
		}
	}
}

func SetupWsRouter(parentRouter chi.Router, h *Handler) {
	go h.hub.Run()

	wsRouter := chi.NewRouter()
	parentRouter.Mount("/ws", wsRouter)

	wsRouter.Get("/", h.handleWs)
}
