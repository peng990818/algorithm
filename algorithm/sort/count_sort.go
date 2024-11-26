package sort

// CountSort 计数排序，统计数量进行排序
func CountSort(arr []int) []int {
    if len(arr) <= 1 {
        return arr
    }

    // 找到数组中的最大值和最小值
    minValue, maxValue := arr[0], arr[0]
    for _, v := range arr {
        if v < minValue {
            minValue = v
        }
        if v > maxValue {
            maxValue = v
        }
    }

    // 统计频次
    count := make([]int, maxValue-minValue+1)
    for _, v := range arr {
        count[v-minValue]++
    }

    // 对计数数组进行累加
    for i := 1; i < len(count); i++ {
        count[i] += count[i-1]
    }

    // 从后向前遍历数组，写入结果数组，保证排序稳定性
    result := make([]int, len(arr))
    for i := len(arr) - 1; i >= 0; i-- {
        val := arr[i]
        result[count[val-minValue]-1] = val
        count[val-minValue]--
    }
    return result
}
