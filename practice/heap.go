package practice

func heapInsert(arr []int, index int) {
    if index >= len(arr) {
        return
    }
    // 找到堆顶，如果插入元素大于堆顶，则交换，继续向上调整
    for arr[index] > arr[(index-1)/2] {
        arr[index], arr[(index-1)/2] = arr[(index-1)/2], arr[index]
        index = (index - 1) / 2
    }
}

func maxHeapify(arr []int, index int, heapSize int) {
    // 向下调整
    left := 2*index + 1
    for left < heapSize {
        largest := left
        if left+1 < heapSize && arr[left+1] > arr[left] {
            largest = left + 1
        }
        if arr[largest] > arr[index] {
            arr[largest], arr[index] = arr[index], arr[largest]
            index = largest
            left = 2*index + 1
        } else {
            break
        }
    }
}

func HeapSort(arr []int) {
    if len(arr) <= 1 {
        return
    }
    // 基于数组构建大根堆
    for i := range arr {
        heapInsert(arr, i)
    }
    // 堆顶作为最大值与最后一个元素交换，然后进行堆调整
    size := len(arr) - 1
    arr[0], arr[size] = arr[size], arr[0]
    for size > 0 {
        maxHeapify(arr, 0, size)
        size--
        arr[0], arr[size] = arr[size], arr[0]
    }
}

func HeapSort1(arr []int) {
    if len(arr) <= 1 {
        return
    }
    for i := len(arr)/2 - 1; i >= 0; i-- {
        maxHeapify(arr, i, len(arr))
    }
    for size := len(arr) - 1; size > 0; size-- {
        arr[0], arr[size] = arr[size], arr[0]
        maxHeapify(arr, 0, size)
    }
}

func minHeapify(arr []int, index int, heapSize int) {
    left := 2*index + 1
    for left < heapSize {
        minimum := left
        if left+1 < heapSize && arr[left] > arr[left+1] {
            minimum = left + 1
        }
        if arr[minimum] < arr[index] {
            arr[minimum], arr[index] = arr[index], arr[minimum]
            index = minimum
            left = 2*index + 1
        } else {
            break
        }
    }
}

func HeapSort2(arr []int) {
    if len(arr) <= 1 {
        return
    }
    for i := len(arr)/2 - 1; i >= 0; i-- {
        minHeapify(arr, i, len(arr))
    }
    for size := len(arr) - 1; size > 0; size-- {
        arr[0], arr[size] = arr[size], arr[0]
        minHeapify(arr, 0, size)
    }
}
