package main

import (
	"ecommerce/cmd"

)

func makeChannel(ch chan string) {
	ch <- "Hello, World!"
}

func main() {

	ch := make(chan string)
	go makeChannel(ch)

	name := <-ch
	println(name)

	servego.ServeGo()
}


// what is channel

// what is peralalizim

// what is context switching

// what is concurrency

// in go explaing all those