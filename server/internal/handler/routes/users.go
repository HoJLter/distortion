package routes

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"io"
	"log/slog"
	"net/http"
	"server/internal/dto"
	"server/internal/repository"

	"github.com/go-chi/chi"
)

func (h Handler) getUsers(w http.ResponseWriter, r *http.Request) {
	usersRecords, err := h.users.GetAll(r.Context())
	if err != nil {
		slog.Error("Error while db request", "err", err)
		http.Error(w, "Db request error", http.StatusInternalServerError)
		return
	}
	var users []dto.GetUserResponse
	for _, record := range usersRecords {
		users = append(users, dto.RecordToUserResponse(record))
	}

	if err = json.NewEncoder(w).Encode(users); err != nil {
		slog.Error("Error while writing request", "err", err)
		http.Error(w, "Error while writing request", http.StatusInternalServerError)
		return
	}
}

func (h Handler) createUser(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		slog.Error("Error while reading body", "err", err)
		http.Error(w, "Error while reading body", http.StatusInternalServerError)
		return
	}

	var requestData dto.CreateUserRequest
	if err = json.Unmarshal(body, &requestData); err != nil {
		slog.Error("Invalid format", "err", err)
		http.Error(w, "Invalid format", http.StatusBadRequest)
		return
	}

	usernameHash := md5.Sum([]byte(requestData.Username))
	usernameHashString := hex.EncodeToString(usernameHash[:])
	record := repository.UserRecord{
		UserUID:    usernameHashString,
		Username:   requestData.Username,
		AvatarPath: requestData.AvatarPath,
	}
	if err = h.users.Create(r.Context(), record); err != nil {
		slog.Error("Db request error", "err", err)
		http.Error(w, "Db request error", http.StatusInternalServerError)
		return
	}
}

func SetupUsersRouter(r chi.Router, h *Handler) {
	usersRouter := chi.NewRouter()
	r.Mount("/users", usersRouter)

	usersRouter.Get("/", h.getUsers)
	usersRouter.Post("/", h.createUser)

}
