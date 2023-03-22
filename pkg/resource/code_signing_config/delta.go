// Copyright Amazon.com Inc. or its affiliates. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"). You may
// not use this file except in compliance with the License. A copy of the
// License is located at
//
//     http://aws.amazon.com/apache2.0/
//
// or in the "license" file accompanying this file. This file is distributed
// on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
// express or implied. See the License for the specific language governing
// permissions and limitations under the License.

// Code generated by ack-generate. DO NOT EDIT.

package code_signing_config

import (
	"bytes"
	"reflect"

	ackcompare "github.com/aws-controllers-k8s/runtime/pkg/compare"
	acktags "github.com/aws-controllers-k8s/runtime/pkg/tags"
)

// Hack to avoid import errors during build...
var (
	_ = &bytes.Buffer{}
	_ = &reflect.Method{}
	_ = &acktags.Tags{}
)

// newResourceDelta returns a new `ackcompare.Delta` used to compare two
// resources
func newResourceDelta(
	a *resource,
	b *resource,
) *ackcompare.Delta {
	delta := ackcompare.NewDelta()
	if (a == nil && b != nil) ||
		(a != nil && b == nil) {
		delta.Add("", a, b)
		return delta
	}

	if ackcompare.HasNilDifference(a.ko.Spec.AllowedPublishers, b.ko.Spec.AllowedPublishers) {
		delta.Add("Spec.AllowedPublishers", a.ko.Spec.AllowedPublishers, b.ko.Spec.AllowedPublishers)
	} else if a.ko.Spec.AllowedPublishers != nil && b.ko.Spec.AllowedPublishers != nil {
		if !ackcompare.SliceStringPEqual(a.ko.Spec.AllowedPublishers.SigningProfileVersionARNs, b.ko.Spec.AllowedPublishers.SigningProfileVersionARNs) {
			delta.Add("Spec.AllowedPublishers.SigningProfileVersionARNs", a.ko.Spec.AllowedPublishers.SigningProfileVersionARNs, b.ko.Spec.AllowedPublishers.SigningProfileVersionARNs)
		}
	}
	if ackcompare.HasNilDifference(a.ko.Spec.CodeSigningPolicies, b.ko.Spec.CodeSigningPolicies) {
		delta.Add("Spec.CodeSigningPolicies", a.ko.Spec.CodeSigningPolicies, b.ko.Spec.CodeSigningPolicies)
	} else if a.ko.Spec.CodeSigningPolicies != nil && b.ko.Spec.CodeSigningPolicies != nil {
		if ackcompare.HasNilDifference(a.ko.Spec.CodeSigningPolicies.UntrustedArtifactOnDeployment, b.ko.Spec.CodeSigningPolicies.UntrustedArtifactOnDeployment) {
			delta.Add("Spec.CodeSigningPolicies.UntrustedArtifactOnDeployment", a.ko.Spec.CodeSigningPolicies.UntrustedArtifactOnDeployment, b.ko.Spec.CodeSigningPolicies.UntrustedArtifactOnDeployment)
		} else if a.ko.Spec.CodeSigningPolicies.UntrustedArtifactOnDeployment != nil && b.ko.Spec.CodeSigningPolicies.UntrustedArtifactOnDeployment != nil {
			if *a.ko.Spec.CodeSigningPolicies.UntrustedArtifactOnDeployment != *b.ko.Spec.CodeSigningPolicies.UntrustedArtifactOnDeployment {
				delta.Add("Spec.CodeSigningPolicies.UntrustedArtifactOnDeployment", a.ko.Spec.CodeSigningPolicies.UntrustedArtifactOnDeployment, b.ko.Spec.CodeSigningPolicies.UntrustedArtifactOnDeployment)
			}
		}
	}
	if ackcompare.HasNilDifference(a.ko.Spec.Description, b.ko.Spec.Description) {
		delta.Add("Spec.Description", a.ko.Spec.Description, b.ko.Spec.Description)
	} else if a.ko.Spec.Description != nil && b.ko.Spec.Description != nil {
		if *a.ko.Spec.Description != *b.ko.Spec.Description {
			delta.Add("Spec.Description", a.ko.Spec.Description, b.ko.Spec.Description)
		}
	}

	return delta
}
