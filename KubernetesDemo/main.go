package main

import (
	"context"
	"fmt"
	"mky.example.com/kubernetes/dynamicclient"
	"time"

	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

func main() {
	//clientset := getClientset()
	//getAllPods(clientset)
	//findPod(clientset)
	//crd()
	//crd1(getClientset())
	//getCRD()
	//updateCRD1()
	dynamicclient.UpdateCRD2(2)
}

func getAllPods(clientset *kubernetes.Clientset) {
	for {
		//if namespace is "", find pods in all namespaces
		pods, err := clientset.CoreV1().Pods("default").List(context.TODO(), metav1.ListOptions{})
		if err != nil {
			panic(err.Error())
		}
		fmt.Printf("There are %d pods in the cluster\n", len(pods.Items))

		podList := pods.Items
		for i := 0; i < len(podList); i++ {
			pod := podList[i]
			name := pod.ObjectMeta.Name
			fmt.Printf("pod %v: %s\n", i, name)
		}

		// Examples for error handling:
		// - Use helper functions like e.g. errors.IsNotFound()
		// - And/or cast to StatusError and use its properties like e.g. ErrStatus.Message
		namespace := "default"
		pod := "example-xxxxx"
		_, err = clientset.CoreV1().Pods(namespace).Get(context.TODO(), pod, metav1.GetOptions{})
		if errors.IsNotFound(err) {
			fmt.Printf("Pod %s in namespace %s not found\n", pod, namespace)
		} else if statusError, isStatus := err.(*errors.StatusError); isStatus {
			fmt.Printf("Error getting pod %s in namespace %s: %v\n",
				pod, namespace, statusError.ErrStatus.Message)
		} else if err != nil {
			panic(err.Error())
		} else {
			fmt.Printf("Found pod %s in namespace %s\n", pod, namespace)
		}

		time.Sleep(10 * time.Second)
	}
}

func findPod(clientset *kubernetes.Clientset){
	pod, err := clientset.CoreV1().Pods("default").Get(context.TODO(), "hello-node-7567d9fdc9-l4fhj", metav1.GetOptions{})
	if err != nil {
		panic(err.Error())
	}
	name := pod.ObjectMeta.Name
	ip := pod.Status.PodIP
	fmt.Printf("pod %s: %s\n", name, ip)
}

func test(clientset *kubernetes.Clientset){
	//pod, err := clientset.CoreV1().
	//if err != nil {
	//	panic(err.Error())
	//}
	//name := pod.ObjectMeta.Name
	//ip := pod.Status.PodIP
	//fmt.Printf("pod %s: %s", name, ip)
}
