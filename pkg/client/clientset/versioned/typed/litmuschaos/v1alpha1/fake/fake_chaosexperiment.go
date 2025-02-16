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

	v1alpha1 "github.com/litmuschaos/chaos-operator/api/litmuschaos/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeChaosExperiments implements ChaosExperimentInterface
type FakeChaosExperiments struct {
	Fake *FakeLitmuschaosV1alpha1
	ns   string
}

var chaosexperimentsResource = schema.GroupVersionResource{Group: "litmuschaos", Version: "v1alpha1", Resource: "chaosexperiments"}

var chaosexperimentsKind = schema.GroupVersionKind{Group: "litmuschaos", Version: "v1alpha1", Kind: "ChaosExperiment"}

// Get takes name of the chaosExperiment, and returns the corresponding chaosExperiment object, and an error if there is any.
func (c *FakeChaosExperiments) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.ChaosExperiment, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(chaosexperimentsResource, c.ns, name), &v1alpha1.ChaosExperiment{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ChaosExperiment), err
}

// List takes label and field selectors, and returns the list of ChaosExperiments that match those selectors.
func (c *FakeChaosExperiments) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.ChaosExperimentList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(chaosexperimentsResource, chaosexperimentsKind, c.ns, opts), &v1alpha1.ChaosExperimentList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.ChaosExperimentList{ListMeta: obj.(*v1alpha1.ChaosExperimentList).ListMeta}
	for _, item := range obj.(*v1alpha1.ChaosExperimentList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested chaosExperiments.
func (c *FakeChaosExperiments) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(chaosexperimentsResource, c.ns, opts))

}

// Create takes the representation of a chaosExperiment and creates it.  Returns the server's representation of the chaosExperiment, and an error, if there is any.
func (c *FakeChaosExperiments) Create(ctx context.Context, chaosExperiment *v1alpha1.ChaosExperiment, opts v1.CreateOptions) (result *v1alpha1.ChaosExperiment, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(chaosexperimentsResource, c.ns, chaosExperiment), &v1alpha1.ChaosExperiment{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ChaosExperiment), err
}

// Update takes the representation of a chaosExperiment and updates it. Returns the server's representation of the chaosExperiment, and an error, if there is any.
func (c *FakeChaosExperiments) Update(ctx context.Context, chaosExperiment *v1alpha1.ChaosExperiment, opts v1.UpdateOptions) (result *v1alpha1.ChaosExperiment, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(chaosexperimentsResource, c.ns, chaosExperiment), &v1alpha1.ChaosExperiment{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ChaosExperiment), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeChaosExperiments) UpdateStatus(ctx context.Context, chaosExperiment *v1alpha1.ChaosExperiment, opts v1.UpdateOptions) (*v1alpha1.ChaosExperiment, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(chaosexperimentsResource, "status", c.ns, chaosExperiment), &v1alpha1.ChaosExperiment{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ChaosExperiment), err
}

// Delete takes name of the chaosExperiment and deletes it. Returns an error if one occurs.
func (c *FakeChaosExperiments) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(chaosexperimentsResource, c.ns, name), &v1alpha1.ChaosExperiment{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeChaosExperiments) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(chaosexperimentsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.ChaosExperimentList{})
	return err
}

// Patch applies the patch and returns the patched chaosExperiment.
func (c *FakeChaosExperiments) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.ChaosExperiment, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(chaosexperimentsResource, c.ns, name, pt, data, subresources...), &v1alpha1.ChaosExperiment{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ChaosExperiment), err
}
