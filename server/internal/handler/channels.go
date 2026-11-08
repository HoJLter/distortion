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

func getMessages(w http.ResponseWriter, r *http.Request) {

}

func postMessage(w http.ResponseWriter, r *http.Request) {

}

func createTextChannel(w http.ResponseWriter, r *http.Request) {

}

func createVoiceChannel(w http.ResponseWriter, r *http.Request) {

}

func connectToVoiceChannel(w http.ResponseWriter, r *http.Request) {

}

func deleteChannel(w http.ResponseWriter, r *http.Request) {

}

func getChannels(w http.ResponseWriter, r *http.Request) {

}

func setupTextChanRouter(parentRouter chi.Router) {
	channelsRouter := chi.NewRouter()
	parentRouter.Mount("/channels", channelsRouter)

	channelsRouter.Get("{channel_uid}/messages", getMessages)
	channelsRouter.Post("{channel_uid}/messages", postMessage)

	channelsRouter.Get("/", getChannels)
	channelsRouter.Post("/", createTextChannel)
	channelsRouter.Post("/", createVoiceChannel)
	channelsRouter.Delete("/{channel_uid}", deleteChannel)
}
