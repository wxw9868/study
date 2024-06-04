package concurrent

import (
	"fmt"
	"testing"
	"time"
)

func TestRunFolder(t *testing.T) {
	RunFolder()
}

func TestTraverseFolder(t *testing.T) {
	start := time.Now()
	TraverseFolder("C:\\WeGameApps")
	r := time.Since(start)
	fmt.Println("r: ", r)

	fmt.Printf("len: %d, cap: %d, r: %v\n", len(paths1), cap(paths1), paths1)
}

func TestFolder(t *testing.T) {
	start := time.Now()
	Folder("C:\\WeGameApps")
	r := time.Since(start)
	fmt.Println("r: ", r)

	fmt.Printf("len: %d, cap: %d, r: %v\n", len(paths), cap(paths), paths)
	//bytes, _ := json.Marshal(paths)
	//ioutil.WriteFile("paths.txt", bytes, 666)
	//for k, path := range paths {
	//	fmt.Println(k, path)
	//}
}
