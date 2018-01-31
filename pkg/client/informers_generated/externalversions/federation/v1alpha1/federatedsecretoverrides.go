/*
Copyright 2018 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// This file was automatically generated by informer-gen

package v1alpha1

import (
	federation_v1alpha1 "github.com/marun/fnord/pkg/apis/federation/v1alpha1"
	clientset "github.com/marun/fnord/pkg/client/clientset_generated/clientset"
	internalinterfaces "github.com/marun/fnord/pkg/client/informers_generated/externalversions/internalinterfaces"
	v1alpha1 "github.com/marun/fnord/pkg/client/listers_generated/federation/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
	time "time"
)

// FederatedSecretOverridesInformer provides access to a shared informer and lister for
// FederatedSecretOverrideses.
type FederatedSecretOverridesInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.FederatedSecretOverridesLister
}

type federatedSecretOverridesInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewFederatedSecretOverridesInformer constructs a new informer for FederatedSecretOverrides type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFederatedSecretOverridesInformer(client clientset.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredFederatedSecretOverridesInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredFederatedSecretOverridesInformer constructs a new informer for FederatedSecretOverrides type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredFederatedSecretOverridesInformer(client clientset.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.FederationV1alpha1().FederatedSecretOverrideses(namespace).List(options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.FederationV1alpha1().FederatedSecretOverrideses(namespace).Watch(options)
			},
		},
		&federation_v1alpha1.FederatedSecretOverrides{},
		resyncPeriod,
		indexers,
	)
}

func (f *federatedSecretOverridesInformer) defaultInformer(client clientset.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredFederatedSecretOverridesInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *federatedSecretOverridesInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&federation_v1alpha1.FederatedSecretOverrides{}, f.defaultInformer)
}

func (f *federatedSecretOverridesInformer) Lister() v1alpha1.FederatedSecretOverridesLister {
	return v1alpha1.NewFederatedSecretOverridesLister(f.Informer().GetIndexer())
}
