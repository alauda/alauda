package main

import (
	"github.com/alauda/alauda/pkg/client"
	"github.com/alauda/alauda/pkg/cmd/alauda"
)

func main() {
	alaudaClient := client.NewClient()
	alaudaClient.Initialize()

	alaudaCmd := alauda.NewAlaudaCmd(alaudaClient)
	alaudaCmd.Execute()
}
