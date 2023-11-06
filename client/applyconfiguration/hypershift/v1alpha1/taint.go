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
)

// TaintApplyConfiguration represents an declarative configuration of the Taint type for use
// with apply.
type TaintApplyConfiguration struct {
	Key    *string         `json:"key,omitempty"`
	Value  *string         `json:"value,omitempty"`
	Effect *v1.TaintEffect `json:"effect,omitempty"`
}

// TaintApplyConfiguration constructs an declarative configuration of the Taint type for use with
// apply.
func Taint() *TaintApplyConfiguration {
	return &TaintApplyConfiguration{}
}

// WithKey sets the Key field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Key field is set to the value of the last call.
func (b *TaintApplyConfiguration) WithKey(value string) *TaintApplyConfiguration {
	b.Key = &value
	return b
}

// WithValue sets the Value field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Value field is set to the value of the last call.
func (b *TaintApplyConfiguration) WithValue(value string) *TaintApplyConfiguration {
	b.Value = &value
	return b
}

// WithEffect sets the Effect field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Effect field is set to the value of the last call.
func (b *TaintApplyConfiguration) WithEffect(value v1.TaintEffect) *TaintApplyConfiguration {
	b.Effect = &value
	return b
}
