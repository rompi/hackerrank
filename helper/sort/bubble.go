package sort

func Bubble(ar []int32) []int32 {
	n := len(ar)
	for i := 0; i < n; i++ {
		for j := i; j < n; j++ {
			if(ar[j] < ar[i]) {
				ar[j], ar[i] = swap(ar[j],ar[i])
			}
		}
	}
	return ar
}