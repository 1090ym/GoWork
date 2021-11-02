package flow

import (
	"fmt"
	"strconv"
	"testing"
	"time"
)

func TestMapToChannel(t *testing.T) {
	values := make([]chan string, 6)
	nums := []int{0, 1, 2, 3, 4, 5}
	for i, _ := range nums {
		values[i] = make(chan string, 3)
	}
	for i := 0; i < len(nums); i++ {
		go WriteMap(values[i], i)
	}

	time.Sleep(5 * time.Second)

	for key, value := range values {
		fmt.Println("key:", key)
		//v :=  <- value
		//fmt.Println("value", v)
		for v := range value {
			v1 := <-value
			fmt.Print("v:", v, " ", "v:", v1, " ")
		}
		fmt.Println()
	}
}

func WriteMap(values chan string, i int) {
	//fmt.Println("i:", strconv.Itoa(i))
	str := strconv.Itoa(i)
	values <- str + str + str
	//fmt.Println("len(values)",len(values))

}
