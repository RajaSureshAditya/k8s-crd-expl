package v1alpha1

import "k8s.io/apimachinery/pkg/runtime"

// DeepCopyInto copies all properties of this object into another object of the
// same type that is provided as a pointer.
func (in *MyProject) DeepCopyInto(out *MyProject) {
	out.TypeMeta = in.TypeMeta
	out.ObjectMeta = in.ObjectMeta
	out.Spec = MyProjectSpec{
		AppId:        in.Spec.AppId,
		Language:     in.Spec.Language,
		Os:           in.Spec.Os,
		InstanceSize: in.Spec.InstanceSize,
	}
}

// DeepCopyObject returns a generically typed copy of an object
func (in *MyProject) DeepCopyObject() runtime.Object {
	out := MyProject{}
	in.DeepCopyInto(&out)

	return &out
}

// DeepCopyObject returns a generically typed copy of an object
func (in *MyProjectList) DeepCopyObject() runtime.Object {
	out := MyProjectList{}
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta

	if in.Items != nil {
		out.Items = make([]MyProject, len(in.Items))
		for i := range in.Items {
			in.Items[i].DeepCopyInto(&out.Items[i])
		}
	}

	return &out
}
