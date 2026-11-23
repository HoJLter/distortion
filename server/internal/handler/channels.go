package handler

import (
	"encoding/json"
	"io"
	"log/slog"
	"net/http"
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

	var messages []MessageDTO
	for i, record := range messagesRecords {
		slog.Debug("", "num", i)
		messages = append(messages, RecordToMessageDTO(record))
	}

	w.Header().Set("Content-Type", "appcilation/json")

	if err := json.NewEncoder(w).Encode(messages); err != nil {
		slog.Error("getMessages() encoding error")
	}
}

/*
SCHEMA:

	{
		user_uid: string,
		content: string
	}
*/

func (h Handler) createMessage(w http.ResponseWriter, r *http.Request) {
	data, err := io.ReadAll(r.Body)
	if err != nil {
		slog.Error("postMessage() error", "err", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var unmarshalledData CreateMessageRequest
	err = json.Unmarshal(data, &unmarshalledData)
	if err != nil {
		slog.Error("Invalid json", "err", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	postMessageParams := repository.PostMessageParams{
		ChannelUID: chi.URLParam(r, "channel_uid"),
		UserUID:    unmarshalledData.UserUID,
		Content:    unmarshalledData.Content,
	}
	slog.Debug("Unmasrhalled struct", "struct:", postMessageParams)
	messageRecord, err := h.messages.Create(r.Context(), postMessageParams)
	if err != nil {
		slog.Error("Post() error", "err", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	resp, err := json.Marshal(CreateMessageResponse{MessageID: messageRecord.MessageID})
	if err != nil {
		slog.Error("Marshal() error", "err", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	_, err = w.Write(resp)
	if err != nil {
		slog.Error("Write() error", "err", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}

func (h Handler) createChannel(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Request reading error", http.StatusBadRequest)
		slog.Error("Request reading error", "body", body)
		return
	}

	var data CreateChannelRequest
	err = json.Unmarshal(body, &data)
	if err != nil {
		http.Error(w, "Invalid format", http.StatusBadRequest)
		slog.Error("Invalid format", "data", data)
	}

	dbParams := repository.CreateMessageParams{}
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
	channelsRouter.Post("/{channel_uid}/messages", h.createMessage)

	channelsRouter.Get("/", h.getChannels)
	channelsRouter.Post("/", h.createChannel)
	channelsRouter.Delete("/{channel_uid}", h.deleteChannel)
}
