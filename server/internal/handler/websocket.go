package handler

import (
	"context"
	"encoding/json"
	"log/slog"
	"net/http"
	"server/internal/repository"

	"github.com/coder/websocket"
	"github.com/go-chi/chi"
)

type wsPostMessageDTO struct {
	ChannelUID string `json:"channel_uid"`
	UserUID    string `json:"user_uid"`
	Content    string `json:"content"`
}

func (h *Handler) handleWs(w http.ResponseWriter, r *http.Request) {

	conn, err := websocket.Accept(w, r, nil)
	if err != nil {
		slog.Error("Websocket upgrade failed", "err", err)
		return
	}
	defer conn.CloseNow()

	ctx := context.Background()

	h.hub.subscribe <- conn
	defer func() { h.hub.unsubscribe <- conn }()

	for {
		typ, data, err := conn.Read(ctx)
		if err != nil {
			status := websocket.CloseStatus(err)
			if status != websocket.StatusNormalClosure &&
				status != websocket.StatusGoingAway {
				slog.Error("ws read error", "err", err)
			}
			return
		}

		var message wsPostMessageDTO
		err = json.Unmarshal(data, &message)
		if err != nil {
			slog.Error("Unmarshalling error", "err", err)
			return
		}
		slog.Debug("Message unmarshalled", "msg", message)

		_, err = h.messages.Post(ctx, repository.PostMessageParams{
			ChannelUID: message.ChannelUID,
			UserUID:    message.UserUID,
			Content:    message.Content,
		})
		if err != nil {
			slog.Error("Database transfer failed", "err", err)
			return
		}
		slog.Debug("Message transfer to database completed")

		h.hub.broadcast <- wsRawMessageDTO{
			ctx:  ctx,
			typ:  typ,
			data: data,
		}
	}
}

func setupWsRouter(parentRouter chi.Router, h *Handler) {
	go h.hub.Run()

	wsRouter := chi.NewRouter()
	parentRouter.Mount("/ws", wsRouter)

	wsRouter.Get("/", h.handleWs)
}
