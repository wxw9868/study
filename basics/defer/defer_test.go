package def

import (
	"fmt"
	"testing"
)

func TestDemo(t *testing.T) {
	fmt.Println("return: ", Demo())
}

func TestReturn(t *testing.T) {
	fmt.Println(Return())
}

func TestDefer(t *testing.T) {
	Defer()
}

func TestR(t *testing.T) {
	Run()
}
