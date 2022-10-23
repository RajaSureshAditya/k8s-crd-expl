package main

import (
	"fmt"
	"os"
	"path/filepath"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/RajaSureshAditya/k8s-crd-expl/apis/types/v1alpha1"
	clientV1alpha1 "github.com/RajaSureshAditya/k8s-crd-expl/clientset/v1alpha1"
	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

// var gvr = schema.GroupVersionResource{
// 	Group:    "",
// 	Version:  "v1",
// 	Resource: "pods",
// }

func Get_k8s_config() (*rest.Config, error) {
	userHomeDir, err := os.UserHomeDir()
	if err != nil {
		fmt.Printf("error getting user home dir: %v\n", err)
		os.Exit(1)
	}
	kubeConfigPath := filepath.Join(userHomeDir, ".kube", "config")
	kubeConfig_cfg, err := clientcmd.BuildConfigFromFlags("", kubeConfigPath)
	if err != nil {
		fmt.Printf("error getting Kubernetes config: %v\n", err)
		os.Exit(1)
	}
	v1alpha1.AddToScheme(scheme.Scheme)
	return kubeConfig_cfg, nil
}

func main() {
	dynamic_config, err := Get_k8s_config()
	if err != nil {
		panic(err)
	}
	clientset, err := clientV1alpha1.NewforConfig(dynamic_config)
	if err != nil {
		panic(err)
	}
	myprojectsList, err := clientset.MyProjects("default").List(metav1.ListOptions{})
	if err != nil {
		panic(err)
	}
	fmt.Printf("projects found: %+v\n", myprojectsList)

}
