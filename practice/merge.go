package practice

// 分治模版
// 归并排序
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
    mid := l + (r-l)>>1
    mergeSort(arr, l, mid)
    mergeSort(arr, mid+1, r)
    merge(arr, l, mid, r)
}

func merge(arr []int, l, mid, r int) {
    help := make([]int, r-l+1)
    i, j, k := l, mid+1, 0
    for i <= mid && j <= r {
        if arr[i] <= arr[j] {
            help[k] = arr[i]
            i++
        } else {
            help[k] = arr[j]
            j++
        }
        k++
    }
    for i <= mid {
        help[k] = arr[i]
        i++
        k++
    }
    for j <= r {
        help[k] = arr[j]
        j++
        k++
    }
    for i := 0; i < len(help); i++ {
        arr[l+i] = help[i]
    }
}

// 计算小和
func SmallSum(arr []int) int {
    if len(arr) <= 1 {
        return 0
    }
    return mergeSort1(arr, 0, len(arr)-1)
}

func mergeSort1(arr []int, l, r int) int {
    if l == r {
        return 0
    }
    mid := l + (r-l)>>1
    return mergeSort1(arr, l, mid) + mergeSort1(arr, mid+1, r) +
        merge1(arr, l, mid, r)
}

func merge1(arr []int, l, mid, r int) int {
    help := make([]int, r-l+1)
    i, j, k, res := l, mid+1, 0, 0
    for i <= mid && j <= r {
        if arr[i] <= arr[j] {
            help[k] = arr[i]
            res += (r - j + 1) * arr[i]
            i++
        } else {
            help[k] = arr[j]
            j++
        }
        k++
    }
    for i <= mid {
        help[k] = arr[i]
        i++
        k++
    }
    for j <= r {
        help[k] = arr[j]
        j++
        k++
    }
    for i := 0; i < len(help); i++ {
        arr[l+i] = help[i]
    }
    return res
}

// 逆序对数量统计
func ReversePair(arr []int) int {
    if len(arr) <= 1 {
        return 0
    }
    return mergeSort2(arr, 0, len(arr)-1)
}

func mergeSort2(arr []int, l, r int) int {
    if l == r {
        return 0
    }
    mid := l + (r-l)>>1
    return mergeSort2(arr, l, mid) + mergeSort2(arr, mid+1, r) +
        merge2(arr, l, mid, r)
}

func merge2(arr []int, l, mid, r int) int {
    help := make([]int, r-l+1)
    i, j, k, res := l, mid+1, 0, 0
    for i <= mid && j <= r {
        if arr[i] <= arr[j] {
            help[k] = arr[i]
            i++
        } else {
            res += mid - i + 1
            help[k] = arr[j]
            j++
        }
        k++
    }
    for i <= mid {
        help[k] = arr[i]
        i++
        k++
    }

    for j <= r {
        help[k] = arr[j]
        j++
        k++
    }

    for i := 0; i < len(help); i++ {
        arr[l+i] = help[i]
    }
    return res
}
