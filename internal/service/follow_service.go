package service

import "microblogging/internal/repository"

// FollowService maneja la lógica de negocio para el seguimiento de usuarios.

type FollowService struct {
	repo repository.FollowRepository
}

// NewFollowService crea una nueva instancia de FollowService.

func NewFollowService(repo repository.FollowRepository) *FollowService {
	return &FollowService{repo: repo}
}

// FollowUser permite a un usuario seguir a otro.

func (s *FollowService) FollowUser(followerID, followedID int) error {
	return s.repo.FollowUser(followerID, followedID)
}

// UnfollowUser permite a un usuario dejar de seguir a otro.

func (s *FollowService) UnfollowUser(followerID, followedID int) error {
	return s.repo.UnfollowUser(followerID, followedID)
}
