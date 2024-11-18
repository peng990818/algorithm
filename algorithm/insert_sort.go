package algorithm

// InsertSortBySwap 基于不停交换的插入排序
func InsertSortBySwap(arr []int) {
    if len(arr) <= 1 {
        return
    }

    for i := 1; i < len(arr); i++ {
        for j := i - 1; j >= 0 && arr[j] > arr[j+1]; j-- {
            arr[j], arr[j+1] = arr[j+1], arr[j]
        }
    }
}

// InsertSortByMove 基于不停移动版本的插入排序
func InsertSortByMove(arr []int) {
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
