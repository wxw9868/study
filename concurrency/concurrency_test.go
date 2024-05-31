package concurrency

import (
	"study/utils"
	"testing"
)

func TestMain(t *testing.T) {
	utils.MethodRuntime(CheckTheNumberOfGoroutines)
	// utils.MethodRuntime(UseChannelsToControlOutputOrder)
	// utils.MethodRuntime(TimeoutExit)
	// utils.MethodRuntime(ProduceAndConsume)
	// utils.MethodRuntime(ManuallyExitTask)
	// utils.MethodRuntime(ReusingGoroutine)
	// utils.MethodRuntime(UsingMutexLock)
	// utils.MethodRuntime(UsingMutexesToResolveDataRaces)
	// utils.MethodRuntime(AtomicOperations)
}
