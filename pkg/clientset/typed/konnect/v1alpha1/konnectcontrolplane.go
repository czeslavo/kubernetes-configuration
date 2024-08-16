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

	v1alpha1 "github.com/kong/kubernetes-configuration/api/konnect/v1alpha1"
	scheme "github.com/kong/kubernetes-configuration/pkg/clientset/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	gentype "k8s.io/client-go/gentype"
)

// KonnectControlPlanesGetter has a method to return a KonnectControlPlaneInterface.
// A group's client should implement this interface.
type KonnectControlPlanesGetter interface {
	KonnectControlPlanes(namespace string) KonnectControlPlaneInterface
}

// KonnectControlPlaneInterface has methods to work with KonnectControlPlane resources.
type KonnectControlPlaneInterface interface {
	Create(ctx context.Context, konnectControlPlane *v1alpha1.KonnectControlPlane, opts v1.CreateOptions) (*v1alpha1.KonnectControlPlane, error)
	Update(ctx context.Context, konnectControlPlane *v1alpha1.KonnectControlPlane, opts v1.UpdateOptions) (*v1alpha1.KonnectControlPlane, error)
	// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
	UpdateStatus(ctx context.Context, konnectControlPlane *v1alpha1.KonnectControlPlane, opts v1.UpdateOptions) (*v1alpha1.KonnectControlPlane, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha1.KonnectControlPlane, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha1.KonnectControlPlaneList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.KonnectControlPlane, err error)
	KonnectControlPlaneExpansion
}

// konnectControlPlanes implements KonnectControlPlaneInterface
type konnectControlPlanes struct {
	*gentype.ClientWithList[*v1alpha1.KonnectControlPlane, *v1alpha1.KonnectControlPlaneList]
}

// newKonnectControlPlanes returns a KonnectControlPlanes
func newKonnectControlPlanes(c *KonnectV1alpha1Client, namespace string) *konnectControlPlanes {
	return &konnectControlPlanes{
		gentype.NewClientWithList[*v1alpha1.KonnectControlPlane, *v1alpha1.KonnectControlPlaneList](
			"konnectcontrolplanes",
			c.RESTClient(),
			scheme.ParameterCodec,
			namespace,
			func() *v1alpha1.KonnectControlPlane { return &v1alpha1.KonnectControlPlane{} },
			func() *v1alpha1.KonnectControlPlaneList { return &v1alpha1.KonnectControlPlaneList{} }),
	}
}
