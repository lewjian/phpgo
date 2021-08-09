package phpgo

import (
	"log"
	"testing"
)

func TestInArray(t *testing.T) {
	item := 123
	data := []int{124, 3543, 23, 12}
	expect := false
	res := InArray(item, data)
	if res != expect {
		t.Fail()
	}
	if false != InArray("12", []int{12}) {
		t.Fail()
	}
	if true != InArray("12", []string{"12", "2"}) {
		t.Fail()
	}
	if true != InArray(12.43, []float64{12.43, 242.12}) {
		t.Fail()
	}
}

func TestArrayUnique(t *testing.T) {
	data := []string{"a", "b","a", "c", "hel", "hel"}
	res := ArrayUnique(data)
	log.Println(res.([]string))
}
