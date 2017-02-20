package main

import "fmt"
import "../echo"

func main() {
	fmt.Println(Hello())
	echo.Run()
}

func Hello() string {
	return "Hello"
}
