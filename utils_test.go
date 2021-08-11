package phpgo

import (
	"log"
	"testing"
)

func TestIf(t *testing.T) {
	a := 1
	b := 2
	log.Println(a > b, a, b)
}