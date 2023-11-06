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

// KubevirtPlatformSpecApplyConfiguration represents an declarative configuration of the KubevirtPlatformSpec type for use
// with apply.
type KubevirtPlatformSpecApplyConfiguration struct {
	BaseDomainPassthrough *bool                                          `json:"baseDomainPassthrough,omitempty"`
	GenerateID            *string                                        `json:"generateID,omitempty"`
	Credentials           *KubevirtPlatformCredentialsApplyConfiguration `json:"credentials,omitempty"`
	StorageDriver         *KubevirtStorageDriverSpecApplyConfiguration   `json:"storageDriver,omitempty"`
}

// KubevirtPlatformSpecApplyConfiguration constructs an declarative configuration of the KubevirtPlatformSpec type for use with
// apply.
func KubevirtPlatformSpec() *KubevirtPlatformSpecApplyConfiguration {
	return &KubevirtPlatformSpecApplyConfiguration{}
}

// WithBaseDomainPassthrough sets the BaseDomainPassthrough field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the BaseDomainPassthrough field is set to the value of the last call.
func (b *KubevirtPlatformSpecApplyConfiguration) WithBaseDomainPassthrough(value bool) *KubevirtPlatformSpecApplyConfiguration {
	b.BaseDomainPassthrough = &value
	return b
}

// WithGenerateID sets the GenerateID field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the GenerateID field is set to the value of the last call.
func (b *KubevirtPlatformSpecApplyConfiguration) WithGenerateID(value string) *KubevirtPlatformSpecApplyConfiguration {
	b.GenerateID = &value
	return b
}

// WithCredentials sets the Credentials field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Credentials field is set to the value of the last call.
func (b *KubevirtPlatformSpecApplyConfiguration) WithCredentials(value *KubevirtPlatformCredentialsApplyConfiguration) *KubevirtPlatformSpecApplyConfiguration {
	b.Credentials = value
	return b
}

// WithStorageDriver sets the StorageDriver field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the StorageDriver field is set to the value of the last call.
func (b *KubevirtPlatformSpecApplyConfiguration) WithStorageDriver(value *KubevirtStorageDriverSpecApplyConfiguration) *KubevirtPlatformSpecApplyConfiguration {
	b.StorageDriver = value
	return b
}
