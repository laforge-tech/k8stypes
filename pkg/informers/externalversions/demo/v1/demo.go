//-----------------------------------------------------------------------------
// Demo API
//-----------------------------------------------------------------------------

// Code generated by informer-gen. DO NOT EDIT.

package v1

import (
	"context"
	time "time"

	demov1 "github.com/laforge-tech/k8stypes/pkg/api/demo/v1"
	client "github.com/laforge-tech/k8stypes/pkg/client"
	internalinterfaces "github.com/laforge-tech/k8stypes/pkg/informers/externalversions/internalinterfaces"
	v1 "github.com/laforge-tech/k8stypes/pkg/listers/demo/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// DemoInformer provides access to a shared informer and lister for
// Demos.
type DemoInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1.DemoLister
}

type demoInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewDemoInformer constructs a new informer for Demo type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewDemoInformer(client client.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredDemoInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredDemoInformer constructs a new informer for Demo type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredDemoInformer(client client.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options metav1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.DemoV1().Demos(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options metav1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.DemoV1().Demos(namespace).Watch(context.TODO(), options)
			},
		},
		&demov1.Demo{},
		resyncPeriod,
		indexers,
	)
}

func (f *demoInformer) defaultInformer(client client.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredDemoInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *demoInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&demov1.Demo{}, f.defaultInformer)
}

func (f *demoInformer) Lister() v1.DemoLister {
	return v1.NewDemoLister(f.Informer().GetIndexer())
}
