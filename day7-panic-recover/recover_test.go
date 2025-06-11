package main

import (
	"fmt"
	"testing"
)

func Test_recover(t *testing.T) {
	defer func() {
		fmt.Println("defer func")
		if err := recover(); err != nil {
			fmt.Println("recover success")
		}
	}()

	arr := []int{1, 2, 3}
	fmt.Println(arr[4])
	fmt.Println("after panic")

}
func Test_main(t *testing.T) {
	Test_recover(t)
	fmt.Print("after recover")
}
