package sort

import (
    "math/rand/v2"
    "sort"
)

// BubbleSort 1、冒泡排序
func BubbleSort(arr []int) {
    if len(arr) <= 1 {
        return
    }
    for i := 0; i < len(arr)-1; i++ {
        for j := 0; j < len(arr)-1-i; j++ {
            if arr[j] > arr[j+1] {
                arr[j], arr[j+1] = arr[j+1], arr[j]
            }
        }
    }
}

// SelectSort 2、选择排序
func SelectSort(arr []int) {
    if len(arr) <= 1 {
        return
    }
    for i := 0; i < len(arr)-1; i++ {
        min := i
        for j := i + 1; j < len(arr); j++ {
            if arr[min] > arr[j] {
                min = j
            }
        }
        if min == i {
            continue
        }
        arr[min], arr[i] = arr[i], arr[min]
    }
}

// InsertSort 3、插入排序 不停交换版本
func InsertSort(arr []int) {
    if len(arr) <= 1 {
        return
    }
    for i := 1; i < len(arr); i++ {
        for j := i - 1; j >= 0 && arr[j] > arr[j+1]; j-- {
            arr[j], arr[j+1] = arr[j+1], arr[j]
        }
    }
}

// InsertSort1 插入排序 不停移动版本
func InsertSort1(arr []int) {
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

// MergeSort 4、归并排序
func MergeSort(arr []int) {
    if len(arr) <= 1 {
        return
    }
    mergeSort(arr, 0, len(arr)-1)
}

func mergeSort(arr []int, start, end int) {
    if start == end {
        return
    }
    mid := start + ((end - start) >> 1)
    mergeSort(arr, start, mid)
    mergeSort(arr, mid+1, end)
    merge(arr, start, mid, end)
}

func merge(arr []int, start, mid, end int) {
    help := make([]int, end-start+1)
    i, p1, p2 := 0, start, mid+1
    for p1 <= mid && p2 <= end {
        if arr[p1] < arr[p2] {
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
        i++
        p1++
    }
    for p2 <= end {
        help[i] = arr[p2]
        i++
        p2++
    }
    for i := 0; i < len(help); i++ {
        arr[start+i] = help[i]
    }
}

// QuickSort 5、快排
func QuickSort(arr []int) {
    if len(arr) <= 1 {
        return
    }
    quickSort(arr, 0, len(arr)-1)
}

func quickSort(arr []int, l, r int) {
    if l < r {
        // 经典快排 r固定为数组中最后一个数
        // 随机快排 将数组中任意一个数交换到最后一个
        index := rand.IntN(r-l+1) + l
        arr[index], arr[r] = arr[r], arr[index]
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
            arr[less], arr[l] = arr[l], arr[less]
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

// HeapSort 6、堆排序
func HeapSort(arr []int) {
    if len(arr) <= 1 {
        return
    }
    // 1、将数组调整为一个堆
    for i := 0; i < len(arr); i++ {
        heapInsert(arr, i)
    }
    // 2、将堆顶调整到数组末尾，减少堆的大小
    size := len(arr)
    for size > 0 {
        size--
        arr[0], arr[size] = arr[size], arr[0]
        heapify(arr, 0, size)
    }
}

func HeapSort1(arr []int) {
    if len(arr) <= 1 {
        return
    }
    for i := len(arr)/2 - 1; i >= 0; i-- {
        heapify(arr, i, len(arr))
    }
    for i := len(arr) - 1; i >= 0; i-- {
        arr[0], arr[i] = arr[i], arr[0]
        heapify(arr, 0, i)
    }
}

func heapInsert(arr []int, index int) {
    for arr[index] > arr[(index-1)/2] {
        arr[index], arr[(index-1)/2] = arr[(index-1)/2], arr[index]
        index = (index - 1) / 2
    }
}

func heapify(arr []int, index, size int) {
    left := 2*index + 1
    for left < size {
        largest := left
        if left+1 < size && arr[left+1] > arr[left] {
            largest = left + 1
        }
        if arr[largest] < arr[index] {
            largest = index
        }
        if largest == index {
            break
        }
        arr[largest], arr[index] = arr[index], arr[largest]
        index = largest
        left = 2*index + 1
    }
}

// BucketSort 7、桶排序
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

// CountingSort 8、计数排序，统计数量进行排序
func CountingSort(arr []int) []int {
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

// RadixSort 9、基数排序，按照十进制的每一位进行排序
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

// ShellSort 10、希尔排序
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
