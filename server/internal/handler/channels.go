package handler

import (
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

	channelsRouter.Get("{channel_uid}/messages", h.getMessages)
	channelsRouter.Post("{channel_uid}/messages", h.postMessage)

	channelsRouter.Get("/", h.getChannels)
	channelsRouter.Post("/", h.createTextChannel)
	channelsRouter.Post("/", h.createVoiceChannel)
	channelsRouter.Delete("/{channel_uid}", h.deleteChannel)
}
