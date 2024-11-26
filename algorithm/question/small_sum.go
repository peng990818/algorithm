package question

// SmallSum
/* 小和问题
在一个数组中，每一个数左边比当前数小的数累加起来，叫做这个数组的小和。求一个数组
的小和。
例子：
[1,3,4,2,5]
1左边比1小的数，没有；
3左边比3小的数，1；
4左边比4小的数，1、3；
2左边比2小的数，1；
5左边比5小的数，1、3、4、2；
所以小和为1+1+3+1+1+3+4+2=16
*/
func SmallSum(arr []int) int {
    if len(arr) == 0 {
        return 0
    }
    return processSum(arr, 0, len(arr)-1)
}

func processSum(arr []int, l, r int) int {
    if l == r {
        return 0
    }
    mid := l + ((r - l) >> 1)
    return processSum(arr, l, mid) + processSum(arr, mid+1, r) + mergeSum(arr, l, mid, r)
}

func mergeSum(arr []int, l, mid, r int) int {
    help := make([]int, r-l+1)
    i, j, k := l, mid+1, 0
    res := 0
    for i <= mid && j <= r {
        if arr[i] < arr[j] {
            // j后面的数都大于arr[i]
            res += (r - j + 1) * arr[i]
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
    return res
}
