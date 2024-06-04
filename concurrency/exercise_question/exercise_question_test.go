package exercise_question

import "testing"

func TestDanGoroutine(t *testing.T) {
	DanGoroutine()
	DuoGoroutine()
	// 两个单协程和多协程的效率差好多
	G()
}

func TestDumpStr(t *testing.T) {
	DumpStr()
}

func TestOrder(t *testing.T) {
	Order()
}

func TestFactorial(t *testing.T) {
	Factorial()
}

func TestAddSlice(t *testing.T) {
	AddSlice()
}

func TestAddSliceMutex(t *testing.T) {
	AddSliceMutex()
}

func TestAddSliceWait(t *testing.T) {
	AddSliceWait()
}

func TestDoType(t *testing.T) {
	DoType()
}
