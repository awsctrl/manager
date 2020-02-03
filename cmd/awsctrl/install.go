/*
Copyright © 2019 AWS Controller authors

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Package main contains all the necessary information to run the controller
package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"sigs.k8s.io/kustomize/api/filesys"
	"sigs.k8s.io/kustomize/api/krusty"
	"sigs.k8s.io/kustomize/api/resmap"

	_ "k8s.io/client-go/plugin/pkg/client/auth/gcp"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/log/zap"
)

var (
	installLog = ctrl.Log.WithName("setup")
	path       string
)

// installCmd represents the start command
var installCmd = &cobra.Command{
	Use:   "install",
	Short: "install will generate the AWS Controller server manifests",
	Long: `AWS Controller install will generate the manifests necessary for installing into
your cluster.

$ awsctrl install manager

To install this into your cluster you can pipe this into kubectl.

$ awsctrl inatall manager | kubectl apply -f -`,
	// Run: installCommand,
}

var installManagerCmd = &cobra.Command{
	Use:   "manager",
	Short: "manager will generate the AWS Controller manager manifests",
	Long: `Install manager will generate the manifests necessary for installing into your cluster.

$ awsctrl install manager

To install this into your cluster you can pipe this into kubectl.

$ awsctrl inatall manager | kubectl apply -f -`,
	Run: func(cmd *cobra.Command, args []string) {
		ctrl.SetLogger(zap.Logger(true))

		options := krusty.MakeDefaultOptions()

		fSys := filesys.MakeFsOnDisk()
		m, err := runKustomize(path, fSys, options)
		if err != nil {
			installLog.Error(errors.New(""), "unable to init installer")
			os.Exit(1)
		}

		y, err := m.AsYaml()
		if err != nil {
			installLog.Error(errors.New(""), "unable to convert YAML")
			os.Exit(1)
		}

		fmt.Print(string(y))
	},
}

var installConfigCmd = &cobra.Command{
	Use:   "config",
	Short: "config will generate the AWS Controller manager self.awsctrl.io/Config CR",
	Long: `Install config will generate the config for your cluster. This uses the aws cli to get your AWS 
Account ID

$ awsctrl install config

To install this into your cluster you can pipe this into kubectl.

$ awsctrl inatall config | kubectl apply -f -`,
	Run: func(cmd *cobra.Command, args []string) {
		ctrl.SetLogger(zap.Logger(true))

		options := krusty.MakeDefaultOptions()

		fSys := filesys.MakeFsOnDisk()
		m, err := runKustomize(path, fSys, options)
		if err != nil {
			installLog.Error(errors.New(""), "unable to init installer")
			os.Exit(1)
		}

		y, err := m.AsYaml()
		if err != nil {
			installLog.Error(errors.New(""), "unable to convert YAML")
			os.Exit(1)
		}

		fmt.Print(string(y))
	},
}

func init() {
	installManagerCmd.Flags().StringVarP(&path, "config-dir", "c", "https://github.com/awsctrl/manager/config/default", "Path to the AWS Controller Manager kustomize configs.")

	installConfigCmd.Flags().StringVarP(&path, "config-dir", "c", "https://github.com/awsctrl/manager/config/self", "Path to the AWS Controller Self Config kustomize configs.")

	installCmd.AddCommand(installManagerCmd)
	installCmd.AddCommand(installConfigCmd)
	rootCmd.AddCommand(installCmd)
}

func runKustomize(directory string, fSys filesys.FileSystem, options *krusty.Options) (resmap.ResMap, error) {
	k := krusty.MakeKustomizer(fSys, options)
	return k.Run(directory)
}
