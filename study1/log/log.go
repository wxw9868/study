package main

import (
	"fmt"
	"study/log/adm"
	"study/log/da"
	"study/log/db"
)

func main() {
	da.Print()
	db.Print()
	fmt.Println(adm.Get())
	//for {
	//	time.Sleep(time.Second)
	//}
}
