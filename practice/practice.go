package practice

func ReversePair(nums []int) int {
    if len(nums) == 0 {
        return 0
    }
    return merge(nums, 0, len(nums)-1)
}

func merge(nums []int, start, end int) int {
    if start >= end {
        return 0
    }
    mid := start + (end-start)/2
    cnt := merge(nums, start, mid) + merge(nums, mid+1, end)
    tmp := []int{}
    i, j := start, mid+1
    for i <= mid && j <= end {
        if nums[i] <= nums[j] {
            tmp = append(tmp, nums[i])
            cnt += j - mid - 1
            i++
        } else {
            tmp = append(tmp, nums[j])
            j++
        }
    }
    for ; i <= mid; i++ {
        tmp = append(tmp, nums[i])
        cnt += end - mid
    }
    for ; j <= end; j++ {
        tmp = append(tmp, nums[j])
    }
    for i := start; i <= end; i++ {
        nums[i] = tmp[i-start]
    }
    return cnt
}
