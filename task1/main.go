package main

import "fmt"

type Human struct {
	Height int
	Weight int
}

func (h *Human) String() {
	fmt.Println("Height", h.Height)
	fmt.Println("Weight", h.Weight)
}

type Action struct {
	Human
}

func (a *Action) Jump() {
	fmt.Println("Jump")
	a.Weight -= 1
}

func main() {
	a := Action{Human{Height: 5, Weight: 10}}
	a.String()
	a.Jump()
	a.String()
}
