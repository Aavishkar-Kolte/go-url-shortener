package utils

import (
	"github.com/google/uuid"
)

func GenerateShortURL() string {
	return uuid.New().String()
}
