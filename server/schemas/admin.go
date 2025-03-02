package schemas

import (

	"github.com/google/uuid"
	"gorm.io/gorm"
)

/*
 * ADMIN, SUPERUSER
 * ORGANIZAÇÃO, ONGs
 */

type Admin struct {
	gorm.Model
	ID        uuid.UUID      `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`

	Usuario   Usuario   `gorm:"foreignKey:UsuarioID" json:"usuario,omitempty"`
	UsuarioID uuid.UUID `gorm:"type:uuid;not null"`

	Nome string `json:"nome"`
}
