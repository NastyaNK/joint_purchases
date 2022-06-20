package router

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
)

type Order struct {
	ID        int
	ProductID int
	UserID    int
	Count     int
}

func (h *Handlers) Buy(w http.ResponseWriter, r *http.Request) {
	data := []Order{}
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		h.Logger.Println("decode order err: ", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("500 - Something bad happened!"))
		return
	}

	for _, val := range data {
		check := Basket{}
		err := h.DB.Get(&check, "SELECT * FROM orders WHERE user_id=$1 and product_id=$2", val.UserID, val.ProductID)
		if err != nil && err != sql.ErrNoRows {
			h.Logger.Println("get one product err: ", err)
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("500 - Something bad happened!"))
			return
		}

		if err == sql.ErrNoRows {
			if _, err := h.DB.Exec("INSERT INTO orders (product_id, user_id, count) VALUES ($1,$2,$3)",
				val.ProductID, val.UserID, val.Count); err != nil {
				fmt.Println("create order err: ", err)
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte("500 - Something bad happened!"))
				return
			}
		} else {
			if _, err := h.DB.Exec("UPDATE orders SET count=count+$1 WHERE id=$2",
				check.Count, check.ID); err != nil {
				fmt.Println("UPDATE orders err: ", err)
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte("500 - Something bad happened!"))
				return
			}
		}

		if _, err := h.DB.Exec("DELETE FROM baskets WHERE user_id=$1 and product_id=$2", val.UserID, val.ProductID); err != nil {
			fmt.Println("DELETE baskets err: ", err)
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("500 - Something bad happened!"))
			return
		}
	}

	w.WriteHeader(http.StatusCreated)
}
