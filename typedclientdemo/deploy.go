package main

import (
	"context"
	"fmt"
	appv1 "k8s.io/api/apps/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	v1 "k8s.io/client-go/kubernetes/typed/apps/v1"
	"typedclientdemo/pkg/client"
)

func getDeployClient() (v1.DeploymentInterface){
	cfg := client.GetCfgOutsideCluster()
	clientset, err := kubernetes.NewForConfig(cfg)
	if err != nil {
		panic(err)
	}
	deployClient := clientset.AppsV1().Deployments("default")
	return deployClient
}

func GetDeploy(deployClient v1.DeploymentInterface, name string) *appv1.Deployment {
	deploy, err := deployClient.Get(context.TODO(), name, metav1.GetOptions{})
	if err != nil {
		panic(err)
	}
	fmt.Println("Get:-----------------", deploy)
	return deploy
}

func UpdateDeployReplica(deployClient v1.DeploymentInterface, deploy *appv1.Deployment, rep int) *appv1.Deployment {
	deploy, err := deployClient.Get(context.TODO(), deploy.Name, metav1.GetOptions{})
	if err != nil {
		panic(err)
	}
	newReplica := int32(rep)
	deploy.Spec.Replicas = &newReplica
	updatedDeploy, err := deployClient.Update(context.TODO(), deploy, metav1.UpdateOptions{})
	if err != nil {
		panic(err)
	}
	fmt.Println("Update:---------------",updatedDeploy)
	return updatedDeploy
}

func WatchDeploy(deployClient v1.DeploymentInterface, deploy *appv1.Deployment) {
	watch, err := deployClient.Watch(context.TODO(), metav1.ListOptions{
		Watch: true,
	})
	if err != nil {
		panic(err)
	}
	go UpdateDeployReplica(deployClient, deploy, 1)
	for {
		select {
		case e, _ := <-watch.ResultChan():
			fmt.Println("--------------------------------------")
			fmt.Println(e.Type, e.Object)
		}
	}
}
