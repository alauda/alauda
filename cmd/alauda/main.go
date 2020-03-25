package main

import (
	"github.com/alauda/alauda/pkg/cmd/alauda"
)

func main() {
	alaudaCmd := alauda.NewAlaudaCmd()
	alaudaCmd.Execute()
}
