package alauda

import (
	"github.com/spf13/cobra"
)

// NewAlaudaCmd creates a new root command for the Alauda CLI.
func NewAlaudaCmd() *cobra.Command {
	alaudaCmd := &cobra.Command{
		Use:   "alauda",
		Short: "Alauda CLI",
		Long:  ``,
	}

	addCommands(alaudaCmd)

	return alaudaCmd
}

func addCommands(cmd *cobra.Command) {
	cmd.AddCommand(
		newVersionCmd(),
	)
}
