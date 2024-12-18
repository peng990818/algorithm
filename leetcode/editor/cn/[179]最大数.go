//给定一组非负整数 nums，重新排列每个数的顺序（每个数不可拆分）使之组成一个最大的整数。 
//
// 注意：输出结果可能非常大，所以你需要返回一个字符串而不是整数。 
//
// 
//
// 示例 1： 
//
// 
//输入：nums = [10,2]
//输出："210" 
//
// 示例 2： 
//
// 
//输入：nums = [3,30,34,5,9]
//输出："9534330"
// 
//
// 
//
// 提示： 
//
// 
// 1 <= nums.length <= 100 
// 0 <= nums[i] <= 10⁹ 
// 
//
// Related Topics 贪心 数组 字符串 排序 👍 1310 👎 0


//leetcode submit region begin(Prohibit modification and deletion)
func largestNumber(nums []int) string {
    ss := make([]string, len(nums))
    for i, num := range nums {
        ss[i] = strconv.Itoa(num)
    }
    sort.Slice(ss, func(i, j int) bool {
        return ss[i]+ss[j] >= ss[j]+ss[i]
    })
    res := ""
    for i:=0;i<len(ss);i++ {
        res += ss[i]
    }
    if res[0] == '0' {
        return "0"
    }
    return res
}
//leetcode submit region end(Prohibit modification and deletion)
