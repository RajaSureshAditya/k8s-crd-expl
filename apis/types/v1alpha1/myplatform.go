package v1alpha1

import meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type MyProject struct {
	meta_v1.TypeMeta   `json:",inline"`
	meta_v1.ObjectMeta `json:"metadata"`
	Spec               MyProjectSpec   `json:"spec"`
	Status             MyProjectStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type MyProjectList struct {
	meta_v1.TypeMeta `json:",inline"`
	meta_v1.ListMeta `json:"metadata"`
	Items            []MyProject `json:"items"`
}
type MyProjectSpec struct {
	AppId           string `json:"appId"`
	EnvironmentType string `json:"environmentType"`
	Language        string `json:"language"`
	Os              string `json:"os"`
	InstanceSize    string `json:"instanceSize"`
	Replicas        int    `json:"replicas"`
}
type MyProjectStatus struct {
	State   string `json:"state,omitempty"`
	Message string `json:"message,omitempty"`
}
