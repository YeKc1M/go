package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genClient: noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// HelloType is a top-level type
type HelloType struct {
	metav1.TypeMeta `json:",inline"`
	// +optional
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +optional
	Status HelloTypeStatus `json:"status,omitempty"`
	// This is where you can define
	// your own custom spec
	Spec HelloSpec `json:"spec,omitempty"`
}

// custom spec
type HelloSpec struct {
	Message string `json:"message,omitempty"`
}

// custom status
type HelloTypeStatus struct {
	Name string
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type HelloTypeList struct {
	metav1.TypeMeta `json:",inline"`
	// +optional
	metav1.ListMeta `son:"metadata,omitempty"`

	Items []HelloType `json:"items"`
}

func (in *HelloType) DeepCopyInto(out *HelloType) {
	out.TypeMeta = in.TypeMeta
	out.ObjectMeta = in.ObjectMeta
	out.Status = in.Status
	out.Spec = HelloSpec{Message: in.Spec.Message}
}

func (in *HelloType) DeepCopyObject() runtime.Object {
	out := HelloType{}
	in.DeepCopyInto(&out)

	return &out
}

func (in *HelloTypeList) DeepCopyObject() runtime.Object {
	out := HelloTypeList{}
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta

	if in.Items != nil {
		out.Items = make([]HelloType, len(in.Items))
		for i := range in.Items {
			in.Items[i].DeepCopyInto(&out.Items[i])
		}
	}

	return &out
}
