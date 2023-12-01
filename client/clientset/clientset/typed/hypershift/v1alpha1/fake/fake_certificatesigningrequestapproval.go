/*


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

	v1alpha1 "github.com/openshift/hypershift/api/hypershift/v1alpha1"
	hypershiftv1alpha1 "github.com/openshift/hypershift/client/applyconfiguration/hypershift/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeCertificateSigningRequestApprovals implements CertificateSigningRequestApprovalInterface
type FakeCertificateSigningRequestApprovals struct {
	Fake *FakeHypershiftV1alpha1
	ns   string
}

var certificatesigningrequestapprovalsResource = v1alpha1.SchemeGroupVersion.WithResource("certificatesigningrequestapprovals")

var certificatesigningrequestapprovalsKind = v1alpha1.SchemeGroupVersion.WithKind("CertificateSigningRequestApproval")

// Get takes name of the certificateSigningRequestApproval, and returns the corresponding certificateSigningRequestApproval object, and an error if there is any.
func (c *FakeCertificateSigningRequestApprovals) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.CertificateSigningRequestApproval, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(certificatesigningrequestapprovalsResource, c.ns, name), &v1alpha1.CertificateSigningRequestApproval{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.CertificateSigningRequestApproval), err
}

// List takes label and field selectors, and returns the list of CertificateSigningRequestApprovals that match those selectors.
func (c *FakeCertificateSigningRequestApprovals) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.CertificateSigningRequestApprovalList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(certificatesigningrequestapprovalsResource, certificatesigningrequestapprovalsKind, c.ns, opts), &v1alpha1.CertificateSigningRequestApprovalList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.CertificateSigningRequestApprovalList{ListMeta: obj.(*v1alpha1.CertificateSigningRequestApprovalList).ListMeta}
	for _, item := range obj.(*v1alpha1.CertificateSigningRequestApprovalList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested certificateSigningRequestApprovals.
func (c *FakeCertificateSigningRequestApprovals) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(certificatesigningrequestapprovalsResource, c.ns, opts))

}

// Create takes the representation of a certificateSigningRequestApproval and creates it.  Returns the server's representation of the certificateSigningRequestApproval, and an error, if there is any.
func (c *FakeCertificateSigningRequestApprovals) Create(ctx context.Context, certificateSigningRequestApproval *v1alpha1.CertificateSigningRequestApproval, opts v1.CreateOptions) (result *v1alpha1.CertificateSigningRequestApproval, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(certificatesigningrequestapprovalsResource, c.ns, certificateSigningRequestApproval), &v1alpha1.CertificateSigningRequestApproval{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.CertificateSigningRequestApproval), err
}

// Update takes the representation of a certificateSigningRequestApproval and updates it. Returns the server's representation of the certificateSigningRequestApproval, and an error, if there is any.
func (c *FakeCertificateSigningRequestApprovals) Update(ctx context.Context, certificateSigningRequestApproval *v1alpha1.CertificateSigningRequestApproval, opts v1.UpdateOptions) (result *v1alpha1.CertificateSigningRequestApproval, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(certificatesigningrequestapprovalsResource, c.ns, certificateSigningRequestApproval), &v1alpha1.CertificateSigningRequestApproval{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.CertificateSigningRequestApproval), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeCertificateSigningRequestApprovals) UpdateStatus(ctx context.Context, certificateSigningRequestApproval *v1alpha1.CertificateSigningRequestApproval, opts v1.UpdateOptions) (*v1alpha1.CertificateSigningRequestApproval, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(certificatesigningrequestapprovalsResource, "status", c.ns, certificateSigningRequestApproval), &v1alpha1.CertificateSigningRequestApproval{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.CertificateSigningRequestApproval), err
}

// Delete takes name of the certificateSigningRequestApproval and deletes it. Returns an error if one occurs.
func (c *FakeCertificateSigningRequestApprovals) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(certificatesigningrequestapprovalsResource, c.ns, name, opts), &v1alpha1.CertificateSigningRequestApproval{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeCertificateSigningRequestApprovals) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(certificatesigningrequestapprovalsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.CertificateSigningRequestApprovalList{})
	return err
}

// Patch applies the patch and returns the patched certificateSigningRequestApproval.
func (c *FakeCertificateSigningRequestApprovals) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.CertificateSigningRequestApproval, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(certificatesigningrequestapprovalsResource, c.ns, name, pt, data, subresources...), &v1alpha1.CertificateSigningRequestApproval{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.CertificateSigningRequestApproval), err
}

// Apply takes the given apply declarative configuration, applies it and returns the applied certificateSigningRequestApproval.
func (c *FakeCertificateSigningRequestApprovals) Apply(ctx context.Context, certificateSigningRequestApproval *hypershiftv1alpha1.CertificateSigningRequestApprovalApplyConfiguration, opts v1.ApplyOptions) (result *v1alpha1.CertificateSigningRequestApproval, err error) {
	if certificateSigningRequestApproval == nil {
		return nil, fmt.Errorf("certificateSigningRequestApproval provided to Apply must not be nil")
	}
	data, err := json.Marshal(certificateSigningRequestApproval)
	if err != nil {
		return nil, err
	}
	name := certificateSigningRequestApproval.Name
	if name == nil {
		return nil, fmt.Errorf("certificateSigningRequestApproval.Name must be provided to Apply")
	}
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(certificatesigningrequestapprovalsResource, c.ns, *name, types.ApplyPatchType, data), &v1alpha1.CertificateSigningRequestApproval{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.CertificateSigningRequestApproval), err
}

// ApplyStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating ApplyStatus().
func (c *FakeCertificateSigningRequestApprovals) ApplyStatus(ctx context.Context, certificateSigningRequestApproval *hypershiftv1alpha1.CertificateSigningRequestApprovalApplyConfiguration, opts v1.ApplyOptions) (result *v1alpha1.CertificateSigningRequestApproval, err error) {
	if certificateSigningRequestApproval == nil {
		return nil, fmt.Errorf("certificateSigningRequestApproval provided to Apply must not be nil")
	}
	data, err := json.Marshal(certificateSigningRequestApproval)
	if err != nil {
		return nil, err
	}
	name := certificateSigningRequestApproval.Name
	if name == nil {
		return nil, fmt.Errorf("certificateSigningRequestApproval.Name must be provided to Apply")
	}
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(certificatesigningrequestapprovalsResource, c.ns, *name, types.ApplyPatchType, data, "status"), &v1alpha1.CertificateSigningRequestApproval{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.CertificateSigningRequestApproval), err
}
