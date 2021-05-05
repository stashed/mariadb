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
	"fmt"
	"path/filepath"

	stash "stash.appscode.dev/apimachinery/client/clientset/versioned"
	"stash.appscode.dev/apimachinery/pkg/restic"

	"github.com/codeskyblue/go-sh"
	core "k8s.io/api/core/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/klog/v2"
	"kmodules.xyz/custom-resources/apis/appcatalog/v1alpha1"
	appcatalog_cs "kmodules.xyz/custom-resources/client/clientset/versioned"
)

const (
	MariaDBUser        = "username"
	MariaDBPassword    = "password"
	MariaDBDumpFile    = "dumpfile.sql"
	MariaDBDumpCMD     = "mysqldump"
	MariaDBRestoreCMD  = "mysql"
	EnvMariaDBPassword = "MYSQL_PWD"
)

type mariadbOptions struct {
	kubeClient    kubernetes.Interface
	stashClient   stash.Interface
	catalogClient appcatalog_cs.Interface

	namespace         string
	backupSessionName string
	appBindingName    string
	myArgs            string
	waitTimeout       int32
	outputDir         string

	setupOptions  restic.SetupOptions
	backupOptions restic.BackupOptions
	dumpOptions   restic.DumpOptions
}

func (opt *mariadbOptions) waitForDBReady(appBinding *v1alpha1.AppBinding, secret *core.Secret) error {
	klog.Infoln("Waiting for the database to be ready.....")
	shell := sh.NewSession()
	shell.SetEnv(EnvMariaDBPassword, string(secret.Data[MariaDBPassword]))
	args := []interface{}{
		"ping",
		"--host", appBinding.Spec.ClientConfig.Service.Name,
		"--user", string(secret.Data[MariaDBUser]),
	}
	if appBinding.Spec.ClientConfig.Service.Port != 0 {
		args = append(args, fmt.Sprintf("--port=%d", appBinding.Spec.ClientConfig.Service.Port))
	}

	if appBinding.Spec.ClientConfig.CABundle != nil {
		args = append(args, fmt.Sprintf("--ssl-ca=%v", filepath.Join(opt.setupOptions.ScratchDir, MariaDBTLSRootCA)))
	}

	return shell.Command("mysqladmin", args...).Run()
}
