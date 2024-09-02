package practice

import (
    "testing"
    "fmt"
)

func TestMergeSort(t *testing.T) {
    arr := []int{1, 3, 4, 2, 5}
    MergeSort(arr)
    fmt.Println(arr)
}

func TestSmallSum(t *testing.T) {
    arr := []int{1, 3, 4, 2, 5}
    fmt.Println(SmallSum(arr))
}

func TestReversePair(t *testing.T) {
    arr := []int{2, 4, 1, 3, 5}
    fmt.Println(ReversePair(arr))
}

func TestBubbleSort(t *testing.T) {
    arr := []int{1, 3, 4, 2, 5}
    BubbleSort(arr)
    fmt.Println(arr)
}

func TestSelectSort(t *testing.T) {
    arr := []int{1, 3, 4, 2, 5}
    SelectSort(arr)
    fmt.Println(arr)
}

func TestInsertSort(t *testing.T) {
    arr := []int{1, 3, 4, 2, 5}
    InsertSort(arr)
    fmt.Println(arr)
}

func TestInsertSort1(t *testing.T) {
    arr := []int{1, 3, 4, 2, 5}
    InsertSort1(arr)
    fmt.Println(arr)
}

func TestHeapSort(t *testing.T) {
    arr := []int{1, 3, 4, 2, 5}
    // HeapSort(arr)
    // HeapSort1(arr)
    HeapSort2(arr)
    fmt.Println(arr)
}

func TestQuickSort(t *testing.T) {
    arr := []int{1, 3, 4, 2, 5}
    QuickSort(arr)
    fmt.Println(arr)
}
