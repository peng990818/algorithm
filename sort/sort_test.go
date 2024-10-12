package sort

import (
    "testing"
    "fmt"
)

func TestBubbleSort(t *testing.T) {
    arr := []int{1, 13, 6, 9, 2}
    BubbleSort(arr)
    fmt.Println(arr)
}

func TestSelectSort(t *testing.T) {
    arr := []int{1, 13, 6, 9, 2}
    SelectSort(arr)
    fmt.Println(arr)
}

func TestInsertSort(t *testing.T) {
    arr := []int{1, 13, 6, 9, 2}
    // InsertSort(arr)
    InsertSort1(arr)
    fmt.Println(arr)
}

func TestMergeSort(t *testing.T) {
    arr := []int{1, 13, 6, 9, 2}
    MergeSort(arr)
    fmt.Println(arr)
}

func TestQuickSort(t *testing.T) {
    arr := []int{1, 13, 6, 9, 2}
    QuickSort(arr)
    fmt.Println(arr)
}

func TestHeapSort(t *testing.T) {
    arr := []int{1, 13, 6, 9, 2}
    // HeapSort(arr)
    HeapSort1(arr)
    fmt.Println(arr)
}

func TestBucketSort(t *testing.T) {
    arr := []int{1, 13, 6, 9, 2}
    res := BucketSort(arr, 2)
    fmt.Println(res)
}

func TestCountingSort(t *testing.T) {
    arr := []int{1, 13, 6, 9, 2}
    res := CountingSort(arr)
    fmt.Println(res)
}

func TestRadixSort(t *testing.T) {
    arr := []int{1, 13, 6, 9, 2}
    RadixSort(arr)
    fmt.Println(arr)
}

func TestShellSort(t *testing.T) {
    arr := []int{1, 13, 6, 9, 2}
    ShellSort(arr, 2)
    fmt.Println(arr)
}
