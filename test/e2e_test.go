package main

import (
	"fmt"
	"os"
	"testing"

	"github.com/alauda/alauda/pkg/cmd/alauda"
)

var cliTests = []struct {
	args []string
}{
	{[]string{"alauda", "version"}},
}

func TestCli(t *testing.T) {

	alaudaCmd := alauda.NewAlaudaCmd()

	for _, tt := range cliTests {
		os.Args = tt.args
		fmt.Println(os.Args)

		err := alaudaCmd.Execute()
		if err != nil {
			t.Error(err)
		}
	}
}
