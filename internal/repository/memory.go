package repository

import (
	"cinema-friend/internal/models"
	"fmt"
	"sync"
	"time"
)

type Storage struct {
	mu          sync.RWMutex
	Movies      map[string]models.Movie
	BookedSeats map[string]map[string]bool
}

func NewStorage() *Storage {
	s := &Storage{
		Movies:      make(map[string]models.Movie),
		BookedSeats: make(map[string]map[string]bool),
	}
	// Моковые данные
	s.Movies["1"] = models.Movie{ID: "1", Title: "Зверополис 2", Genre: "Мультфильм", ImageURL: "https://images.unsplash.com/photo-1534447677768-be436bb09401?w=300"}
	s.Movies["2"] = models.Movie{ID: "2", Title: "Человек-бензопила", Genre: "Аниме", ImageURL: "https://images.unsplash.com/photo-1607604276583-eef5d076aa5f?w=300"}
	return s
}

func (s *Storage) GetMovies() []models.Movie {
	s.mu.RLock()
	defer s.mu.RUnlock()
	var list []models.Movie
	for _, m := range s.Movies {
		list = append(list, m)
	}
	return list
}

func (s *Storage) BookSeat(req models.BookingRequest) (string, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if s.BookedSeats[req.SessionID] == nil {
		s.BookedSeats[req.SessionID] = make(map[string]bool)
	}

	seatKey := fmt.Sprintf("%d_%d", req.Row, req.Seat)
	if s.BookedSeats[req.SessionID][seatKey] {
		return "", fmt.Errorf("место уже занято")
	}

	s.BookedSeats[req.SessionID][seatKey] = true
	return fmt.Sprintf("TICK-%d", time.Now().Unix()), nil
}
