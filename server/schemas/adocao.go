package schemas

import (

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Adocao struct {
	gorm.Model
	ID        uuid.UUID `gorm:"type:uuid;primary_key;default:uuid_generate_v4()"`

	AnimalID  uuid.UUID `gorm:"type:uuid;not null"`
	TutorID   uuid.UUID `gorm:"type:uuid;not null"`
	Status    string    `gorm:"not null"`
}
