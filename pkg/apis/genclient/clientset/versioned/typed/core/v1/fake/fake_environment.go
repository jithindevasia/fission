/*
Copyright The Fission Authors.

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

	corev1 "github.com/fission/fission/pkg/apis/core/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeEnvironments implements EnvironmentInterface
type FakeEnvironments struct {
	Fake *FakeCoreV1
	ns   string
}

var environmentsResource = schema.GroupVersionResource{Group: "fission.io", Version: "v1", Resource: "environments"}

var environmentsKind = schema.GroupVersionKind{Group: "fission.io", Version: "v1", Kind: "Environment"}

// Get takes name of the _environment, and returns the corresponding environment object, and an error if there is any.
func (c *FakeEnvironments) Get(ctx context.Context, name string, options v1.GetOptions) (result *corev1.Environment, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(environmentsResource, c.ns, name), &corev1.Environment{})

	if obj == nil {
		return nil, err
	}
	return obj.(*corev1.Environment), err
}

// List takes label and field selectors, and returns the list of Environments that match those selectors.
func (c *FakeEnvironments) List(ctx context.Context, opts v1.ListOptions) (result *corev1.EnvironmentList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(environmentsResource, environmentsKind, c.ns, opts), &corev1.EnvironmentList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &corev1.EnvironmentList{ListMeta: obj.(*corev1.EnvironmentList).ListMeta}
	for _, item := range obj.(*corev1.EnvironmentList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested environments.
func (c *FakeEnvironments) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(environmentsResource, c.ns, opts))

}

// Create takes the representation of a _environment and creates it.  Returns the server's representation of the environment, and an error, if there is any.
func (c *FakeEnvironments) Create(ctx context.Context, _environment *corev1.Environment, opts v1.CreateOptions) (result *corev1.Environment, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(environmentsResource, c.ns, _environment), &corev1.Environment{})

	if obj == nil {
		return nil, err
	}
	return obj.(*corev1.Environment), err
}

// Update takes the representation of a _environment and updates it. Returns the server's representation of the environment, and an error, if there is any.
func (c *FakeEnvironments) Update(ctx context.Context, _environment *corev1.Environment, opts v1.UpdateOptions) (result *corev1.Environment, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(environmentsResource, c.ns, _environment), &corev1.Environment{})

	if obj == nil {
		return nil, err
	}
	return obj.(*corev1.Environment), err
}

// Delete takes name of the _environment and deletes it. Returns an error if one occurs.
func (c *FakeEnvironments) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(environmentsResource, c.ns, name), &corev1.Environment{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeEnvironments) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(environmentsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &corev1.EnvironmentList{})
	return err
}

// Patch applies the patch and returns the patched environment.
func (c *FakeEnvironments) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *corev1.Environment, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(environmentsResource, c.ns, name, pt, data, subresources...), &corev1.Environment{})

	if obj == nil {
		return nil, err
	}
	return obj.(*corev1.Environment), err
}
