package pen_questions

import (
	"fmt"
	"math/rand"
	"runtime"
	"strings"
	"sync"
	"time"
)

// IfPowerOfTwo
// 判断一个数是不是2的幂数
// 例如：8是，6不是
func IfPowerOfTwo(i int) (bool, int) {
	if i >= 2 {
		j := 1
		for {
			j++

			a := i % 2
			i = i / 2
			if i == 2 && a == 0 {
				return true, j
			}

			if a == 1 {
				return false, 0
			}
		}
	}
	return false, 0
}

// AlphanumericExchange1 方法一
// 实现一个交替打印字母和数字的简单程序
func AlphanumericExchange1() {
	var zm, sz = make(chan struct{}), make(chan struct{})

	go func() {
		for {
			<-zm
			fmt.Println("a")
			sz <- struct{}{}
		}

	}()

	go func() {
		for {
			<-sz
			fmt.Println("1")
			zm <- struct{}{}
		}
	}()

	zm <- struct{}{}

	select {}
}

// AlphanumericExchange2 方法二
// 实现一个交替打印字母和数字的简单程序，呈现的效果是：
// a1b2c3d4e5f6g7h8i9j10k11l12m13n14o15p16q17r18s19t20u21v22w23x24y25z26
func AlphanumericExchange2() {
	var letters, numbers = make(chan struct{}), make(chan struct{})
	var exit, done = make(chan struct{}), make(chan struct{})
	var wg sync.WaitGroup
	wg.Add(2)

	num := 1
	str := 'a'

	go func() {
		defer wg.Done()
	Loop:
		for {
			select {
			case <-letters:
				fmt.Printf("%c", str)
				if str <= 'z' {
					str++
					numbers <- struct{}{}
				}
			case <-exit:
				break Loop
			}
		}
	}()
	go func() {
		defer wg.Done()
	Loop:
		for {
			select {
			case <-numbers:
				fmt.Printf("%d", num)
				if num < 26 {
					num++
					letters <- struct{}{}
				} else {
					done <- struct{}{}
				}
			case <-exit:
				break Loop
			}
		}
	}()

	letters <- struct{}{}

	<-done
	close(exit)

	fmt.Print("\n")
	wg.Wait()
}

// Printout n为层数
/*
使用代码输出下面的形状
         *
        ***
       *****
      *******
     *********
    ***********
   *************
  ***************
 *****************
*******************
*/
func Printout(n int) {
	for i := 1; i <= n; i++ {
		for m := n - i; m > 0; m-- {
			fmt.Print(" ")
		}
		for j := 0; j < 2*i-1; j++ {
			fmt.Print("*")
		}
		fmt.Print("\n")
	}
}

// MappingInversion
/*
定义一个字母映射如下: a 变成 z, b 变成 y, c 变成 x,..., n 变成 m,
m 变成 n,...,z 变成 a。
将字符串反射定义为将字母反射应用于其每个字符的结果。
请反射给定的字符串。
示例1:
输入:"name"
输出:"mznv"
示例2:
输入:"abyz"
输出:"zyba"

func solution(inputString string) string {
     // 在这⾥写代码

     return
}
*/
func MappingInversion(inputString string) string {
	// 在这⾥写代码
	var am, zn string
	var a = 'a'
	var z = 'z'
	for i := 0; i < 13; i++ {
		am += fmt.Sprintf("%c", a)
		a++

		zn += fmt.Sprintf("%c", z)
		z--
	}

	var outString = ""
	for i := 0; i < len(inputString); i++ {
		for j := 0; j < 13; j++ {
			if string(am[j]) == string(inputString[i]) {
				outString += string(zn[j])

			} else {
				if string(zn[j]) == string(inputString[i]) {
					outString += string(am[j])
				}
			}
		}
	}
	return outString
}

// Games
/*
编写一个程序来计算体育联赛的比赛场数，
给定联赛的球队数量为n,其中每支球队都与
除自己之外的所有球队进行一场比赛。有2支
球队，则打一场，有3支球队则打3场。那么
n支球队需要打多少场比赛呢?
注:输入参数为球队数量;输出为场数
示例1
输入:4
输出:6
示例2
输入:20
输出:190

func solution(team int) int {
   // 在这⾥写代码

   return
}
*/
func Games(team int) int {
	// 在这⾥写代码
	if team < 2 {
		fmt.Println("条件不成立")
		return 0
	}
	count := 0
	doGames(team, &count)
	return count
}

