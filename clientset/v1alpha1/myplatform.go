package v1alpha1

import (
	"context"
	"time"

	"github.com/RajaSureshAditya/k8s-crd-expl/apis/types/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/rest"
)

type MyProjectInterface interface {
	List(opts metav1.ListOptions) (*v1alpha1.MyProjectList, error)
	Get(name string, options metav1.GetOptions) (*v1alpha1.MyProject, error)
	Create(*v1alpha1.MyProject) (*v1alpha1.MyProject, error)
	// Watch(opts metav1.ListOptions) (watch.Interface, error)
}

type MyProjectClient struct {
	restClient rest.Interface
	ns         string
}

func (c *MyProjectClient) List(opts metav1.ListOptions) (*v1alpha1.MyProjectList, error) {
	result := v1alpha1.MyProjectList{}
	ctx := context.Background()
	// defer cancel()
	err := c.restClient.
		Get().
		Namespace(c.ns).
		Resource("myprojects").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do(ctx).
		Into(&result)

	return &result, err
}

func (c *MyProjectClient) Get(name string, opts metav1.GetOptions) (*v1alpha1.MyProject, error) {
	result := v1alpha1.MyProject{}
	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(300*time.Second))
	defer cancel()
	err := c.restClient.
		Get().
		Namespace(c.ns).
		Resource("myprojects").
		Name(name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Do(ctx).
		Into(&result)

	return &result, err
}

func (c *MyProjectClient) Create(MyProject *v1alpha1.MyProject) (*v1alpha1.MyProject, error) {
	result := v1alpha1.MyProject{}
	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(300*time.Second))
	defer cancel()
	err := c.restClient.
		Post().
		Namespace(c.ns).
		Resource("myprojects").
		Body(MyProject).
		Do(ctx).
		Into(&result)

	return &result, err
}

// func (c *projectClient) Watch(opts metav1.ListOptions) (watch.Interface, error) {
// 	opts.Watch = true
// 	return c.restClient.
// 		Get().
// 		Namespace(c.ns).
// 		Resource("projects").
// 		VersionedParams(&opts, scheme.ParameterCodec).
// 		Watch()
// }
