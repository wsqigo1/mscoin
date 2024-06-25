package tools

import (
	"fmt"
	"math/rand"
)

func Rand4Num() string {
	return fmt.Sprintf("%04d", rand.Intn(9999))
}
