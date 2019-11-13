package sort

import (
	"math/rand"
	"sort"
	"testing"
	"time"
)

var max = 100000
var data = make([]int, max)

func initData() {
	rand.Seed(time.Now().Unix())
	for i := 0; i < max; i++ {
		var tmp = rand.Intn(10000)
		data[i] = tmp
	}
}

func TestBubbleSort(t *testing.T) {
	var a = []int{3, 5, 10, 1, 5, 4, 6, 2, 9, 7, 8}
	bubbleSort(a)
	t.Log("bubbleSort: ", a)
}

func TestInsertionSort(t *testing.T) {
	var a = []int{3, 5, 10, 1, 5, 4, 6, 2, 9, 7, 8}
	insertionSort(a)
	t.Log("insertionSort: ", a)
}

func TestSelectionSort(t *testing.T) {
	var a = []int{3, 5, 10, 1, 5, 4, 6, 2, 9, 7, 8}
	selectionSort(a)
	t.Log("selectionSort: ", a)
}

func TestMergeSort(t *testing.T) {
	var a = []int{3, 5, 10, 1, 5, 4, 6, 2, 9, 7, 8}
	mergeSort(a)
	t.Log("mergeSort: ", a)
}

func TestQuickSort(t *testing.T) {
	var a = []int{3, 5, 10, 1, 5, 4, 6, 2, 9, 7, 8}
	quickSort(a)
	t.Log("quickSort: ", a)
}

func TestCountingSort(t *testing.T) {
	var a = []int{3, 5, 10, 1, 5, 4, 6, 2, 9, 7, 8}
	countingSort(a)
	t.Log("countingSort: ", a)
}

//func BenchmarkBubbleSort(b *testing.B) {
//	initData()
//	b.ResetTimer()
//	for i := 0; i < b.N; i++ {
//		bubbleSort(data)
//	}
//}
//
//func BenchmarkInsertionSort(b *testing.B) {
//	initData()
//	b.ResetTimer()
//	for i := 0; i < b.N; i++ {
//		insertionSort(data)
//	}
//}
//
//func BenchmarkSelectionSort(b *testing.B) {
//	initData()
//	b.ResetTimer()
//	for i := 0; i < b.N; i++ {
//		selectionSort(data)
//	}
//}

func BenchmarkMergeSort(b *testing.B) {
	initData()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		mergeSort(data)
	}
}

func BenchmarkQuickSort(b *testing.B) {
	initData()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		quickSort(data)
	}
}

func BenchmarkBuiltinSort(b *testing.B) {
	initData()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		sort.Ints(data)
	}
}
