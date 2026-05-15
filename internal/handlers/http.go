package handlers

import (
	"cinema-friend/internal/models"
	"cinema-friend/internal/repository"
	"encoding/json"
	"net/http"
)

type Handler struct {
	Repo *repository.Storage
}

func enableCORS(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
	(*w).Header().Set("Access-Control-Allow-Headers", "Content-Type")
}

func (h *Handler) GetMoviesHandler(w http.ResponseWriter, r *http.Request) {
	enableCORS(&w)
	if r.Method == "OPTIONS" {
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(h.Repo.GetMovies())
}

func (h *Handler) BookSeatHandler(w http.ResponseWriter, r *http.Request) {
	enableCORS(&w)
	if r.Method == "OPTIONS" {
		return
	}

	var req models.BookingRequest
	json.NewDecoder(r.Body).Decode(&req)

	ticketID, err := h.Repo.BookSeat(req)
	w.Header().Set("Content-Type", "application/json")

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(models.BookingResponse{Success: false, Message: err.Error()})
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(models.BookingResponse{Success: true, TicketID: ticketID, Message: "Успех!"})
}
