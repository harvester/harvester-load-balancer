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

	v1beta1 "github.com/harvester/harvester/pkg/apis/harvesterhci.io/v1beta1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeUpgradeLogs implements UpgradeLogInterface
type FakeUpgradeLogs struct {
	Fake *FakeHarvesterhciV1beta1
	ns   string
}

var upgradelogsResource = schema.GroupVersionResource{Group: "harvesterhci.io", Version: "v1beta1", Resource: "upgradelogs"}

var upgradelogsKind = schema.GroupVersionKind{Group: "harvesterhci.io", Version: "v1beta1", Kind: "UpgradeLog"}

// Get takes name of the upgradeLog, and returns the corresponding upgradeLog object, and an error if there is any.
func (c *FakeUpgradeLogs) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1beta1.UpgradeLog, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(upgradelogsResource, c.ns, name), &v1beta1.UpgradeLog{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.UpgradeLog), err
}

// List takes label and field selectors, and returns the list of UpgradeLogs that match those selectors.
func (c *FakeUpgradeLogs) List(ctx context.Context, opts v1.ListOptions) (result *v1beta1.UpgradeLogList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(upgradelogsResource, upgradelogsKind, c.ns, opts), &v1beta1.UpgradeLogList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1beta1.UpgradeLogList{ListMeta: obj.(*v1beta1.UpgradeLogList).ListMeta}
	for _, item := range obj.(*v1beta1.UpgradeLogList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested upgradeLogs.
func (c *FakeUpgradeLogs) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(upgradelogsResource, c.ns, opts))

}

// Create takes the representation of a upgradeLog and creates it.  Returns the server's representation of the upgradeLog, and an error, if there is any.
func (c *FakeUpgradeLogs) Create(ctx context.Context, upgradeLog *v1beta1.UpgradeLog, opts v1.CreateOptions) (result *v1beta1.UpgradeLog, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(upgradelogsResource, c.ns, upgradeLog), &v1beta1.UpgradeLog{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.UpgradeLog), err
}

// Update takes the representation of a upgradeLog and updates it. Returns the server's representation of the upgradeLog, and an error, if there is any.
func (c *FakeUpgradeLogs) Update(ctx context.Context, upgradeLog *v1beta1.UpgradeLog, opts v1.UpdateOptions) (result *v1beta1.UpgradeLog, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(upgradelogsResource, c.ns, upgradeLog), &v1beta1.UpgradeLog{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.UpgradeLog), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeUpgradeLogs) UpdateStatus(ctx context.Context, upgradeLog *v1beta1.UpgradeLog, opts v1.UpdateOptions) (*v1beta1.UpgradeLog, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(upgradelogsResource, "status", c.ns, upgradeLog), &v1beta1.UpgradeLog{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.UpgradeLog), err
}

// Delete takes name of the upgradeLog and deletes it. Returns an error if one occurs.
func (c *FakeUpgradeLogs) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(upgradelogsResource, c.ns, name, opts), &v1beta1.UpgradeLog{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeUpgradeLogs) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(upgradelogsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1beta1.UpgradeLogList{})
	return err
}

// Patch applies the patch and returns the patched upgradeLog.
func (c *FakeUpgradeLogs) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta1.UpgradeLog, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(upgradelogsResource, c.ns, name, pt, data, subresources...), &v1beta1.UpgradeLog{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.UpgradeLog), err
}
