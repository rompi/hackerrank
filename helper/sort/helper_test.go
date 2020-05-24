package sort_test

import "reflect"

func isEqual(a, b []int32) bool {
	return reflect.DeepEqual(a, b)
}

func benchmarkData() []int32 {
	return []int32{5,91,0,2,3,51,23,24,22,5,1,5,1,231,43,531,535,1,343,12,1,3,53,515,426,2452,524,64,2,123}
}

func exampleData() []int32 {
	return []int32{1, 2, 4, 3}
}