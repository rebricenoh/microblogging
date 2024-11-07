package repository

import (
	"microblogging/internal/domain"

	"gorm.io/gorm"
)

// TweetRepository define los métodos para manejar los tweets.
type TweetRepository interface {
	CreateTweet(tweet domain.Tweet) error
	GetTimeline(userID int) ([]domain.Tweet, error)
}

// GormTweetRepository implementa TweetRepository usando GORM.
type GormTweetRepository struct {
	db *gorm.DB
}

// NewGormTweetRepository crea una nueva instancia de GormTweetRepository.
func NewGormTweetRepository(db *gorm.DB) *GormTweetRepository {
	return &GormTweetRepository{db: db}
}

// CreateTweet permite crear un nuevo tweet en la base de datos.
func (r *GormTweetRepository) CreateTweet(tweet domain.Tweet) error {
	return r.db.Create(&tweet).Error
}

// GetTimeline obtiene los tweets de los usuarios seguidos por un usuario específico.
func (r *GormTweetRepository) GetTimeline(userID int) ([]domain.Tweet, error) {
	var tweets []domain.Tweet
	query := `
        SELECT t.*
        FROM tweets t
        JOIN follows f ON t.user_id = f.followed_id
        WHERE f.follower_id = ?
        ORDER BY t.created_at DESC
    `
	err := r.db.Raw(query, userID).Scan(&tweets).Error
	return tweets, err
}
