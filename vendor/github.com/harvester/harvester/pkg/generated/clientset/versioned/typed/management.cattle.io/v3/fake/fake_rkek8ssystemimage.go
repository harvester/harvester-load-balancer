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

	v3 "github.com/rancher/rancher/pkg/apis/management.cattle.io/v3"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeRkeK8sSystemImages implements RkeK8sSystemImageInterface
type FakeRkeK8sSystemImages struct {
	Fake *FakeManagementV3
	ns   string
}

var rkek8ssystemimagesResource = schema.GroupVersionResource{Group: "management.cattle.io", Version: "v3", Resource: "rkek8ssystemimages"}

var rkek8ssystemimagesKind = schema.GroupVersionKind{Group: "management.cattle.io", Version: "v3", Kind: "RkeK8sSystemImage"}

// Get takes name of the rkeK8sSystemImage, and returns the corresponding rkeK8sSystemImage object, and an error if there is any.
func (c *FakeRkeK8sSystemImages) Get(ctx context.Context, name string, options v1.GetOptions) (result *v3.RkeK8sSystemImage, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(rkek8ssystemimagesResource, c.ns, name), &v3.RkeK8sSystemImage{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v3.RkeK8sSystemImage), err
}

// List takes label and field selectors, and returns the list of RkeK8sSystemImages that match those selectors.
func (c *FakeRkeK8sSystemImages) List(ctx context.Context, opts v1.ListOptions) (result *v3.RkeK8sSystemImageList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(rkek8ssystemimagesResource, rkek8ssystemimagesKind, c.ns, opts), &v3.RkeK8sSystemImageList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v3.RkeK8sSystemImageList{ListMeta: obj.(*v3.RkeK8sSystemImageList).ListMeta}
	for _, item := range obj.(*v3.RkeK8sSystemImageList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested rkeK8sSystemImages.
func (c *FakeRkeK8sSystemImages) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(rkek8ssystemimagesResource, c.ns, opts))

}

// Create takes the representation of a rkeK8sSystemImage and creates it.  Returns the server's representation of the rkeK8sSystemImage, and an error, if there is any.
func (c *FakeRkeK8sSystemImages) Create(ctx context.Context, rkeK8sSystemImage *v3.RkeK8sSystemImage, opts v1.CreateOptions) (result *v3.RkeK8sSystemImage, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(rkek8ssystemimagesResource, c.ns, rkeK8sSystemImage), &v3.RkeK8sSystemImage{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v3.RkeK8sSystemImage), err
}

// Update takes the representation of a rkeK8sSystemImage and updates it. Returns the server's representation of the rkeK8sSystemImage, and an error, if there is any.
func (c *FakeRkeK8sSystemImages) Update(ctx context.Context, rkeK8sSystemImage *v3.RkeK8sSystemImage, opts v1.UpdateOptions) (result *v3.RkeK8sSystemImage, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(rkek8ssystemimagesResource, c.ns, rkeK8sSystemImage), &v3.RkeK8sSystemImage{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v3.RkeK8sSystemImage), err
}

// Delete takes name of the rkeK8sSystemImage and deletes it. Returns an error if one occurs.
func (c *FakeRkeK8sSystemImages) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(rkek8ssystemimagesResource, c.ns, name, opts), &v3.RkeK8sSystemImage{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeRkeK8sSystemImages) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(rkek8ssystemimagesResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v3.RkeK8sSystemImageList{})
	return err
}

// Patch applies the patch and returns the patched rkeK8sSystemImage.
func (c *FakeRkeK8sSystemImages) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v3.RkeK8sSystemImage, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(rkek8ssystemimagesResource, c.ns, name, pt, data, subresources...), &v3.RkeK8sSystemImage{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v3.RkeK8sSystemImage), err
}
