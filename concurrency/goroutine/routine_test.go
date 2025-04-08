package routine

import "testing"

func TestStudyContext(t *testing.T) {
	StudyContext()
}

func TestGoDone(t *testing.T) {
	GoDone()
}

func BenchmarkGoDone(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GoDone()
	}
}

func TestDoOrder(t *testing.T) {
	DoOrder()
}

func TestDoOrder1(t *testing.T) {
	DoOrder1()
}

func TestPutMode(t *testing.T) {
	PutMode()
}

func TestWaitGroup2(t *testing.T) {
	WaitGroup2()
}

func TestWaitGroup(t *testing.T) {
	WaitGroup()
}

func TestCount(t *testing.T) {
	Count()
}
