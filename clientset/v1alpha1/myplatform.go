package v1alpha1

import (
	"context"
	"time"

	"github.com/RajaSureshAditya/k8s-crd-expl/apis/types/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/rest"
)

type MyplatformInterface interface {
	List(opts metav1.ListOptions) (*v1alpha1.MyplatformList, error)
	Get(name string, options metav1.GetOptions) (*v1alpha1.Myplatform, error)
	Create(*v1alpha1.Myplatform) (*v1alpha1.Myplatform, error)
	// Watch(opts metav1.ListOptions) (watch.Interface, error)
}

type MyplatformClient struct {
	restClient rest.Interface
	ns         string
}

func (c *MyplatformClient) List(opts metav1.ListOptions) (*v1alpha1.MyplatformList, error) {
	result := v1alpha1.MyplatformList{}
	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(300*time.Second))
	defer cancel()
	err := c.restClient.
		Get().
		Namespace(c.ns).
		Resource("myplatform").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do(ctx).
		Into(&result)

	return &result, err
}

func (c *MyplatformClient) Get(name string, opts metav1.GetOptions) (*v1alpha1.Myplatform, error) {
	result := v1alpha1.Myplatform{}
	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(300*time.Second))
	defer cancel()
	err := c.restClient.
		Get().
		Namespace(c.ns).
		Resource("myplatforms").
		Name(name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Do(ctx).
		Into(&result)

	return &result, err
}

func (c *MyplatformClient) Create(myplatform *v1alpha1.Myplatform) (*v1alpha1.Myplatform, error) {
	result := v1alpha1.Myplatform{}
	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(300*time.Second))
	defer cancel()
	err := c.restClient.
		Post().
		Namespace(c.ns).
		Resource("myplatforms").
		Body(myplatform).
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
