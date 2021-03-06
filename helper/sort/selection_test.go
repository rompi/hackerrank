package sort_test

import (
	"fmt"
	"github.com/rompi/hackerrank/helper/sort"
	"testing"
)

func TestSelection(t *testing.T) {
	assertValidation := func(t *testing.T, got []int32, expected []int32) {
		t.Helper()
		if !isEqual(got, expected) {
			t.Errorf("expected: %v, got: %v", expected, got)
		}
	}
	t.Run("ok", func(t *testing.T) {
		initialData := []int32{1,5,2,9,4,1,5,6}
		got := sort.Selection(initialData)
		expected := []int32{1,1,2,4,5,5,6,9}
		assertValidation(t, got, expected)
	})
}

func BenchmarkSelection(b *testing.B) {
	arrs := benchmarkData()
	for i := 0; i < b.N; i++ {
		sort.Selection(arrs)
	}
}

func ExampleSelection() {
	arrs := exampleData()
	result := sort.Selection(arrs)
	fmt.Println(result)
	// Output: [1 2 3 4]
}