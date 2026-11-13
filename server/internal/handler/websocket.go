package handler

import (
	"context"
	"log/slog"
	"net/http"
	"time"

	"github.com/coder/websocket"
	"github.com/go-chi/chi"
)

func (h *Handler) handleWs(w http.ResponseWriter, r *http.Request) {
	conn, err := websocket.Accept(w, r, nil)
	if err != nil {
		slog.Error("Websocket upgrade failed", "err", err)
		return
	}
	defer conn.CloseNow()

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	for {
		typ, data, err := conn.Read(ctx)
		if err != nil {
			if websocket.CloseStatus(err) == websocket.StatusNormalClosure {
				return
			}
			slog.Error("ws read error", "err", err)
		}
		err = conn.Write(ctx, typ, data)
		if err != nil {
			slog.Error("ws write error", "err", err)
			return
		}
	}
}

func setupWsRouter(parentRouter chi.Router, h *Handler) {
	wsRouter := chi.NewRouter()
	parentRouter.Mount("/ws", wsRouter)

	wsRouter.Get("/", h.handleWs)
}
