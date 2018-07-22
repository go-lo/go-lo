package main

import (
	"fmt"

	"github.com/satori/go.uuid"
)

func main() {
	s := uuid.NewV5(uuid.NamespaceDNS, "loadtests-are-great.example.com").String()

	fmt.Println(s)
}
