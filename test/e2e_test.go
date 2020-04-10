package main

import (
	"fmt"
	"os"
	"testing"

	"github.com/alauda/alauda/pkg/client"
	"github.com/alauda/alauda/pkg/cmd/alauda"
)

var cliTests = []struct {
	args []string
}{
	{[]string{"alauda", "version"}},
	{[]string{"alauda", "kubectl", "get", "ns", "--insecure-skip-tls-verify=true"}},
	{[]string{"alauda", "k", "get", "pod", "--insecure-skip-tls-verify=true"}},
}

func TestCli(t *testing.T) {
	alaudaClient := client.NewClient()
	alaudaClient.Initialize()

	alaudaCmd := alauda.NewAlaudaCmd(alaudaClient)

	for _, tt := range cliTests {
		os.Args = tt.args
		fmt.Println(os.Args)

		err := alaudaCmd.Execute()
		if err != nil {
			t.Error(err)
		}
	}
}
