package basic_knowledge

import (
	"fmt"
	"time"
)

func StudySlice() {
	fmt.Println("make slice: ")
	slice := make([]int, 3, 5)
	fmt.Printf("r: %v, p: %p, len: %d, cap: %d\n", slice, slice, len(slice), cap(slice))
	fmt.Println("append slice: ")
	slice = append(slice, 1, 2)
	fmt.Printf("r: %v, p: %p, len: %d, cap: %d\n", slice, slice, len(slice), cap(slice))
	slice = append(slice, 3)
	fmt.Printf("r: %v, p: %p, len: %d, cap: %d\n", slice, slice, len(slice), cap(slice))

	fmt.Println("definition slice1: ")
	var slice1 []string
	fmt.Printf("r: %v, p: %p, len: %d, cap: %d\n", slice1, slice1, len(slice1), cap(slice1))
	fmt.Println("append slice1: ")
	slice1 = append(slice1, "w", "e")
	fmt.Printf("r: %v, p: %p, len: %d, cap: %d\n", slice1, slice1, len(slice1), cap(slice1))
	slice1 = append(slice1, "q")
	fmt.Printf("r: %v, p: %p, len: %d, cap: %d\n", slice1, slice1, len(slice1), cap(slice1))

	fmt.Println("new slice2: ")
	slice2 := new([]int)
	fmt.Printf("r: %v, p: %p, len: %d, cap: %d\n", slice2, slice2, len(*slice2), cap(*slice2))
	fmt.Println("append slice2: ")
	*slice2 = append(*slice2, 1, 2)
	fmt.Printf("r: %v, p: %p, len: %d, cap: %d\n", *slice2, *slice2, len(*slice2), cap(*slice2))
	*slice2 = append(*slice2, 5)
	fmt.Printf("r: %v, p: %p, len: %d, cap: %d\n", *slice2, *slice2, len(*slice2), cap(*slice2))

	fmt.Println("make slice3: ")
	slice3 := make([]int, 0)
	fmt.Printf("r: %v, p: %p, len: %d, cap: %d\n", slice3, slice3, len(slice3), cap(slice3))
}

func StudyCopy() {
	slice1 := []int{1, 2, 3, 4, 5}
	slice2 := slice1
	slice3 := []int{0, 0, 0, 0, 0, 11, 22, 33, 44, 55}
	copy(slice3, slice1)
	slice1[1] = 100
	fmt.Println("copy: ", slice1)
	fmt.Println("copy: ", slice2)
	fmt.Println("copy: ", slice3)

	// copy函数
	// 通过copy函数可以把一个切片内容复制到另一个切片中
	// copy(目标切片, 源切片)
	// 拷贝时严格按照脚标进行拷贝
	// 出处：https://blog.csdn.net/weixin_42677653/article/details/105130929

	copySlice()
	//copyDel()
	s13 := copyDelEle([]int{1, 2, 3, 4, 5, 6, 7, 8}, 2)
	fmt.Println(s13)

}

func copySlice() {
	// 1.把短切片拷贝到长切片中
	s1 := []int{1, 2}
	s2 := []int{3, 4, 5, 6, 7}
	copy(s2, s1)
	fmt.Println(s1) // [1 2]
	fmt.Println(s2) // [1 2 5 6 7]

	// 2.把长切片拷贝到短切片中
	s3 := []int{1, 2}
	s4 := []int{3, 4, 5, 6, 7}
	copy(s3, s4)
	fmt.Println(s3) // [3 4]
	fmt.Println(s4) // [3 4 5 6 7]

	// 3.把切片片段拷贝到切片中
	s5 := []int{1, 2}
	s6 := []int{3, 4, 5, 6, 7}
	copy(s5, s6[1:])
	fmt.Println(s5) // [4 5]
	fmt.Println(s6) // [3 4 5 6 7]

	s7 := []int{1, 2}
	s8 := []int{3, 4, 5, 6, 7}
	copy(s7, s8[3:])
	fmt.Println(s7) // [6 7]
	fmt.Println(s8) // [3 4 5 6 7]

	s9 := []int{1, 2}
	s10 := []int{3, 4, 5, 6, 7}
	copy(s9, s10[4:])
	fmt.Println(s9)  // [7 2]  // 如果后面没有了就不替换,用自己的元素
	fmt.Println(s10) // [3 4 5 6 7]

}

