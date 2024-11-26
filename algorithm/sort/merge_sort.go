package sort

// MergeSort 归并排序
func MergeSort(arr []int) {
    if len(arr) <= 1 {
        return
    }
    mergeSort(arr, 0, len(arr)-1)
}

func mergeSort(arr []int, l, r int) {
    if l == r {
        return
    }
    mid := l + ((r - l) >> 1)
    mergeSort(arr, l, mid)
    mergeSort(arr, mid+1, r)
    merge(arr, l, mid, r)
}

func merge(arr []int, l, mid, r int) {
    help := make([]int, r-l+1)
    i, p1, p2 := 0, l, mid+1
    for p1 <= mid && p2 <= r {
        if arr[p1] <= arr[p2] {
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
        p1++
        i++
    }
    for p2 <= r {
        help[i] = arr[p2]
        p2++
        i++
    }
    for i := 0; i < len(help); i++ {
        arr[l+i] = help[i]
    }
}
