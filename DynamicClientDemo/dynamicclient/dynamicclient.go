package dynamicclient

import (
	"context"
	"dynamicclientdemo/cfg"
	"fmt"
	"k8s.io/apimachinery/pkg/api/meta"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/runtime/serializer/yaml"
	"k8s.io/client-go/discovery"
	memory "k8s.io/client-go/discovery/cached"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/restmapper"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/retry"
	"log"
)

const metaCRD = `
apiVersion: "example.sealyun.com/v1alpha1"
kind: "Project"
metadata:
  name: "example-project"
  namespace: "default"
spec:
  replicas: 3
`


func GetCfgOutsideCluster() (*rest.Config, error) {
	config, err := clientcmd.BuildConfigFromFlags("", *cfg.GetCfgPathFromHome())
	if err != nil {
		panic(err)
	}
	return config, err
}

func GetCfgInsideCluster() (*rest.Config, error)  {
	return rest.InClusterConfig()
}

func GetGVRdyClient(gvk *schema.GroupVersionKind,namespace string) (dynamic.ResourceInterface, error)  {
	var dr dynamic.ResourceInterface
	//rest.InClusterConfig() used for in-cluster service account auth, it may work, maybe.
	config,err := GetCfgOutsideCluster()
	if err != nil {
		panic(err.Error())
	}

	// 创建discovery客户端
	discoveryClient,err := discovery.NewDiscoveryClientForConfig(config)
	if err != nil {
		panic(err.Error())
	}

	// 获取GVK GVR 映射
	mapperGVRGVK := restmapper.NewDeferredDiscoveryRESTMapper(memory.NewMemCacheClient(discoveryClient))

	// 根据资源GVK 获取资源的GVR GVK映射
	resourceMapper,err := mapperGVRGVK.RESTMapping(gvk.GroupKind(),gvk.Version)
	if err != nil {
		panic(err.Error())
	}

	// 创建动态客户端
	dynamicClient,err := dynamic.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}

	if resourceMapper.Scope.Name() == meta.RESTScopeNameNamespace {
		// 获取gvr对应的动态客户端
		dr = dynamicClient.Resource(resourceMapper.Resource).Namespace(namespace)
	} else {
		// 获取gvr对应的动态客户端
		dr = dynamicClient.Resource(resourceMapper.Resource)
	}

	return dr, err
}

func dynamicclient() (dynamic.ResourceInterface, *unstructured.Unstructured){
	obj := &unstructured.Unstructured{}
	_, gvk, err := yaml.NewDecodingSerializer(unstructured.UnstructuredJSONScheme).Decode([]byte(metaCRD), nil, obj)
	if err != nil {
		panic(err.Error())
	}

	dr, err := GetGVRdyClient(gvk, obj.GetNamespace())
	if err != nil {
		panic(err.Error())
	}

	return dr, obj

	//创建
	//objCreate, err = dr.Create(obj, metav1.CreateOptions{})
	//if err != nil {
	//	//panic(fmt.Errorf("Create resource ERROR: %v", err))
	//	log.Print(err)
	//}
	//log.Print("Create: : ", objCreate)

	// 删除
	//err = dr.Delete(obj.GetName(),&metav1.DeleteOptions{})
	//if err != nil {
	//	panic(fmt.Errorf("delete resource ERROR : %v", err))
	//} else {
	//	log.Print("删除成功")
	//}
}

func GetCRD(crname string) {
	dr, _ := dynamicclient()
	objGet, err := dr.Get(context.TODO(), crname, metav1.GetOptions{})
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(objGet)

	spec, found, err := unstructured.NestedMap(objGet.Object, "spec")
	if err != nil || !found || spec == nil {
		panic(err.Error())
	}
	fmt.Println(spec)
}

func UpdateCRD2(replicas int){
	dr, obj := dynamicclient()
	retryErr := retry.RetryOnConflict(retry.DefaultRetry, func() (err error) {
		// check if resource exists
		result, err := dr.Get(context.TODO(), obj.GetName(),metav1.GetOptions{})
		if err != nil {
			panic(fmt.Errorf("failed to get latest version of : %v", err))
		}
		// update spec
		spec := make(map[string]interface{})
		spec["replicas"] = int64(replicas)//work
		if err := unstructured.SetNestedMap(result.Object, spec, "spec", ); err != nil {
			panic(err)
		}
		// update resource
		objUpdate, err := dr.Update(context.TODO(), result,metav1.UpdateOptions{})
		log.Print("update : ",objUpdate)
		return err
	})
	if retryErr != nil {
		panic(fmt.Errorf("update failed: %v", retryErr))
	} else {
		log.Print("更新成功")
	}
}
