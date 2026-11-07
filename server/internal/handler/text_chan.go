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

func createChannel(w http.ResponseWriter, r *http.Request) {

}

func deleteChannel(w http.ResponseWriter, r *http.Request) {

}

func getChannels(w http.ResponseWriter, r *http.Request) {

}

func setupTextChanRouter(parentRouter chi.Router) {
	textChanRouter := chi.NewRouter()
	parentRouter.Mount("/text-channels", textChanRouter)

	parentRouter.Get("{channel_uid}/messages", getMessages)
	parentRouter.Post("{channel_uid}/messages", postMessage)
	parentRouter.Delete("/{channel_uid}", deleteChannel)
	parentRouter.Post("/", createChannel)
	parentRouter.Post("/", getChannels)
}
