package router

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
)

type Basket struct {
	ID        int
	ProductID int `db:"product_id"`
	UserID    int `db:"user_id"`
	Count     int
	AddedTime time.Time `db:"added_time"`
}

func (h *Handlers) AddBasket(w http.ResponseWriter, r *http.Request) {
	data := Basket{}
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		h.Logger.Println("decode basket err: ", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("500 - Something bad happened!"))
		return
	}

	check := Basket{}
	err := h.DB.Get(&check, "SELECT * FROM baskets WHERE user_id=$1 and product_id=$2", data.UserID, data.ProductID)
	if err != nil && err != sql.ErrNoRows {
		h.Logger.Println("get one product err: ", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("500 - Something bad happened!"))
		return
	}

	if err == sql.ErrNoRows {
		data.AddedTime = time.Now()

		if _, err := h.DB.Exec("INSERT INTO baskets (product_id, user_id, count, added_time) VALUES ($1,$2,$3,$4)",
			data.ProductID, data.UserID, data.Count, data.AddedTime); err != nil {
			fmt.Println("create basket err: ", err)
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("500 - Something bad happened!"))
			return
		}
	} else {
		if _, err := h.DB.Exec("UPDATE baskets SET count=$1 WHERE id=$2",
			check.Count+1, check.ID); err != nil {
			fmt.Println("create basket err: ", err)
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("500 - Something bad happened!"))
			return
		}
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("{'message': 'Добавлено в корзину'}"))
}

func (h *Handlers) GetBasket(w http.ResponseWriter, r *http.Request) {
	UserID := chi.URLParam(r, "userID")
	res := []Basket{}

	if err := h.DB.Select(&res, "SELECT * FROM baskets WHERE user_id = $1", UserID); err != nil {
		fmt.Println("get basket err: ", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("500 - Something bad happened!"))
		return
	}

	byteData, err := json.Marshal(&res)
	if err != nil {
		h.Logger.Println("marshal basket err: ", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("500 - Something bad happened!"))
		return
	}

	if _, err := w.Write(byteData); err != nil {
		h.Logger.Println("get basket err: ", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("500 - Something bad happened!"))
		return
	}

	w.WriteHeader(http.StatusOK)
}
