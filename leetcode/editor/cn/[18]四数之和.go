//给你一个由 n 个整数组成的数组 nums ，和一个目标值 target 。请你找出并返回满足下述全部条件且不重复的四元组 [nums[a], nums[
//b], nums[c], nums[d]] （若两个四元组元素一一对应，则认为两个四元组重复）： 
//
// 
// 0 <= a, b, c, d < n 
// a、b、c 和 d 互不相同 
// nums[a] + nums[b] + nums[c] + nums[d] == target 
// 
//
// 你可以按 任意顺序 返回答案 。 
//
// 
//
// 示例 1： 
//
// 
//输入：nums = [1,0,-1,0,-2,2], target = 0
//输出：[[-2,-1,1,2],[-2,0,0,2],[-1,0,0,1]]
// 
//
// 示例 2： 
//
// 
//输入：nums = [2,2,2,2,2], target = 8
//输出：[[2,2,2,2]]
// 
//
// 
//
// 提示： 
//
// 
// 1 <= nums.length <= 200 
// -10⁹ <= nums[i] <= 10⁹ 
// -10⁹ <= target <= 10⁹ 
// 
//
// Related Topics 数组 双指针 排序 👍 2073 👎 0


//leetcode submit region begin(Prohibit modification and deletion)
// func fourSum(nums []int, target int) [][]int {
// if len(nums) < 4 {
// return nil
// }
// var res [][]int
// sort.Ints(nums)
// for k := 0; k < len(nums)-3; k++ {
// if nums[k] > target && nums[k] >= 0 {
// break
// }
// if k > 0 && nums[k] == nums[k-1] {
// continue
// }
// for i := k + 1; i < len(nums)-2; i++ {
// if nums[i]+nums[k] >= 0 && nums[i]+nums[k] > target {
// break
// }
// if i > k+1 && nums[i] == nums[i-1] {
// continue
// }
// l, r := i+1, len(nums)-1
// for l < r {
// if nums[k]+nums[i]+nums[l]+nums[r] < target {
// l++
// } else if nums[k]+nums[i]+nums[l]+nums[r] > target {
// r--
// } else {
// res = append(res, []int{nums[k], nums[i], nums[l], nums[r]})
// for l < r && nums[l] == nums[l+1] {
// l++
// }
// for l < r && nums[r] == nums[r-1] {
// r--
// }
// l++
// r--
// }
// }
// }
// }
// return res
// }

// func fourSum(nums []int, target int) [][]int {
// sort.Ints(nums)
// var res [][]int
// for i:=0;i<len(nums);i++ {
// if nums[i] > target && nums[i] >= 0 {
// break
// }
// if i>0 && nums[i] == nums[i-1] {
// continue
// }
// for j:=i+1;j<len(nums);j++ {
// if nums[i] + nums[j] > target && nums[i]+nums[j] >= 0 {
// break
// }
// if j>i+1 && nums[j] == nums[j-1] {
// continue
// }
// l, r := j+1, len(nums)-1
// for l < r {
// if nums[i] + nums[j] + nums[l] + nums[r] < target {
// l++
// } else if nums[i] + nums[j] + nums[l] + nums[r] > target {
// r--
// } else {
// res = append(res, []int{nums[i], nums[j], nums[l], nums[r]})
// for r > l && nums[r] == nums[r-1] {
// r--
// }
// for l < r && nums[l] == nums[l+1] {
// l++
// }
// l++
// r--
// }
// }
// }
// }
// return res
// }


func fourSum(nums []int, target int) [][]int {
var res [][]int
sort.Ints(nums)
for i:=0;i<len(nums);i++ {
if nums[i] > target && nums[i] >= 0 {
break
}
if i>0 && nums[i] == nums[i-1] {
continue
}
for j:=i+1;j<len(nums);j++ {
if nums[i] + nums[j] > target && nums[i] + nums[j] >= 0 {
break
}
if j > i+1 && nums[j] == nums[j-1] {
continue
}
l, r := j+1, len(nums)-1
for l < r {
if nums[i] + nums[j] + nums[l] + nums[r] == target {
res = append(res, []int{nums[i], nums[j], nums[l], nums[r]})
for r > l && nums[r] == nums[r-1] {
r--
}
for l < r && nums[l] == nums[l+1] {
l++
}
l++
r--
} else if nums[i] + nums[j] + nums[l] + nums[r] < target {
l++
} else {
r--
}
}
}
}
return res
}
//leetcode submit region end(Prohibit modification and deletion)
