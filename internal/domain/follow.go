package domain

// Follow representa la relación de seguimiento entre usuarios.

type Follow struct {
	ID         int `gorm:"primaryKey"`
	FollowerID int `gorm:"index;not null"`
	FollowedID int `gorm:"index;not null"`
}
