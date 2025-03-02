package util

import (
	"time"

	"github.com/google/uuid"
)

func ParseStringToUUID(id string) (uuid.UUID, error) {
	return uuid.Parse(id)
}

func ParseStringToTime(value string) (time.Time, error) {
	return time.Parse(time.RFC3339, value)
}
