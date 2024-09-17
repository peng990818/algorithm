package sort

import (
    "math/rand/v2"
)

// BubbleSort 1、冒泡排序
func BubbleSort(arr []int) {
    if len(arr) <= 1 {
        return
    }
    for i := 0; i < len(arr)-1; i++ {
        for j := 0; j < len(arr)-1-i; j++ {
            if arr[j] > arr[j+1] {
                arr[j], arr[j+1] = arr[j+1], arr[j]
            }
        }
    }
}

// SelectSort 2、选择排序
func SelectSort(arr []int) {
    if len(arr) <= 1 {
        return
    }
    for i := 0; i < len(arr)-1; i++ {
        min := i
        for j := i + 1; j < len(arr); j++ {
            if arr[min] > arr[j] {
                min = j
            }
        }
        if min == i {
            continue
        }
        arr[min], arr[i] = arr[i], arr[min]
    }
}

// InsertSort 3、插入排序 不停交换版本
func InsertSort(arr []int) {
    if len(arr) <= 1 {
        return
    }
    for i := 1; i < len(arr); i++ {
        for j := i - 1; j >= 0 && arr[j] > arr[j+1]; j-- {
            arr[j], arr[j+1] = arr[j+1], arr[j]
        }
    }
}

// InsertSort1 插入排序 不停移动版本
func InsertSort1(arr []int) {
    if len(arr) <= 1 {
        return
    }
    for i := 1; i < len(arr); i++ {
        key := arr[i]
        j := i - 1
        for j >= 0 && arr[j] > key {
            arr[j+1] = arr[j]
            j--
        }
        arr[j+1] = key
    }
}

// MergeSort 4、归并排序
func MergeSort(arr []int) {
    if len(arr) <= 1 {
        return
    }
    mergeSort(arr, 0, len(arr)-1)
}

func mergeSort(arr []int, start, end int) {
    if start == end {
        return
    }
    mid := start + ((end - start) >> 1)
    mergeSort(arr, start, mid)
    mergeSort(arr, mid+1, end)
    merge(arr, start, mid, end)
}

func merge(arr []int, start, mid, end int) {
    help := make([]int, end-start+1)
    i, p1, p2 := 0, start, mid+1
    for p1 <= mid && p2 <= end {
        if arr[p1] < arr[p2] {
            help[i] = arr[p1]
            p1++
        } else {
            help[i] = arr[p2]
            p2++
        }
        i++
    }
    for p1 <= mid {
        help[i] = arr[p1]
        i++
        p1++
    }
    for p2 <= end {
        help[i] = arr[p2]
        i++
        p2++
    }
    for i := 0; i < len(help); i++ {
        arr[start+i] = help[i]
    }
}

// QuickSort 5、快排
func QuickSort(arr []int) {
    if len(arr) <= 1 {
        return
    }
    quickSort(arr, 0, len(arr)-1)
}

func quickSort(arr []int, l, r int) {
    if l < r {
        // 经典快排 r固定为数组中最后一个数
        // 随即快排 将数组中任意一个数交换到最后一个
        index := rand.IntN(r-l+1) + l
        arr[index], arr[r] = arr[r], arr[index]
        p := partition(arr, l, r)
        quickSort(arr, l, p[0]-1)
        quickSort(arr, p[1]+1, r)
    }
}

func partition(arr []int, l, r int) []int {
    less := l - 1
    more := r
    for l < more {
        if arr[l] < arr[r] {
            less++
            arr[less], arr[l] = arr[l], arr[less]
            l++
        } else if arr[l] > arr[r] {
            more--
            arr[more], arr[l] = arr[l], arr[more]
        } else {
            l++
        }
    }
    arr[more], arr[r] = arr[r], arr[more]
    return []int{less + 1, more}
}

// HeapSort 6、堆排序
func HeapSort(arr []int) {
    if len(arr) <= 1 {
        return
    }
    // 1、将数组调整为一个堆
    for i := 0; i < len(arr); i++ {
        heapInsert(arr, i)
    }
    // 2、将堆顶调整到数组末尾，减少堆的大小
    size := len(arr)
    for size > 0 {
        size--
        arr[0], arr[size] = arr[size], arr[0]
        heapify(arr, 0, size)
    }
}

func HeapSort1(arr []int) {
    if len(arr) <= 1 {
        return
    }
    for i := len(arr)/2 - 1; i >= 0; i-- {
        heapify(arr, i, len(arr))
    }
    for i := len(arr) - 1; i >= 0; i-- {
        arr[0], arr[i] = arr[i], arr[0]
        heapify(arr, 0, i)
    }
}

func heapInsert(arr []int, index int) {
    for arr[index] > arr[(index-1)/2] {
        arr[index], arr[(index-1)/2] = arr[(index-1)/2], arr[index]
        index = (index - 1) / 2
    }
}

func heapify(arr []int, index, size int) {
    left := 2*index + 1
    for left < size {
        largest := left
        if left+1 < size && arr[left+1] > arr[left] {
            largest = left + 1
        }
        if arr[largest] < arr[index] {
            largest = index
        }
        if largest == index {
            break
        }
        arr[largest], arr[index] = arr[index], arr[largest]
        index = largest
        left = 2*index + 1
    }
}
