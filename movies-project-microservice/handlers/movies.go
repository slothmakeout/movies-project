package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/slothmakeout/movies-project/pkg/movies"
)

type Movies struct {
	l *log.Logger
	moviesService *moviesService.MoviesService
}

func NewMovies(l *log.Logger, moviesService *moviesService.MoviesService) *Movies {
	return &Movies{l, moviesService}
}

func (p *Movies) GetMovies(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("Handled GET Movies")

	// Получение списка всех фильмов с использованием сервиса
	movieList, err := p.moviesService.GetAllMovies()
	if err != nil {
		http.Error(rw, "Unable to get movie list", http.StatusBadRequest)
		return
	}

	// Отправка списка в формате JSON в ResponseWriter
	err = movieList.ToJSON(rw)
	if err != nil {
		http.Error(rw, "Unable to marshal json", http.StatusInternalServerError)
		return
	}
}

func (p *Movies) GetMovieById(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("Handled GET Movie by ID")

	// Получение значения id из переменных пути
	vars := mux.Vars(r)
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

	// Получение фильма по id с использованием сервиса
	movie, err := p.moviesService.GetMovieById(id)
	if err != nil {
		http.Error(rw, "Unable to get movie by ID", http.StatusNotFound)
		return
	}

	// Отправка фильма в формате JSON в ResponseWriter
	rw.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(rw).Encode(movie); err != nil {
		http.Error(rw, "Unable to encode JSON", http.StatusInternalServerError)
		return
	}
}