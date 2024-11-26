package sort

// ShellSort 希尔排序
// gap表示每隔几步
func ShellSort(src []int, gap int) {
    l := len(src)
    for step := l / gap; step >= 1; step /= gap {
        for i := 0; i < step; i++ {
            InsertionSortByStep(src, step)
        }
    }
}

func InsertionSortByStep(src []int, step int) {
    for j := step; j < len(src); j += step {
        key := src[j]
        i := j - step
        for i >= 0 && src[i] > key {
            src[i+step] = src[i]
            i -= step
        }
        src[i+step] = key
    }
}
