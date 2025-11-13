package item

import (
	"time"
)

type Item struct {
	ID         uint       `json:"id" gorm:"primaryKey"`
	Name       string     `json:"name" gorm:"type:varchar(100);not null"`
	Quantity   float64    `json:"quantity" gorm:"not null;default 0"`
	Unit       string     `json:"unit" gorm:"type:varchar(20);not null"`
	ExpiryDate *time.Time `json:"expiry_date,omitempty"`
	CategoryID *uint      `json:"category_id,omitempty"`
	Notes      string     `json:"notes,omitempty" gorm:"type:text"`
	CreatedAt  time.Time  `json:"created_at"`
	UpdatedAt  time.Time  `json:"updated_at"`
}
