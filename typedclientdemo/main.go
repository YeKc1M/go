package main

import (
	"context"
	"fmt"
	v12 "k8s.io/apimachinery/pkg/apis/meta/v1"
	v1 "typedclientdemo/pkg/apis/foo/v1"
	"typedclientdemo/pkg/client"
)

func main() {
	//test1()
	//testHtClientGet()
	//testHtClientUpdate()
	//testInformer()
	//testWatch()

	deployClient := getDeployClient()
	deploy := GetDeploy(deployClient, "hello-node")
	//UpdateDeployReplica(deployClient, deploy, 1)
	WatchDeploy(deployClient, deploy)
}

func test1() {
	crdClient := client.GetClientByCfg(client.GetCfgOutsideCluster())
	result := v1.HelloTypeList{}
	crdClient.Get().Resource("hellotypes").Do(context.TODO()).Into(&result)
	fmt.Println(result)
}

func testHtClientGet() {
	htclient := client.NewHelloTypeClient("default")
	result, err := htclient.Get("hello", v12.GetOptions{})
	if err != nil {
		panic(err)
	}
	fmt.Println(result)
	fmt.Println(result.Spec)
	fmt.Println(result.Status)
	fmt.Println(result.ObjectMeta)
	fmt.Println(result.TypeMeta)
}

func testHtClientUpdate(){
	htclient := client.NewHelloTypeClient("default")
	before, err := htclient.Get("hello", v12.GetOptions{})
	if err != nil {
		panic(err)
	}
	//var update v1.HelloType
	//before.DeepCopyInto(&update)
	//update.Name = "updated"
	//update.Spec.Message = "new update function"
	//update.Status
	//updatedRes, err := htclient.Update(before, &update, v12.UpdateOptions{})
	before.Spec.Message = "update1321313"
	updatedRes, err := htclient.Update(before, v12.UpdateOptions{})
	if err != nil {
		panic(err)
	}
	fmt.Println(updatedRes)
}

func testWatch() {
	htclient := client.NewHelloTypeClient("default")
	w, err := htclient.Watch(v12.ListOptions{})
	if err != nil {
		panic(err)
	}
	for true {
		select {
		case event, b := <-w.ResultChan():
			if b {
				fmt.Println(event, b)
			}
		}
	}
}


func testInformer(){
	htclient := client.NewHelloTypeClient("default")
	store := client.WatchResources(htclient)
	res, exst, err := store.GetByKey("default/hellotype")
	if err != nil {
		panic(err)
	}
	if exst {
		fmt.Println(res)
	}
	fmt.Println(store)
}
