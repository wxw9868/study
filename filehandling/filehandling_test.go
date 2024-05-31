package filehandling

import (
	"fmt"
	"log"
	"path/filepath"
	"sort"
	"study/utils"
	"testing"
)

func TestMain(t *testing.T) {
	utils.MethodRuntime(func() {
		absolutePath, err := filepath.Abs("actress.json")
		if err != nil {
			log.Fatal(err)
		}
		var actress = make(map[string]struct{})
		ReadFileToMap(absolutePath, &actress)
		var slice []string
		for k, _ := range actress {
			slice = append(slice, k)
		}
		sort.Strings(slice)
		fmt.Printf("%d, %+v\n", len(slice), slice)
	})

	utils.MethodRuntime(TraverseFolders)
	utils.MethodRuntime(ReadWriteFile)
}
