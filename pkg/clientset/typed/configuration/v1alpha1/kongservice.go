/*
Copyright 2021 Kong, Inc.

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

// Code generated by client-gen. DO NOT EDIT.

package v1alpha1

import (
	"context"

	v1alpha1 "github.com/kong/kubernetes-configuration/api/configuration/v1alpha1"
	scheme "github.com/kong/kubernetes-configuration/pkg/clientset/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	gentype "k8s.io/client-go/gentype"
)

// KongServicesGetter has a method to return a KongServiceInterface.
// A group's client should implement this interface.
type KongServicesGetter interface {
	KongServices(namespace string) KongServiceInterface
}

// KongServiceInterface has methods to work with KongService resources.
type KongServiceInterface interface {
	Create(ctx context.Context, kongService *v1alpha1.KongService, opts v1.CreateOptions) (*v1alpha1.KongService, error)
	Update(ctx context.Context, kongService *v1alpha1.KongService, opts v1.UpdateOptions) (*v1alpha1.KongService, error)
	// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
	UpdateStatus(ctx context.Context, kongService *v1alpha1.KongService, opts v1.UpdateOptions) (*v1alpha1.KongService, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha1.KongService, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha1.KongServiceList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.KongService, err error)
	KongServiceExpansion
}

// kongServices implements KongServiceInterface
type kongServices struct {
	*gentype.ClientWithList[*v1alpha1.KongService, *v1alpha1.KongServiceList]
}

// newKongServices returns a KongServices
func newKongServices(c *ConfigurationV1alpha1Client, namespace string) *kongServices {
	return &kongServices{
		gentype.NewClientWithList[*v1alpha1.KongService, *v1alpha1.KongServiceList](
			"kongservices",
			c.RESTClient(),
			scheme.ParameterCodec,
			namespace,
			func() *v1alpha1.KongService { return &v1alpha1.KongService{} },
			func() *v1alpha1.KongServiceList { return &v1alpha1.KongServiceList{} }),
	}
}
