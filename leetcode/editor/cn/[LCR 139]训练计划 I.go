//教练使用整数数组 actions 记录一系列核心肌群训练项目编号。为增强训练趣味性，需要将所有奇数编号训练项目调整至偶数编号训练项目之前。请将调整后的训练项
//目编号以 数组 形式返回。 
//
// 
//
// 示例 1： 
//
// 
//输入：actions = [1,2,3,4,5]
//输出：[1,3,5,2,4] 
//解释：为正确答案之一 
//
// 
//
// 提示： 
//
// 
// 0 <= actions.length <= 50000 
// 0 <= actions[i] <= 10000 
// 
//
// 
//
// Related Topics 数组 双指针 排序 👍 339 👎 0


//leetcode submit region begin(Prohibit modification and deletion)
func trainingPlan(actions []int) []int {
    if len(actions) <= 1 {
        return actions
    }
    less, more := -1, len(actions)
    l := 0
    for l < more {
        if actions[l] % 2 == 1 {
            less++
            actions[l], actions[less] = actions[less], actions[l]
            l++
        } else {
            more--
            actions[l], actions[more] = actions[more], actions[l]
        }
    }
    return actions
}
//leetcode submit region end(Prohibit modification and deletion)
