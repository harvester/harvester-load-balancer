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

// FakeClusterAlertRules implements ClusterAlertRuleInterface
type FakeClusterAlertRules struct {
	Fake *FakeManagementV3
	ns   string
}

var clusteralertrulesResource = schema.GroupVersionResource{Group: "management.cattle.io", Version: "v3", Resource: "clusteralertrules"}

var clusteralertrulesKind = schema.GroupVersionKind{Group: "management.cattle.io", Version: "v3", Kind: "ClusterAlertRule"}

// Get takes name of the clusterAlertRule, and returns the corresponding clusterAlertRule object, and an error if there is any.
func (c *FakeClusterAlertRules) Get(ctx context.Context, name string, options v1.GetOptions) (result *v3.ClusterAlertRule, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(clusteralertrulesResource, c.ns, name), &v3.ClusterAlertRule{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v3.ClusterAlertRule), err
}

// List takes label and field selectors, and returns the list of ClusterAlertRules that match those selectors.
func (c *FakeClusterAlertRules) List(ctx context.Context, opts v1.ListOptions) (result *v3.ClusterAlertRuleList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(clusteralertrulesResource, clusteralertrulesKind, c.ns, opts), &v3.ClusterAlertRuleList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v3.ClusterAlertRuleList{ListMeta: obj.(*v3.ClusterAlertRuleList).ListMeta}
	for _, item := range obj.(*v3.ClusterAlertRuleList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested clusterAlertRules.
func (c *FakeClusterAlertRules) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(clusteralertrulesResource, c.ns, opts))

}

// Create takes the representation of a clusterAlertRule and creates it.  Returns the server's representation of the clusterAlertRule, and an error, if there is any.
func (c *FakeClusterAlertRules) Create(ctx context.Context, clusterAlertRule *v3.ClusterAlertRule, opts v1.CreateOptions) (result *v3.ClusterAlertRule, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(clusteralertrulesResource, c.ns, clusterAlertRule), &v3.ClusterAlertRule{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v3.ClusterAlertRule), err
}

// Update takes the representation of a clusterAlertRule and updates it. Returns the server's representation of the clusterAlertRule, and an error, if there is any.
func (c *FakeClusterAlertRules) Update(ctx context.Context, clusterAlertRule *v3.ClusterAlertRule, opts v1.UpdateOptions) (result *v3.ClusterAlertRule, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(clusteralertrulesResource, c.ns, clusterAlertRule), &v3.ClusterAlertRule{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v3.ClusterAlertRule), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeClusterAlertRules) UpdateStatus(ctx context.Context, clusterAlertRule *v3.ClusterAlertRule, opts v1.UpdateOptions) (*v3.ClusterAlertRule, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(clusteralertrulesResource, "status", c.ns, clusterAlertRule), &v3.ClusterAlertRule{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v3.ClusterAlertRule), err
}

// Delete takes name of the clusterAlertRule and deletes it. Returns an error if one occurs.
func (c *FakeClusterAlertRules) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(clusteralertrulesResource, c.ns, name, opts), &v3.ClusterAlertRule{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeClusterAlertRules) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(clusteralertrulesResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v3.ClusterAlertRuleList{})
	return err
}

// Patch applies the patch and returns the patched clusterAlertRule.
func (c *FakeClusterAlertRules) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v3.ClusterAlertRule, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(clusteralertrulesResource, c.ns, name, pt, data, subresources...), &v3.ClusterAlertRule{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v3.ClusterAlertRule), err
}
