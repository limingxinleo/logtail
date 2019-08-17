package handler

import "fmt"

type Stdout struct {
}

func (Stdout) Handle(text string) {
	fmt.Println(text)
}
