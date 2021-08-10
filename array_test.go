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

func TestArrayChunkInt(t *testing.T) {
	data := []int{1,2,3,4,5,6,7,8}
	log.Println(ArrayChunkInt(data, 3))
	log.Println(ArrayChunkInt(data, 0))
	log.Println(ArrayChunkInt(data, 5))
	log.Println(ArrayChunkInt(data, 10))
}

func TestArrayChunkString(t *testing.T) {
	data := []string{"1","2","2", "5","hel", "45sfd"}
	log.Println(ArrayChunkString(data, 3))
	log.Println(ArrayChunkString(data, 0))
	log.Println(ArrayChunkString(data, 5))
	log.Println(ArrayChunkString(data, 10))
}

func TestArrayMerge(t *testing.T) {
	data, err := ArrayMerge([]string{"aksfh", "s"}, []string{"ascgjhag"})
	log.Println(data.([]string), err)
}

func TestArraySearch(t *testing.T) {
	log.Println(ArraySearch(1, []int{2,3,4,1,6}))
	log.Println(ArraySearch("abc", []int{2,3,4,1,6}))
	log.Println(ArraySearch("abc", []string{"ashd", "salchj", "abc"}))
	type S struct {
		ID int
		Name string
	}
	log.Println(ArraySearch(S{
		ID:   2,
		Name: "liujian",
	}, []S{{2,"liujian"}}))
}
func TestArrayWalk(t *testing.T) {
	data := []int{123,456,234,8, 45,23}
	var result []int
	ArrayWalk(data, func(item interface{}, index int) bool {
		if index > 2 {
			return false
		}
		num := item.(int)
		result = append(result, num*3)
		return true
	})
	log.Println(data, result)

}
