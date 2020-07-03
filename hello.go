package module

import (
	"fmt"

	uuid "github.com/satori/go.uuid"
)

func Hello() {
	fmt.Println(uuid.NewV4())
}
