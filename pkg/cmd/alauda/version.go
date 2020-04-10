package alauda

import (
	"fmt"

	"github.com/alauda/alauda/pkg/client"
	"github.com/spf13/cobra"
)

const version = "0.0.1"

func newVersionCmd(alauda client.APIClient) *cobra.Command {
	versionCmd := &cobra.Command{
		Use:   "version",
		Short: "Display version of Alauda CLI",
		Long:  ``,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("[alauda]", version)
		},
	}

	return versionCmd
}
