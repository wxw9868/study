package pen_questions

import (
	"fmt"
	"testing"
)

func TestMain(t *testing.T) {
	Main()
}

func TestIfPowerOfTwo(t *testing.T) {
	i := 32
	ok, power := IfPowerOfTwo(i)
	t.Logf("ok=%t,power=%d", ok, power)
}

func TestAlphanumericExchange1(t *testing.T) {
	AlphanumericExchange1()
}

func TestAlphanumericExchange2(t *testing.T) {
	AlphanumericExchange2()
}

func TestPrintout(t *testing.T) {
	Printout(10)
}

func TestMappingInversion(t *testing.T) {
	s := MappingInversion("name")
	fmt.Println(s)
}

func TestGames(t *testing.T) {
	games := Games(20)
	t.Logf("%d", games)
}

func TestProcessString(t *testing.T) {
	ProcessString()
}

func TestMaxAvg(t *testing.T) {
	MaxAvg()
}

func TestGetGroupList(t *testing.T) {
	GetGroupList()
}

func TestDoSearch(t *testing.T) {
	slice := []int{23, 32, 78, 43, 76, 65, 345, 762, 915, 86}
	DoSearch(slice)
}
