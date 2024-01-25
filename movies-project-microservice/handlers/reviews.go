// handlers/reviews.go
package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/slothmakeout/movies-project/data/models"
	services "github.com/slothmakeout/movies-project/pkg/reviews"
)

type Reviews struct {
	l              *log.Logger
	reviewsService *services.ReviewsService
}

func NewReviews(l *log.Logger, reviewsService *services.ReviewsService) *Reviews {
	return &Reviews{l, reviewsService}
}

func (r *Reviews) GetReviews(rw http.ResponseWriter, req *http.Request) {
	r.l.Println("Handled GET Reviews")

	// Получение списка всех отзывов с использованием сервиса
	reviewsList, err := r.reviewsService.GetAllReviews()
	if err != nil {
		http.Error(rw, "Unable to get reviews list", http.StatusBadRequest)
		return
	}

	// Отправка списка в формате JSON в ResponseWriter
	rw.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(rw).Encode(reviewsList); err != nil {
		http.Error(rw, "Unable to encode JSON", http.StatusInternalServerError)
		return
	}
}

func (r *Reviews) GetReviewById(rw http.ResponseWriter, req *http.Request) {
	r.l.Println("Handled GET Review by ID")

	// Получение значения id из переменных пути
	vars := mux.Vars(req)
	idStr, ok := vars["id"]
	if !ok {
		http.Error(rw, "ID not provided in the request", http.StatusBadRequest)
		return
	}

	// Преобразование строки id в число
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(rw, "Invalid ID format", http.StatusBadRequest)
		return
	}

	// Получение отзыва по id с использованием сервиса
	review, err := r.reviewsService.GetReviewById(id)
	if err != nil {
		http.Error(rw, "Unable to get review by ID", http.StatusNotFound)
		return
	}

	// Отправка отзыва в формате JSON в ResponseWriter
	rw.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(rw).Encode(review); err != nil {
		http.Error(rw, "Unable to encode JSON", http.StatusInternalServerError)
		return
	}
}

func (r *Reviews) GetReviewsByMovieId(rw http.ResponseWriter, req *http.Request) {
	r.l.Println("Handled GET Reviews by Movie ID")

	// Получение значения movieId из переменных пути
	vars := mux.Vars(req)
	movieIdStr, ok := vars["id"]
	if !ok {
		http.Error(rw, "Movie ID not provided in the request", http.StatusBadRequest)
		return
	}

	// Преобразование строки movieId в число
	movieId, err := strconv.Atoi(movieIdStr)
	if err != nil {
		http.Error(rw, "Invalid Movie ID format", http.StatusBadRequest)
		return
	}

	// Получение отзывов по movieId с использованием сервиса
	reviewsList, err := r.reviewsService.GetReviewsByMovieId(movieId)
	if err != nil {
		http.Error(rw, "Unable to get reviews by Movie ID", http.StatusNotFound)
		return
	}

	// Отправка списка в формате JSON в ResponseWriter
	rw.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(rw).Encode(reviewsList); err != nil {
		http.Error(rw, "Unable to encode JSON", http.StatusInternalServerError)
		return
	}
}

func (r *Reviews) AddReview(rw http.ResponseWriter, req *http.Request) {
	r.l.Println("Handled POST Review")

	// Декодирование JSON-данных из тела запроса
	var review models.Review
	err := json.NewDecoder(req.Body).Decode(&review)
	if err != nil {
		http.Error(rw, "Unable to decode JSON", http.StatusBadRequest)
		return
	}

	// Добавление отзыва с использованием сервиса
	err = r.reviewsService.AddReview(&review)
	if err != nil {
		http.Error(rw, "Unable to add review", http.StatusInternalServerError)
		return
	}

	// Отправка ответа об успешном добавлении
	rw.WriteHeader(http.StatusCreated)
}

func (r *Reviews) UpdateReview(rw http.ResponseWriter, req *http.Request) {
	r.l.Println("Handled PUT Review")

	// Получение значения id из переменных пути
	vars := mux.Vars(req)
	idStr, ok := vars["id"]
	if !ok {
		http.Error(rw, "ID not provided in the request", http.StatusBadRequest)
		return
	}

	// Преобразование строки id в число
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(rw, "Invalid ID format", http.StatusBadRequest)
		return
	}

	// Декодирование JSON-данных из тела запроса
	var updatedReview models.Review
	err = json.NewDecoder(req.Body).Decode(&updatedReview)
	if err != nil {
		http.Error(rw, "Unable to decode JSON", http.StatusBadRequest)
		return
	}

	// Обновление отзыва с использованием сервиса
	err = r.reviewsService.UpdateReview(id, &updatedReview)
	if err != nil {
		http.Error(rw, "Unable to update review", http.StatusInternalServerError)
		return
	}

	// Отправка ответа об успешном обновлении
	rw.WriteHeader(http.StatusOK)
}

func (r *Reviews) DeleteReview(rw http.ResponseWriter, req *http.Request) {
	r.l.Println("Handled DELETE Review")

	// Получение значения id из переменных пути
	vars := mux.Vars(req)
	idStr, ok := vars["id"]
	if !ok {
		http.Error(rw, "ID not provided in the request", http.StatusBadRequest)
		return
	}

	// Преобразование строки id в число
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(rw, "Invalid ID format", http.StatusBadRequest)
		return
	}

	// Удаление отзыва с использованием сервиса
	err = r.reviewsService.DeleteReview(id)
	if err != nil {
		http.Error(rw, "Unable to delete review", http.StatusInternalServerError)
		return
	}

	// Отправка ответа об успешном удалении
	rw.WriteHeader(http.StatusOK)
}
