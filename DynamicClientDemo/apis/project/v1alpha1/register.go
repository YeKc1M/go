package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

//register type
const GroupName = "example.sealyun.com"
const GroupVersion = "v1alpha1"

var SchemaGroupVersion = schema.GroupVersion{
	Group:   GroupName,
	Version: GroupVersion,
}

var(
	SchemeBuilder runtime.SchemeBuilder
	localSchemeBuilder = &SchemeBuilder
	AddToScheme = localSchemeBuilder.AddToScheme
)

func init() {
	localSchemeBuilder.Register(addKnownTypes)
}

func Resource(resource string) schema.GroupResource {
	return SchemaGroupVersion.wthResource(resource).GroupResource()
}

func addKnownTypes(scheme *runtime.Scheme) error {
	scheme.AddKnownTypes(SchemaGroupVersion, &Project{}, &ProjectList{})
	scheme.AddKnownTypes(SchemaGroupVersion, &metav1.Status{})
	metav1.AddToGroupVersion(scheme, SchemaGroupVersion)
	return nil
}
