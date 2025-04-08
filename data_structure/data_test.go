package datastructure

import (
	"fmt"
	"testing"

	"github.com/wxw9868/study/utils"
)

func TestPrint(t *testing.T) {
	utils.MethodRuntime(func() {
		head := NewHeroNode()
		head.Insert(&HeroNode{Name: "曹操", No: 1})
		head.Insert(&HeroNode{Name: "刘备", No: 2})
		// head.Insert(&HeroNode{Name: "孙权", No: 3})
		// head.Remove(1)
		head.Print()
	})
}

func TestFindFatherNode(t *testing.T) {
	fmt.Println("start test")
	t.Log("wxw")
}

func TestListNode_Insert(t *testing.T) {
	l := NewLinklist()
	for i := range 10 {
		l.Insert(i)
	}
	l.Print()
}
