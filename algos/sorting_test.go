package algos

import "testing"

func TestSelection(t *testing.T) {
	arr := []int{9, 8, 7, 6, 5, 4, 3, 2, 1}
	test1 := arr[3]
	t.Log(arr)
	SelectionSort(arr)
	t.Log(arr)
	if test1 == arr[3] {
		t.Fail()
	}
}

func BenchmarkSelection(b *testing.B) {
	arr := []int{9, 8, 7, 6, 5, 4, 3, 2, 1}
	for i := 0; i < b.N; i++ {
		arr[5] = i
		SelectionSort(arr)
	}
}
