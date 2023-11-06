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
// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1alpha1

import (
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// HostedClusterStatusApplyConfiguration represents an declarative configuration of the HostedClusterStatus type for use
// with apply.
type HostedClusterStatusApplyConfiguration struct {
	Version                  *ClusterVersionStatusApplyConfiguration `json:"version,omitempty"`
	KubeConfig               *v1.LocalObjectReference                `json:"kubeconfig,omitempty"`
	KubeadminPassword        *v1.LocalObjectReference                `json:"kubeadminPassword,omitempty"`
	IgnitionEndpoint         *string                                 `json:"ignitionEndpoint,omitempty"`
	ControlPlaneEndpoint     *APIEndpointApplyConfiguration          `json:"controlPlaneEndpoint,omitempty"`
	OAuthCallbackURLTemplate *string                                 `json:"oauthCallbackURLTemplate,omitempty"`
	Conditions               []metav1.Condition                      `json:"conditions,omitempty"`
	Platform                 *PlatformStatusApplyConfiguration       `json:"platform,omitempty"`
}

// HostedClusterStatusApplyConfiguration constructs an declarative configuration of the HostedClusterStatus type for use with
// apply.
func HostedClusterStatus() *HostedClusterStatusApplyConfiguration {
	return &HostedClusterStatusApplyConfiguration{}
}

// WithVersion sets the Version field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Version field is set to the value of the last call.
func (b *HostedClusterStatusApplyConfiguration) WithVersion(value *ClusterVersionStatusApplyConfiguration) *HostedClusterStatusApplyConfiguration {
	b.Version = value
	return b
}

// WithKubeConfig sets the KubeConfig field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the KubeConfig field is set to the value of the last call.
func (b *HostedClusterStatusApplyConfiguration) WithKubeConfig(value v1.LocalObjectReference) *HostedClusterStatusApplyConfiguration {
	b.KubeConfig = &value
	return b
}

// WithKubeadminPassword sets the KubeadminPassword field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the KubeadminPassword field is set to the value of the last call.
func (b *HostedClusterStatusApplyConfiguration) WithKubeadminPassword(value v1.LocalObjectReference) *HostedClusterStatusApplyConfiguration {
	b.KubeadminPassword = &value
	return b
}

// WithIgnitionEndpoint sets the IgnitionEndpoint field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the IgnitionEndpoint field is set to the value of the last call.
func (b *HostedClusterStatusApplyConfiguration) WithIgnitionEndpoint(value string) *HostedClusterStatusApplyConfiguration {
	b.IgnitionEndpoint = &value
	return b
}

// WithControlPlaneEndpoint sets the ControlPlaneEndpoint field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ControlPlaneEndpoint field is set to the value of the last call.
func (b *HostedClusterStatusApplyConfiguration) WithControlPlaneEndpoint(value *APIEndpointApplyConfiguration) *HostedClusterStatusApplyConfiguration {
	b.ControlPlaneEndpoint = value
	return b
}

// WithOAuthCallbackURLTemplate sets the OAuthCallbackURLTemplate field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the OAuthCallbackURLTemplate field is set to the value of the last call.
func (b *HostedClusterStatusApplyConfiguration) WithOAuthCallbackURLTemplate(value string) *HostedClusterStatusApplyConfiguration {
	b.OAuthCallbackURLTemplate = &value
	return b
}

// WithConditions adds the given value to the Conditions field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the Conditions field.
func (b *HostedClusterStatusApplyConfiguration) WithConditions(values ...metav1.Condition) *HostedClusterStatusApplyConfiguration {
	for i := range values {
		b.Conditions = append(b.Conditions, values[i])
	}
	return b
}

// WithPlatform sets the Platform field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Platform field is set to the value of the last call.
func (b *HostedClusterStatusApplyConfiguration) WithPlatform(value *PlatformStatusApplyConfiguration) *HostedClusterStatusApplyConfiguration {
	b.Platform = value
	return b
}
