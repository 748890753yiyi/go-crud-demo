package main

import "fmt"


type A struct {
	valA int
}

type InfA interface {
	greeting()
}

func (a *A) greeting() {
	fmt.Println("Hello, A")
}

type B struct {
	*A     // or A?
	valB int
}

func main() {
	a := new(A)
	b := new(B)
	a.greeting()
	b.greeting()
}