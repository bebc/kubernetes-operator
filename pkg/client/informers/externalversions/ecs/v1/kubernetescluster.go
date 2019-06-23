/*
Copyright The Kubernetes Authors.

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

// Code generated by informer-gen. DO NOT EDIT.

package v1

import (
	time "time"

	ecsv1 "github.com/gosoon/kubernetes-operator/pkg/apis/ecs/v1"
	versioned "github.com/gosoon/kubernetes-operator/pkg/client/clientset/versioned"
	internalinterfaces "github.com/gosoon/kubernetes-operator/pkg/client/informers/externalversions/internalinterfaces"
	v1 "github.com/gosoon/kubernetes-operator/pkg/client/listers/ecs/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// KubernetesClusterInformer provides access to a shared informer and lister for
// KubernetesClusters.
type KubernetesClusterInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1.KubernetesClusterLister
}

type kubernetesClusterInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// NewKubernetesClusterInformer constructs a new informer for KubernetesCluster type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewKubernetesClusterInformer(client versioned.Interface, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredKubernetesClusterInformer(client, resyncPeriod, indexers, nil)
}

// NewFilteredKubernetesClusterInformer constructs a new informer for KubernetesCluster type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredKubernetesClusterInformer(client versioned.Interface, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options metav1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.EcsV1().KubernetesClusters().List(options)
			},
			WatchFunc: func(options metav1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.EcsV1().KubernetesClusters().Watch(options)
			},
		},
		&ecsv1.KubernetesCluster{},
		resyncPeriod,
		indexers,
	)
}

func (f *kubernetesClusterInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredKubernetesClusterInformer(client, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *kubernetesClusterInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&ecsv1.KubernetesCluster{}, f.defaultInformer)
}

func (f *kubernetesClusterInformer) Lister() v1.KubernetesClusterLister {
	return v1.NewKubernetesClusterLister(f.Informer().GetIndexer())
}
