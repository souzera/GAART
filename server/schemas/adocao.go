package schemas

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Adocao struct {
	gorm.Model
	ID uuid.UUID `gorm:"type:uuid;primary_key;default:uuid_generate_v4()"`

	AnimalID uuid.UUID `gorm:"type:uuid;not null"`
	Animal   Animal    `gorm:"foreignKey:AnimalID"`

	TutorID uuid.UUID `gorm:"type:uuid;not null"`
	Tutor   Tutor     `gorm:"foreignKey:TutorID"`

	Descricao string `gorm:"type:varchar(255);"`
	Status    bool   `gorm:"type:boolean;default:true"`
}

// REQUESTS

type AdocaoRequest struct {
	AnimalID  string `json:"animal" binding:"required"`
	TutorID   string `json:"tutor" binding:"required"`
	Descricao *string   `json:"descricao"`
	Status    *bool     `json:"status"`
}
