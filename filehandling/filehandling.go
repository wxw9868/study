package filehandling

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func MergeData() {
	// var actress1 = make(map[string]struct{})
	// var actress2 = make(map[string]struct{})
	// ReadFileToMap("actress.json", &actress1)
	// ReadFileToMap("new.json", &actress2)
	// for k, _ := range actress2 {
	// 	if _, ok := actress1[k]; !ok {
	// 		actress1[k] = struct{}{}
	// 		delete(actress2, k)
	// 	}
	// }
	// WriteMapToFile("MergeData.json", &actress1)

	// fmt.Printf("%+v\n", actress1)
	// fmt.Printf("%+v\n", actress2)

}

// 遍历文件夹
func TraverseFolders() {
	path := "/Users/v_weixiongwei/go/src/video/assets/image/poster"
	recursive(path)
}

func recursive(dir string) {
	var actress = make(map[string]struct{})
	var data = make(map[string]struct{})

	files, err := os.ReadDir(dir)
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		if file.IsDir() {
			recursive(filepath.Join(dir, "/", file.Name()))
		}
		filename := file.Name()
		i := 4
		if filepath.Ext(filename) == ".jpg" {
			arr := strings.Split(strings.Split(filename, ".")[0], "_")
			actress[arr[len(arr)-1]] = struct{}{}
			if len(arr) > i {
				if len(arr[len(arr)-i]) < 20 && len(arr[len(arr)-i]) > 6 {
					_, ok := actress[arr[len(arr)-i]]
					if !ok {
						data[arr[len(arr)-i]] = struct{}{}
					}
				}
			}
		}
	}
	WriteMapToFile("actress.json", &actress)
	WriteMapToFile("data.json", &data)
}

// 读取文件数据到 map
func ReadFileToMap(name string, v any) error {
	bytes, err := os.ReadFile(name)
	if err != nil {
		return err
	}
	if err = json.Unmarshal(bytes, &v); err != nil {
		return err
	}
	return nil
}

// 读取 map 数据到文件
func WriteMapToFile(name string, v any) error {
	bytes, err := json.Marshal(v)
	if err != nil {
		return err
	}
	err = os.WriteFile(name, bytes, 0666)
	if err != nil {
		return err
	}
	return nil
}

// ReadWriteFile 读写文件
func ReadWriteFile() {
	file, err := os.Open("query.txt") // For read access.
	if err != nil {
		log.Fatal(err)
	}

	var buffer bytes.Buffer
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		buffer.WriteString("'")
		buffer.WriteString(scanner.Text())
		buffer.WriteString("'")
		buffer.WriteString(",")
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}

	if err = os.WriteFile("result.txt", buffer.Bytes(), os.ModePerm); err != nil {
		log.Fatal(err)
	}
}

// MoveFile 移动文件并修改文件名称
func MoveFile(oldPath, newPath string) error {
	if err := os.Rename(oldPath, newPath); err != nil {
		log.Fatalf("%v", err)
		return err
	}

	// if runtime.GOOS == "windows" { //跨卷移动
	// 	from, err := syscall.UTF16PtrFromString(oldPath)
	// 	if err != nil {
	// 		log.Fatalf("%v", err)
	// 		return err
	// 	}
	// 	to, err := syscall.UTF16PtrFromString(newPath)
	// 	if err != nil {
	// 		log.Fatalf("%v", err)
	// 		return err
	// 	}
	// 	//windows API
	// 	if err = syscall.MoveFile(from, to); err != nil {
	// 		log.Fatalf("%v", err)
	// 		return err
	// 	}
	// } else {
	// 	if err := os.Rename(oldPath, newPath); err != nil {
	// 		log.Fatalf("%v", err)
	// 		return err
	// 	}
	// }
	return nil
}

// const (
// 	FILE_PATH = "D:\\浏览器下载"
// 	MOVE_PATH = "K:\\Git项目备份"
// )

// var lists = map[string]string{
// 	"bbs":                  "\\BBS",
// 	"bbsMas":               "\\BBS\\VUE\\后台",
// 	"bbsMan":               "\\BBS\\VUE\\前台",
// 	"community":            "\\community",
// 	"crm":                  "\\CRM",
// 	"VIM":                  "\\VIM",
// 	"hyjr":                 "\\会员金融",
// 	"cmf":                  "\\企业联盟",
// 	"business_card":        "\\企业联盟\\名片推广",
// 	"wenku":                "\\文库",
// 	"java":                 "\\小程序",
// 	"xiaoochengxu":         "\\小程序",
// 	"shop":                 "\\艺券商城",
// 	"yipiao_mail":          "\\艺券商城",
// 	"bank":                 "\\银行",
// 	"bank_boot":            "\\银行\\boot",
// 	"bank_bus":             "\\银行\\bus",
// 	"bank_client":          "\\银行\\client",
// 	"bank_com":             "\\银行\\com",
// 	"bank_cus":             "\\银行\\cus",
// 	"bank_cus-master-vue":  "\\银行\\cus\\VUE",
// 	"bank_dubbo":           "\\银行\\dubbo",
// 	"bank_loan":            "\\银行\\loan",
// 	"bank_OA":              "\\银行\\OA",
// 	"bank_customer_system": "\\银行\\customer_system",
// 	"banniban_bbs":         "\\帮你办社区",
// 	"stock":                "\\股票",
// }

// // 遍历文件夹中的文件
// func ReadDir() error {
// 	list, err := ioutil.ReadDir(FILE_PATH)
// 	if err != nil {
// 		log.Fatalf("%v", err)
// 	}
// 	for _, fi := range list {
// 		if !fi.IsDir() {
// 			filename := fi.Name()
// 			for k, v := range lists {
// 				if strings.Contains(filename, k) {
// 					s := strings.LastIndex(filename, "-")
// 					e := strings.LastIndex(filename, ".")
// 					content := filename[s+1 : e]
// 					date := time.Now().Format("2006.01.02")
// 					rename := strings.Replace(filename, content, date, -1)
// 					oldPath := FILE_PATH + "\\" + filename
// 					newPath := MOVE_PATH + v + "\\" + rename
// 					if err = moveFile(oldPath, newPath); err != nil {
// 						log.Fatalf("%v", err)
// 						return err
// 					}
// 					fmt.Println(oldPath + " >> " + newPath)
// 				}
// 			}
// 		}
// 	}
// 	return nil
// }
