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

package function_url_config

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

	if ackcompare.HasNilDifference(a.ko.Spec.AuthType, b.ko.Spec.AuthType) {
		delta.Add("Spec.AuthType", a.ko.Spec.AuthType, b.ko.Spec.AuthType)
	} else if a.ko.Spec.AuthType != nil && b.ko.Spec.AuthType != nil {
		if *a.ko.Spec.AuthType != *b.ko.Spec.AuthType {
			delta.Add("Spec.AuthType", a.ko.Spec.AuthType, b.ko.Spec.AuthType)
		}
	}
	if ackcompare.HasNilDifference(a.ko.Spec.CORS, b.ko.Spec.CORS) {
		delta.Add("Spec.CORS", a.ko.Spec.CORS, b.ko.Spec.CORS)
	} else if a.ko.Spec.CORS != nil && b.ko.Spec.CORS != nil {
		if ackcompare.HasNilDifference(a.ko.Spec.CORS.AllowCredentials, b.ko.Spec.CORS.AllowCredentials) {
			delta.Add("Spec.CORS.AllowCredentials", a.ko.Spec.CORS.AllowCredentials, b.ko.Spec.CORS.AllowCredentials)
		} else if a.ko.Spec.CORS.AllowCredentials != nil && b.ko.Spec.CORS.AllowCredentials != nil {
			if *a.ko.Spec.CORS.AllowCredentials != *b.ko.Spec.CORS.AllowCredentials {
				delta.Add("Spec.CORS.AllowCredentials", a.ko.Spec.CORS.AllowCredentials, b.ko.Spec.CORS.AllowCredentials)
			}
		}
		if !ackcompare.SliceStringPEqual(a.ko.Spec.CORS.AllowHeaders, b.ko.Spec.CORS.AllowHeaders) {
			delta.Add("Spec.CORS.AllowHeaders", a.ko.Spec.CORS.AllowHeaders, b.ko.Spec.CORS.AllowHeaders)
		}
		if !ackcompare.SliceStringPEqual(a.ko.Spec.CORS.AllowMethods, b.ko.Spec.CORS.AllowMethods) {
			delta.Add("Spec.CORS.AllowMethods", a.ko.Spec.CORS.AllowMethods, b.ko.Spec.CORS.AllowMethods)
		}
		if !ackcompare.SliceStringPEqual(a.ko.Spec.CORS.AllowOrigins, b.ko.Spec.CORS.AllowOrigins) {
			delta.Add("Spec.CORS.AllowOrigins", a.ko.Spec.CORS.AllowOrigins, b.ko.Spec.CORS.AllowOrigins)
		}
		if !ackcompare.SliceStringPEqual(a.ko.Spec.CORS.ExposeHeaders, b.ko.Spec.CORS.ExposeHeaders) {
			delta.Add("Spec.CORS.ExposeHeaders", a.ko.Spec.CORS.ExposeHeaders, b.ko.Spec.CORS.ExposeHeaders)
		}
		if ackcompare.HasNilDifference(a.ko.Spec.CORS.MaxAge, b.ko.Spec.CORS.MaxAge) {
			delta.Add("Spec.CORS.MaxAge", a.ko.Spec.CORS.MaxAge, b.ko.Spec.CORS.MaxAge)
		} else if a.ko.Spec.CORS.MaxAge != nil && b.ko.Spec.CORS.MaxAge != nil {
			if *a.ko.Spec.CORS.MaxAge != *b.ko.Spec.CORS.MaxAge {
				delta.Add("Spec.CORS.MaxAge", a.ko.Spec.CORS.MaxAge, b.ko.Spec.CORS.MaxAge)
			}
		}
	}
	if ackcompare.HasNilDifference(a.ko.Spec.FunctionName, b.ko.Spec.FunctionName) {
		delta.Add("Spec.FunctionName", a.ko.Spec.FunctionName, b.ko.Spec.FunctionName)
	} else if a.ko.Spec.FunctionName != nil && b.ko.Spec.FunctionName != nil {
		if *a.ko.Spec.FunctionName != *b.ko.Spec.FunctionName {
			delta.Add("Spec.FunctionName", a.ko.Spec.FunctionName, b.ko.Spec.FunctionName)
		}
	}
	if !reflect.DeepEqual(a.ko.Spec.FunctionRef, b.ko.Spec.FunctionRef) {
		delta.Add("Spec.FunctionRef", a.ko.Spec.FunctionRef, b.ko.Spec.FunctionRef)
	}
	if ackcompare.HasNilDifference(a.ko.Spec.Qualifier, b.ko.Spec.Qualifier) {
		delta.Add("Spec.Qualifier", a.ko.Spec.Qualifier, b.ko.Spec.Qualifier)
	} else if a.ko.Spec.Qualifier != nil && b.ko.Spec.Qualifier != nil {
		if *a.ko.Spec.Qualifier != *b.ko.Spec.Qualifier {
			delta.Add("Spec.Qualifier", a.ko.Spec.Qualifier, b.ko.Spec.Qualifier)
		}
	}

	return delta
}
