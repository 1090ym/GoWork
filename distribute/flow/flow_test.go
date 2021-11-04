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
	row := make([]interface{}, 1024)
	for i := 0; i < 1024; i++ {
		row[i] = i
	}

	// 添加work函数到每个step
	DisManager.AddStepToFlow(Step1Func)

}

// 需要传入到step的封装后的函数
func Step1Func(shard InputDataShard) {
	ptr := uintptr((*emptyInterface)(unsafe.Pointer(&shard.data)).word)
	input := *(*[]int)(unsafe.Pointer(ptr))
	Power(input)
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
