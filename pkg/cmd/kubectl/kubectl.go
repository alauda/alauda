package kubectl

import (
	"github.com/alauda/alauda/pkg/client"
	"github.com/spf13/cobra"
	"k8s.io/kubernetes/pkg/kubectl/cmd"
)

// NewKubectlCmd creates a new embedded kubectl command.
func NewKubectlCmd(alauda client.APIClient) *cobra.Command {
	kubectlCmd := cmd.NewDefaultKubectlCommand()

	kubectlCmd.Aliases = []string{"k", "k8s"}

	return kubectlCmd
}
