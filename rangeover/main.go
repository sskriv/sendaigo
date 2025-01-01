package main

import (
	"flag"
	"fmt"
	"iter"
)

var (
	sep = "-----------------------------"
)

func main() {
	fmt.Println("Hello, world!!")
	flag.Parse()
	switch flag.Arg(0) {
	case "chan":
		rangeChannel()
	case "int":
		rangeInteger()
	case "func":
		rangeFunction()
	case "iter":
		iterator()
	}
}

func rangeChannel() {
	fmt.Println("range over channel")

	c := make(chan int, 5)
	go rc(cap(c), c)
	for v := range c { // channelがcloseするまで
		fmt.Println(v)
	}
}

func rc(cap int, c chan int) {
	for i := 0; i < cap; i++ {
		c <- i
	}
	close(c)
}

func rangeInteger() {
	fmt.Println("for loop")
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	fmt.Println(sep)

	fmt.Println("range over integer")
	for i := range 5 {
		fmt.Println(i)
	}
}

func rangeFunction() {
	fmt.Println("range over function")

	for v := range f { // range over function
		fmt.Println(v)
	}
}

func f(yield func(int) bool) {
	for n := range 5 {
		if !yield(n) {
			return
		}
	}
}

func iterator() {
	fmt.Println("iterator")

	s := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for i := range double(s) {
		if i > 10 {
			break
		}
		fmt.Println(i)
	}
}
func double(s []int) iter.Seq[int] {
	return func(yield func(int) bool) {
		for _, v := range s {
			if !yield(v * 2) { // yieldの返り値boolをハンドリングする
				return
			}

			// yield(v * 2) //panicになる可能性があるので注意
			// panic: runtime error: range function continued iteration after function for loop body returned false
		}
	}
}
