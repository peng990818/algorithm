package sort

import "math/rand/v2"

// QuickSort 经典快排
func QuickSort(arr []int) {
    if len(arr) <= 1 {
        return
    }
    quickSort(arr, 0, len(arr)-1)
}

func quickSort(arr []int, l, r int) {
    if l < r {
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
            arr[l], arr[less] = arr[less], arr[l]
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

// RandomQuickSort 随机快排
func RandomQuickSort(arr []int) {
    if len(arr) <= 1 {
        return
    }
    randomQuickSort(arr, 0, len(arr)-1)
}

func randomQuickSort(arr []int, l, r int) {
    if l < r {
        index := rand.IntN(r-l+1) + l
        arr[index], arr[r] = arr[r], arr[index]
        p := partition(arr, l, r)
        randomQuickSort(arr, l, p[0]-1)
        randomQuickSort(arr, p[1]+1, r)
    }
}
