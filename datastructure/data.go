package datastructure

import "fmt"

type HeroNode struct {
	No   int       `json:"no"`
	Name string    `json:"name"`
	Next *HeroNode `json:"next"`
}

var length int

func NewHeroNode() *HeroNode {
	node := new(HeroNode)
	return node
}

// 增加节点
func (h *HeroNode) Insert(node *HeroNode) {
	if node == nil {
		return
	}
	temp := h
	for {
		if temp.Next == nil {
			break
		}
		temp = h.Next
	}
	temp.Next = node
	length++
}

// 移除指定节点
func (h *HeroNode) Remove(heroNo int) {
	if heroNo == 0 { // 头部
		return
	}
	prev := h
	temp := prev.Next
	isRemove := false
	for {
		if temp.No == heroNo {
			isRemove = true
			break
		}
		prev = temp
		temp = temp.Next
	}
	if isRemove {
		if temp.Next == nil {
			prev.Next = nil
		} else {
			prev.Next = temp.Next
		}
	}
}

func (h *HeroNode) Print() {
	for {
		if h == nil {
			break
		}
		fmt.Printf("no=%d,name=%s,,next=%p,len=%d\n", h.No, h.Name, h.Next, GetLength())
		h = h.Next
	}
}

func GetLength() int {
	return length
}
