package sort

func Selection(arr []int32) []int32 {
	n := len(arr)
	for i := 0; i < n; i++ {
		min_idx := i
		for j := i; j < n; j++ {
			if arr[min_idx] > arr[j] {
				min_idx = j
			}
		}
		arr[min_idx], arr[i] = swap(arr[min_idx], arr[i])
	}
	return arr
}