func doGames(team int, count *int) {
	if team >= 2 {
		team--
		for i := 1; i <= team; i++ {
			*count++
		}
		doGames(team, count)
	}
}

// ProcessString
/*
请完成以下任务：
1. 请勿将外部IDE中编写好的代码粘贴到此Pad中，需在此Pad中完成代码编写和调试。
2. 面试结束后，请点击右下角的“End Interview”按钮。
3. 字符输入规则如下：
    a. 每行表示一条记录，字段之间以逗号（,）分隔。
    b. 如果字段内容包含逗号（,），则需使用双引号（"）包裹。
    c. 如果字段内容包含双引号（"），则需使用两个双引号（""）进行转义并用双引号包裹。
4. 编写解析程序，将解析后的内容按行输出，字段之间以制表符（\t）分隔。
5. 以下是一个示例：
    1）输入：Linda,47,"旅游,""攀岩",New Job
    2）输出：Linda 47 旅游,"攀岩 New Job
请根据以上要求编写解析程序，并将解析后的内容按照规则输出。
*/
func ProcessString() {
	rows := `2,Tina,37,"足球,""篮球",Old Job
3,Alice Job,66,"""看电影"",旅游","上海,上海市"
4,John,44,"洗衣机101,""","LA""CITY"""
5,"Jane,li",55,Hiking,Canada`
	execute(rows)
}

func execute(rows string) {
	Str := strings.Replace(rows, `""`, `-`, -1)
	var status = true
	var s1 string
	for _, v := range Str {
		if v == ',' && status {
			s1 += "\t"
		} else if v == '"' && status {
			status = false
		} else if v == '"' && !status {
			status = true
		} else {
			s1 += string(v)
		}
	}
	result := strings.Replace(s1, `-`, `"`, -1)
	println(result)
}

// func execute(rows string) {
// 	for _, row := range strings.Split(rows, "\n") {
// 		result := ""
// 		inPuote := false
// 		for i := 0; i < len(row); i++ {
// 			ch := string(row[i])
// 			if ch == "\"" {
// 				if inPuote {
// 					if i+1 < len(row) && string(row[i+1]) == "\"" {
// 						result += ch
// 						i++
// 					} else {
// 						inPuote = false
// 					}
// 				} else {
// 					inPuote = true
// 				}
// 			} else if ch == "," {
// 				if inPuote {
// 					result += ch
// 				} else {
// 					result += "\t"
// 				}
// 			} else {
// 				result += ch
// 			}
// 		}
// 		println(result)
// 	}
// }
//func execute(rows string) {
//	for _, row1 := range strings.Split(rows, "\n") {
//		row := []rune(row1)
//		result := ""
//		inQuote := false
//		for i := 0; i < len(row); i++ {
//			ch := string(row[i])
//			if ch == "\"" {
//				if inQuote {
//					if i+1 < len(row) && string(row[i+1]) == "\"" {
//						result += ch
//						i++
//					} else {
//						inQuote = false
//					}
//				} else {
//					inQuote = true
//				}
//			} else if ch == "," {
//				if inQuote {
//					result += ch
//				} else {
//					result += "\t"
//				}
//			} else {
//				result += ch
//			}
//		}
//		println(result)
//	}
//}

// MaxAvg 给出个两组，求出最高平均分
func MaxAvg() {
	arr := [5]string{"a", "b", "b", "c", "c"}
	brr := [5]int{80, 100, 20, 100, 90}

	data := make(map[string]int)
	count := make(map[string]int)
	for i := 0; i < len(arr); i++ {
		if _, ok := data[arr[i]]; ok {
			data[arr[i]] = data[arr[i]] + brr[i]
			count[arr[i]]++
		} else {
			count[arr[i]] = 1
			data[arr[i]] = brr[i]
		}
	}

	res := make(map[int]string)
	var ar []int
	for k, v := range count {
		if v > 1 {
			data[k] = data[k] / v
			res[data[k]] = k
		} else {
			res[data[k]] = k
		}
		ar = append(ar, data[k])
	}

	for i := 0; i < len(ar); i++ {
		for j := i + 1; j < len(ar); j++ {
			if ar[i] < ar[j] {
				ar[i], ar[j] = ar[j], ar[i]
			}
		}
	}

	fmt.Println(res[ar[0]], ar[0])
}

