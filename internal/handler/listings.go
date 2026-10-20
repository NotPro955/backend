package handler

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"time"
)

type listing struct {
	ID          string    `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Price       string    `json:"price"`
	CreatedAt   time.Time `json:"created_at"`
	City        string    `json:"city"`
	Status      string    `json:"status"`
}

type ListingHandler struct {
	db *sql.DB
}

func NewListingHandler(db *sql.DB) *ListingHandler {
	return &ListingHandler{
		db: db,
	}
}

func (lh ListingHandler) List(w http.ResponseWriter, r *http.Request) {
	rows, err := lh.db.Query(`SELECT title,id,description,price,city,created_at,status
		FROM listings
		ORDER BY created_at
		LIMIT 10`)
	if err != nil {
		log.Printf("db query: %v", err)
		http.Error(w, "internal error", http.StatusInternalServerError)
		return
	}

	defer rows.Close()

	listings := []listing{}
	for rows.Next() {
		var l listing
		if err := rows.Scan(&l.Title, &l.ID, &l.Description, &l.Price, &l.City, &l.CreatedAt, &l.Status); err != nil {
			log.Printf("rows.scan: %v", err)
			http.Error(w, "internal error", http.StatusInternalServerError)
			return
		}

		listings = append(listings, l)
	}

	if rows.Err() != nil {
		log.Printf("rows err: %v", rows.Err())
		http.Error(w, "internal error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(listings)
}

func (lh ListingHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	_, err := lh.db.Exec(`DELETE FROM listings WHERE id = $1`, id)

	if err != nil {
		log.Printf("delete: %v", err)
		http.Error(w, "internal error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)

}
