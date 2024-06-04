package sort

// 插入排序
func InsertSort(nums []int) []int {
    if len(nums) == 0 {
        return nil
    }
    for i := 1; i < len(nums); i++ {
        tmp := nums[i]
        j := i - 1
        for ; j >= 0 && nums[j] > tmp; j-- {
            nums[j+1] = nums[j]
        }
        nums[j+1] = tmp
    }
    return nums
}

func InsertSortByDes(nums []int) []int {
    if len(nums) == 0 {
        return nil
    }
    for i := 1; i < len(nums); i++ {
        tmp := nums[i]
        j := i - 1
        for ; j >= 0 && nums[j] < tmp; j-- {
            nums[j+1] = nums[j]
        }
        nums[j+1] = tmp
    }
    return nums
}
