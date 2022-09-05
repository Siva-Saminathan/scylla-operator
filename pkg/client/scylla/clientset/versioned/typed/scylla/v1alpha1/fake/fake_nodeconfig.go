// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"

	v1alpha1 "github.com/scylladb/scylla-operator/pkg/api/scylla/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeNodeConfigs implements NodeConfigInterface
type FakeNodeConfigs struct {
	Fake *FakeScyllaV1alpha1
}

var nodeconfigsResource = schema.GroupVersionResource{Group: "scylla.scylladb.com", Version: "v1alpha1", Resource: "nodeconfigs"}

var nodeconfigsKind = schema.GroupVersionKind{Group: "scylla.scylladb.com", Version: "v1alpha1", Kind: "NodeConfig"}

// Get takes name of the nodeConfig, and returns the corresponding nodeConfig object, and an error if there is any.
func (c *FakeNodeConfigs) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.NodeConfig, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(nodeconfigsResource, name), &v1alpha1.NodeConfig{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.NodeConfig), err
}

// List takes label and field selectors, and returns the list of NodeConfigs that match those selectors.
func (c *FakeNodeConfigs) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.NodeConfigList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(nodeconfigsResource, nodeconfigsKind, opts), &v1alpha1.NodeConfigList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.NodeConfigList{ListMeta: obj.(*v1alpha1.NodeConfigList).ListMeta}
	for _, item := range obj.(*v1alpha1.NodeConfigList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested nodeConfigs.
func (c *FakeNodeConfigs) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(nodeconfigsResource, opts))
}

// Create takes the representation of a nodeConfig and creates it.  Returns the server's representation of the nodeConfig, and an error, if there is any.
func (c *FakeNodeConfigs) Create(ctx context.Context, nodeConfig *v1alpha1.NodeConfig, opts v1.CreateOptions) (result *v1alpha1.NodeConfig, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(nodeconfigsResource, nodeConfig), &v1alpha1.NodeConfig{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.NodeConfig), err
}

// Update takes the representation of a nodeConfig and updates it. Returns the server's representation of the nodeConfig, and an error, if there is any.
func (c *FakeNodeConfigs) Update(ctx context.Context, nodeConfig *v1alpha1.NodeConfig, opts v1.UpdateOptions) (result *v1alpha1.NodeConfig, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(nodeconfigsResource, nodeConfig), &v1alpha1.NodeConfig{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.NodeConfig), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeNodeConfigs) UpdateStatus(ctx context.Context, nodeConfig *v1alpha1.NodeConfig, opts v1.UpdateOptions) (*v1alpha1.NodeConfig, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(nodeconfigsResource, "status", nodeConfig), &v1alpha1.NodeConfig{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.NodeConfig), err
}

// Delete takes name of the nodeConfig and deletes it. Returns an error if one occurs.
func (c *FakeNodeConfigs) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteActionWithOptions(nodeconfigsResource, name, opts), &v1alpha1.NodeConfig{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeNodeConfigs) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(nodeconfigsResource, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.NodeConfigList{})
	return err
}

// Patch applies the patch and returns the patched nodeConfig.
func (c *FakeNodeConfigs) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.NodeConfig, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(nodeconfigsResource, name, pt, data, subresources...), &v1alpha1.NodeConfig{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.NodeConfig), err
}
