package algorithm

import "sort"

// BucketSort 桶排序
func BucketSort(arr []int, bucketSize ...int) []int {
    if len(arr) <= 1 {
        return arr
    }

    // bucket大小默认值设置为10
    size := 10
    if len(bucketSize) > 0 && bucketSize[0] > 0 {
        size = bucketSize[0]
    }

    // 找到待排序数组的最大值和最小值
    minValue, maxValue := arr[0], arr[0]
    for i := 0; i < len(arr); i++ {
        if maxValue < arr[i] {
            maxValue = arr[i]
        }
        if minValue > arr[i] {
            minValue = arr[i]
        }
    }

    // 创建对应数量的桶
    bucketCount := (maxValue-minValue)/size + 1
    buckets := make([][]int, bucketCount)

    // 将数组中的值放入对应的桶中
    for _, v := range arr {
        index := (v - minValue) / size
        buckets[index] = append(buckets[index], v)
    }

    // 对每个桶中的数据进行排序，写入原数组
    result := make([]int, 0, len(arr))
    for _, bucket := range buckets {
        if len(bucket) > 0 {
            sort.Ints(bucket)
            result = append(result, bucket...)
        }
    }
    return result
}
