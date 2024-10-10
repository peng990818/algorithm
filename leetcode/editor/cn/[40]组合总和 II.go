//给定一个候选人编号的集合 candidates 和一个目标数 target ，找出 candidates 中所有可以使数字和为 target 的组合。 
//
// candidates 中的每个数字在每个组合中只能使用 一次 。 
//
// 注意：解集不能包含重复的组合。 
//
// 
//
// 示例 1: 
//
// 
//输入: candidates = [10,1,2,7,6,1,5], target = 8,
//输出:
//[
//[1,1,6],
//[1,2,5],
//[1,7],
//[2,6]
//] 
//
// 示例 2: 
//
// 
//输入: candidates = [2,5,2,1,2], target = 5,
//输出:
//[
//[1,2,2],
//[5]
//] 
//
// 
//
// 提示: 
//
// 
// 1 <= candidates.length <= 100 
// 1 <= candidates[i] <= 50 
// 1 <= target <= 30 
// 
//
// Related Topics 数组 回溯 👍 1594 👎 0


//leetcode submit region begin(Prohibit modification and deletion)
// 纯递归版本
func combinationSum21(candidates []int, target int) (ans [][]int) {
    if len(candidates) == 0 {
        return
    }
    sort.Ints(candidates)
    for i:=0; i<len(candidates); i++ {
        if candidates[i] > target {
            candidates = candidates[:i]
            break
        }
    }
    res := make([][]int, 0)
    path := make([]int, 0)
    used := make([]bool, len(candidates))
    process(candidates, target, 0, &res, path, used)
    return res
}

// 纯递归版本
func process1(src []int, target int, index int, res *[][]int, path []int, used []bool) {
    if target == 0 {
        tmp := make([]int, len(path))
        copy(tmp, path)
        *res = append(*res, tmp)
        return
    }

    if index >= len(src) {
        return
    }

    if index > 0 && src[index] == src[index-1] && !used[index-1] {
        goto label
    }

    path = append(path, src[index])
    used[index] = true
    process1(src, target-src[index], index+1, res, path, used)

    path = path[:len(path)-1]
    used[index] = false
label:
    process1(src, target, index+1, res, path, used)
}

func combinationSum2(candidates []int, target int) (ans [][]int) {
    if len(candidates) == 0 {
        return
    }
    sort.Ints(candidates)
    res := make([][]int, 0)
    path := make([]int, 0)
    // 去重使用
    used := make([]bool, len(candidates))
    process(candidates, target, 0, &res, path, used)
    return res
}

// 循环+递归
func process(src []int, target int, index int, res *[][]int, path []int, used []bool) {
    if target == 0 {
        tmp := make([]int, len(path))
        copy(tmp, path)
        *res = append(*res, tmp)
        return
    }
    for i := index; i < len(src); i++ {
        if src[i] > target {
            break
        }
        // used[i-1]为true标记上一次递归使用过
        // used[i-1]为false标记当前层使用过
        if i > 0 && src[i] == src[i-1] && !used[i-1] {
            continue
        }
        path = append(path, src[i])
        used[i] = true
        process(src, target-src[i], i+1, res, path, used)
        path = path[:len(path)-1]
        used[i] = false
    }
}


//leetcode submit region end(Prohibit modification and deletion)
