package hellomod

import "fmt"

// SayHello say hello to the world
func SayHello() {
	fmt.Println("Hellomod world! Version v0.8.0")
}

func NewSayHello() {
	fmt.Println("New Hellomod world!!! Version v0.8.0")
}
