package main

import (
	"context"
	"fmt"
	"k8s.io/api/node/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/rest"
)

//define type
type ProjectSpec struct {
	Replicas int `json:"replicas"`
}

type Project struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec ProjectSpec `json:"spec"`
}

type ProjectList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`

	Items []Project `json:"items"`
}

func (in *Project) DeepCopyInto(out *Project) {
	out.TypeMeta = in.TypeMeta
	out.ObjectMeta = in.ObjectMeta
	out.Spec = ProjectSpec{
		Replicas:in.Spec.Replicas,
	}
}

func (in *Project) DeepCopyObject() runtime.Object {
	out := Project{}
	in.DeepCopyInto(&out)
	return &out
}

func (in *ProjectList) DeepCopyObject() runtime.Object {
	out := ProjectList{}
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		out.Items = make([]Project, len(in.Items))
		for item := range in.Items {
			in.Items[item].DeepCopyInto(&out.Items[item])
		}
	}
	return &out
}

//register type
const GroupName = "example.sealyun.com"
const GroupVersion = "v1alpha1"

var SchemaGroupVersion = schema.GroupVersion{
	Group:   GroupName,
	Version: GroupVersion,
}

func addKnownTypes(scheme *runtime.Scheme) error {
	scheme.AddKnownTypes(SchemaGroupVersion,
		&Project{},
		&ProjectList{},
		)
	metav1.AddToGroupVersion(scheme, SchemaGroupVersion)
	return nil
}

//construct client
func crd(){
	crdCfg := *getConfigFromPath(getCfgPathFromHome())
	crdCfg.ContentConfig.GroupVersion = &schema.GroupVersion{
		Group:v1alpha1.GroupName,
		Version:"v1alpha1",
	}
	crdCfg.APIPath = "/apis"
	crdCfg.NegotiatedSerializer = scheme.Codecs
	crdCfg.UserAgent = rest.DefaultKubernetesUserAgent()
	exampleRestClient, err := rest.UnversionedRESTClientFor(&crdCfg)
	if err != nil {
		panic(err.Error())
	}
	result := ProjectList{}
	err = exampleRestClient.Get().Resource("projects").Do(context.TODO()).Into(&result)
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(len(result.Items))
}


func crd1(clientset *kubernetes.Clientset){
	restClient := clientset.CoreV1().RESTClient()
	result := ProjectList{}
	restClient.Get().Resource("projects").Do(context.TODO()).Into(&result)
	fmt.Println(len(result.Items))//0
}
