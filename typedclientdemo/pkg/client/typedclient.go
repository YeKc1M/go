package client

import (
	"flag"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/runtime/serializer"
	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/tools/clientcmd"
	"path/filepath"
	v1 "typedclientdemo/pkg/apis/foo/v1"

	restclient "k8s.io/client-go/rest"
	"k8s.io/client-go/util/homedir"
)

func getCfgPathFromHome() *string{
	var kubeconfig *string
	home := homedir.HomeDir()
	//fmt.Printf("homeDir: %s\n", home)
	if home != "" {
		kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
	}else {
		kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	}
	return kubeconfig
}

func getConfigFromPath(path *string) *restclient.Config {
	config, err := clientcmd.BuildConfigFromFlags("", *path)
	if err != nil {
		panic(err.Error())
	}
	return config
}

func GetCfgOutsideCluster() *restclient.Config {
	return getConfigFromPath(getCfgPathFromHome())
}

func GetClientByCfg(cfg *restclient.Config) *restclient.RESTClient {
	v1.AddToScheme(scheme.Scheme)
	crdConfig := *cfg
	crdConfig.ContentConfig.GroupVersion = &schema.GroupVersion{
		Group:   v1.SchemeGroupVersion.Group,
		Version: v1.SchemeGroupVersion.Version,
	}
	crdConfig.APIPath = "/apis"
	crdConfig.NegotiatedSerializer = serializer.NewCodecFactory(scheme.Scheme)
	crdConfig.UserAgent = restclient.DefaultKubernetesUserAgent()

	crdRestClient, err := restclient.UnversionedRESTClientFor(&crdConfig)
	if err != nil {
		panic(err)
	}
	return crdRestClient
}
