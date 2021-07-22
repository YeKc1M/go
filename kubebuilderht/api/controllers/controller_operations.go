package controllers

import (
	"context"
	"flag"
	"fmt"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
	clientgoscheme "k8s.io/client-go/kubernetes/scheme"
	foocomv1 "kubebuilderht/api/v1"
	"os"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log/zap"
	"time"
)

var (
	scheme       = runtime.NewScheme()
	setupLog     = ctrl.Log.WithName("setup")
	metricsHost               = "0.0.0.0"
	metricsPort         int32 = 8383
	operatorMetricsPort int32 = 8686
	HTReconsiler *HelloTypeReconciler
)

func init() {
	utilruntime.Must(clientgoscheme.AddToScheme(scheme))

	utilruntime.Must(foocomv1.AddToScheme(scheme))
	//+kubebuilder:scaffold:scheme
}

func WatchTemplate() {
	//var metricsAddr string
	//var enableLeaderElection bool
	//var probeAddr string
	//flag.StringVar(&metricsAddr, "metrics-bind-address", ":8080", "The address the metric endpoint binds to.")
	//flag.StringVar(&probeAddr, "health-probe-bind-address", ":8081", "The address the probe endpoint binds to.")
	//flag.BoolVar(&enableLeaderElection, "leader-elect", false,
	//	"Enable leader election for controller manager. "+
	//		"Enabling this will ensure there is only one active controller manager.")
	opts := zap.Options{
		Development: true,
	}
	opts.BindFlags(flag.CommandLine)
	flag.Parse()

	ctrl.SetLogger(zap.New(zap.UseFlagOptions(&opts)))

	mgr, err := ctrl.NewManager(ctrl.GetConfigOrDie(), ctrl.Options{
		Scheme: scheme,
		MetricsBindAddress:     fmt.Sprintf("%s:%d", metricsHost, metricsPort),
		//Port:                   9443,
		//HealthProbeBindAddress: probeAddr,
		//LeaderElection:         enableLeaderElection,
		//LeaderElectionID:       "2ac1ce84.my.domain",
	})
	if err != nil {
		setupLog.Error(err, "unable to start manager")
		os.Exit(1)
	}

	HTReconsiler = &HelloTypeReconciler{
		Client: mgr.GetClient(),
		Scheme: mgr.GetScheme(),
	}

	if err = HTReconsiler.SetupWithManager(mgr); err != nil {
		setupLog.Error(err, "unable to create controller", "controller", "HelloType")
		os.Exit(1)
	}
	//+kubebuilder:scaffold:builder

	//if err := mgr.AddHealthzCheck("healthz", healthz.Ping); err != nil {
	//	setupLog.Error(err, "unable to set up health check")
	//	os.Exit(1)
	//}
	//if err := mgr.AddReadyzCheck("readyz", healthz.Ping); err != nil {
	//	setupLog.Error(err, "unable to set up ready check")
	//	os.Exit(1)
	//}

	setupLog.Info("starting manager")
	if err := mgr.Start(ctrl.SetupSignalHandler()); err != nil {
		setupLog.Error(err, "problem running manager")
		os.Exit(1)
	}
}

func (r *HelloTypeReconciler) GetHelloType(namespace string, name string) *foocomv1.HelloType {
	getRes := &foocomv1.HelloType{}
	if err := r.Get(context.TODO(), types.NamespacedName{
		Namespace: namespace,
		Name:      name,
	}, getRes); err != nil {
		panic(err)
	}
	fmt.Println(time.Now(), "get")
	fmt.Println(getRes)
	return getRes
}

func (r *HelloTypeReconciler) ListHelloType()  {
	htList := foocomv1.HelloTypeList{}
	err := r.List(context.TODO(), &htList, client.InNamespace("default"))
	if err != nil {
		panic(err)
	}
	fmt.Println(time.Now() ,"list")
	fmt.Println(htList)
}
