package reviewsService

import (
	"sync"

	"github.com/slothmakeout/movies-project/data/models"
	"gorm.io/gorm"
)

type ReviewsService struct {
	db *gorm.DB
}

var reviewsServiceInstance *ReviewsService
var reviewsServiceOnce sync.Once

func GetReviewsService(db *gorm.DB) *ReviewsService {
	reviewsServiceOnce.Do(func() {
		reviewsServiceInstance = NewReviewsService(db)
	})
	return reviewsServiceInstance
}

func NewReviewsService(db *gorm.DB) *ReviewsService {
	return &ReviewsService{
		db: db,
	}
}

func (rs *ReviewsService) GetAllReviews() ([]models.Review, error) {
	var reviews []models.Review
	result := rs.db.Find(&reviews)
	if result.Error != nil {
		return nil, result.Error
	}
	return reviews, nil
}

func (rs *ReviewsService) GetReviewById(id int) (*models.Review, error) {
	var review models.Review
	result := rs.db.First(&review, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &review, nil
}

func (rs *ReviewsService) GetReviewsByMovieId(movieId int) ([]models.Review, error) {
	var reviews []models.Review
	result := rs.db.Where("movie_id = ?", movieId).Find(&reviews)
	if result.Error != nil {
		return nil, result.Error
	}
	return reviews, nil
}

func (rs *ReviewsService) AddReview(review *models.Review) error {
	result := rs.db.Create(review)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (rs *ReviewsService) UpdateReview(id int, updatedReview *models.Review) error {
	// Поиск отзыва по ID
	existingReview, err := rs.GetReviewById(id)
	if err != nil {
		return err
	}

	// Обновление полей отзыва
	existingReview.Rating = updatedReview.Rating
	existingReview.ReviewText = updatedReview.ReviewText

	// Сохранение обновленного отзыва
	result := rs.db.Save(existingReview)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (rs *ReviewsService) DeleteReview(id int) error {
	// Поиск отзыва по ID
	existingReview, err := rs.GetReviewById(id)
	if err != nil {
		return err
	}

	// Удаление отзыва
	result := rs.db.Delete(existingReview)
	if result.Error != nil {
		return result.Error
	}

	return nil
}
