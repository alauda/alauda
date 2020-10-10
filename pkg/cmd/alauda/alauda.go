package alauda

import (
	"github.com/alauda/alauda/pkg/client"
	"github.com/alauda/alauda/pkg/cmd/kubectl"
	"github.com/alauda/alauda/pkg/cmd/term"
	"github.com/spf13/cobra"
)

// NewAlaudaCmd creates a new root command for the Alauda CLI.
func NewAlaudaCmd(alauda client.APIClient) *cobra.Command {
	alaudaCmd := &cobra.Command{
		Use:   "alauda",
		Short: "Alauda CLI",
		Long:  ``,
	}

	addCommands(alaudaCmd, alauda)

	return alaudaCmd
}

func addCommands(cmd *cobra.Command, alauda client.APIClient) {
	cmd.AddCommand(
		newVersionCmd(alauda),

		kubectl.NewKubectlCmd(alauda),
		term.NewTermCmd(alauda),
	)
}
