package phpgo

import (
	"reflect"
	"testing"
)

type M struct {
	x int
	y int
}

func TestInArray(t *testing.T) {
	// 整形
	item1 := 325
	arrayData1 := []int{1, 325, 57, 324}
	if got := InArray(item1, arrayData1); !got {
		t.Errorf("item=%v, arrayData=%v, got=%v", item1, arrayData1, got)
	}

	item1 = 325
	arrayData1 = []int{1, 3251, 57, 324}
	if got := InArray(item1, arrayData1); got {
		t.Errorf("item=%v, arrayData=%v, got=%v", item1, arrayData1, got)
	}

	// 浮点型
	item2 := 12.24
	arrayData2 := []float64{12.124, 45, 3532.24, 12.24, 4642.2}
	if got := InArray(item2, arrayData2); !got {
		t.Errorf("item=%v, arrayData=%v, got=%v", item2, arrayData2, got)
	}
	item2 = 12.243
	arrayData2 = []float64{12.124, 45, 3532.24, 12.24, 4642.2}
	if got := InArray(item2, arrayData2); got {
		t.Errorf("item=%v, arrayData=%v, got=%v", item2, arrayData2, got)
	}
	// 字符串
	item3 := "hello"
	arrayData3 := []string{"a", "c", "e", "hello"}
	if got := InArray(item3, arrayData3); !got {
		t.Errorf("item=%v, arrayData=%v, got=%v", item3, arrayData3, got)
	}
	item3 = "ahfoah"
	arrayData3 = []string{"a", "c", "e", "hello"}
	if got := InArray(item2, arrayData2); got {
		t.Errorf("item=%v, arrayData=%v, got=%v", item3, arrayData3, got)
	}

}

func TestArrayUnique(t *testing.T) {
	s := []int{1, 2, 5, 2, 5, 6, 1}
	got := ArrayUnique(s)
	if !reflect.DeepEqual([]int{1, 2, 5, 6}, got) {
		t.Errorf("item=%v, got=%v", s, got)
	}
	s1 := []float64{1, 2, 5, 2, 5, 6, 1}
	got1 := ArrayUnique(s1)
	if !reflect.DeepEqual([]float64{1, 2, 5, 6}, got1) {
		t.Errorf("item=%v, got=%v", s1, got1)
	}
	s2 := []string{"a", "a", "b", "b", "c", "e", "c", "e"}

	got2 := ArrayUnique(s2)
	if !reflect.DeepEqual([]string{"a", "b", "c", "e"}, got2) {
		t.Errorf("item=%v, got=%v", s2, got2)
	}
}

func TestArraySum(t *testing.T) {
	s := []int{1, 2, 5, 2}
	got := ArraySum(s)
	want := 10
	if want != got {
		t.Errorf("item=%v, want=%v, got=%v", s, want, got)
	}

	s1 := []float64{1.5, 2.5, 5, 2}
	got1 := ArraySum(s1)
	want1 := 11.0
	if want1 != got1 {
		t.Errorf("item=%v, want=%v, got=%v", s1, want1, got1)
	}
}

func TestArrayChunk(t *testing.T) {
	d := []M{
		{1, 2},
		{2, 3},
		{3, 4},
		{4, 5},
		{5, 6},
	}
	got := ArrayChunk(d, 2)
	want := [][]M{
		{{1, 2}, {2, 3}},
		{{3, 4}, {4, 5}},
		{{5, 6}},
	}
	if !reflect.DeepEqual(want, got) {
		t.Errorf("item=%v, want=%v, got=%v", d, want, got)
	}
}

func TestArrayDiff(t *testing.T) {
	d := []int{1, 2, 3, 4, 5}
	s1 := []int{2, 3}
	s2 := []int{5}
	want := []int{1, 4}
	got := ArrayDiff(d, s1, s2)
	if !reflect.DeepEqual(want, got) {
		t.Errorf("item=%v, want=%v, got=%v", d, want, got)
	}
}

func TestArrayIntersect(t *testing.T) {
	s1 := []int{1, 2, 3}
	s2 := []int{2}
	want := s2
	got := ArrayIntersect(s1, s2)
	if !reflect.DeepEqual(want, got) {
		t.Errorf("item=%v, want=%v, got=%v", s1, want, got)
	}
}

func TestArrayMerge(t *testing.T) {
	s1 := []int{1, 2, 3}
	s2 := []int{2}
	want := []int{1, 2, 3, 2}
	got := ArrayMerge(s1, s2)
	if !reflect.DeepEqual(want, got) {
		t.Errorf("item=%v, want=%v, got=%v", s1, want, got)
	}
}

func TestArraySearch(t *testing.T) {
	s1 := []int{1, 2, 3}
	n := 3
	want := 2
	got := ArraySearch[int](n, s1)
	if !reflect.DeepEqual(want, got) {
		t.Errorf("item=%v, want=%v, got=%v", s1, want, got)
	}
}

func TestArrayProduct(t *testing.T) {
	s1 := []int{1, 2, 3}
	want := 6
	got := ArrayProduct(s1)
	if !reflect.DeepEqual(want, got) {
		t.Errorf("item=%v, want=%v, got=%v", s1, want, got)
	}
}

func TestArrayWalk(t *testing.T) {
	s1 := []int{1, 2, 3}
	ArrayWalk(s1, func(item, index int) bool {
		if item == 2 {
			s1[index] = item * 10
			return false
		}
		return true
	})
	want := []int{1, 20, 3}
	if !reflect.DeepEqual(want, s1) {
		t.Errorf("item=%v, want=%v, got=%v", s1, want, s1)
	}
}

func TestJoin(t *testing.T) {
	s1 := []int{1, 2, 3}
	got := Join(s1, "|")
	want := "1|2|3"
	if !reflect.DeepEqual(want, got) {
		t.Errorf("item=%v, want=%v, got=%v", s1, want, got)
	}

	s2 := []string{"a", "c", "d"}
	got2 := Join(s2, ",")
	want2 := "a,c,d"
	if !reflect.DeepEqual(want2, got2) {
		t.Errorf("item=%v, want=%v, got=%v", s2, want2, got2)
	}
}

func TestArrayReverse(t *testing.T) {
	s := []int{1, 2, 3}
	want := []int{3, 2, 1}
	ArrayReverse(s)
	if !reflect.DeepEqual(s, want) {
		t.Errorf("item=%v, want=%v, got=%v", s, want, s)
	}

	s1 := []string{"1", "s", "e", "2"}
	want1 := []string{"2", "e", "s", "1"}
	ArrayReverse(s1)

	if !reflect.DeepEqual(s1, want1) {
		t.Errorf("item=%v, want=%v, got=%v", s1, want1, s1)
	}
}
