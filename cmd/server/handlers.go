package main

import (
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"strings"

	"github.com/go-chi/chi/v5"
	"isustud.com/m/internal/models"
	"isustud.com/m/internal/sl"
	"isustud.com/m/storage/mariadb"
)

func TagsGetHandler(log *slog.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		op := "TagsGetHandler"
		data := make([]map[string]string, 2)
		data[0] = map[string]string{
			"tag_id": "c96b47d7-c537-412b-a480-e8e793196650",
			"pass":   "H@EGOHDSKMV",
		}
		data[1] = map[string]string{
			"tag_id": "38091eee-25b4-43f0-b0b7-79355bfc2901",
			"pass":   "SDKFJASD!@#",
		}
		if err := encode(w, r, 200, data); err != nil {
			log.Error(op+": error encode", sl.Err(err))
		}
	}
}

func TagsPostHandler(
	log *slog.Logger,
	storage *mariadb.Storage,
) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		op := "TagsPostHandler "

		reqToken := r.Header.Get("Authorization")
		splitToken := strings.Split(reqToken, "Bearer")
		if len(splitToken) != 2 {
			log.Error(op + "err dplit token")
			http.Error(w, "Incorrect token", http.StatusBadRequest)
		}

		reqToken = strings.TrimSpace(splitToken[1])
		log.Info(reqToken)
		key, err := storage.GetApiKey(reqToken)

		if err != nil {
			log.Error(op+"bad token", sl.Err(err))
			http.Error(w, "Bad Token", http.StatusUnauthorized)
		}
		tag, err := decode[models.Tag](r)
		log.Info(fmt.Sprintf("%+v", tag))
		if err != nil {
			log.Error(op+"err decode body", sl.Err(err))
			w.WriteHeader(http.StatusInternalServerError)
		}
		tag, err = storage.GetTag(tag)
		if err != nil {
			log.Error(op+"err decode body", sl.Err(err))
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		}
		if err = storage.SetLogs(key.Login, tag.Alias); err != nil {
			log.Error(op+"err loging activity", sl.Err(err))
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		}
		encode(w, r, 200, tag)
	}
}

func KeyGetHandler(
	log *slog.Logger,
	storage *mariadb.Storage,
) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		op := "KeyGetHandler "
		login := chi.URLParam(r, "login")
		token, err := GenerateRandomStringURLSafe(32)
		if err != nil {
			http.Error(w, http.StatusText(500), http.StatusInternalServerError)
			return
		}
		err = storage.SetApiKey(login, token)
		if err != nil {
			if errors.Is(err, mariadb.ErrDuplicate) {
				http.Error(w, "Имя уже занято", http.StatusConflict)
				return
			}
			fmt.Printf("%T\n", err)
			log.Error(op+"err db", sl.Err(err))
			http.Error(w, http.StatusText(400), http.StatusBadRequest)
		}
		encode(w, r, 200, models.Key{Login: login, APIKey: token})
	}
}
