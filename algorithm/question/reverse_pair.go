package question

/*
逆序对问题
在一个数组中，左边的数如果比右边的数大，则这两个数构成一个逆序对，请打印所有逆序
对。这里仅统计数量，如需打印，在统计时打印即可
*/

func ReversePair(arr []int) int {
    if len(arr) == 0 {
        return 0
    }
    return processReversePair(arr, 0, len(arr)-1)
}

func processReversePair(arr []int, l, r int) int {
    if l == r {
        return 0
    }
    mid := l + ((r - l) >> 1)
    return processReversePair(arr, l, mid) + processReversePair(arr, mid+1, r) + mergeReversePair(arr, l, mid, r)
}

func mergeReversePair(arr []int, l, mid, r int) int {
    help := make([]int, r-l+1)
    i, j, k := l, mid+1, 0
    res := 0
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
