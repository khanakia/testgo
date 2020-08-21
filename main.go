package main

import (
	"fmt"

	"github.com/khanakia/testgo/pkg/hello"
	"github.com/khanakia/testgo/pkg/hello2"
)

func main() {
	fmt.Println(hello.Say)
	fmt.Println(hello2.GetV())
}
