package slice_test

import "testing"

func TestSliceInit(t *testing.T) {
	s1 := []int{1, 2, 3, 4}
	t.Log(len(s1), cap(s1))
}

func TestSliceGrowing(t *testing.T) {
	s := []int{}
	for i := 0; i < 100; i++ {
		s = append(s, i)
		t.Log(len(s), cap(s))
	}
}
