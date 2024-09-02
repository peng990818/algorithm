package practice

import "math/rand/v2"

// BubbleSort 冒泡排序
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

// SelectSort 选择排序
func SelectSort(arr []int) {
    if len(arr) <= 1 {
        return
    }

    for i := 0; i < len(arr)-1; i++ {
        minIndex := i
        for j := i + 1; j < len(arr); j++ {
            if arr[minIndex] > arr[j] {
                minIndex = j
            }
        }
        if i != minIndex {
            arr[minIndex], arr[i] = arr[i], arr[minIndex]
        }
    }
}

// 交换版本的插入排序
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

func QuickSort(arr []int) {
    if len(arr) <= 1 {
        return
    }
    // 经典快排
    // classicQuickSort(arr, 0, len(arr)-1)
    // 优化等于区快排
    // quickSort(arr, 0, len(arr)-1)
    // 最终优化，随机快排
    randQuickSort(arr, 0, len(arr)-1)
}

// 经典快排
func classicQuickSort(arr []int, l, r int) {
    if l < r {
        p := classicPartition(arr, l, r)
        classicQuickSort(arr, l, p-1)
        classicQuickSort(arr, p+1, r)
    }
}

func classicPartition(arr []int, l, r int) int {
    less := l - 1
    tmp := arr[r]
    for j := l; j < r; j++ {
        if arr[j] <= tmp {
            less++
            arr[less], arr[j] = arr[j], arr[less]
        }
    }
    arr[less+1], arr[r] = arr[r], arr[less+1]
    return less + 1
}

// 优化等于区快排
func quickSort(arr []int, l, r int) {
    if l < r {
        pa := partition(arr, l, r)
        quickSort(arr, l, pa[0]-1)
        quickSort(arr, pa[1]+1, r)
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

// 随机+优化等于区快排
func randQuickSort(arr []int, l, r int) {
    if l < r {
        index := rand.IntN(r-l+1) + l
        arr[index], arr[r] = arr[r], arr[index]
        pa := partition(arr, l, r)
        randQuickSort(arr, l, pa[0]-1)
        randQuickSort(arr, pa[1]+1, r)
    }
}
