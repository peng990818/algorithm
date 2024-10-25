package main

func main() {

}

func MaxNum(arr []int) int {
    if len(arr) == 0 {
        return -1
    }
    i, j := 0, len(arr)-1
    for i <= j {
        if arr[i] > arr[j] {
            j--
        } else if arr[i] < arr[j] {
            i++
        } else {
            i++
            j--
        }
    }
    return arr[i]
}
