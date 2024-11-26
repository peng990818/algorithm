package algorithm

import (
    "testing"
    "fmt"
    "algorithm/algorithm/sort"
    "algorithm/algorithm/question"
)

func TestGenerateRandomArray(t *testing.T) {
    for i := 0; i < 10; i++ {
        fmt.Println(GenerateRandomArray(10))
    }
}

func TestBubbleSort(t *testing.T) {
    arr := GenerateRandomArray(10)
    fmt.Println(arr)
    sort.BubbleSort(arr)
    fmt.Println(arr)
}

func TestSelectSort(t *testing.T) {
    arr := GenerateRandomArray(10)
    fmt.Println(arr)
    sort.SelectSort(arr)
    fmt.Println(arr)
}

func TestInsertSort(t *testing.T) {
    arr := GenerateRandomArray(10)
    fmt.Println(arr)
    sort.InsertSortBySwap(arr)
    fmt.Println(arr)
    arr = GenerateRandomArray(10)
    fmt.Println(arr)
    sort.InsertSortByMove(arr)
    fmt.Println(arr)
}

func TestHeapSort(t *testing.T) {
    arr := GenerateRandomArray(10)
    fmt.Println(arr)
    sort.HeapSort(arr)
    fmt.Println(arr)
    arr = GenerateRandomArray(10)
    fmt.Println(arr)
    sort.HeapSortReverse(arr)
    fmt.Println(arr)
}

func TestMergeSort(t *testing.T) {
    arr := GenerateRandomArray(10)
    fmt.Println(arr)
    sort.MergeSort(arr)
    fmt.Println(arr)
}

func TestQuickSort(t *testing.T) {
    arr := GenerateRandomArray(10)
    fmt.Println(arr)
    sort.QuickSort(arr)
    fmt.Println(arr)
    arr = GenerateRandomArray(10)
    fmt.Println(arr)
    sort.RandomQuickSort(arr)
    fmt.Println(arr)
}

func TestBucketSort(t *testing.T) {
    arr := GenerateRandomArray(10)
    fmt.Println(arr)
    arr = sort.BucketSort(arr)
    fmt.Println(arr)
}

func TestCountSort(t *testing.T) {
    arr := GenerateRandomArray(10)
    fmt.Println(arr)
    arr = sort.CountSort(arr)
    fmt.Println(arr)
}

func TestRadixSort(t *testing.T) {
    arr := GenerateRandomArray(10)
    fmt.Println(arr)
    sort.RadixSort(arr)
    fmt.Println(arr)
}

func TestShellSort(t *testing.T) {
    arr := GenerateRandomArray(10)
    fmt.Println(arr)
    sort.ShellSort(arr, 2)
    fmt.Println(arr)
}

func TestSmallSum(t *testing.T) {
    arr := []int{1, 3, 4, 2, 5}
    fmt.Println(question.SmallSum(arr))
}

func TestReversePair(t *testing.T) {
    arr := []int{1, 3, 4, 2, 5}
    fmt.Println(question.ReversePair(arr))
}
