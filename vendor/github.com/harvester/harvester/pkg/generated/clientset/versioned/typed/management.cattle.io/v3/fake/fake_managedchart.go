/*
Copyright 2024 Rancher Labs, Inc.

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

// FakeManagedCharts implements ManagedChartInterface
type FakeManagedCharts struct {
	Fake *FakeManagementV3
	ns   string
}

var managedchartsResource = schema.GroupVersionResource{Group: "management.cattle.io", Version: "v3", Resource: "managedcharts"}

var managedchartsKind = schema.GroupVersionKind{Group: "management.cattle.io", Version: "v3", Kind: "ManagedChart"}

// Get takes name of the managedChart, and returns the corresponding managedChart object, and an error if there is any.
func (c *FakeManagedCharts) Get(ctx context.Context, name string, options v1.GetOptions) (result *v3.ManagedChart, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(managedchartsResource, c.ns, name), &v3.ManagedChart{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v3.ManagedChart), err
}

// List takes label and field selectors, and returns the list of ManagedCharts that match those selectors.
func (c *FakeManagedCharts) List(ctx context.Context, opts v1.ListOptions) (result *v3.ManagedChartList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(managedchartsResource, managedchartsKind, c.ns, opts), &v3.ManagedChartList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v3.ManagedChartList{ListMeta: obj.(*v3.ManagedChartList).ListMeta}
	for _, item := range obj.(*v3.ManagedChartList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested managedCharts.
func (c *FakeManagedCharts) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(managedchartsResource, c.ns, opts))

}

// Create takes the representation of a managedChart and creates it.  Returns the server's representation of the managedChart, and an error, if there is any.
func (c *FakeManagedCharts) Create(ctx context.Context, managedChart *v3.ManagedChart, opts v1.CreateOptions) (result *v3.ManagedChart, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(managedchartsResource, c.ns, managedChart), &v3.ManagedChart{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v3.ManagedChart), err
}

// Update takes the representation of a managedChart and updates it. Returns the server's representation of the managedChart, and an error, if there is any.
func (c *FakeManagedCharts) Update(ctx context.Context, managedChart *v3.ManagedChart, opts v1.UpdateOptions) (result *v3.ManagedChart, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(managedchartsResource, c.ns, managedChart), &v3.ManagedChart{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v3.ManagedChart), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeManagedCharts) UpdateStatus(ctx context.Context, managedChart *v3.ManagedChart, opts v1.UpdateOptions) (*v3.ManagedChart, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(managedchartsResource, "status", c.ns, managedChart), &v3.ManagedChart{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v3.ManagedChart), err
}

// Delete takes name of the managedChart and deletes it. Returns an error if one occurs.
func (c *FakeManagedCharts) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(managedchartsResource, c.ns, name, opts), &v3.ManagedChart{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeManagedCharts) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(managedchartsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v3.ManagedChartList{})
	return err
}

// Patch applies the patch and returns the patched managedChart.
func (c *FakeManagedCharts) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v3.ManagedChart, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(managedchartsResource, c.ns, name, pt, data, subresources...), &v3.ManagedChart{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v3.ManagedChart), err
}
