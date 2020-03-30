package alauda

import (
	"fmt"

	"github.com/spf13/cobra"
)

const version = "0.0.1"

func newVersionCmd() *cobra.Command {
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
