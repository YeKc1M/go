package client

import (
	"context"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/rest"
	v1 "typedclientdemo/pkg/apis/foo/v1"
)

type helloTypeInterface interface {
	List(opts metav1.ListOptions) (*v1.HelloTypeList, error)
	Get(name string, opts metav1.GetOptions) (*v1.HelloType, error)
	Create(helloType *v1.HelloType) (*v1.HelloType, error)
	Watch(opts metav1.ListOptions) (watch.Interface, error)
}

type helloTypeClient struct {
	restClient rest.Interface
	namespace string
}

func NewHelloTypeClient(namespace string) *helloTypeClient {
	return &helloTypeClient{
		restClient: GetClientByCfg(GetCfgOutsideCluster()),
		namespace:  namespace,
	}
}

func (c *helloTypeClient) List(opts metav1.ListOptions) (*v1.HelloTypeList, error) {
	result := v1.HelloTypeList{}
	err := c.restClient.
		Get().
		Namespace(c.namespace).
		Resource("hellotypes").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do(context.TODO()).
		Into(&result)
	return &result, err
}

func (c *helloTypeClient) Get(name string, opts metav1.GetOptions) (*v1.HelloType, error) {
	result := v1.HelloType{}
	err := c.restClient.
		Get().
		Namespace(c.namespace).
		Resource("hellotypes").
		Name(name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Do(context.TODO()).
		Into(&result)
	return &result, err
}

func (c *helloTypeClient) Create(hellotype *v1.HelloType) (*v1.HelloType, error) {
	result := v1.HelloType{}
	err := c.restClient.
		Post().
		Namespace(c.namespace).
		Resource("hellotypes").
		Body(hellotype).
		Do(context.TODO()).
		Into(&result)
	return &result, err
}

func (c *helloTypeClient) Watch(opts metav1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.restClient.
		Get().
		Namespace(c.namespace).
		Resource("projects").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch(context.Background())
}

func (c *helloTypeClient) Update(updatedHt *v1.HelloType) (*v1.HelloType, error) {
	result := v1.HelloType{}
	err := c.restClient.Put().Namespace(c.namespace).Resource("hellotypes").Name(updatedHt.Name).Body(updatedHt).Do(context.TODO()).Into(&result)
	return &result, err
}