// 公司：流体网络
// 题目：公司有n个组，每组人数相同，>=1人，需要进行随机的组队吃饭。
// 要求：
// 1. 两两一队，不能落单，落单则三人一队
// 2. 一个人只出现一次
// 3. 队伍中至少包含两个组
// 4. 随机组队，重复执行程序得到的结果不一样
// 注：要同时满足条件1.2.3.4
// 举例：
// GroupList = [  # 小组列表
// 		['小名', '小红', '小马', '小丽', '小强'],
//		['大壮', '大力', '大1', '大2', '大3'],
//		['阿花', '阿朵', '阿蓝', '阿紫', '阿红'],
//		['A', 'B', 'C', 'D', 'E'],
//		['一', '二', '三', '四', '五'],
//		['建国', '建军', '建民', '建超', '建跃'],
//		['爱民', '爱军', '爱国', '爱辉', '爱月']
//	]
// 输入：GroupList
// 输出：(A, 小名)，（B, 小红）。

var (
	GroupList = [][]string{
		{"小名", "小红", "小马", "小丽", "小强"},
		{"大壮", "大力", "大1", "大2", "大3"},
		{"阿花", "阿朵", "阿蓝", "阿紫", "阿红"},
		{"A", "B", "C", "D", "E"},
		{"一", "二", "三", "四", "五"},
		{"建国", "建军", "建民", "建超", "建跃"},
		{"爱民", "爱军", "爱国", "爱辉", "爱月"},
	}

	result [][]string

	gwMap   = make(map[int]struct{})
	dwMap   = make(map[int]map[int]struct{})
	doneMap = make(map[int]struct{})
)

// GetGroupList 存在问题
func GetGroupList() {
	gwMap = map[int]struct{}{}
	dwMap = map[int]map[int]struct{}{}
	doneMap = map[int]struct{}{}
	result = [][]string{}
	list := doGroupList(GroupList, result)
	fmt.Println(list)
}

// doGroupList
// 组号：0 - n-1；队员号：0 - 4
// 随机一个组号 gw 和 队员号 dw
// 创建组号记录组：key为组号 value为空结构体
// gwMap := make(map[int]struct{})
// 创建组中队员的匹配情况：key为组号，value为队员号
// dwMap := make(map[int]map[int]struct{})
// 组中的人已经匹配完毕
// doneMap := make(map[int]struct{})
func doGroupList(data, result [][]string) [][]string {
	n := len(data)
	if n <= 1 {
		panic("队伍太少了")
	}

	var array []string
	g1 := 9
	i := 1
Loop2:
Loop1:
	gw := rand.Intn(n)
	if i == 1 {
		g1 = gw
	}
	if i == 2 {
		if gw == g1 {
			goto Loop1
		}
	}
	if _, ok := doneMap[gw]; ok {
		goto Loop1
	}
	if _, ok := gwMap[gw]; !ok {
		gwMap[gw] = struct{}{}
	}

Loop3:
	dw := rand.Intn(5)
	if len(dwMap[gw]) <= 0 {
		dwMap[gw] = make(map[int]struct{})
	}
	if _, ok := dwMap[gw][dw]; !ok {
		dwMap[gw][dw] = struct{}{}
	} else {
		goto Loop3
	}

	if len(dwMap[gw]) == 5 {
		doneMap[gw] = struct{}{}
	}

	array = append(array, GroupList[gw][dw])
	if i == 1 {
		i++
		goto Loop2
	}
	result = append(result, array)

	if len(doneMap) == n {
		return result
	}

	fmt.Println(result, len(result))
	return doGroupList(data, result)
}

// DoSearch 搜索切片中的某个数据
func DoSearch(slice []int) {
	ch := make(chan int, runtime.NumCPU())
	t := time.NewTimer(time.Second * 2)
	not, yes := make(chan struct{}), make(chan struct{})

	go func() {
		for _, value := range slice {
			ch <- value
		}
	}()

	for i := 0; i < runtime.NumCPU(); i++ {
		go func() {
			for {
				select {
				case <-t.C:
					close(not)
				case r := <-ch:
					if r == 345 {
						close(yes)
					}
				}
			}
		}()
	}

	select {
	case <-not:
		fmt.Println("Timeout Not Found")
		return
	case <-yes:
		fmt.Println("Found it!")
		return
	}
}

// BubbleSort 冒泡排序
func BubbleSort(arr []int) []int {
	count := len(arr)
	for i := 0; i < count; i++ {
		for j := i + 1; j < count; j++ {
			if arr[i] > arr[j] {
				arr[i], arr[j] = arr[j], arr[i]
			}
		}
	}
	return arr
}
