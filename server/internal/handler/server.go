package handler

import (
	"log/slog"
	"net/http"
	"server/internal/config"
	"strconv"

	"github.com/go-chi/chi"
)

/*
 *
 * ТЕКСТОВЫЕ КАНАЛЫ
 * GET  /api/channels/{channel_uid}/messages - получить сообщения
 * POST /api/channels/{channel_uid}/messages - отправить сообщение
 * POST /api/channels/ - добавить канал
 * DELETE /api/channels/{channel_uid} - удалить канал
 *
 * ГОЛОСОВЫЕ КАНАЛЫ
 * POST   /api/channels/{channel_uid}/connection - подключится к голосовому каналу
 * DELETE /api/channels/{channel_uid}/connection/{connection_uid}
 *
 * ПОЛЬЗОВАТЕЛИ
 * GET   /api/users - получить всех пользователей
 * POST  /api/users - добавить пользователя
 * GET   /api/users/{user_uid} - получить конкретного пользователя
 * PATCH /api/users/{user_uid} - изменить пользователя
 * POST  /api/users/{user_uid}/ban - забанить пользователя
 * POST  /api/users/{user_uid}/pardon - разбанить пользователя
 *
 * МАРШРУТИЗАЦИЯ В СЕТИ
 * GET    /api/neigh - получить список серверов-соседей
 * POST   /api/neigh - добавить сервер-соседа
 * PATCH  /api/neigh/{record_id} - изменить запись в таблице маршрутизации
 * DELETE /api/neigh/{record_id} - удалить запись в таблице маршрутизации
 *
 */

func SlogLoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		slog.Debug("HTTP request achieved",
			"method", request.Method,
			"url", request.URL.Path)

		next.ServeHTTP(writer, request)
	})
}

func RunServer(cfg config.ServerConfig) error {
	rootRouter := chi.NewRouter()
	rootRouter.Use(SlogLoggingMiddleware)

	apiRouter := chi.NewRouter()
	rootRouter.Mount("/api", apiRouter)

	usersRouter := chi.NewRouter()
	voiceChanRouter := chi.NewRouter()
	netRoutingRouter := chi.NewRouter()

	apiRouter.Mount("/users", usersRouter)
	apiRouter.Mount("/voice-channels", voiceChanRouter)
	apiRouter.Mount("/routing", netRoutingRouter)

	addressToListen := ":" + strconv.Itoa(cfg.Port)
	err := http.ListenAndServe(addressToListen, rootRouter)
	if err != nil {
		slog.Error("Server starting error", "err", err)
		return err
	}
	return nil
}
