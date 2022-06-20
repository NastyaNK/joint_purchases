package router

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type Product struct {
	ID          int
	Name        string
	Description string
	Required    int
	Price       float32
	StartTime   string `db:"start_time"`
	EndTime     string `db:"end_time"`
	Organizer   int
	Image       string
}

func (h *Handlers) List(w http.ResponseWriter, r *http.Request) {
	var res []Product
	if err := h.DB.Select(&res, "SELECT * FROM products"); err != nil {
		h.Logger.Println("get list products err: ", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("500 - Something bad happened!"))
		return
	}

	byteData, err := json.Marshal(&res)
	if err != nil {
		h.Logger.Println("marshal list products err: ", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("500 - Something bad happened!"))
		return
	}

	if _, err := w.Write(byteData); err != nil {
		h.Logger.Println("write list products err: ", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("500 - Something bad happened!"))
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (h *Handlers) Search(w http.ResponseWriter, r *http.Request) {
	text := "%" + chi.URLParam(r, "name") + "%"
	var res []Product
	if err := h.DB.Select(&res, "SELECT * FROM products WHERE name ILIKE $1", text); err != nil {
		h.Logger.Println("get list products err: ", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("500 - Something bad happened!"))
		return
	}

	byteData, err := json.Marshal(&res)
	if err != nil {
		h.Logger.Println("marshal list products err: ", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("500 - Something bad happened!"))
		return
	}

	if _, err := w.Write(byteData); err != nil {
		h.Logger.Println("write list products err: ", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("500 - Something bad happened!"))
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (h *Handlers) One(w http.ResponseWriter, r *http.Request) {
	productID := chi.URLParam(r, "productID")
	var res Product
	if err := h.DB.Get(&res, "SELECT * FROM products WHERE id=$1", productID); err != nil {
		h.Logger.Println("get one product err: ", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("500 - Something bad happened!"))
		return
	}

	byteData, err := json.Marshal(&res)
	if err != nil {
		h.Logger.Println("marshal one product err: ", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("500 - Something bad happened!"))
		return
	}

	if _, err := w.Write(byteData); err != nil {
		h.Logger.Println("write one product err: ", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("500 - Something bad happened!"))
		return
	}
	w.WriteHeader(http.StatusOK)
}
