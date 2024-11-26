package sort

// BubbleSort 冒泡排序
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
