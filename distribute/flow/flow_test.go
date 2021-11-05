package flow

import (
	"fmt"
	"testing"
	"unsafe"
)

// 测试分布式计算每个数的平方
func TestLocalDistribute(t *testing.T) {
	// 初始化虚拟节点
	DisManager.InitDistributor([]int{1, 2, 3, 4})
	// 设置数据
	row := make([]interface{}, 10)
	for i := 0; i < 10; i++ {
		row[i] = i
	}

	// 添加work函数到每个step
	DisManager.AddStepToFlow(Step1Func)
	DisManager.ReceiveRpcMsg(row, 1, 0)
	DisManager.ReceiveRpcMsg(row, 2, 0)
	DisManager.ReceiveRpcMsg(row, 3, 0)
	DisManager.ReceiveRpcMsg(row, 4, 0)
}

// 需要传入到step的封装后的函数
func Step1Func(input InputDataShard) {
	fmt.Println("Func InputData:", input)
	ptr := uintptr((*emptyInterface)(unsafe.Pointer(&input.data)).word)
	data := *(*[]int)(unsafe.Pointer(ptr))
	Power(data)
}

// work函数
func Power(numbers []int) {
	for _, number := range numbers {
		fmt.Println(number, ":", number*number)
	}
}

// 测试work
func TestPower(t *testing.T) {
	Power([]int{1, 2, 3, 4})
}
