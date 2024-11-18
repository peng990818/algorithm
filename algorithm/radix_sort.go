package algorithm

// RadixSort 基数排序，按照十进制的每一位进行排序
func RadixSort(src []int) {
    if len(src) == 0 {
        return
    }
    // 找到最大值
    maxValue := src[0]
    for i := 1; i < len(src); i++ {
        if src[i] > maxValue {
            maxValue = src[i]
        }
    }
    // 找到最大位数
    c := 1
    for maxValue >= 10 {
        maxValue /= 10
        c++
    }
    d := 1
    for i := 0; i < c; i++ {
        InsertionSortByDigit(src, d)
        d *= 10
    }
}

func InsertionSortByDigit(src []int, d int) {
    if len(src) == 0 {
        return
    }
    for j := 1; j < len(src); j++ {
        key := src[j]
        i := j - 1
        for i >= 0 && src[i]/d%10 > key/d%10 {
            src[i+1] = src[i]
            i--
        }
        src[i+1] = key
    }
}
