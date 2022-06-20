package router

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type User struct {
	ID       int
	Name     string
	Email    string
	Password string
	Role     string
}

func (h *Handlers) Login(w http.ResponseWriter, r *http.Request) {
	data := User{}
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		h.Logger.Println("decode login err: ", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("500 - Something bad happened!"))
		return
	}

	var user User
	if err := h.DB.Get(&user, "SELECT * FROM users WHERE email=$1", data.Email); err != nil {
		h.Logger.Println("get user by login err: ", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("500 - Something bad happened!"))
		return
	}

	hash := sha256.Sum256([]byte(data.Password))
	data.Password = fmt.Sprintf("%x", hash)

	if user.Password != data.Password {
		h.Logger.Println("invalid password")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("500 - Something bad happened!"))
		return
	}

	w.Write([]byte(fmt.Sprintf("%v", user.ID)))
	w.WriteHeader(http.StatusOK)
}

func (h *Handlers) Register(w http.ResponseWriter, r *http.Request) {
	data := User{}
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		h.Logger.Println("decode user register err: ", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("500 - Something bad happened!"))
		return
	}

	hash := sha256.Sum256([]byte(data.Password))
	data.Password = fmt.Sprintf("%x", hash)

	if _, err := h.DB.Exec(`INSERT INTO users ("name", "email", "password", "role") VALUES ($1,$2,$3,$4)`,
		data.Name, data.Email, data.Password, data.Role); err != nil {
		h.Logger.Println("create user err: ", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("500 - Something bad happened!"))
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (h *Handlers) getUser(w http.ResponseWriter, r *http.Request) {
	UserID := chi.URLParam(r, "userID")
	data := User{}

	if err := h.DB.Get(&data, "SELECT * FROM users WHERE id = $1", UserID); err != nil {
		fmt.Println("get user err: ", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("500 - Something bad happened!"))
		return
	}

	byteData, err := json.Marshal(&data)
	if err != nil {
		h.Logger.Println("marshal user err: ", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("500 - Something bad happened!"))
		return
	}

	if _, err := w.Write(byteData); err != nil {
		h.Logger.Println("get user err: ", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("500 - Something bad happened!"))
		return
	}

	w.WriteHeader(http.StatusOK)
}
