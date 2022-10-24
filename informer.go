package main

import (
	"time"

	"github.com/RajaSureshAditya/k8s-crd-expl/apis/types/v1alpha1"
	client_v1alpha1 "github.com/RajaSureshAditya/k8s-crd-expl/clientset/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/util/wait"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/tools/cache"
)

func WatchResources(clientSet client_v1alpha1.MyProjectV1alphainterface) cache.Store {
	projectStore, projectController := cache.NewInformer(
		&cache.ListWatch{
			ListFunc: func(lo metav1.ListOptions) (result runtime.Object, err error) {
				return clientSet.MyProjects("default").List(lo)
			},
			WatchFunc: func(lo metav1.ListOptions) (watch.Interface, error) {
				return clientSet.MyProjects("default").Watch(lo)
			},
		},
		&v1alpha1.Project{},
		3*time.Minute,
		cache.ResourceEventHandlerFuncs{},
	)

	go projectController.Run(wait.NeverStop)
	return projectStore
}
