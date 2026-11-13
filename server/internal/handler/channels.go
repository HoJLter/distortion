package handler

import (
	"encoding/json"
	"errors"
	"io"
	"log/slog"
	"net/http"
	"server/internal/repository"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/go-playground/validator/v10"
)

/*
* ТЕКСТОВЫЕ КАНАЛЫ
* GET  /api/channels/{channel_uid}/messages - получить сообщения
* POST /api/channels/{channel_uid}/messages - отправить сообщение
* POST /api/channels/ - добавить канал
* DELETE /api/channels/{channel_uid} - удалить канал
 */

// VALIDATION

type ValidationError struct {
	Field string `json:"field"`
	Tag   string `json:"tag"`
	Param any    `json:"param"`
}

func validateStruct(structToValidate any) []ValidationError {
	validate := validator.New()

	var structuredErrors []ValidationError

	if err := validate.Struct(structToValidate); err != nil {
		var validationErrors validator.ValidationErrors
		if errors.As(err, &validationErrors) {
			for _, e := range validationErrors {
				structuredErrors = append(structuredErrors, ValidationError{
					Field: e.Field(),
					Tag:   e.Tag(),
					Param: e.Param(),
				})
			}
		}
	}
	return structuredErrors
}

// VALIDATION END

// HANDLERS

func (h Handler) getMessages(w http.ResponseWriter, r *http.Request) {
	channelUID := chi.URLParam(r, "channel_uid")
	offset, err := strconv.Atoi(r.URL.Query().Get("offset"))
	limit, err := strconv.Atoi(r.URL.Query().Get("limit"))
	if err != nil {
		slog.Error("getMessages() query params failed to fetch", "err", err, "channelUID", channelUID)
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	messages, err := h.messages.GetMessagesForChannel(r.Context(), channelUID, offset, limit)
	if err != nil {
		slog.Error("getMessages() error", "err", err, "channelUID", channelUID)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	slog.Debug("Acquired messages: ", "messages", messages)

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
func (h Handler) postMessage(w http.ResponseWriter, r *http.Request) {
	type request struct {
		UserUID string `json:"user_uid" validate:"required"`
		Content string `json:"content" validate:"required"`
	}
	type PostMessageResponse struct {
		MessageID int `json:"message_id"`
	}

	data, err := io.ReadAll(r.Body)
	if err != nil {
		slog.Error("postMessage() error", "err", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var unmarshalledData request
	err = json.Unmarshal(data, &unmarshalledData)
	if err != nil {
		slog.Error("Invalid json", "err", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// validateErrors := validateStruct(data)
	// if len(validateErrors) != 0 {
	//	for valErr := range validateErrors {
	//
	//	}
	//	return
	//}

	postMessageParams := repository.PostMessageParams{
		ChannelUID: chi.URLParam(r, "channel_uid"),
		UserUID:    unmarshalledData.UserUID,
		Content:    unmarshalledData.Content,
	}
	slog.Debug("Unmasrhalled struct", "struct:", postMessageParams)
	messageID, err := h.messages.Post(r.Context(), postMessageParams)
	if err != nil {
		slog.Error("Post() error", "err", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	resp, err := json.Marshal(PostMessageResponse{MessageID: messageID})
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
