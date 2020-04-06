package kubectl

import (
	"github.com/spf13/cobra"
	"k8s.io/kubernetes/pkg/kubectl/cmd"
)

// NewKubectlCmd creates a new embedded kubectl command.
func NewKubectlCmd() *cobra.Command {
	kubectlCmd := cmd.NewDefaultKubectlCommand()

	kubectlCmd.Aliases = []string{"k", "k8s"}

	return kubectlCmd
}
