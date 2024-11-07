package domain

import "time"

// Tweet representa un mensaje corto publicado por un usuario.

type Tweet struct {
	ID        int       `gorm:"primaryKey"`
	UserID    int       `gorm:"index;not null"` // índice para búsquedas rápidas
	Content   string    `gorm:"type:varchar(280);not null"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
}
