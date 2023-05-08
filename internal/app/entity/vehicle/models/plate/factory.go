package plate

import "math/rand"

func New() int {
	return rand.Intn(1000)
}
