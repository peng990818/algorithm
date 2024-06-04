package sort

import (
    "testing"
    "fmt"
)

func TestInsertSort(t *testing.T) {
    input := []int{5, 4, 3, 2, 1, 6, 7}
    fmt.Println(InsertSort(input))
    fmt.Println(InsertSortByDes(input))
}
