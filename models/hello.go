package models

import "fmt"

type HelloModel struct {
	Name string
}

func (h HelloModel) SayHello() string {
	message := fmt.Sprintf("Hello %s", h.Name)

	return message
}
