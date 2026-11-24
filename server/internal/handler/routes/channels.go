package routes

import (
	"encoding/json"
	"io"
	"log/slog"
	"net/http"
	"server/internal/dto"
	"server/internal/repository"
	"strconv"

	"github.com/go-chi/chi"
)

/*
* ТЕКСТОВЫЕ КАНАЛЫ
* GET  /api/channels/{channel_uid}/messages - получить сообщения
* POST /api/channels/{channel_uid}/messages - отправить сообщение
* POST /api/channels/ - добавить канал
* DELETE /api/channels/{channel_uid} - удалить канал
 */

// HANDLERS

func (h Handler) getMessages(w http.ResponseWriter, r *http.Request) {
	channelUID := chi.URLParam(r, "channel_uid")
	offset, err := strconv.Atoi(r.URL.Query().Get("offset"))
	limit, err := strconv.Atoi(r.URL.Query().Get("limit"))
	if err != nil {
		slog.Error("getMessages() query params failed to fetch", "err", err, "channelUID", channelUID)
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	messagesRecords, err := h.messages.GetForChannel(r.Context(), channelUID, offset, limit)
	if err != nil {
		slog.Error("getMessages() error", "err", err, "channelUID", channelUID)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var messages []dto.MessageDTO
	for _, record := range messagesRecords {
		messages = append(messages, dto.RecordToMessageDTO(record))
	}

	w.Header().Set("Content-Type", "appcilation/json")

	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(messages); err != nil {
		slog.Error("getMessages() encoding error")
	}
}

func (h Handler) createChannel(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Request reading error", http.StatusBadRequest)
		slog.Error("Request reading error", "body", body, "err", err)
		return
	}

	var data dto.CreateChannelRequest
	err = json.Unmarshal(body, &data)
	if err != nil {
		http.Error(w, "Invalid format", http.StatusBadRequest)
		slog.Error("Invalid format", "data", data, "err", err)
		return
	}

	dbParams := repository.CreateChannelParams{
		UID:  data.ChannelUID,
		Type: data.Type,
	}
	if err := h.channels.Create(r.Context(), dbParams); err != nil {
		http.Error(w, "DB request failed", http.StatusInternalServerError)
		slog.Error("Error while db request", "params", dbParams, "err", err)
		return
	}
}

func (h Handler) getChannels(w http.ResponseWriter, r *http.Request) {
	channelsRecords, err := h.channels.GetAll(r.Context())
	if err != nil {
		http.Error(w, "DB request failed", http.StatusInternalServerError)
		slog.Error("Error while db request", "err", err)
		return
	}

	var channels []dto.GetChannelResponse
	for _, record := range channelsRecords {
		channels = append(channels, dto.RecordToChannelDTO(record))
	}

	if err := json.NewEncoder(w).Encode(channels); err != nil {
		http.Error(w, "Error while writing request", http.StatusInternalServerError)
		slog.Error("Error while writing request")
		return
	}
}

func (h Handler) deleteChannel(w http.ResponseWriter, r *http.Request) {
	channel_uid := chi.URLParam(r, "channel_uid")

	if err := h.channels.Delete(r.Context(), channel_uid); err != nil {
		http.Error(w, "DB request failed", http.StatusInternalServerError)
		slog.Error("Error while db request", "err", err)
		return
	}
}

func (h Handler) connectToVoiceChannel(w http.ResponseWriter, r *http.Request) {

}

func SetupChannelsRouter(parentRouter chi.Router, h *Handler) {
	channelsRouter := chi.NewRouter()
	parentRouter.Mount("/channels", channelsRouter)

	channelsRouter.Get("/{channel_uid}/messages", h.getMessages)

	channelsRouter.Get("/", h.getChannels)
	channelsRouter.Post("/", h.createChannel)
	channelsRouter.Delete("/{channel_uid}", h.deleteChannel)
}
