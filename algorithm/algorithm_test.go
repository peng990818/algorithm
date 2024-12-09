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

func sortArray(nums []int) []int {
    quickSort(nums, 0, len(nums)-1)
    return nums
}

func quickSort(nums []int, l, r int) {
    if l < r {
        p := partition(nums, l, r)
        quickSort(nums, l, p[0]-1)
        quickSort(nums, p[1]+1, r)
    }
}

func partition(nums []int, l, r int) []int {
    less, more := l-1, r
    for l < more {
        if nums[l] < nums[r] {
            less++
            nums[l], nums[less] = nums[less], nums[l]
            l++
        } else if nums[l] > nums[r] {
            more--
            nums[l], nums[more] = nums[more], nums[l]
        } else {
            l++
        }
    }
    nums[more], nums[r] = nums[r], nums[more]
    return []int{less + 1, more}
}

func TestName(t *testing.T) {
    arr := sortArray([]int{110, 100, 0})
    fmt.Println(arr)
}
