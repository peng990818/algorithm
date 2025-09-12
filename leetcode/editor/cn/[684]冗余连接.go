//树可以看成是一个连通且 无环 的 无向 图。 
//
// 给定一个图，该图从一棵 n 个节点 (节点值 1～n) 的树中添加一条边后获得。添加的边的两个不同顶点编号在 1 到 n 中间，且这条附加的边不属于树中已
//存在的边。图的信息记录于长度为 n 的二维数组 edges ，edges[i] = [ai, bi] 表示图中在 ai 和 bi 之间存在一条边。 
//
// 请找出一条可以删去的边，删除后可使得剩余部分是一个有着 n 个节点的树。如果有多个答案，则返回数组 edges 中最后出现的那个。 
//
// 
//
// 示例 1： 
//
// 
//
// 
//输入: edges = [[1,2], [1,3], [2,3]]
//输出: [2,3]
// 
//
// 示例 2： 
//
// 
//
// 
//输入: edges = [[1,2], [2,3], [3,4], [1,4], [1,5]]
//输出: [1,4]
// 
//
// 
//
// 提示: 
//
// 
// n == edges.length 
// 3 <= n <= 1000 
// edges[i].length == 2 
// 1 <= ai < bi <= edges.length 
// ai != bi 
// edges 中无重复元素 
// 给定的图是连通的 
// 
//
// Related Topics 深度优先搜索 广度优先搜索 并查集 图 👍 718 👎 0


//leetcode submit region begin(Prohibit modification and deletion)
// 并查集
type BSet struct {
n int
father []int
}

func NewBSet(n int) *BSet {
b := &BSet{
n: n,
father: make([]int, n),
}
for i := range b.father {
b.father[i] = i
}
return b
}

func (b *BSet) Find(u int) int {
if u == b.father[u] {
return u
}
b.father[u] = b.Find(b.father[u])
return b.father[u]
}

func (b *BSet) IsSame(u, v int) bool {
u = b.Find(u)
v = b.Find(v)
return u == v
}

func (b *BSet) Join(u, v int) {
u = b.Find(u)
v = b.Find(v)
if u == v {
return
}
b.father[v] = u
}

func findRedundantConnection(edges [][]int) []int {
bSet := NewBSet(1005)
for i:=0;i<len(edges);i++ {
if bSet.IsSame(edges[i][0], edges[i][1]) {
return edges[i]
}
bSet.Join(edges[i][0], edges[i][1])
}
return nil
}
//leetcode submit region end(Prohibit modification and deletion)
