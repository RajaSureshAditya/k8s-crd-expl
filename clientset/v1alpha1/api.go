package v1alpha1

import (
	"fmt"
	"os"

	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/runtime/serializer"
	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/rest"
)

type MyProjectV1alphaclient struct {
	client rest.Interface
}

type MyProjectV1alphainterface interface {
	MyProjects(ns string) MyProjectInterface
}

func NewforConfig(c *rest.Config) (*MyProjectV1alphaclient, error) {
	kubeConfig_cfg := *c
	kubeConfig_cfg.ContentConfig.GroupVersion = &schema.GroupVersion{Group: "contoso.com", Version: "v1alpha1"}
	kubeConfig_cfg.APIPath = "/apis"
	// kubeConfig_cfg.NegotiatedSerializer = serializer.NewCodecFactory(scheme.Scheme)
	kubeConfig_cfg.NegotiatedSerializer = serializer.WithoutConversionCodecFactory{CodecFactory: scheme.Codecs}
	// kubeConfig_cfg.NegotiatedSerializer = scheme.Codecs.WithoutConversion()
	kubeConfig_cfg.UserAgent = rest.DefaultKubernetesUserAgent()
	MyProjectclientset, err := rest.RESTClientFor(&kubeConfig_cfg)
	// MyProjectclientset, err := rest.UnversionedRESTClientFor(&kubeConfig_cfg)
	if err != nil {
		fmt.Printf("error creating dynamic client: %v\n", err)
		os.Exit(1)
	}

	return &MyProjectV1alphaclient{client: MyProjectclientset}, nil
}

func (mpl *MyProjectV1alphaclient) MyProjects(namespace string) MyProjectInterface {
	return &MyProjectClient{
		restClient: mpl.client,
		ns:         namespace,
	}
}
