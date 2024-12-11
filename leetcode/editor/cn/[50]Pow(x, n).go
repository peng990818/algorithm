//实现 pow(x, n) ，即计算 x 的整数 n 次幂函数（即，xⁿ ）。 
//
// 
//
// 示例 1： 
//
// 
//输入：x = 2.00000, n = 10
//输出：1024.00000
// 
//
// 示例 2： 
//
// 
//输入：x = 2.10000, n = 3
//输出：9.26100
// 
//
// 示例 3： 
//
// 
//输入：x = 2.00000, n = -2
//输出：0.25000
//解释：2-2 = 1/22 = 1/4 = 0.25
// 
//
// 
//
// 提示： 
//
// 
// -100.0 < x < 100.0 
// -231 <= n <= 231-1 
// n 是一个整数 
// 要么 x 不为零，要么 n > 0 。 
// -104 <= xⁿ <= 104 
// 
//
// Related Topics 递归 数学 👍 1395 👎 0


//leetcode submit region begin(Prohibit modification and deletion)
// 递归版本
// func myPow(x float64, n int) float64 {
//     if n == 0 {
//         return 1
//     }
//     if n > 0 {
//         return quickMul(x, n)
//     }
//     return 1.0 / quickMul(x, -n)
// }
// 快速幂
// func quickMul(x float64, n int) float64 {
//     if n == 0 {
//         return 1
//     }
//     y := quickMul(x, n/2)
//     if n % 2 == 0 {
//         return y * y
//     }
//     return y*y*x
// }

// func myPow(x float64, n int) float64 {
//     res, b := 1.0, false
//     if n < 0 {
//         n = -n
//         b = true
//     }
//     for n > 0 {
//         if n % 2 == 1 {
//             res *= x
//         }
//         x *= x
//         n /= 2
//     }
//     if b {
//         return 1.0/res
//     }
//     return res
// }

// 递归
// func myPow(x float64, n int) float64 {
//     if n == 0 {
//         return 1.0
//     }
//     if n > 0 {
//         return quick(x, n)
//     }
//     return 1.0/quick(x, n)
// }
//
// func quick(x float64, n int) float64 {
//     if n == 0 {
//         return 1
//     }
//     y := quick(x, n/2)
//     if n%2 == 0 {
//         return y*y
//     }
//     return y*y*x
// }

// 非递归
func myPow(x float64, n int) float64 {
    if n == 0 {
        return 1.0
    }
    res := 1.0
    b := false
    if n < 0 {
        n = -n
        b = true
    }
    for n > 0 {
        if n % 2 != 0 {
            res *= x
        }
        x*=x
        n /= 2
    }
    if b {
        return 1.0/res
    }
    return res
}
//leetcode submit region end(Prohibit modification and deletion)
