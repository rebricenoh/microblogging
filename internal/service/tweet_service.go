package service

import (
	"errors"
	"microblogging/internal/domain"
	"microblogging/internal/repository"
	"time"
)

// TweetService maneja la lógica de negocio relacionada con los tweets.

type TweetService struct {
	repo repository.TweetRepository
}

// NewTweetService crea una nueva instancia de TweetService.

func NewTweetService(repo repository.TweetRepository) *TweetService {
	return &TweetService{repo: repo}
}

// PostTweet permite a un usuario publicar un tweet.

func (s *TweetService) PostTweet(userID int, content string) error {
	if len(content) > 280 {
		return errors.New("el tweet excede el límite de caracteres")
	}
	tweet := domain.Tweet{
		UserID:    userID,
		Content:   content,
		CreatedAt: time.Now(),
	}
	return s.repo.CreateTweet(tweet)
}

// GetTimeline obtiene el timeline de un usuario.

func (s *TweetService) GetTimeline(userID int) ([]domain.Tweet, error) {
	return s.repo.GetTimeline(userID)
}
