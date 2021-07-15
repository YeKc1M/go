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
	testHtClientUpdate()
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
	before.Spec.Message = "update2"
	updatedRes, err := htclient.Update(before)
	if err != nil {
		panic(err)
	}
	fmt.Println(updatedRes)
}
