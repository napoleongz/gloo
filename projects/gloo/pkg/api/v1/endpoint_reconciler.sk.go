// Code generated by solo-kit. DO NOT EDIT.

package v1

import (
	"github.com/solo-io/go-utils/contextutils"
	"github.com/solo-io/solo-kit/pkg/api/v1/clients"
	"github.com/solo-io/solo-kit/pkg/api/v1/reconcile"
	"github.com/solo-io/solo-kit/pkg/api/v1/resources"
)

// Option to copy anything from the original to the desired before writing. Return value of false means don't update
type TransitionEndpointFunc func(original, desired *Endpoint) (bool, error)

type EndpointReconciler interface {
	Reconcile(namespace string, desiredResources EndpointList, transition TransitionEndpointFunc, opts clients.ListOpts) error
}

func endpointsToResources(list EndpointList) resources.ResourceList {
	var resourceList resources.ResourceList
	for _, endpoint := range list {
		resourceList = append(resourceList, endpoint)
	}
	return resourceList
}

func NewEndpointReconciler(client EndpointClient, statusSetter resources.StatusSetter) EndpointReconciler {
	return &endpointReconciler{
		base: reconcile.NewReconciler(client.BaseClient(), statusSetter),
	}
}

type endpointReconciler struct {
	base reconcile.Reconciler
}

func (r *endpointReconciler) Reconcile(namespace string, desiredResources EndpointList, transition TransitionEndpointFunc, opts clients.ListOpts) error {
	opts = opts.WithDefaults()
	opts.Ctx = contextutils.WithLogger(opts.Ctx, "endpoint_reconciler")
	var transitionResources reconcile.TransitionResourcesFunc
	if transition != nil {
		transitionResources = func(original, desired resources.Resource) (bool, error) {
			return transition(original.(*Endpoint), desired.(*Endpoint))
		}
	}
	return r.base.Reconcile(namespace, endpointsToResources(desiredResources), transitionResources, opts)
}
