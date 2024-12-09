//仓库管理员以数组 stock 形式记录商品库存表，其中 stock[i] 表示对应商品库存余量。请返回库存余量最少的 cnt 个商品余量，返回 顺序不限。 
//
//
// 
//
// 示例 1： 
//
// 
//输入：stock = [2,5,7,4], cnt = 1
//输出：[2]
// 
//
// 示例 2： 
//
// 
//输入：stock = [0,2,3,6], cnt = 2
//输出：[0,2] 或 [2,0] 
//
// 
//
// 提示： 
//
// 
// 0 <= cnt <= stock.length <= 10000 0 <= stock[i] <= 10000 
// 
//
// 
//
// Related Topics 数组 分治 快速选择 排序 堆（优先队列） 👍 595 👎 0


//leetcode submit region begin(Prohibit modification and deletion)
func inventoryManagement(stock []int, cnt int) []int {
    if len(stock) == 0 || cnt <= 0 || cnt > len(stock) {
        return nil
    }
    buildHeap(stock)
    res := make([]int, 0, cnt)
    for size:=len(stock) - 1; size>=0&&cnt > 0;size-- {
        stock[0], stock[size] = stock[size], stock[0]
        res = append(res, stock[size])
        heapify(stock, 0, size)
        cnt--
    }
    return res
}

// 小根堆
func buildHeap(arr []int) {
    for i:=len(arr)/2-1;i>=0;i-- {
        heapify(arr, i, len(arr))
    }
}

func heapify(arr []int, index, size int) {
    left := 2*index + 1
    for left < size {
        smallest := index
        if arr[smallest] > arr[left] {
            smallest = left
        }
        if left + 1 < size && arr[smallest] > arr[left+1] {
            smallest = left+1
        }
        if smallest == index {
            break
        }
        arr[smallest], arr[index] = arr[index], arr[smallest]
        index = smallest
        left = 2*index+1
    }
}
//leetcode submit region end(Prohibit modification and deletion)
