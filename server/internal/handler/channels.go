package handler

import (
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/go-chi/chi"
)

/*
* ТЕКСТОВЫЕ КАНАЛЫ
* GET  /api/channels/{channel_uid}/messages - получить сообщения
* POST /api/channels/{channel_uid}/messages - отправить сообщение
* POST /api/channels/ - добавить канал
* DELETE /api/channels/{channel_uid} - удалить канал
 */

func (h Handler) getMessages(w http.ResponseWriter, r *http.Request) {
	channelUID := chi.URLParam(r, "channel_uid")
	messages, err := h.messages.GetMessagesForChannel(r.Context(), channelUID)
	if err != nil {
		slog.Error("getMessages() error", "err", err, "channelUID", channelUID)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	slog.Debug("Acquired messages: ", "messages", messages)
	for _, message := range messages {
		marsh, err := json.Marshal(message)
		if err != nil {
			slog.Error("getMessages() marshalling error", "err", err, "message", message)
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		_, err = w.Write(marsh)
		if err != nil {
			slog.Error("getMessages() writing response error", "err", err)
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
	}
}

func (h Handler) postMessage(w http.ResponseWriter, r *http.Request) {

}

func (h Handler) createTextChannel(w http.ResponseWriter, r *http.Request) {

}

func (h Handler) createVoiceChannel(w http.ResponseWriter, r *http.Request) {

}

func (h Handler) connectToVoiceChannel(w http.ResponseWriter, r *http.Request) {

}

func (h Handler) deleteChannel(w http.ResponseWriter, r *http.Request) {

}

func (h Handler) getChannels(w http.ResponseWriter, r *http.Request) {}

func setupChannelsRouter(parentRouter chi.Router, h *Handler) {
	channelsRouter := chi.NewRouter()
	parentRouter.Mount("/channels", channelsRouter)

	channelsRouter.Get("/{channel_uid}/messages", h.getMessages)
	channelsRouter.Post("/{channel_uid}/messages", h.postMessage)

	channelsRouter.Get("/", h.getChannels)
	channelsRouter.Post("/", h.createTextChannel)
	channelsRouter.Post("/", h.createVoiceChannel)
	channelsRouter.Delete("/{channel_uid}", h.deleteChannel)
}
