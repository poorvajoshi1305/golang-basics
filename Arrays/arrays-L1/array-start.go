package main

import "fmt"

func main() {
	// var nums [3]int;
	var nums [3]int
	fmt.Printf("%v\n", nums)
	fmt.Printf("%#v\n", nums)

	var arr = [5]int{1, 2, 3}
	fmt.Printf("%v\n", arr)

	var names = [2]string{"a", "b"}
	fmt.Printf("%v\n", names)

	//elipsis operator
	var n = [...]int{1, 2}
	fmt.Printf("%v\n", n)
}
