package main

import "fmt"

// Глобальная переменная будет изменяться при каждом вызове 
// someFunc, тяжело будет отследить состояние строки
var justString string

// Варианты улучшения - вынести переменную в main, передавать пропсом в someFunc
// func main() {
//   var justString string
// 	 someFunc(&justString)
// }

// вариант два - создавать строку в самой функции 

func createHugeString(ln int) (res string) {
	for range ln {
		res += "a"
	}
	return
}

func someFunc() {
  v := createHugeString(1 << 10)
  justString = v[:100]
	fmt.Printf(justString)
}

func main() {
  someFunc()
}