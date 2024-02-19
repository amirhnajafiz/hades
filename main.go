package main

import (
	"fmt"
	"os"

	"k8s.io/apimachinery/pkg/runtime"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
	clientgoscheme "k8s.io/client-go/kubernetes/scheme"
	_ "k8s.io/client-go/plugin/pkg/client/auth"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/healthz"
	//+kubebuilder:scaffold:imports

	hadesv1beta1 "github.com/amirhnajafiz/hades/api/v1beta1"
	"github.com/amirhnajafiz/hades/internal/config"
	"github.com/amirhnajafiz/hades/internal/controllers/soles"
	"github.com/amirhnajafiz/hades/internal/logger"
)

var (
	scheme   = runtime.NewScheme()
	setupLog = ctrl.Log.WithName("setup")
)

func init() {
	utilruntime.Must(clientgoscheme.AddToScheme(scheme))
	utilruntime.Must(hadesv1beta1.AddToScheme(scheme))
	//+kubebuilder:scaffold:scheme
}

func main() {
	// load configs
	cfg := config.Config{}

	// create a new logger
	ctrl.SetLogger(logger.New(cfg.Logger.Level))

	// create a new manager
	mgr, err := ctrl.NewManager(ctrl.GetConfigOrDie(), ctrl.Options{
		Scheme:                 scheme,
		Port:                   9443,
		MetricsBindAddress:     fmt.Sprintf(":%d", cfg.Operator.Metrics),
		HealthProbeBindAddress: fmt.Sprintf(":%d", cfg.Operator.Probe),
		LeaderElection:         cfg.Operator.LeaderElect,
		LeaderElectionID:       "ae19d348.github.com",
	})
	if err != nil {
		setupLog.Error(err, "unable to start manager")

		os.Exit(1)
	}

	if err = (&soles.SoleReconciler{
		Client: mgr.GetClient(),
		Scheme: mgr.GetScheme(),
	}).SetupWithManager(mgr); err != nil {
		setupLog.Error(err, "unable to create controller", "controller", "Sole")
		os.Exit(1)
	}

	//+kubebuilder:scaffold:builder

	if err := mgr.AddHealthzCheck("healthz", healthz.Ping); err != nil {
		setupLog.Error(err, "unable to set up health check")

		os.Exit(1)
	}

	if err := mgr.AddReadyzCheck("readyz", healthz.Ping); err != nil {
		setupLog.Error(err, "unable to set up ready check")

		os.Exit(1)
	}

	setupLog.Info("starting manager")

	if err := mgr.Start(ctrl.SetupSignalHandler()); err != nil {
		setupLog.Error(err, "problem running manager")

		os.Exit(1)
	}
}
