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

func TestSliceShareMemory(t *testing.T) {
	year := []string{"Jan", "Feb", "Mar", "Apr", "May", "Jun", "Jul", "Aug", "Sep",
		"Oct", "Nov", "Dec"}
	Q2 := year[3:6]
	t.Log(Q2, len(Q2), cap(Q2))
	summer := year[5:10]
	t.Log(summer, len(summer), cap(summer))

}
