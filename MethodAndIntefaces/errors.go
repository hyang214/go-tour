package main

import (
	"fmt"
	"strconv"
)

func Atoi(s string) (int, error) {
	i, err := strconv.Atoi(s)
	if err != nil {
		fmt.Printf("couldn't convert number: %v\n", err)
		return 0, err
	}
	fmt.Println("Converted integer:", i)
	return i, err
}

func main() {
	i1, err1 := Atoi("a")
	if err1 != nil {
		fmt.Printf("error  %v\n", err1)
	}
	fmt.Println(i1)

	i2, err2 := Atoi("43")
	if err2 != nil {
		fmt.Printf("error  %v\n", err2)
	}
	fmt.Println(i2)
}
