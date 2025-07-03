//你可以选择使用单链表或者双链表，设计并实现自己的链表。 
//
// 单链表中的节点应该具备两个属性：val 和 next 。val 是当前节点的值，next 是指向下一个节点的指针/引用。 
//
// 如果是双向链表，则还需要属性 prev 以指示链表中的上一个节点。假设链表中的所有节点下标从 0 开始。 
//
// 实现 MyLinkedList 类： 
//
// 
// MyLinkedList() 初始化 MyLinkedList 对象。 
// int get(int index) 获取链表中下标为 index 的节点的值。如果下标无效，则返回 -1 。 
// void addAtHead(int val) 将一个值为 val 的节点插入到链表中第一个元素之前。在插入完成后，新节点会成为链表的第一个节点。 
// void addAtTail(int val) 将一个值为 val 的节点追加到链表中作为链表的最后一个元素。 
// void addAtIndex(int index, int val) 将一个值为 val 的节点插入到链表中下标为 index 的节点之前。如果 
//index 等于链表的长度，那么该节点会被追加到链表的末尾。如果 index 比长度更大，该节点将 不会插入 到链表中。 
// void deleteAtIndex(int index) 如果下标有效，则删除链表中下标为 index 的节点。 
// 
//
// 
//
// 示例： 
//
// 
//输入
//["MyLinkedList", "addAtHead", "addAtTail", "addAtIndex", "get", 
//"deleteAtIndex", "get"]
//[[], [1], [3], [1, 2], [1], [1], [1]]
//输出
//[null, null, null, null, 2, null, 3]
//
//解释
//MyLinkedList myLinkedList = new MyLinkedList();
//myLinkedList.addAtHead(1);
//myLinkedList.addAtTail(3);
//myLinkedList.addAtIndex(1, 2);    // 链表变为 1->2->3
//myLinkedList.get(1);              // 返回 2
//myLinkedList.deleteAtIndex(1);    // 现在，链表变为 1->3
//myLinkedList.get(1);              // 返回 3
// 
//
// 
//
// 提示： 
//
// 
// 0 <= index, val <= 1000 
// 请不要使用内置的 LinkedList 库。 
// 调用 get、addAtHead、addAtTail、addAtIndex 和 deleteAtIndex 的次数不超过 2000 。 
// 
//
// Related Topics 设计 链表 👍 1149 👎 0


//leetcode submit region begin(Prohibit modification and deletion)
type MyLinkedList struct {
head   *node
tail   *node
length int
}

type node struct {
val  int
next *node
prev *node
}

func Constructor() MyLinkedList {
return MyLinkedList{}
}

func (this *MyLinkedList) Get(index int) int {
if index < 0 || index >= this.length {
return -1
}
p := this.head
for index > 0 {
p = p.next
index--
}
if p == nil {
return -1
}
return p.val
}

func (this *MyLinkedList) AddAtHead(val int) {
if this.head == nil {
this.head = &node{val: val}
this.tail = this.head
this.length++
return
}
n := &node{val: val}
n.next = this.head
this.head.prev = n
this.length++
this.head = n
}

func (this *MyLinkedList) AddAtTail(val int) {
if this.tail == nil {
this.tail = &node{val: val}
this.head = this.tail
this.length++
return
}
n := &node{val: val}
n.prev = this.tail
this.tail.next = n
this.tail = n
this.length++
}

func (this *MyLinkedList) AddAtIndex(index int, val int) {
if index > this.length || index < 0 {
return
}
if index == this.length {
this.AddAtTail(val)
return
}
if index == 0 {
this.AddAtHead(val)
return
}
n := &node{val: val}
p := this.head
for index > 0 {
p = p.next
index--
}
n.next = p
n.prev = p.prev
p.prev.next = n
p.prev = n
this.length++
}

func (this *MyLinkedList) DeleteAtIndex(index int) {
if index < 0 || index >= this.length {
return
}
p := this.head
for index > 0 {
p = p.next
index--
}
if p.prev != nil {
p.prev.next = p.next
} else {
this.head = p.next
}
if p.next != nil {
p.next.prev = p.prev
} else {
this.tail = p.prev
}
this.length--
p = nil
}


/**
 * Your MyLinkedList object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Get(index);
 * obj.AddAtHead(val);
 * obj.AddAtTail(val);
 * obj.AddAtIndex(index,val);
 * obj.DeleteAtIndex(index);
 */
//leetcode submit region end(Prohibit modification and deletion)
