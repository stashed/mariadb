/*
Copyright AppsCode Inc. and Contributors

Licensed under the AppsCode Free Trial License 1.0.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    https://github.com/appscode/licenses/raw/1.0.0/AppsCode-Free-Trial-1.0.0.md

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package pkg

import (
	"stash.appscode.dev/apimachinery/client/clientset/versioned/scheme"

	"github.com/spf13/cobra"
	v "gomodules.xyz/x/version"
	clientsetscheme "k8s.io/client-go/kubernetes/scheme"
)

var SupportedProducts = []string{"stash-enterprise", "kubedb-ext-stash"}

var licenseApiService string

func NewRootCmd() *cobra.Command {
	rootCmd := &cobra.Command{
		Use:               "stash-mariadb",
		Short:             `MariaDB backup & restore plugin for Stash by AppsCode`,
		Long:              `MariaDB backup & restore plugin for Stash by AppsCode. For more information, visit here: https://appscode.com/products/stash`,
		DisableAutoGenTag: true,
		PersistentPreRunE: func(c *cobra.Command, args []string) error {
			return scheme.AddToScheme(clientsetscheme.Scheme)
		},
	}
	rootCmd.PersistentFlags().StringVar(&licenseApiService, "license-apiservice", "", "Name of ApiService used to expose License endpoint")

	rootCmd.AddCommand(v.NewCmdVersion())
	rootCmd.AddCommand(NewCmdBackup())
	rootCmd.AddCommand(NewCmdRestore())

	return rootCmd
}
