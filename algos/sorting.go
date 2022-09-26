package algos

// The selection sort algorithm sorts an array by repeatedly finding the
// smallest element from the unsorted part and swapping the elements.
//
// Complexity O(n^2) as we have 2 nested loops.
// one loop to get elements out of the array.
// one loop (nested) to compare first element w/ every other element in the array.
//
// Space O(1) as we use only 1 variable of extra temp memory the rest we just swap.
func SelectionSort(arr []int) {
	for i := 0; i < len(arr)-1; i++ {
		var index = i
		for j := i + 1; j < len(arr); j++ {
			if arr[j] < arr[index] {
				index = j
			}

			if index != i {
				arr[index], arr[i] = arr[i], arr[index] // swap
			}
		}
	}
}
