//给定两个字符串形式的非负整数 num1 和num2 ，计算它们的和并同样以字符串形式返回。 
//
// 你不能使用任何內建的用于处理大整数的库（比如 BigInteger）， 也不能直接将输入的字符串转换为整数形式。 
//
// 
//
// 示例 1： 
//
// 
//输入：num1 = "11", num2 = "123"
//输出："134"
// 
//
// 示例 2： 
//
// 
//输入：num1 = "456", num2 = "77"
//输出："533"
// 
//
// 示例 3： 
//
// 
//输入：num1 = "0", num2 = "0"
//输出："0"
// 
//
// 
//
// 
//
// 提示： 
//
// 
// 1 <= num1.length, num2.length <= 10⁴ 
// num1 和num2 都只包含数字 0-9 
// num1 和num2 都不包含任何前导零 
// 
//
// Related Topics 数学 字符串 模拟 👍 853 👎 0


//leetcode submit region begin(Prohibit modification and deletion)
func addStrings(num1 string, num2 string) string {
    lens := len(num1)
    if len(num2) > lens {
        lens = len(num2)
    }
    res := make([]byte, lens+1)
    i, j, k, b := len(num1) - 1, len(num2) - 1, lens, false
    for i >= 0 && j >= 0 {
        num := num1[i] - '0' + num2[j] - '0'
        // 进位
        if b {
            num++
            b = false
        }
        if num >= 10 {
            num = num - 10 + '0'
            b = true
        } else {
            num += '0'
        }
        res[k] = num
        k--
        i--
        j--
    }
    for i>=0 {
        tmp := num1[i] - '0'
        if b {
            tmp++
            b = false
        }
        if tmp >= 10 {
            tmp = tmp - 10 + '0'
            b = true
        } else {
            tmp+='0'
        }
        res[k] = tmp
        k--
        i--
    }
    for j >= 0 {
        tmp := num2[j] - '0'
        if b {
            tmp++
            b = false
        }
        if tmp >= 10 {
            tmp = tmp - 10 + '0'
            b = true
        } else {
            tmp+='0'
        }
        res[k] = tmp
        k--
        j--
    }
    if b {
        res[k] = '1'
    } else {
        res = res[1:]
    }
    return string(res)
}
//leetcode submit region end(Prohibit modification and deletion)
