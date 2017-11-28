// This file was automatically generated by informer-gen

package v1

import (
	internalinterfaces "github.com/openshift/client-go/route/informers/externalversions/internalinterfaces"
)

// Interface provides access to all the informers in this group version.
type Interface interface {
	// Routes returns a RouteInformer.
	Routes() RouteInformer
}

type version struct {
	internalinterfaces.SharedInformerFactory
}

// New returns a new Interface.
func New(f internalinterfaces.SharedInformerFactory) Interface {
	return &version{f}
}

// Routes returns a RouteInformer.
func (v *version) Routes() RouteInformer {
	return &routeInformer{factory: v.SharedInformerFactory}
}
