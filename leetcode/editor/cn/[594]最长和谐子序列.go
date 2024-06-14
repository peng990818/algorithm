//和谐数组是指一个数组里元素的最大值和最小值之间的差别 正好是 1 。 
//
// 现在，给你一个整数数组 nums ，请你在所有可能的子序列中找到最长的和谐子序列的长度。 
//
// 数组的子序列是一个由数组派生出来的序列，它可以通过删除一些元素或不删除元素、且不改变其余元素的顺序而得到。 
//
// 
//
// 示例 1： 
//
// 
//输入：nums = [1,3,2,2,5,2,3,7]
//输出：5
//解释：最长的和谐子序列是 [3,2,2,2,3]
// 
//
// 示例 2： 
//
// 
//输入：nums = [1,2,3,4]
//输出：2
// 
//
// 示例 3： 
//
// 
//输入：nums = [1,1,1,1]
//输出：0
// 
//
// 
//
// 提示： 
//
// 
// 1 <= nums.length <= 2 * 10⁴ 
// -10⁹ <= nums[i] <= 10⁹ 
// 
//
// Related Topics 数组 哈希表 计数 排序 滑动窗口 👍 402 👎 0


//leetcode submit region begin(Prohibit modification and deletion)
func findLHS(nums []int) int {
    // 排序+滑动窗口
    sort.Ints(nums)
    res := 0
    i:=0
    for j:=0;j<len(nums);j++{
        // 找到开始的位置
        for i<j && nums[j]-nums[i] > 1 {
            i++
        }
        // 统计频次并比较记录最大值
        if nums[j]-nums[i] == 1 && j-i+1 > res {
            res = j-i+1
        }
    }
    return res



    // 哈希
    // 对每个数出现的次数进行计数
    // mp := make(map[int]int)
    // for _, num := range nums {
    //     mp[num]++
    // }
    //
    // // 遍历hash，找到相差为1的两个数出现次数最大的值即为答案
    // res := 0
    // for k, v := range mp {
    //     if val, ok := mp[k+1]; ok {
    //         sum := v + val
    //         if sum > res {
    //             res = sum
    //         }
    //     }
    // }
    // return res
}
//leetcode submit region end(Prohibit modification and deletion)
