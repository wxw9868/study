package datastructure

import (
	"fmt"
	"testing"
)

func TestPrint(t *testing.T) {
	head := NewHeroNode()
	head.Insert(&HeroNode{Name: "曹操", No: 1})
	head.Insert(&HeroNode{Name: "刘备", No: 2})
	// head.Insert(&HeroNode{Name: "孙权", No: 3})
	// head.Remove(1)
	head.Print()
}

func TestFindFatherNode(t *testing.T) {
	fmt.Println("start test")
	t.Log("wxw")
}
