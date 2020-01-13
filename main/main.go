package main

import (
	"belajargolang"
	"fmt"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(2)
	num := 23
	nam := "Haloo"
	go belajargolang.ReflectInt(num)

	belajargolang.ReflectString(nam)

	var input string
	fmt.Scanln(&input)
}
