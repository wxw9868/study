package filehandling

import (
	"testing"

	"github.com/wxw9868/study/utils"
)

func TestMain(t *testing.T) {
	// 1 249 2 248
	// utils.MethodRuntime(func() {
	// 	absolutePath, err := filepath.Abs("actress.json")
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// 	var actress = make(map[string]struct{})
	// 	ReadFileToMap(absolutePath, &actress)
	// 	var slice []string
	// 	for k, _ := range actress {
	// 		slice = append(slice, k)
	// 	}
	// 	sort.Strings(slice)
	// 	fmt.Printf("%d, %+v\n", len(slice), slice)
	// })

	// utils.MethodRuntime(TraverseFolders)
	utils.MethodRuntime(MergeData)

	// utils.MethodRuntime(ReadWriteFile)
}
