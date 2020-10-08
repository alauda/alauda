package term

import (
	"github.com/alauda/alauda/pkg/client"
	"github.com/alauda/alauda/pkg/terminal"
	"github.com/spf13/cobra"
)

// NewTermCmd creates a new terminal command.
func NewTermCmd(alauda client.APIClient) *cobra.Command {
	termCmd := &cobra.Command{
		Use:   "term",
		Short: "Start a terminal window for alauda",
		Long:  ``,
		Run: func(cmd *cobra.Command, args []string) {
			startTerminal()
		},
	}

	return termCmd
}

func startTerminal() error {
	term := terminal.NewTerminal()

	err := term.Start()
	if err != nil {
		return err
	}

	return nil
}
