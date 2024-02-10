package main

import (
	"fmt"
	"runtime"
	"strings"
	"unicode/utf8"
	"unsafe"
)

var justString string

func someFunc() {
	v := strings.Repeat("♬", 100000)
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Printf("Alloc = %v\n", m.Alloc)
	fmt.Println(utf8.RuneCountInString(v[:10]))
	// fmt.Println(v)
	a := make([]rune, utf8.RuneCountInString(v[:10]))
	fmt.Println(unsafe.StringData(v))
	// copy(a, []rune(v[:10]))
	justString = string(a)
}

func main() {
	a := "\U0010FFFF"
	fmt.Println(a)
}
