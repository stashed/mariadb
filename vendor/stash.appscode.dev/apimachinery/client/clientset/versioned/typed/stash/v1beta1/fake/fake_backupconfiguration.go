/*
Copyright AppsCode Inc. and Contributors

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

package fake

import (
	"context"

	v1beta1 "stash.appscode.dev/apimachinery/apis/stash/v1beta1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeBackupConfigurations implements BackupConfigurationInterface
type FakeBackupConfigurations struct {
	Fake *FakeStashV1beta1
	ns   string
}

var backupconfigurationsResource = v1beta1.SchemeGroupVersion.WithResource("backupconfigurations")

var backupconfigurationsKind = v1beta1.SchemeGroupVersion.WithKind("BackupConfiguration")

// Get takes name of the backupConfiguration, and returns the corresponding backupConfiguration object, and an error if there is any.
func (c *FakeBackupConfigurations) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1beta1.BackupConfiguration, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(backupconfigurationsResource, c.ns, name), &v1beta1.BackupConfiguration{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.BackupConfiguration), err
}

// List takes label and field selectors, and returns the list of BackupConfigurations that match those selectors.
func (c *FakeBackupConfigurations) List(ctx context.Context, opts v1.ListOptions) (result *v1beta1.BackupConfigurationList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(backupconfigurationsResource, backupconfigurationsKind, c.ns, opts), &v1beta1.BackupConfigurationList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1beta1.BackupConfigurationList{ListMeta: obj.(*v1beta1.BackupConfigurationList).ListMeta}
	for _, item := range obj.(*v1beta1.BackupConfigurationList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested backupConfigurations.
func (c *FakeBackupConfigurations) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(backupconfigurationsResource, c.ns, opts))

}

// Create takes the representation of a backupConfiguration and creates it.  Returns the server's representation of the backupConfiguration, and an error, if there is any.
func (c *FakeBackupConfigurations) Create(ctx context.Context, backupConfiguration *v1beta1.BackupConfiguration, opts v1.CreateOptions) (result *v1beta1.BackupConfiguration, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(backupconfigurationsResource, c.ns, backupConfiguration), &v1beta1.BackupConfiguration{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.BackupConfiguration), err
}

// Update takes the representation of a backupConfiguration and updates it. Returns the server's representation of the backupConfiguration, and an error, if there is any.
func (c *FakeBackupConfigurations) Update(ctx context.Context, backupConfiguration *v1beta1.BackupConfiguration, opts v1.UpdateOptions) (result *v1beta1.BackupConfiguration, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(backupconfigurationsResource, c.ns, backupConfiguration), &v1beta1.BackupConfiguration{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.BackupConfiguration), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeBackupConfigurations) UpdateStatus(ctx context.Context, backupConfiguration *v1beta1.BackupConfiguration, opts v1.UpdateOptions) (*v1beta1.BackupConfiguration, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(backupconfigurationsResource, "status", c.ns, backupConfiguration), &v1beta1.BackupConfiguration{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.BackupConfiguration), err
}

// Delete takes name of the backupConfiguration and deletes it. Returns an error if one occurs.
func (c *FakeBackupConfigurations) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(backupconfigurationsResource, c.ns, name, opts), &v1beta1.BackupConfiguration{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeBackupConfigurations) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(backupconfigurationsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1beta1.BackupConfigurationList{})
	return err
}

// Patch applies the patch and returns the patched backupConfiguration.
func (c *FakeBackupConfigurations) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta1.BackupConfiguration, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(backupconfigurationsResource, c.ns, name, pt, data, subresources...), &v1beta1.BackupConfiguration{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.BackupConfiguration), err
}
