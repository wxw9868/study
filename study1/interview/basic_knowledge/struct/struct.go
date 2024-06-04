package struc

import (
	"encoding/json"
	"fmt"
)

type Human struct {
	age  int
	name string
	sex  string
}

func DoStruct() {
	msg := Human{age: 25, name: "卫雄伟", sex: "男"}
	humanBytes, err := json.Marshal(msg)
	if err != nil {
		fmt.Println("json序列化失败！")
	}

	hu := &Human{}
	err = json.Unmarshal(humanBytes, &hu)
	if err != nil {
		fmt.Println("json反序列化失败！")
	} else {
		fmt.Println("json反序列化成功。")
	}
}

type Girl struct {
	Name       string `json:"name"`
	DressColor string `json:"dress_color"`
}

func (g Girl) SetColor(color string) {
	g.DressColor = color
	fmt.Printf("set &g: %p\n", &g)
}

func (g Girl) JSON() string {
	fmt.Printf("JSON &g: %p\n", &g)
	data, _ := json.Marshal(&g)
	return string(data)
}

func DoSet() {
	g := Girl{Name: "menglu"}
	fmt.Printf("&g: %p\n", &g)
	g.SetColor("white")
	fmt.Println(g.JSON())
}

type People struct {
	Name string
	Age  int
}

func StudyStruct() {
	i := new(int)
	fmt.Println("int: ", i)
	people := new(People)
	fmt.Println("struct：", people)
}
