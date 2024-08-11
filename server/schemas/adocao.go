package schemas

import (
	"time"

	"github.com/google/uuid"
)

type Adocao struct {
	ID       uuid.UUID
	AnimalID uuid.UUID
	TutorID  uuid.UUID
	Data     time.Time
}
