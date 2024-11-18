package algorithm

// HeapSort 堆排序 从小到大
func HeapSort(arr []int) {
    // 构建堆
    n := len(arr)
    if n <= 1 {
        return
    }
    for i := n/2 - 1; i >= 0; i-- {
        maxHeapify(arr, i, n)
    }
    // 不停移除堆顶，进行堆调整
    for size := n - 1; size > 0; size-- {
        arr[0], arr[size] = arr[size], arr[0]
        maxHeapify(arr, 0, size)
    }
}

func maxHeapify(arr []int, index int, size int) {
    left := 2*index + 1
    for left < size {
        largest := index
        if arr[left] > arr[largest] {
            largest = left
        }
        right := left + 1
        if right < size && arr[right] > arr[largest] {
            largest = right
        }
        if largest == index {
            break
        }
        arr[largest], arr[index] = arr[index], arr[largest]
        index = largest
        left = 2*index + 1
    }
}

func minHeapify(arr []int, index int, size int) {
    left := 2*index + 1
    for left < size {
        smallest := index
        if arr[left] < arr[smallest] {
            smallest = left
        }
        right := left + 1
        if right < size && arr[right] < arr[smallest] {
            smallest = right
        }
        if smallest == index {
            break
        }
        arr[smallest], arr[index] = arr[index], arr[smallest]
        index = smallest
        left = 2*index + 1
    }
}

// HeapSortReverse 堆排序 从大到小
func HeapSortReverse(arr []int) {
    n := len(arr)
    if n <= 1 {
        return
    }
    for i := n/2 - 1; i >= 0; i-- {
        minHeapify(arr, i, n)
    }
    for size := n - 1; size > 0; size-- {
        arr[size], arr[0] = arr[0], arr[size]
        minHeapify(arr, 0, size)
    }
}
