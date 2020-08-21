package hello2

import (
	"fmt"

	"github.com/google/uuid"
)

const Say = "Hello2"

func getV() {
	fmt.Println(uuid.New().String())
}
