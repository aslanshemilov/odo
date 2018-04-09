package fake

import (
	quota_v1 "github.com/openshift/api/quota/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeClusterResourceQuotas implements ClusterResourceQuotaInterface
type FakeClusterResourceQuotas struct {
	Fake *FakeQuotaV1
}

var clusterresourcequotasResource = schema.GroupVersionResource{Group: "quota.openshift.io", Version: "v1", Resource: "clusterresourcequotas"}

var clusterresourcequotasKind = schema.GroupVersionKind{Group: "quota.openshift.io", Version: "v1", Kind: "ClusterResourceQuota"}

// Get takes name of the clusterResourceQuota, and returns the corresponding clusterResourceQuota object, and an error if there is any.
func (c *FakeClusterResourceQuotas) Get(name string, options v1.GetOptions) (result *quota_v1.ClusterResourceQuota, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(clusterresourcequotasResource, name), &quota_v1.ClusterResourceQuota{})
	if obj == nil {
		return nil, err
	}
	return obj.(*quota_v1.ClusterResourceQuota), err
}

// List takes label and field selectors, and returns the list of ClusterResourceQuotas that match those selectors.
func (c *FakeClusterResourceQuotas) List(opts v1.ListOptions) (result *quota_v1.ClusterResourceQuotaList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(clusterresourcequotasResource, clusterresourcequotasKind, opts), &quota_v1.ClusterResourceQuotaList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &quota_v1.ClusterResourceQuotaList{}
	for _, item := range obj.(*quota_v1.ClusterResourceQuotaList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested clusterResourceQuotas.
func (c *FakeClusterResourceQuotas) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(clusterresourcequotasResource, opts))
}

// Create takes the representation of a clusterResourceQuota and creates it.  Returns the server's representation of the clusterResourceQuota, and an error, if there is any.
func (c *FakeClusterResourceQuotas) Create(clusterResourceQuota *quota_v1.ClusterResourceQuota) (result *quota_v1.ClusterResourceQuota, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(clusterresourcequotasResource, clusterResourceQuota), &quota_v1.ClusterResourceQuota{})
	if obj == nil {
		return nil, err
	}
	return obj.(*quota_v1.ClusterResourceQuota), err
}

// Update takes the representation of a clusterResourceQuota and updates it. Returns the server's representation of the clusterResourceQuota, and an error, if there is any.
func (c *FakeClusterResourceQuotas) Update(clusterResourceQuota *quota_v1.ClusterResourceQuota) (result *quota_v1.ClusterResourceQuota, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(clusterresourcequotasResource, clusterResourceQuota), &quota_v1.ClusterResourceQuota{})
	if obj == nil {
		return nil, err
	}
	return obj.(*quota_v1.ClusterResourceQuota), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeClusterResourceQuotas) UpdateStatus(clusterResourceQuota *quota_v1.ClusterResourceQuota) (*quota_v1.ClusterResourceQuota, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(clusterresourcequotasResource, "status", clusterResourceQuota), &quota_v1.ClusterResourceQuota{})
	if obj == nil {
		return nil, err
	}
	return obj.(*quota_v1.ClusterResourceQuota), err
}

// Delete takes name of the clusterResourceQuota and deletes it. Returns an error if one occurs.
func (c *FakeClusterResourceQuotas) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(clusterresourcequotasResource, name), &quota_v1.ClusterResourceQuota{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeClusterResourceQuotas) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(clusterresourcequotasResource, listOptions)

	_, err := c.Fake.Invokes(action, &quota_v1.ClusterResourceQuotaList{})
	return err
}

// Patch applies the patch and returns the patched clusterResourceQuota.
func (c *FakeClusterResourceQuotas) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *quota_v1.ClusterResourceQuota, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(clusterresourcequotasResource, name, data, subresources...), &quota_v1.ClusterResourceQuota{})
	if obj == nil {
		return nil, err
	}
	return obj.(*quota_v1.ClusterResourceQuota), err
}