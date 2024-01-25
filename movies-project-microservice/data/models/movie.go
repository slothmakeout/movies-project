package models

import (
	"encoding/json"
	"io"
	"time"
	"gorm.io/gorm"
)

type Movie struct {
	gorm.Model
	Title       string    `json:"title"`
	ReleaseDate time.Time `json:"release_date"`
	// Другие поля фильма, если необходимо
}

type Movies []*Movie

func (m Movies) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(m)
}

func GetMovies(db *gorm.DB) (Movies, error) {
	var movies Movies
	result := db.Find(&movies)
	if result.Error != nil {
		return nil, result.Error
	}
	return movies, nil
}
