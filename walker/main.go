package main

import "fmt"

type Walker interface {
	Walk() string
}

type Dog string

func (d Dog) Walk() string {
	return "Dog walking"
}

func main() {
	var dog Dog
	callWalker(dog)
}

func callWalker(w Walker) {
	fmt.Println(w.Walk())
}
