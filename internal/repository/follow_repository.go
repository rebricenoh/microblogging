package repository

import (
	"microblogging/internal/domain"

	"gorm.io/gorm"
)

// FollowRepository define los métodos para manejar las relaciones de seguimiento de usuarios.

type FollowRepository interface {
	FollowUser(followerID, followedID int) error
	UnfollowUser(followerID, followedID int) error
	GetFollowedUserIDs(userID int) ([]int, error)
}

// GormFollowRepository implementa FollowRepository usando GORM.

type GormFollowRepository struct {
	db *gorm.DB
}

// NewGormFollowRepository crea una nueva instancia de GormFollowRepository.

func NewGormFollowRepository(db *gorm.DB) *GormFollowRepository {
	return &GormFollowRepository{db: db}
}

// FollowUser permite a un usuario seguir a otro.

func (r *GormFollowRepository) FollowUser(followerID, followedID int) error {
	follow := domain.Follow{
		FollowerID: followerID,
		FollowedID: followedID,
	}
	return r.db.Create(&follow).Error
}

// UnfollowUser permite a un usuario dejar de seguir a otro.

func (r *GormFollowRepository) UnfollowUser(followerID, followedID int) error {
	return r.db.Where("follower_id = ? AND followed_id = ?", followerID, followedID).Delete(&domain.Follow{}).Error
}

// GetFollowedUserIDs obtiene los IDs de los usuarios seguidos por un usuario.

func (r *GormFollowRepository) GetFollowedUserIDs(userID int) ([]int, error) {
	var followedIDs []int
	err := r.db.Table("follows").Select("followed_id").Where("follower_id = ?", userID).Scan(&followedIDs).Error
	return followedIDs, err
}