// 使用copy()函数完成删除切片的元素,且保证源切片内容不变
func copyDel() {
	s11 := []int{1, 2, 3, 4, 5, 6, 7, 8}
	// 要删除的元素脚标为n
	n := 2
	s12 := make([]int, n)
	copy(s12, s11[0:n])
	s12 = append(s12, s11[n+1:]...)
	fmt.Println(s11) // [1 2 3 4 5 6 7 8]
	fmt.Println(s12) // [1 2 4 5 6 7 8]
}

func copyDelEle(slice []int, n int) []int { // slice为源切片, n即要删除的元素的脚标, 调用时传参
	//s11 := []int {1, 2, 3, 4, 5, 6, 7, 8}
	s11 := slice
	s12 := make([]int, n)
	copy(s12, s11[0:n])
	s12 = append(s12, s11[n+1:]...)
	//fmt.Println(s11)
	//fmt.Println(s12)
	return s12
}

func StudyMap() {
	var m0 map[string]int
	fmt.Println(m0)
	fmt.Println(m0 == nil)
	fmt.Printf("r: %v, p: %p, len: %d\n", m0, m0, len(m0))

	m0 = make(map[string]int)
	fmt.Printf("r: %v, p: %p, len: %d\n", m0, m0, len(m0))

	m0["age"] = 10
	m0["year"] = 2023
	fmt.Printf("r: %v, p: %p, len: %d\n", m0, m0, len(m0))
}

func StudyChannel() {
	fmt.Println("channel")
	var ch chan int
	fmt.Printf("r: %v, p: %p, len: %d, cap: %d\n", ch, ch, len(ch), cap(ch))

	ch = make(chan int)
	fmt.Printf("r: %v, p: %p, len: %d, cap: %d\n", ch, ch, len(ch), cap(ch))

	go func() {
		ch <- 1
	}()
	fmt.Printf("r: %v, p: %p, len: %d, cap: %d\n", ch, ch, len(ch), cap(ch))
	time.Sleep(time.Millisecond)
	<-ch

	fmt.Println("ch1")
	ch1 := make(chan int, 10)
	fmt.Printf("r: %v, p: %p, len: %d, cap: %d\n", ch1, ch1, len(ch1), cap(ch1))

	ch1 <- 1
	fmt.Printf("r: %v, p: %p, len: %d, cap: %d\n", ch1, ch1, len(ch1), cap(ch1))
	r := <-ch1
	fmt.Println("data1: ", r)
	fmt.Printf("r: %v, p: %p, len: %d, cap: %d\n", ch1, ch1, len(ch1), cap(ch1))

	for i := 0; i < 10; i++ {
		ch1 <- i
	}
	close(ch1)
	fmt.Printf("r: %v, p: %p, len: %d, cap: %d\n", ch1, ch1, len(ch1), cap(ch1))
	//for i := range ch1 {
	//	fmt.Println("data2: ", i)
	//}
End:
	for {
		select {
		case i, ok := <-ch1:
			if !ok {
				fmt.Println("data2: ", i, ok)
				break End
			}
			fmt.Println("data2: ", i, ok)
		}
	}
	time.Sleep(time.Second)
}

// StudyReturnAndDefer 理解return和defer的执行逻辑关系
func StudyReturnAndDefer() {
	fmt.Println(a())
	fmt.Println(b())
	fmt.Println(c())
}

func a() int {
	var i int
	defer func() {
		i++
	}()
	return i
}

func b() (i int) {
	defer func() {
		i++
	}()
	return i
}

func c() int {
	var i int
	defer func() {
		i++
	}()
	return i + 1
}
