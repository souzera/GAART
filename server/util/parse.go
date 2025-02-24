package util

import "github.com/google/uuid"

func ParseStringToUUID(id string) (uuid.UUID, error) {
	return uuid.Parse(id)
}
