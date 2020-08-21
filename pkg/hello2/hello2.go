package hello2

import (
	"github.com/google/uuid"
)

const Say = "Hello2"

func GetV() string {
	return uuid.New().String()
}
