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

// 遍历文件夹
func TraverseFolders() {
	path := "/Users/v_weixiongwei/go/src/video/assets/image/poster"
	recursive(path)
}

func recursive(dir string) {
	var data = make(map[string]struct{})
	var actress = make(map[string]struct{})

	files, err := os.ReadDir(dir)
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		if file.IsDir() {
			recursive(dir + "/" + file.Name())
		}
		name := file.Name()
		if filepath.Ext(name) == ".jpg" {
			arr := strings.Split(strings.Split(name, ".")[0], "_")
			actress[arr[len(arr)-1]] = struct{}{}
			if len(arr[len(arr)-2]) < 20 && len(arr[len(arr)-2]) > 6 {
				_, ok := actress[arr[len(arr)-2]]
				if !ok {
					data[arr[len(arr)-2]] = struct{}{}
				}
			}
		}
	}
	// WriteMapToFile("actress.json", &actress)
	// WriteMapToFile("data.json", &data)
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
