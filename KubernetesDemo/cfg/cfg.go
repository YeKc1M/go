package cfg

import (
	"flag"
	"fmt"
	"k8s.io/client-go/kubernetes"
	restclient "k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
	"path/filepath"
)

func GetCfgPathFromHome() *string{
	var kubeconfig *string
	home := homedir.HomeDir()
	fmt.Printf("homeDir: %s\n", home)
	if home != "" {
		kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
	}else {
		kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	}
	return kubeconfig
}

func GetConfigFromPath(path *string) *restclient.Config {
	config, err := clientcmd.BuildConfigFromFlags("", *path)
	if err != nil {
		panic(err.Error())
	}
	return config
}

func getClientset() *kubernetes.Clientset{
	kubeconfig := GetCfgPathFromHome()
	fmt.Println(*kubeconfig)
	flag.Parse()

	//use current context in kubeconfig
	config := GetConfigFromPath(kubeconfig)

	//create clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}
	return clientset
}
