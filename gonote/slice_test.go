package gonote

import (
	"fmt"
	"testing"
	"unsafe"
)

func TestSliceMem(t *testing.T) {
	slice := "hellohhhhhhh"

	str := slice[2:3]

	fmt.Println("str size:", unsafe.Sizeof(str))
	fmt.Println("slice size:", unsafe.Sizeof(slice))

}
