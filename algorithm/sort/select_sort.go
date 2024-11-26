package sort

func SelectSort(arr []int) {
    if len(arr) <= 1 {
        return
    }
    for i := 0; i < len(arr)-1; i++ {
        minIndex := i
        for j := i + 1; j < len(arr); j++ {
            if arr[minIndex] > arr[j] {
                minIndex = j
            }
        }
        arr[minIndex], arr[i] = arr[i], arr[minIndex]
    }
}
