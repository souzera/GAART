package schemas

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Especie struct {
	gorm.Model
	ID        uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	DeletedAt time.Time `json:"deteledAt,omitempty"`

	Nome      string `json:"nome"`
	Genero    string `json:"genero"`
	Domesitco bool   `json:"domestico"`
}
