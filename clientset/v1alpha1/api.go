package v1alpha1

import (
	"fmt"
	"os"

	v1alpha1 "github.com/RajaSureshAditya/k8s-crd-expl/apis/types/v1alpha1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/rest"
)

type MyplatformV1alphaclient struct {
	client rest.Interface
}

type MyplatformV1alphainterface interface {
	Myplatforms(ns string) MyplatformInterface
}

func NewforConfig(c *rest.Config) (*MyplatformV1alphaclient, error) {
	kubeConfig_cfg := *c
	kubeConfig_cfg.ContentConfig.GroupVersion = &schema.GroupVersion{Group: v1alpha1.GroupName, Version: v1alpha1.GroupVersion}
	kubeConfig_cfg.APIPath = "/apis"
	// kubeConfig_cfg.NegotiatedSerializer = serializer.NewCodecFactory(scheme.Scheme)
	kubeConfig_cfg.NegotiatedSerializer = scheme.Codecs.WithoutConversion()
	kubeConfig_cfg.UserAgent = rest.DefaultKubernetesUserAgent()
	// myplatformclientset, err := rest.RESTClientFor(&kubeConfig_cfg)
	myplatformclientset, err := rest.UnversionedRESTClientFor(&kubeConfig_cfg)
	if err != nil {
		fmt.Printf("error creating dynamic client: %v\n", err)
		os.Exit(1)
	}

	return &MyplatformV1alphaclient{client: myplatformclientset}, nil
}

func (mpl *MyplatformV1alphaclient) Myplatforms(namespace string) MyplatformInterface {
	return &MyplatformClient{
		restClient: mpl.client,
		ns:         namespace,
	}
}
