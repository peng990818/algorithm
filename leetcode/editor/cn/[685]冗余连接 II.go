//在本问题中，有根树指满足以下条件的 有向 图。该树只有一个根节点，所有其他节点都是该根节点的后继。该树除了根节点之外的每一个节点都有且只有一个父节点，而根节
//点没有父节点。 
//
// 输入一个有向图，该图由一个有着 n 个节点（节点值不重复，从 1 到 n）的树及一条附加的有向边构成。附加的边包含在 1 到 n 中的两个不同顶点间，这条
//附加的边不属于树中已存在的边。 
//
// 结果图是一个以边组成的二维数组 edges 。 每个元素是一对 [ui, vi]，用以表示 有向 图中连接顶点 ui 和顶点 vi 的边，其中 ui 是 
//vi 的一个父节点。 
//
// 返回一条能删除的边，使得剩下的图是有 n 个节点的有根树。若有多个答案，返回最后出现在给定二维数组的答案。 
//
// 
//
// 示例 1： 
// 
// 
//输入：edges = [[1,2],[1,3],[2,3]]
//输出：[2,3]
// 
//
// 示例 2： 
// 
// 
//输入：edges = [[1,2],[2,3],[3,4],[4,1],[1,5]]
//输出：[4,1]
// 
//
// 
//
// 提示： 
//
// 
// n == edges.length 
// 3 <= n <= 1000 
// edges[i].length == 2 
// 1 <= ui, vi <= n 
// 
//
// Related Topics 深度优先搜索 广度优先搜索 并查集 图 👍 468 👎 0


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

func getRemoveEdge(edges [][]int) []int {
b := NewBSet(1010)
for i:=0;i<len(edges);i++ {
if b.IsSame(edges[i][0], edges[i][1]) {
return edges[i]
}
b.Join(edges[i][0], edges[i][1])
}
return nil
}

func isTreeAfterRemoveEdge(edges [][]int, del int) bool {
b := NewBSet(1010)
for i:=0;i<len(edges);i++ {
if i == del {
continue
}
if b.IsSame(edges[i][0], edges[i][1]) {
return false
}
b.Join(edges[i][0], edges[i][1])
}
return true
}

func findRedundantDirectedConnection(edges [][]int) []int {
inDegree := make([]int, 1010)
n := len(edges)
for i:=0;i<n;i++ {
inDegree[edges[i][1]]++
}
vec := make([]int, 0)
for i:=n-1;i>=0;i-- {
if inDegree[edges[i][1]] == 2 {
vec = append(vec, i)
}
}
if len(vec) > 0 {
if isTreeAfterRemoveEdge(edges, vec[0]) {
return edges[vec[0]]
}
return edges[vec[1]]
}
return getRemoveEdge(edges)
}
//leetcode submit region end(Prohibit modification and deletion)
