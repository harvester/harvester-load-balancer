/*
Copyright 2023 Rancher Labs, Inc.

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

// Code generated by main. DO NOT EDIT.

package fake

import (
	"context"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
	corev1 "kubevirt.io/api/core/v1"
)

// FakeKubeVirts implements KubeVirtInterface
type FakeKubeVirts struct {
	Fake *FakeKubevirtV1
	ns   string
}

var kubevirtsResource = schema.GroupVersionResource{Group: "kubevirt.io", Version: "v1", Resource: "kubevirts"}

var kubevirtsKind = schema.GroupVersionKind{Group: "kubevirt.io", Version: "v1", Kind: "KubeVirt"}

// Get takes name of the kubeVirt, and returns the corresponding kubeVirt object, and an error if there is any.
func (c *FakeKubeVirts) Get(ctx context.Context, name string, options v1.GetOptions) (result *corev1.KubeVirt, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(kubevirtsResource, c.ns, name), &corev1.KubeVirt{})

	if obj == nil {
		return nil, err
	}
	return obj.(*corev1.KubeVirt), err
}

// List takes label and field selectors, and returns the list of KubeVirts that match those selectors.
func (c *FakeKubeVirts) List(ctx context.Context, opts v1.ListOptions) (result *corev1.KubeVirtList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(kubevirtsResource, kubevirtsKind, c.ns, opts), &corev1.KubeVirtList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &corev1.KubeVirtList{ListMeta: obj.(*corev1.KubeVirtList).ListMeta}
	for _, item := range obj.(*corev1.KubeVirtList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested kubeVirts.
func (c *FakeKubeVirts) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(kubevirtsResource, c.ns, opts))

}

// Create takes the representation of a kubeVirt and creates it.  Returns the server's representation of the kubeVirt, and an error, if there is any.
func (c *FakeKubeVirts) Create(ctx context.Context, kubeVirt *corev1.KubeVirt, opts v1.CreateOptions) (result *corev1.KubeVirt, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(kubevirtsResource, c.ns, kubeVirt), &corev1.KubeVirt{})

	if obj == nil {
		return nil, err
	}
	return obj.(*corev1.KubeVirt), err
}

// Update takes the representation of a kubeVirt and updates it. Returns the server's representation of the kubeVirt, and an error, if there is any.
func (c *FakeKubeVirts) Update(ctx context.Context, kubeVirt *corev1.KubeVirt, opts v1.UpdateOptions) (result *corev1.KubeVirt, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(kubevirtsResource, c.ns, kubeVirt), &corev1.KubeVirt{})

	if obj == nil {
		return nil, err
	}
	return obj.(*corev1.KubeVirt), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeKubeVirts) UpdateStatus(ctx context.Context, kubeVirt *corev1.KubeVirt, opts v1.UpdateOptions) (*corev1.KubeVirt, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(kubevirtsResource, "status", c.ns, kubeVirt), &corev1.KubeVirt{})

	if obj == nil {
		return nil, err
	}
	return obj.(*corev1.KubeVirt), err
}

// Delete takes name of the kubeVirt and deletes it. Returns an error if one occurs.
func (c *FakeKubeVirts) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(kubevirtsResource, c.ns, name, opts), &corev1.KubeVirt{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeKubeVirts) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(kubevirtsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &corev1.KubeVirtList{})
	return err
}

// Patch applies the patch and returns the patched kubeVirt.
func (c *FakeKubeVirts) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *corev1.KubeVirt, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(kubevirtsResource, c.ns, name, pt, data, subresources...), &corev1.KubeVirt{})

	if obj == nil {
		return nil, err
	}
	return obj.(*corev1.KubeVirt), err
}
