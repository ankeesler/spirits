package main

import (
	"flag"
	"fmt"
	"os"

	// Import all Kubernetes client auth plugins (e.g. Azure, GCP, OIDC, etc.)
	// to ensure that exec-entrypoint and run can make use of them.
	_ "k8s.io/client-go/plugin/pkg/client/auth"

	"k8s.io/apimachinery/pkg/runtime"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
	clientgoscheme "k8s.io/client-go/kubernetes/scheme"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/healthz"
	"sigs.k8s.io/controller-runtime/pkg/log/zap"

	"github.com/ankeesler/spirits/internal/actionchannel"
	spiritsinternal "github.com/ankeesler/spirits/internal/apis/spirits"
	spiritsv1alpha1 "github.com/ankeesler/spirits/pkg/apis/spirits/v1alpha1"
	"github.com/ankeesler/spirits/pkg/controller"
	"github.com/go-logr/logr"
	//+kubebuilder:scaffold:imports
)

var (
	scheme   = runtime.NewScheme()
	setupLog = ctrl.Log.WithName("setup")
)

func init() {
	utilruntime.Must(clientgoscheme.AddToScheme(scheme))

	utilruntime.Must(spiritsinternal.AddToScheme(scheme))
	utilruntime.Must(spiritsv1alpha1.AddToScheme(scheme))
	//+kubebuilder:scaffold:scheme
}

func main() {
	var metricsAddr string
	var enableLeaderElection bool
	var probeAddr string
	flag.StringVar(&metricsAddr, "metrics-bind-address", ":8080", "The address the metric endpoint binds to.")
	flag.StringVar(&probeAddr, "health-probe-bind-address", ":8081", "The address the probe endpoint binds to.")
	flag.BoolVar(&enableLeaderElection, "leader-elect", false,
		"Enable leader election for controller manager. "+
			"Enabling this will ensure there is only one active controller manager.")
	opts := zap.Options{
		Development: true,
	}
	opts.BindFlags(flag.CommandLine)
	flag.Parse()

	ctrl.SetLogger(zap.New(zap.UseFlagOptions(&opts)))

	mgr, err := ctrl.NewManager(ctrl.GetConfigOrDie(), ctrl.Options{
		Scheme:                 scheme,
		MetricsBindAddress:     metricsAddr,
		Port:                   9443,
		HealthProbeBindAddress: probeAddr,
		LeaderElection:         enableLeaderElection,
		LeaderElectionID:       "62f0cb35.ankeesler.github.com",
	})
	if err != nil {
		setupLog.Error(err, "unable to start manager")
		os.Exit(1)
	}

	actionChannel := actionchannel.ActionChannel{}
	defer actionChannel.Close()

	if err = (&controller.SpiritReconciler{
		Client:       mgr.GetClient(),
		Scheme:       mgr.GetScheme(),
		ActionSource: &actionChannel,
	}).SetupWithManager(mgr); err != nil {
		setupLog.Error(err, "unable to create controller", "controller", "Spirit")
		os.Exit(1)
	}
	if err = (&controller.BattleReconciler{
		Client:     mgr.GetClient(),
		Scheme:     mgr.GetScheme(),
		ActionSink: &actionChannel,
	}).SetupWithManager(mgr); err != nil {
		setupLog.Error(err, "unable to create controller", "controller", "Battle")
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

	ctx := ctrl.SetupSignalHandler()

	setupLog.Info("starting apiserver")
	// if err := (&apiserver.APIServer{
	// 	Port:       9444,
	// 	DNSName:    getDNSName(setupLog),
	// 	ActionSink: &actionChannel,
	// 	PostStartHook: func() error {
	// 		go func() {
	setupLog.Info("starting managers")
	if err := mgr.Start(ctx); err != nil {
		setupLog.Error(err, "problem running manager")
		os.Exit(1)
	}
	// 		}()
	// 		return nil
	// 	},
	// }).Start(ctx); err != nil {
	// 	setupLog.Error(err, "problem running apiserver")
	// 	os.Exit(1)
	// }
}

func getDNSName(log logr.Logger) string {
	namespace := os.Getenv("POD_NAMESPACE")
	if len(namespace) == 0 {
		namespace = "default"
	}
	name := os.Getenv("POD_NAME")
	if len(name) == 0 {
		name = "spirits-controller-manager"
	}
	return fmt.Sprintf("%s.%s.svc.cluster.internal", name, namespace)
}
