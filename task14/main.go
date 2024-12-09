package main

import (
	"fmt"
	"reflect"
)

func do(i interface{}) {
	switch i.(type) {
	case int:
		fmt.Printf("int")
	case string:
		fmt.Printf("string")
	case bool:
		fmt.Printf("bool")
	case chan int:
		fmt.Printf("channel")
	default:
		fmt.Printf("???")
	}
}

func do2(i interface{}) {
	v := reflect.ValueOf(i)
	switch v.Kind() {
	case reflect.Int:
		fmt.Printf("int")
	case reflect.String:
		fmt.Printf("string")
	case reflect.Bool:
		fmt.Printf("bool")
	case reflect.Chan:
		fmt.Printf("channel")
	default:
		fmt.Printf("???")
	}
}

func main() {
	do(21)
	do("hello")
	do(true)
}