package main

import (
	"flag"
	"os"

	"awsctrl.io/pkg/apis"
	"awsctrl.io/pkg/controller"
	"awsctrl.io/pkg/controller/stack"
	"awsctrl.io/pkg/queue"
	"awsctrl.io/pkg/webhook"

	_ "k8s.io/client-go/plugin/pkg/client/auth/gcp"
	"sigs.k8s.io/controller-runtime/pkg/client/config"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	logf "sigs.k8s.io/controller-runtime/pkg/runtime/log"
	"sigs.k8s.io/controller-runtime/pkg/runtime/signals"

	"awsctrl.io/pkg/aws"
	"awsctrl.io/pkg/token"
)

func main() {
	var metricsAddr string
	flag.StringVar(&metricsAddr, "metrics-addr", ":8080", "The address the metric endpoint binds to.")
	flag.Parse()
	logf.SetLogger(logf.ZapLogger(false))
	log := logf.Log.WithName("entrypoint")

	// Get a config to talk to the apiserver
	log.Info("setting up client for manager")
	cfg, err := config.GetConfig()
	if err != nil {
		log.Error(err, "unable to set up client config")
		os.Exit(1)
	}

	// Create a new Cmd to provide shared dependencies and start components
	log.Info("setting up manager")
	mgr, err := manager.New(cfg, manager.Options{MetricsBindAddress: metricsAddr})
	if err != nil {
		log.Error(err, "unable to set up overall controller manager")
		os.Exit(1)
	}

	log.Info("registering components.")

	// Setup Scheme for all resources
	log.Info("setting up scheme")
	if err := apis.AddToScheme(mgr.GetScheme()); err != nil {
		log.Error(err, "unable add APIs to scheme")
		os.Exit(1)
	}

	// Setup all Controllers
	log.Info("setting up controller")
	if err := controller.AddToManager(mgr); err != nil {
		log.Error(err, "unable to register controllers to the manager")
		os.Exit(1)
	}

	log.Info("setting up webhooks")
	if err := webhook.AddToManager(mgr); err != nil {
		log.Error(err, "unable to register webhooks to the manager")
		os.Exit(1)
	}

	signalCh := signals.SetupSignalHandler()

	go func() {
		log.Info("starting the queue.")
		if err := queue.New(mgr, stack.NewReconciler(mgr, aws.New(), token.New(), os.Getenv("POD_NAMESPACE")).Reconcile).Start(signalCh); err != nil {
			log.Error(err, "unable to run the queue")
			os.Exit(1)
		}
	}()

	// Start the Cmd
	log.Info("starting the cmd.")
	if err := mgr.Start(signalCh); err != nil {
		log.Error(err, "unable to run the manager")
		os.Exit(1)
	}
}
