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

// Package cmd contains all the necessary information to run the controller
package cmd

import (
	"os"
	"time"

	"github.com/spf13/cobra"

	"k8s.io/apimachinery/pkg/runtime"
	clientgoscheme "k8s.io/client-go/kubernetes/scheme"
	_ "k8s.io/client-go/plugin/pkg/client/auth/gcp"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/log/zap"

	"go.awsctrl.io/manager/aws"
	"go.awsctrl.io/manager/controllers/cloudformation"
	"go.awsctrl.io/manager/controllers/controllermanager"
	"go.awsctrl.io/manager/controllers/self"
	"go.awsctrl.io/manager/token"

	cloudformationv1alpha1 "go.awsctrl.io/manager/apis/cloudformation/v1alpha1"
	selfv1alpha1 "go.awsctrl.io/manager/apis/self/v1alpha1"
)

var (
	resyncTimeout time.Duration = 15 * time.Minute
	scheme                      = runtime.NewScheme()
	setupLog                    = ctrl.Log.WithName("setup")

	awsclient            aws.AWS
	configname           string
	metricsaddr          string
	enableleaderelection bool
)

// startCmd represents the start command
var startCmd = &cobra.Command{
	Use:   "start",
	Short: "start will run the AWS Controller server",
	Long: `AWS Controller start will run the server which listens to a Kubernetes backend
for changes to the Custom Resources under the domain awsctrl.io and will create
update or delete resources as controlled by the CRs.

$ awsctrl start

This will run the server and connect to whatever Kubernetes API Server you
point it at.`,
	Run: runCommand,
}

func runCommand(cmd *cobra.Command, args []string) {

	ctrl.SetLogger(zap.Logger(true))

	mgr, err := ctrl.NewManager(ctrl.GetConfigOrDie(), ctrl.Options{
		Scheme:             scheme,
		MetricsBindAddress: metricsaddr,
		LeaderElection:     enableleaderelection,
		Port:               9443,
		SyncPeriod:         &resyncTimeout,
	})
	if err != nil {
		setupLog.Error(err, "unable to start manager")
		os.Exit(1)
	}

	awsclient = aws.New()

	if err = (&self.ConfigReconciler{
		Client:       mgr.GetClient(),
		Log:          ctrl.Log.WithName("controllers").WithName("self").WithName("config"),
		Scheme:       mgr.GetScheme(),
		ConfigName:   configname,
		PodNamespace: os.Getenv("POD_NAMESPACE"),
		AWSClient:    awsclient,
	}).SetupWithManager(mgr); err != nil {
		setupLog.Error(err, "unable to create controller", "controller", "Config")
		os.Exit(1)
	}

	if err = (&cloudformation.StackReconciler{
		Client:       mgr.GetClient(),
		Log:          ctrl.Log.WithName("controllers").WithName("cloudformation").WithName("stack"),
		Scheme:       mgr.GetScheme(),
		ConfigName:   configname,
		PodNamespace: os.Getenv("POD_NAMESPACE"),
		AWSClient:    awsclient,
		TokenClient:  token.New(),
	}).SetupWithManager(mgr); err != nil {
		setupLog.Error(err, "unable to create controller", "controller", "Stack")
		os.Exit(1)
	}

	if name, err := controllermanager.SetupControllers(mgr); err != nil {
		setupLog.Error(err, "unable to create controller", "controller", name)
		os.Exit(1)
	}

	// +kubebuilder:scaffold:builder

	setupLog.Info("starting manager")
	if err := mgr.Start(ctrl.SetupSignalHandler()); err != nil {
		setupLog.Error(err, "problem running manager")
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(startCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// startCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// startCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	startCmd.PersistentFlags().StringVar(&configname, "config-name", "config", "Name of the self.awsctrl.io/config to use.")
	startCmd.PersistentFlags().StringVar(&metricsaddr, "metrics-addr", ":8080", "The address the metric endpoint binds to.")
	startCmd.PersistentFlags().BoolVar(&enableleaderelection, "enable-leader-election", false,
		"Enable leader election for controller manager. Enabling this will ensure there is only one active controller manager.")

	_ = clientgoscheme.AddToScheme(scheme)
	_ = selfv1alpha1.AddToScheme(scheme)
	_ = cloudformationv1alpha1.AddToScheme(scheme)

	_ = controllermanager.AddAllSchemes(scheme)
	// +kubebuilder:scaffold:scheme
}
