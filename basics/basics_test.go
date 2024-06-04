package basics

import (
	"github.com/wxw9868/study/utils"
	"testing"
)

func TestMain(t *testing.T) {
	utils.MethodRuntime(Basics)
	utils.MethodRuntime(StructToJson)
}

func TestStudySlice(t *testing.T) {
	StudySlice()
}

func TestStudyCopy(t *testing.T) {
	StudyCopy()
}

func TestStudyMap(t *testing.T) {
	StudyMap()
}

func TestStudyChannel(t *testing.T) {
	StudyChannel()
}

func TestStudyReturnAndDefer(t *testing.T) {
	StudyReturnAndDefer()
}
