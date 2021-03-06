// Code generated by client-gen. DO NOT EDIT.

package versioned

import (
	openshiftapiserverv1alpha1 "github.com/openshift/cluster-openshift-apiserver-operator/pkg/generated/clientset/versioned/typed/openshiftapiserver/v1alpha1"
	discovery "k8s.io/client-go/discovery"
	rest "k8s.io/client-go/rest"
	flowcontrol "k8s.io/client-go/util/flowcontrol"
)

type Interface interface {
	Discovery() discovery.DiscoveryInterface
	OpenshiftapiserverV1alpha1() openshiftapiserverv1alpha1.OpenshiftapiserverV1alpha1Interface
	// Deprecated: please explicitly pick a version if possible.
	Openshiftapiserver() openshiftapiserverv1alpha1.OpenshiftapiserverV1alpha1Interface
}

// Clientset contains the clients for groups. Each group has exactly one
// version included in a Clientset.
type Clientset struct {
	*discovery.DiscoveryClient
	openshiftapiserverV1alpha1 *openshiftapiserverv1alpha1.OpenshiftapiserverV1alpha1Client
}

// OpenshiftapiserverV1alpha1 retrieves the OpenshiftapiserverV1alpha1Client
func (c *Clientset) OpenshiftapiserverV1alpha1() openshiftapiserverv1alpha1.OpenshiftapiserverV1alpha1Interface {
	return c.openshiftapiserverV1alpha1
}

// Deprecated: Openshiftapiserver retrieves the default version of OpenshiftapiserverClient.
// Please explicitly pick a version.
func (c *Clientset) Openshiftapiserver() openshiftapiserverv1alpha1.OpenshiftapiserverV1alpha1Interface {
	return c.openshiftapiserverV1alpha1
}

// Discovery retrieves the DiscoveryClient
func (c *Clientset) Discovery() discovery.DiscoveryInterface {
	if c == nil {
		return nil
	}
	return c.DiscoveryClient
}

// NewForConfig creates a new Clientset for the given config.
func NewForConfig(c *rest.Config) (*Clientset, error) {
	configShallowCopy := *c
	if configShallowCopy.RateLimiter == nil && configShallowCopy.QPS > 0 {
		configShallowCopy.RateLimiter = flowcontrol.NewTokenBucketRateLimiter(configShallowCopy.QPS, configShallowCopy.Burst)
	}
	var cs Clientset
	var err error
	cs.openshiftapiserverV1alpha1, err = openshiftapiserverv1alpha1.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}

	cs.DiscoveryClient, err = discovery.NewDiscoveryClientForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	return &cs, nil
}

// NewForConfigOrDie creates a new Clientset for the given config and
// panics if there is an error in the config.
func NewForConfigOrDie(c *rest.Config) *Clientset {
	var cs Clientset
	cs.openshiftapiserverV1alpha1 = openshiftapiserverv1alpha1.NewForConfigOrDie(c)

	cs.DiscoveryClient = discovery.NewDiscoveryClientForConfigOrDie(c)
	return &cs
}

// New creates a new Clientset for the given RESTClient.
func New(c rest.Interface) *Clientset {
	var cs Clientset
	cs.openshiftapiserverV1alpha1 = openshiftapiserverv1alpha1.New(c)

	cs.DiscoveryClient = discovery.NewDiscoveryClient(c)
	return &cs
}
