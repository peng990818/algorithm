package algorithm

import (
    "testing"
    "fmt"
)

func TestGenerateRandomArray(t *testing.T) {
    for i := 0; i < 10; i++ {
        fmt.Println(GenerateRandomArray())
    }
}

func TestBubbleSort(t *testing.T) {
    arr := GenerateRandomArray()
    fmt.Println(arr)
    BubbleSort(arr)
    fmt.Println(arr)
}

func TestSelectSort(t *testing.T) {
    arr := GenerateRandomArray()
    fmt.Println(arr)
    SelectSort(arr)
    fmt.Println(arr)
}

func TestInsertSort(t *testing.T) {
    arr := GenerateRandomArray()
    fmt.Println(arr)
    InsertSortBySwap(arr)
    fmt.Println(arr)
    arr = GenerateRandomArray()
    fmt.Println(arr)
    InsertSortByMove(arr)
    fmt.Println(arr)
}

func TestHeapSort(t *testing.T) {
    arr := GenerateRandomArray()
    fmt.Println(arr)
    HeapSort(arr)
    fmt.Println(arr)
    arr = GenerateRandomArray()
    fmt.Println(arr)
    HeapSortReverse(arr)
    fmt.Println(arr)
}

func TestMergeSort(t *testing.T) {
    arr := GenerateRandomArray()
    fmt.Println(arr)
    MergeSort(arr)
    fmt.Println(arr)
}

func TestQuickSort(t *testing.T) {
    arr := GenerateRandomArray()
    fmt.Println(arr)
    QuickSort(arr)
    fmt.Println(arr)
    arr = GenerateRandomArray()
    fmt.Println(arr)
    RandomQuickSort(arr)
    fmt.Println(arr)
}

func TestBucketSort(t *testing.T) {
    arr := GenerateRandomArray()
    fmt.Println(arr)
    arr = BucketSort(arr)
    fmt.Println(arr)
}

func TestCountSort(t *testing.T) {
    arr := GenerateRandomArray()
    fmt.Println(arr)
    arr = CountSort(arr)
    fmt.Println(arr)
}

func TestRadixSort(t *testing.T) {
    arr := GenerateRandomArray()
    fmt.Println(arr)
    RadixSort(arr)
    fmt.Println(arr)
}

func TestShellSort(t *testing.T) {
    arr := GenerateRandomArray()
    fmt.Println(arr)
    ShellSort(arr, 2)
    fmt.Println(arr)
}
