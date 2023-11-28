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

package v1beta1

import (
	v1beta1 "github.com/openshift/hypershift/api/hypershift/v1beta1"
)

// PlatformSpecApplyConfiguration represents an declarative configuration of the PlatformSpec type for use
// with apply.
type PlatformSpecApplyConfiguration struct {
	Type      *v1beta1.PlatformType                    `json:"type,omitempty"`
	AWS       *AWSPlatformSpecApplyConfiguration       `json:"aws,omitempty"`
	Agent     *AgentPlatformSpecApplyConfiguration     `json:"agent,omitempty"`
	IBMCloud  *IBMCloudPlatformSpecApplyConfiguration  `json:"ibmcloud,omitempty"`
	Azure     *AzurePlatformSpecApplyConfiguration     `json:"azure,omitempty"`
	PowerVS   *PowerVSPlatformSpecApplyConfiguration   `json:"powervs,omitempty"`
	Kubevirt  *KubevirtPlatformSpecApplyConfiguration  `json:"kubevirt,omitempty"`
	OpenStack *OpenStackPlatformSpecApplyConfiguration `json:"openstack,omitempty"`
}

// PlatformSpecApplyConfiguration constructs an declarative configuration of the PlatformSpec type for use with
// apply.
func PlatformSpec() *PlatformSpecApplyConfiguration {
	return &PlatformSpecApplyConfiguration{}
}

// WithType sets the Type field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Type field is set to the value of the last call.
func (b *PlatformSpecApplyConfiguration) WithType(value v1beta1.PlatformType) *PlatformSpecApplyConfiguration {
	b.Type = &value
	return b
}

// WithAWS sets the AWS field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the AWS field is set to the value of the last call.
func (b *PlatformSpecApplyConfiguration) WithAWS(value *AWSPlatformSpecApplyConfiguration) *PlatformSpecApplyConfiguration {
	b.AWS = value
	return b
}

// WithAgent sets the Agent field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Agent field is set to the value of the last call.
func (b *PlatformSpecApplyConfiguration) WithAgent(value *AgentPlatformSpecApplyConfiguration) *PlatformSpecApplyConfiguration {
	b.Agent = value
	return b
}

// WithIBMCloud sets the IBMCloud field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the IBMCloud field is set to the value of the last call.
func (b *PlatformSpecApplyConfiguration) WithIBMCloud(value *IBMCloudPlatformSpecApplyConfiguration) *PlatformSpecApplyConfiguration {
	b.IBMCloud = value
	return b
}

// WithAzure sets the Azure field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Azure field is set to the value of the last call.
func (b *PlatformSpecApplyConfiguration) WithAzure(value *AzurePlatformSpecApplyConfiguration) *PlatformSpecApplyConfiguration {
	b.Azure = value
	return b
}

// WithPowerVS sets the PowerVS field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the PowerVS field is set to the value of the last call.
func (b *PlatformSpecApplyConfiguration) WithPowerVS(value *PowerVSPlatformSpecApplyConfiguration) *PlatformSpecApplyConfiguration {
	b.PowerVS = value
	return b
}

// WithKubevirt sets the Kubevirt field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Kubevirt field is set to the value of the last call.
func (b *PlatformSpecApplyConfiguration) WithKubevirt(value *KubevirtPlatformSpecApplyConfiguration) *PlatformSpecApplyConfiguration {
	b.Kubevirt = value
	return b
}

// WithOpenStack sets the OpenStack field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the OpenStack field is set to the value of the last call.
func (b *PlatformSpecApplyConfiguration) WithOpenStack(value *OpenStackPlatformSpecApplyConfiguration) *PlatformSpecApplyConfiguration {
	b.OpenStack = value
	return b
}
