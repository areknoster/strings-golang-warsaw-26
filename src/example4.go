package main

import "fmt"

func main(){
	const Hello = "Hello, Мир"
	type myType string
	var k myType = Hello
	fmt.Println(k)

	// WYSIWYG literals - backquotes!
	backQuoteStr := `This string
	will have
	tabs in it`
	fmt.Println(backQuoteStr)
}
