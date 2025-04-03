package main

// func main() {
// gonum := runtime.NumGoroutine()
// fmt.Printf("NumGoroutine: %d\n", gonum)
// go func() {
// 	time.Sleep(time.Second)
// 	fmt.Println("1")
// }()

// go func() {
// 	time.Sleep(time.Second)
// 	fmt.Println("2")
// }()

// gonum = runtime.NumGoroutine()
// fmt.Printf("NumGoroutine: %d\n", gonum)
// time.Sleep(time.Second * 2)
// gonum = runtime.NumGoroutine()
// fmt.Printf("NumGoroutine: %d\n", gonum)
// select {}

// i := -5
// j := +5
// fmt.Printf("%+d %+d", i, j)

// arr := [3]int{1, 2, 3}
// fmt.Printf("arr: p=%p,len=%d,cap=%d\n", &arr, len(arr), cap(arr))
// fmt.Printf("arr: p=%p\n", &arr[0])
// fmt.Printf("arr: p=%p\n", &arr[1])
// fmt.Printf("arr: p=%p\n", &arr[2])
// s1 := []int{1, 2, 3}
// s2 := []int{4, 5}
// fmt.Printf("p=%p\n", &s1[0])
// fmt.Printf("p=%p\n", &s1[1])
// fmt.Printf("p=%p\n", &s1[2])
// fmt.Printf("s1: p=%p,len=%d,cap=%d\n", &s1, len(s1), cap(s1))
// fmt.Printf("s2: p=%p,len=%d,cap=%d\n", &s2, len(s2), cap(s2))
// s1 = append(s1, s2...)
// fmt.Printf("s1: p=%p,len=%d,cap=%d\n", &s1, len(s1), cap(s1))
// fmt.Printf("s2: p=%p,len=%d,cap=%d\n", &s2, len(s2), cap(s2))
// fmt.Println(s1)
// }

// type S struct {
// 	A int
// 	B *int
// 	C float64
// 	D func() string
// 	E chan struct{}
// }

// func main() {
// 	s := S{
// 		A: 1,
// 		B: nil,
// 		C: 12.15,
// 		D: func() string {
// 			return "NowCoder"
// 		},
// 		E: make(chan struct{}),
// 	}

// 	_, err := json.Marshal(s)
// 	if err != nil {
// 		log.Printf("err occurred..")
// 		return
// 	}
// 	log.Printf("everything is ok.")
// 	return
// }
