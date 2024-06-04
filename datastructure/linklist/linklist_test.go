package linklist

import "testing"

func TestListNode_Insert(t *testing.T) {
	l := NewLinklist()
	for i := 0; i < 10; i++ {
		l.Insert(i)
	}
	l.Print()
}
