package moviesService

import (
	"sync"

	"github.com/slothmakeout/movies-project/data/models"
	"gorm.io/gorm"
)

type MoviesService struct {
	db *gorm.DB
}

var moviesServiceInstance *MoviesService
var moviesServiceOnce sync.Once

func GetService(db *gorm.DB) *MoviesService {
	moviesServiceOnce.Do(func() {
		moviesServiceInstance = NewMoviesService(db)
	})
	return moviesServiceInstance
}

func NewMoviesService(db *gorm.DB) *MoviesService {
	return &MoviesService{
		db: db,
	}
}

func (ms *MoviesService) GetAllMovies() (*models.Movies, error) {
	var movies *models.Movies
	result := ms.db.Find(&movies)
	if result.Error != nil {
		return nil, result.Error
	}
	return movies, nil
}

func (ms *MoviesService) GetMovieById(id int) (*models.Movie, error) {
	var movie models.Movie
	result := ms.db.First(&movie, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &movie, nil
}
