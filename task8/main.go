package main

import (
	"fmt"
	"strconv"
)

func SetZero(v int64, k uint) int64 {
	return v & ^(1 << int64(k))
}

func SetOne(v int64, k uint) int64 {
	return v | (1 << int64(k))
}

func main() {
	var n int64 = 61
	fmt.Println(strconv.FormatInt(n, 2))
	fmt.Println(strconv.FormatInt(SetZero(n, 5), 2))
	fmt.Println(strconv.FormatInt(SetOne(n, 1), 2))
}