package schemas

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Adocao struct {
	gorm.Model
	ID        uuid.UUID `gorm:"type:uuid;primary_key;default:uuid_generate_v4()"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	DeletedAt time.Time `json:"deteledAt,omitempty"`
	AnimalID  uuid.UUID `gorm:"type:uuid;not null"`
	TutorID   uuid.UUID `gorm:"type:uuid;not null"`
	Status    string    `gorm:"not null"`
}
