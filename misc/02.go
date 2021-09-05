package myProject

import "fmt"

func Arraytest() {
	var myInt int64
	fmt.Println("for loop type 01")
	for j := 0; j < 20; j++ {
		myInt += 1
	}
	fmt.Println(myInt)
}
func Arraytest01() {
	var myInt01 int32
	fmt.Println("Example for loop 02")
	i := 1
	for i <= 20 {
		myInt01 = myInt01 + 1
		fmt.Println(myInt01)
		i++
	}
}
