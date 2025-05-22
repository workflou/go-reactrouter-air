package handler

import (
	"encoding/json"
	"log/slog"
	"net/http"
	"template/database"
	"template/store"
	"time"

	"golang.org/x/crypto/bcrypt"
)

func AdminCount(w http.ResponseWriter, r *http.Request) {
	count, err := store.New(database.DB).CountAdmins(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]int64{
		"count": count,
	})
}

func CreateAdmin(w http.ResponseWriter, r *http.Request) {
	time.Sleep(3 * time.Second)

	type request struct {
		Name     string `json:"name"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	count, err := store.New(database.DB).CountAdmins(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if count > 0 {
		http.Error(w, http.StatusText(http.StatusConflict), http.StatusConflict)
		return
	}

	var req request
	json.NewDecoder(r.Body).Decode(&req)

	slog.Info("CreateAdmin", slog.String("name", req.Name), slog.String("email", req.Email))

	if req.Name == "" || req.Password == "" || req.Email == "" {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = store.New(database.DB).CreateAdmin(r.Context(), store.CreateAdminParams{
		Name:     req.Name,
		Email:    req.Email,
		Password: string(hash),
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
