package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/rajatjindal/kubectl-remove-in-allns/pkg/k8s"
	"github.com/spf13/cobra"
	"k8s.io/cli-runtime/pkg/genericclioptions"
	"k8s.io/client-go/kubernetes"
)

//Version is set during build time
var Version = "unknown"

//RemoveAllNamespacesOptions is struct for removing a resource from all namespaces
type RemoveAllNamespacesOptions struct {
	configFlags *genericclioptions.ConfigFlags
	IOStreams   genericclioptions.IOStreams

	args         []string
	kubeclient   *kubernetes.Clientset
	resourceName string
	resourceType string
	printVersion bool
}

// NewRemoveAllNamespacesOptions provides an instance of RemoveAllNamespacesOptions with default values
func NewRemoveAllNamespacesOptions(streams genericclioptions.IOStreams) *RemoveAllNamespacesOptions {
	return &RemoveAllNamespacesOptions{
		configFlags: genericclioptions.NewConfigFlags(true),
		IOStreams:   streams,
	}
}

// NewCmdModifySecret provides a cobra command wrapping RemoveAllNamespacesOptions
func NewCmdModifySecret(streams genericclioptions.IOStreams) *cobra.Command {
	o := NewRemoveAllNamespacesOptions(streams)

	cmd := &cobra.Command{
		Use:          "remove-allns [resource-type] [resource-name] [flags]",
		Short:        "Removes the given resource from all namespaces",
		SilenceUsage: true,
		RunE: func(c *cobra.Command, args []string) error {
			if o.printVersion {
				fmt.Println(Version)
				os.Exit(0)
			}

			if err := o.Complete(c, args); err != nil {
				return err
			}
			if err := o.Validate(); err != nil {
				return err
			}
			if err := o.Run(); err != nil {
				return err
			}

			return nil
		},
	}

	cmd.Flags().BoolVar(&o.printVersion, "version", false, "prints version of plugin")
	o.configFlags.AddFlags(cmd.Flags())

	return cmd
}

// Complete sets all information required for updating the current context
func (o *RemoveAllNamespacesOptions) Complete(cmd *cobra.Command, args []string) error {
	o.args = args

	if len(args) == 2 {
		o.resourceType = args[0]
		o.resourceName = args[1]
	}

	config, err := o.configFlags.ToRESTConfig()
	if err != nil {
		return err
	}

	o.kubeclient, err = kubernetes.NewForConfig(config)
	if err != nil {
		return err
	}

	return nil
}

// Validate ensures that all required arguments and flag values are provided
func (o *RemoveAllNamespacesOptions) Validate() error {
	if len(o.args) != 2 {
		return fmt.Errorf("exactly 2 arguments are needed")
	}

	return nil
}

// Run fetches the given secret manifest from the cluster, decodes the payload, opens an editor to make changes, and applies the modified manifest when done
func (o *RemoveAllNamespacesOptions) Run() error {
	namespaces, err := k8s.GetAllNamespaces(o.kubeclient)
	if err != nil {
		return err
	}

	switch strings.ToLower(o.resourceType) {
	case "cm", "configmap":
		err := k8s.DeleteConfigMaps(o.kubeclient, o.resourceName, namespaces)
		if err != nil {
			return err
		}
	case "secret":
		err := k8s.DeleteSecrets(o.kubeclient, o.resourceName, namespaces)
		if err != nil {
			return err
		}
	case "dep", "deployment":
		err := k8s.DeleteDeployments(o.kubeclient, o.resourceName, namespaces)
		if err != nil {
			return err
		}
	default:
		return fmt.Errorf("unsupported resource type %q", o.resourceType, namespaces)
	}

	return nil
}
