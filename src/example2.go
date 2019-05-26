package main

import (
	"fmt"
	"reflect"
)

func main(){
	var str string =  "Hello"
	fmt.Println("string's element type:",reflect.TypeOf(str[0]))
	ruStr := "привет"
	fmt.Printf("% x\n", str)
	fmt.Printf("% x\n", ruStr)
	var b []byte = []byte(ruStr)
	fmt.Println(b)
	ruStr = string(b)
	fmt.Println(ruStr)
	//str[0] = 0x01;



}
