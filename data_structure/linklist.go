package datastructure

// ListNode 单链表
type ListNode struct {
	Val  int
	Next *ListNode
}

// NewLinklist 创建链表
func NewLinklist() *ListNode {
	return &ListNode{}
}

// Insert 插入链表
func (head *ListNode) Insert(val int) {
	node := &ListNode{
		Val:  val,
		Next: nil,
	}
	if head.Next == nil {
		head.Next = node
		return
	}
	cur := head.Next
	for cur.Next != nil {
		cur = cur.Next
	}
	cur.Next = node
}

// Find 查找链表
func (head *ListNode) Find(val int) *ListNode {
	cur := head.Next
	for cur != nil {
		if cur.Val == val {
			return cur
		}
		cur = cur.Next
	}
	return nil
}

// Delete 删除链表
func (head *ListNode) Delete(val int) {
	pre := head
	cur := head.Next
	for cur != nil {
		if cur.Val == val {
			pre.Next = cur.Next
			return
		}
		pre = cur
		cur = cur.Next
	}
}

// Print 遍历链表
func (head *ListNode) Print() {
	cur := head.Next
	for cur != nil {
		println(cur.Val)
		cur = cur.Next
	}
}
