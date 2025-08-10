package main

import (
	"encoding/json"
	"fmt"

	"github.com/JustinRoc/pkg/util"
)

type A struct {
	Name string
}

func main() {
	a := A{
		Name: "a",
	}
	b, err := json.Marshal(a)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(b))
	a2, err := util.FromJSON[A]([]byte(b))
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(a2.Name)
}
