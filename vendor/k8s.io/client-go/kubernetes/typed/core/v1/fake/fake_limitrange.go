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

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"
	json "encoding/json"
	"fmt"

	corev1 "k8s.io/api/core/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	applyconfigurationscorev1 "k8s.io/client-go/applyconfigurations/core/v1"
	testing "k8s.io/client-go/testing"
)

// FakeLimitRanges implements LimitRangeInterface
type FakeLimitRanges struct {
	Fake *FakeCoreV1
	ns   string
}

var limitrangesResource = schema.GroupVersionResource{Group: "", Version: "v1", Resource: "limitranges"}

var limitrangesKind = schema.GroupVersionKind{Group: "", Version: "v1", Kind: "LimitRange"}

// Get takes name of the limitRange, and returns the corresponding limitRange object, and an error if there is any.
func (c *FakeLimitRanges) Get(ctx context.Context, name string, options v1.GetOptions) (result *corev1.LimitRange, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(limitrangesResource, c.ns, name), &corev1.LimitRange{})

	if obj == nil {
		return nil, err
	}
	return obj.(*corev1.LimitRange), err
}

// List takes label and field selectors, and returns the list of LimitRanges that match those selectors.
func (c *FakeLimitRanges) List(ctx context.Context, opts v1.ListOptions) (result *corev1.LimitRangeList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(limitrangesResource, limitrangesKind, c.ns, opts), &corev1.LimitRangeList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &corev1.LimitRangeList{ListMeta: obj.(*corev1.LimitRangeList).ListMeta}
	for _, item := range obj.(*corev1.LimitRangeList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested limitRanges.
func (c *FakeLimitRanges) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(limitrangesResource, c.ns, opts))

}

// Create takes the representation of a limitRange and creates it.  Returns the server's representation of the limitRange, and an error, if there is any.
func (c *FakeLimitRanges) Create(ctx context.Context, limitRange *corev1.LimitRange, opts v1.CreateOptions) (result *corev1.LimitRange, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(limitrangesResource, c.ns, limitRange), &corev1.LimitRange{})

	if obj == nil {
		return nil, err
	}
	return obj.(*corev1.LimitRange), err
}

// Update takes the representation of a limitRange and updates it. Returns the server's representation of the limitRange, and an error, if there is any.
func (c *FakeLimitRanges) Update(ctx context.Context, limitRange *corev1.LimitRange, opts v1.UpdateOptions) (result *corev1.LimitRange, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(limitrangesResource, c.ns, limitRange), &corev1.LimitRange{})

	if obj == nil {
		return nil, err
	}
	return obj.(*corev1.LimitRange), err
}

// Delete takes name of the limitRange and deletes it. Returns an error if one occurs.
func (c *FakeLimitRanges) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(limitrangesResource, c.ns, name), &corev1.LimitRange{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeLimitRanges) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(limitrangesResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &corev1.LimitRangeList{})
	return err
}

// Patch applies the patch and returns the patched limitRange.
func (c *FakeLimitRanges) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *corev1.LimitRange, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(limitrangesResource, c.ns, name, pt, data, subresources...), &corev1.LimitRange{})

	if obj == nil {
		return nil, err
	}
	return obj.(*corev1.LimitRange), err
}

// Apply takes the given apply declarative configuration, applies it and returns the applied limitRange.
func (c *FakeLimitRanges) Apply(ctx context.Context, limitRange *applyconfigurationscorev1.LimitRangeApplyConfiguration, opts v1.ApplyOptions) (result *corev1.LimitRange, err error) {
	if limitRange == nil {
		return nil, fmt.Errorf("limitRange provided to Apply must not be nil")
	}
	data, err := json.Marshal(limitRange)
	if err != nil {
		return nil, err
	}
	name := limitRange.Name
	if name == nil {
		return nil, fmt.Errorf("limitRange.Name must be provided to Apply")
	}
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(limitrangesResource, c.ns, *name, types.ApplyPatchType, data), &corev1.LimitRange{})

	if obj == nil {
		return nil, err
	}
	return obj.(*corev1.LimitRange), err
}
