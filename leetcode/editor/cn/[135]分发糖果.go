//n 个孩子站成一排。给你一个整数数组 ratings 表示每个孩子的评分。 
//
// 你需要按照以下要求，给这些孩子分发糖果： 
//
// 
// 每个孩子至少分配到 1 个糖果。 
// 相邻两个孩子中，评分更高的那个会获得更多的糖果。 
// 
//
// 请你给每个孩子分发糖果，计算并返回需要准备的 最少糖果数目 。 
//
// 
//
// 示例 1： 
//
// 
//输入：ratings = [1,0,2]
//输出：5
//解释：你可以分别给第一个、第二个、第三个孩子分发 2、1、2 颗糖果。
// 
//
// 示例 2： 
//
// 
//输入：ratings = [1,2,2]
//输出：4
//解释：你可以分别给第一个、第二个、第三个孩子分发 1、2、1 颗糖果。
//     第三个孩子只得到 1 颗糖果，这满足题面中的两个条件。 
//
// 
//
// 提示： 
//
// 
// n == ratings.length 
// 1 <= n <= 2 * 10⁴ 
// 0 <= ratings[i] <= 2 * 10⁴ 
// 
//
// Related Topics 贪心 数组 👍 1696 👎 0


//leetcode submit region begin(Prohibit modification and deletion)
func candy(ratings []int) int {
if len(ratings) == 0 {
return 0
}
candies := make([]int, len(ratings))
for i := range candies {
candies[i] = 1
}
for i := 1; i < len(ratings); i++ {
// 相邻的右大于左
if ratings[i] > ratings[i-1] {
candies[i] = candies[i-1] + 1
}
}
for j := len(ratings) - 2; j >= 0; j-- {
// 相邻的左大于右
if ratings[j] > ratings[j+1] {
candies[j] = max(candies[j], candies[j+1]+1)
}
}
res := 0
for i := range candies {
res += candies[i]
}
return res
}
//leetcode submit region end(Prohibit modification and deletion)
