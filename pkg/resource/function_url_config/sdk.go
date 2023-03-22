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
	"context"
	"errors"
	"fmt"
	"reflect"
	"strings"

	ackv1alpha1 "github.com/aws-controllers-k8s/runtime/apis/core/v1alpha1"
	ackcompare "github.com/aws-controllers-k8s/runtime/pkg/compare"
	ackcondition "github.com/aws-controllers-k8s/runtime/pkg/condition"
	ackerr "github.com/aws-controllers-k8s/runtime/pkg/errors"
	ackrequeue "github.com/aws-controllers-k8s/runtime/pkg/requeue"
	ackrtlog "github.com/aws-controllers-k8s/runtime/pkg/runtime/log"
	"github.com/aws/aws-sdk-go/aws"
	svcsdk "github.com/aws/aws-sdk-go/service/lambda"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	svcapitypes "github.com/aws-controllers-k8s/lambda-controller/apis/v1alpha1"
)

// Hack to avoid import errors during build...
var (
	_ = &metav1.Time{}
	_ = strings.ToLower("")
	_ = &aws.JSONValue{}
	_ = &svcsdk.Lambda{}
	_ = &svcapitypes.FunctionURLConfig{}
	_ = ackv1alpha1.AWSAccountID("")
	_ = &ackerr.NotFound
	_ = &ackcondition.NotManagedMessage
	_ = &reflect.Value{}
	_ = fmt.Sprintf("")
	_ = &ackrequeue.NoRequeue{}
)

// sdkFind returns SDK-specific information about a supplied resource
func (rm *resourceManager) sdkFind(
	ctx context.Context,
	r *resource,
) (latest *resource, err error) {
	rlog := ackrtlog.FromContext(ctx)
	exit := rlog.Trace("rm.sdkFind")
	defer func() {
		exit(err)
	}()
	// If any required fields in the input shape are missing, AWS resource is
	// not created yet. Return NotFound here to indicate to callers that the
	// resource isn't yet created.
	if rm.requiredFieldsMissingFromReadOneInput(r) {
		return nil, ackerr.NotFound
	}

	input, err := rm.newDescribeRequestPayload(r)
	if err != nil {
		return nil, err
	}

	var resp *svcsdk.GetFunctionUrlConfigOutput
	resp, err = rm.sdkapi.GetFunctionUrlConfigWithContext(ctx, input)
	rm.metrics.RecordAPICall("READ_ONE", "GetFunctionUrlConfig", err)
	if err != nil {
		if reqErr, ok := ackerr.AWSRequestFailure(err); ok && reqErr.StatusCode() == 404 {
			return nil, ackerr.NotFound
		}
		if awsErr, ok := ackerr.AWSError(err); ok && awsErr.Code() == "ResourceNotFoundException" {
			return nil, ackerr.NotFound
		}
		return nil, err
	}

	// Merge in the information we read from the API call above to the copy of
	// the original Kubernetes object we passed to the function
	ko := r.ko.DeepCopy()

	if resp.AuthType != nil {
		ko.Spec.AuthType = resp.AuthType
	} else {
		ko.Spec.AuthType = nil
	}
	if resp.Cors != nil {
		f1 := &svcapitypes.CORS{}
		if resp.Cors.AllowCredentials != nil {
			f1.AllowCredentials = resp.Cors.AllowCredentials
		}
		if resp.Cors.AllowHeaders != nil {
			f1f1 := []*string{}
			for _, f1f1iter := range resp.Cors.AllowHeaders {
				var f1f1elem string
				f1f1elem = *f1f1iter
				f1f1 = append(f1f1, &f1f1elem)
			}
			f1.AllowHeaders = f1f1
		}
		if resp.Cors.AllowMethods != nil {
			f1f2 := []*string{}
			for _, f1f2iter := range resp.Cors.AllowMethods {
				var f1f2elem string
				f1f2elem = *f1f2iter
				f1f2 = append(f1f2, &f1f2elem)
			}
			f1.AllowMethods = f1f2
		}
		if resp.Cors.AllowOrigins != nil {
			f1f3 := []*string{}
			for _, f1f3iter := range resp.Cors.AllowOrigins {
				var f1f3elem string
				f1f3elem = *f1f3iter
				f1f3 = append(f1f3, &f1f3elem)
			}
			f1.AllowOrigins = f1f3
		}
		if resp.Cors.ExposeHeaders != nil {
			f1f4 := []*string{}
			for _, f1f4iter := range resp.Cors.ExposeHeaders {
				var f1f4elem string
				f1f4elem = *f1f4iter
				f1f4 = append(f1f4, &f1f4elem)
			}
			f1.ExposeHeaders = f1f4
		}
		if resp.Cors.MaxAge != nil {
			f1.MaxAge = resp.Cors.MaxAge
		}
		ko.Spec.CORS = f1
	} else {
		ko.Spec.CORS = nil
	}
	if resp.CreationTime != nil {
		ko.Status.CreationTime = resp.CreationTime
	} else {
		ko.Status.CreationTime = nil
	}
	if resp.FunctionArn != nil {
		ko.Status.FunctionARN = resp.FunctionArn
	} else {
		ko.Status.FunctionARN = nil
	}
	if resp.FunctionUrl != nil {
		ko.Status.FunctionURL = resp.FunctionUrl
	} else {
		ko.Status.FunctionURL = nil
	}

	rm.setStatusDefaults(ko)
	return &resource{ko}, nil
}

// requiredFieldsMissingFromReadOneInput returns true if there are any fields
// for the ReadOne Input shape that are required but not present in the
// resource's Spec or Status
func (rm *resourceManager) requiredFieldsMissingFromReadOneInput(
	r *resource,
) bool {
	return r.ko.Spec.FunctionName == nil

}

// newDescribeRequestPayload returns SDK-specific struct for the HTTP request
// payload of the Describe API call for the resource
func (rm *resourceManager) newDescribeRequestPayload(
	r *resource,
) (*svcsdk.GetFunctionUrlConfigInput, error) {
	res := &svcsdk.GetFunctionUrlConfigInput{}

	if r.ko.Spec.FunctionName != nil {
		res.SetFunctionName(*r.ko.Spec.FunctionName)
	}
	if r.ko.Spec.Qualifier != nil {
		res.SetQualifier(*r.ko.Spec.Qualifier)
	}

	return res, nil
}

// sdkCreate creates the supplied resource in the backend AWS service API and
// returns a copy of the resource with resource fields (in both Spec and
// Status) filled in with values from the CREATE API operation's Output shape.
func (rm *resourceManager) sdkCreate(
	ctx context.Context,
	desired *resource,
) (created *resource, err error) {
	rlog := ackrtlog.FromContext(ctx)
	exit := rlog.Trace("rm.sdkCreate")
	defer func() {
		exit(err)
	}()
	input, err := rm.newCreateRequestPayload(ctx, desired)
	if err != nil {
		return nil, err
	}

	var resp *svcsdk.CreateFunctionUrlConfigOutput
	_ = resp
	resp, err = rm.sdkapi.CreateFunctionUrlConfigWithContext(ctx, input)
	rm.metrics.RecordAPICall("CREATE", "CreateFunctionUrlConfig", err)
	if err != nil {
		return nil, err
	}
	// Merge in the information we read from the API call above to the copy of
	// the original Kubernetes object we passed to the function
	ko := desired.ko.DeepCopy()

	if resp.AuthType != nil {
		ko.Spec.AuthType = resp.AuthType
	} else {
		ko.Spec.AuthType = nil
	}
	if resp.Cors != nil {
		f1 := &svcapitypes.CORS{}
		if resp.Cors.AllowCredentials != nil {
			f1.AllowCredentials = resp.Cors.AllowCredentials
		}
		if resp.Cors.AllowHeaders != nil {
			f1f1 := []*string{}
			for _, f1f1iter := range resp.Cors.AllowHeaders {
				var f1f1elem string
				f1f1elem = *f1f1iter
				f1f1 = append(f1f1, &f1f1elem)
			}
			f1.AllowHeaders = f1f1
		}
		if resp.Cors.AllowMethods != nil {
			f1f2 := []*string{}
			for _, f1f2iter := range resp.Cors.AllowMethods {
				var f1f2elem string
				f1f2elem = *f1f2iter
				f1f2 = append(f1f2, &f1f2elem)
			}
			f1.AllowMethods = f1f2
		}
		if resp.Cors.AllowOrigins != nil {
			f1f3 := []*string{}
			for _, f1f3iter := range resp.Cors.AllowOrigins {
				var f1f3elem string
				f1f3elem = *f1f3iter
				f1f3 = append(f1f3, &f1f3elem)
			}
			f1.AllowOrigins = f1f3
		}
		if resp.Cors.ExposeHeaders != nil {
			f1f4 := []*string{}
			for _, f1f4iter := range resp.Cors.ExposeHeaders {
				var f1f4elem string
				f1f4elem = *f1f4iter
				f1f4 = append(f1f4, &f1f4elem)
			}
			f1.ExposeHeaders = f1f4
		}
		if resp.Cors.MaxAge != nil {
			f1.MaxAge = resp.Cors.MaxAge
		}
		ko.Spec.CORS = f1
	} else {
		ko.Spec.CORS = nil
	}
	if resp.CreationTime != nil {
		ko.Status.CreationTime = resp.CreationTime
	} else {
		ko.Status.CreationTime = nil
	}
	if resp.FunctionArn != nil {
		ko.Status.FunctionARN = resp.FunctionArn
	} else {
		ko.Status.FunctionARN = nil
	}
	if resp.FunctionUrl != nil {
		ko.Status.FunctionURL = resp.FunctionUrl
	} else {
		ko.Status.FunctionURL = nil
	}

	rm.setStatusDefaults(ko)
	return &resource{ko}, nil
}

// newCreateRequestPayload returns an SDK-specific struct for the HTTP request
// payload of the Create API call for the resource
func (rm *resourceManager) newCreateRequestPayload(
	ctx context.Context,
	r *resource,
) (*svcsdk.CreateFunctionUrlConfigInput, error) {
	res := &svcsdk.CreateFunctionUrlConfigInput{}

	if r.ko.Spec.AuthType != nil {
		res.SetAuthType(*r.ko.Spec.AuthType)
	}
	if r.ko.Spec.CORS != nil {
		f1 := &svcsdk.Cors{}
		if r.ko.Spec.CORS.AllowCredentials != nil {
			f1.SetAllowCredentials(*r.ko.Spec.CORS.AllowCredentials)
		}
		if r.ko.Spec.CORS.AllowHeaders != nil {
			f1f1 := []*string{}
			for _, f1f1iter := range r.ko.Spec.CORS.AllowHeaders {
				var f1f1elem string
				f1f1elem = *f1f1iter
				f1f1 = append(f1f1, &f1f1elem)
			}
			f1.SetAllowHeaders(f1f1)
		}
		if r.ko.Spec.CORS.AllowMethods != nil {
			f1f2 := []*string{}
			for _, f1f2iter := range r.ko.Spec.CORS.AllowMethods {
				var f1f2elem string
				f1f2elem = *f1f2iter
				f1f2 = append(f1f2, &f1f2elem)
			}
			f1.SetAllowMethods(f1f2)
		}
		if r.ko.Spec.CORS.AllowOrigins != nil {
			f1f3 := []*string{}
			for _, f1f3iter := range r.ko.Spec.CORS.AllowOrigins {
				var f1f3elem string
				f1f3elem = *f1f3iter
				f1f3 = append(f1f3, &f1f3elem)
			}
			f1.SetAllowOrigins(f1f3)
		}
		if r.ko.Spec.CORS.ExposeHeaders != nil {
			f1f4 := []*string{}
			for _, f1f4iter := range r.ko.Spec.CORS.ExposeHeaders {
				var f1f4elem string
				f1f4elem = *f1f4iter
				f1f4 = append(f1f4, &f1f4elem)
			}
			f1.SetExposeHeaders(f1f4)
		}
		if r.ko.Spec.CORS.MaxAge != nil {
			f1.SetMaxAge(*r.ko.Spec.CORS.MaxAge)
		}
		res.SetCors(f1)
	}
	if r.ko.Spec.FunctionName != nil {
		res.SetFunctionName(*r.ko.Spec.FunctionName)
	}
	if r.ko.Spec.Qualifier != nil {
		res.SetQualifier(*r.ko.Spec.Qualifier)
	}

	return res, nil
}

// sdkUpdate patches the supplied resource in the backend AWS service API and
// returns a new resource with updated fields.
func (rm *resourceManager) sdkUpdate(
	ctx context.Context,
	desired *resource,
	latest *resource,
	delta *ackcompare.Delta,
) (updated *resource, err error) {
	rlog := ackrtlog.FromContext(ctx)
	exit := rlog.Trace("rm.sdkUpdate")
	defer func() {
		exit(err)
	}()
	input, err := rm.newUpdateRequestPayload(ctx, desired, delta)
	if err != nil {
		return nil, err
	}

	var resp *svcsdk.UpdateFunctionUrlConfigOutput
	_ = resp
	resp, err = rm.sdkapi.UpdateFunctionUrlConfigWithContext(ctx, input)
	rm.metrics.RecordAPICall("UPDATE", "UpdateFunctionUrlConfig", err)
	if err != nil {
		return nil, err
	}
	// Merge in the information we read from the API call above to the copy of
	// the original Kubernetes object we passed to the function
	ko := desired.ko.DeepCopy()

	if resp.AuthType != nil {
		ko.Spec.AuthType = resp.AuthType
	} else {
		ko.Spec.AuthType = nil
	}
	if resp.Cors != nil {
		f1 := &svcapitypes.CORS{}
		if resp.Cors.AllowCredentials != nil {
			f1.AllowCredentials = resp.Cors.AllowCredentials
		}
		if resp.Cors.AllowHeaders != nil {
			f1f1 := []*string{}
			for _, f1f1iter := range resp.Cors.AllowHeaders {
				var f1f1elem string
				f1f1elem = *f1f1iter
				f1f1 = append(f1f1, &f1f1elem)
			}
			f1.AllowHeaders = f1f1
		}
		if resp.Cors.AllowMethods != nil {
			f1f2 := []*string{}
			for _, f1f2iter := range resp.Cors.AllowMethods {
				var f1f2elem string
				f1f2elem = *f1f2iter
				f1f2 = append(f1f2, &f1f2elem)
			}
			f1.AllowMethods = f1f2
		}
		if resp.Cors.AllowOrigins != nil {
			f1f3 := []*string{}
			for _, f1f3iter := range resp.Cors.AllowOrigins {
				var f1f3elem string
				f1f3elem = *f1f3iter
				f1f3 = append(f1f3, &f1f3elem)
			}
			f1.AllowOrigins = f1f3
		}
		if resp.Cors.ExposeHeaders != nil {
			f1f4 := []*string{}
			for _, f1f4iter := range resp.Cors.ExposeHeaders {
				var f1f4elem string
				f1f4elem = *f1f4iter
				f1f4 = append(f1f4, &f1f4elem)
			}
			f1.ExposeHeaders = f1f4
		}
		if resp.Cors.MaxAge != nil {
			f1.MaxAge = resp.Cors.MaxAge
		}
		ko.Spec.CORS = f1
	} else {
		ko.Spec.CORS = nil
	}
	if resp.CreationTime != nil {
		ko.Status.CreationTime = resp.CreationTime
	} else {
		ko.Status.CreationTime = nil
	}
	if resp.FunctionArn != nil {
		ko.Status.FunctionARN = resp.FunctionArn
	} else {
		ko.Status.FunctionARN = nil
	}
	if resp.FunctionUrl != nil {
		ko.Status.FunctionURL = resp.FunctionUrl
	} else {
		ko.Status.FunctionURL = nil
	}

	rm.setStatusDefaults(ko)
	return &resource{ko}, nil
}

// newUpdateRequestPayload returns an SDK-specific struct for the HTTP request
// payload of the Update API call for the resource
func (rm *resourceManager) newUpdateRequestPayload(
	ctx context.Context,
	r *resource,
	delta *ackcompare.Delta,
) (*svcsdk.UpdateFunctionUrlConfigInput, error) {
	res := &svcsdk.UpdateFunctionUrlConfigInput{}

	if r.ko.Spec.AuthType != nil {
		res.SetAuthType(*r.ko.Spec.AuthType)
	}
	if r.ko.Spec.CORS != nil {
		f1 := &svcsdk.Cors{}
		if r.ko.Spec.CORS.AllowCredentials != nil {
			f1.SetAllowCredentials(*r.ko.Spec.CORS.AllowCredentials)
		}
		if r.ko.Spec.CORS.AllowHeaders != nil {
			f1f1 := []*string{}
			for _, f1f1iter := range r.ko.Spec.CORS.AllowHeaders {
				var f1f1elem string
				f1f1elem = *f1f1iter
				f1f1 = append(f1f1, &f1f1elem)
			}
			f1.SetAllowHeaders(f1f1)
		}
		if r.ko.Spec.CORS.AllowMethods != nil {
			f1f2 := []*string{}
			for _, f1f2iter := range r.ko.Spec.CORS.AllowMethods {
				var f1f2elem string
				f1f2elem = *f1f2iter
				f1f2 = append(f1f2, &f1f2elem)
			}
			f1.SetAllowMethods(f1f2)
		}
		if r.ko.Spec.CORS.AllowOrigins != nil {
			f1f3 := []*string{}
			for _, f1f3iter := range r.ko.Spec.CORS.AllowOrigins {
				var f1f3elem string
				f1f3elem = *f1f3iter
				f1f3 = append(f1f3, &f1f3elem)
			}
			f1.SetAllowOrigins(f1f3)
		}
		if r.ko.Spec.CORS.ExposeHeaders != nil {
			f1f4 := []*string{}
			for _, f1f4iter := range r.ko.Spec.CORS.ExposeHeaders {
				var f1f4elem string
				f1f4elem = *f1f4iter
				f1f4 = append(f1f4, &f1f4elem)
			}
			f1.SetExposeHeaders(f1f4)
		}
		if r.ko.Spec.CORS.MaxAge != nil {
			f1.SetMaxAge(*r.ko.Spec.CORS.MaxAge)
		}
		res.SetCors(f1)
	}
	if r.ko.Spec.FunctionName != nil {
		res.SetFunctionName(*r.ko.Spec.FunctionName)
	}
	if r.ko.Spec.Qualifier != nil {
		res.SetQualifier(*r.ko.Spec.Qualifier)
	}

	return res, nil
}

// sdkDelete deletes the supplied resource in the backend AWS service API
func (rm *resourceManager) sdkDelete(
	ctx context.Context,
	r *resource,
) (latest *resource, err error) {
	rlog := ackrtlog.FromContext(ctx)
	exit := rlog.Trace("rm.sdkDelete")
	defer func() {
		exit(err)
	}()
	input, err := rm.newDeleteRequestPayload(r)
	if err != nil {
		return nil, err
	}
	var resp *svcsdk.DeleteFunctionUrlConfigOutput
	_ = resp
	resp, err = rm.sdkapi.DeleteFunctionUrlConfigWithContext(ctx, input)
	rm.metrics.RecordAPICall("DELETE", "DeleteFunctionUrlConfig", err)
	return nil, err
}

// newDeleteRequestPayload returns an SDK-specific struct for the HTTP request
// payload of the Delete API call for the resource
func (rm *resourceManager) newDeleteRequestPayload(
	r *resource,
) (*svcsdk.DeleteFunctionUrlConfigInput, error) {
	res := &svcsdk.DeleteFunctionUrlConfigInput{}

	if r.ko.Spec.FunctionName != nil {
		res.SetFunctionName(*r.ko.Spec.FunctionName)
	}
	if r.ko.Spec.Qualifier != nil {
		res.SetQualifier(*r.ko.Spec.Qualifier)
	}

	return res, nil
}

// setStatusDefaults sets default properties into supplied custom resource
func (rm *resourceManager) setStatusDefaults(
	ko *svcapitypes.FunctionURLConfig,
) {
	if ko.Status.ACKResourceMetadata == nil {
		ko.Status.ACKResourceMetadata = &ackv1alpha1.ResourceMetadata{}
	}
	if ko.Status.ACKResourceMetadata.Region == nil {
		ko.Status.ACKResourceMetadata.Region = &rm.awsRegion
	}
	if ko.Status.ACKResourceMetadata.OwnerAccountID == nil {
		ko.Status.ACKResourceMetadata.OwnerAccountID = &rm.awsAccountID
	}
	if ko.Status.Conditions == nil {
		ko.Status.Conditions = []*ackv1alpha1.Condition{}
	}
}

// updateConditions returns updated resource, true; if conditions were updated
// else it returns nil, false
func (rm *resourceManager) updateConditions(
	r *resource,
	onSuccess bool,
	err error,
) (*resource, bool) {
	ko := r.ko.DeepCopy()
	rm.setStatusDefaults(ko)

	// Terminal condition
	var terminalCondition *ackv1alpha1.Condition = nil
	var recoverableCondition *ackv1alpha1.Condition = nil
	var syncCondition *ackv1alpha1.Condition = nil
	for _, condition := range ko.Status.Conditions {
		if condition.Type == ackv1alpha1.ConditionTypeTerminal {
			terminalCondition = condition
		}
		if condition.Type == ackv1alpha1.ConditionTypeRecoverable {
			recoverableCondition = condition
		}
		if condition.Type == ackv1alpha1.ConditionTypeResourceSynced {
			syncCondition = condition
		}
	}
	var termError *ackerr.TerminalError
	if rm.terminalAWSError(err) || err == ackerr.SecretTypeNotSupported || err == ackerr.SecretNotFound || errors.As(err, &termError) {
		if terminalCondition == nil {
			terminalCondition = &ackv1alpha1.Condition{
				Type: ackv1alpha1.ConditionTypeTerminal,
			}
			ko.Status.Conditions = append(ko.Status.Conditions, terminalCondition)
		}
		var errorMessage = ""
		if err == ackerr.SecretTypeNotSupported || err == ackerr.SecretNotFound || errors.As(err, &termError) {
			errorMessage = err.Error()
		} else {
			awsErr, _ := ackerr.AWSError(err)
			errorMessage = awsErr.Error()
		}
		terminalCondition.Status = corev1.ConditionTrue
		terminalCondition.Message = &errorMessage
	} else {
		// Clear the terminal condition if no longer present
		if terminalCondition != nil {
			terminalCondition.Status = corev1.ConditionFalse
			terminalCondition.Message = nil
		}
		// Handling Recoverable Conditions
		if err != nil {
			if recoverableCondition == nil {
				// Add a new Condition containing a non-terminal error
				recoverableCondition = &ackv1alpha1.Condition{
					Type: ackv1alpha1.ConditionTypeRecoverable,
				}
				ko.Status.Conditions = append(ko.Status.Conditions, recoverableCondition)
			}
			recoverableCondition.Status = corev1.ConditionTrue
			awsErr, _ := ackerr.AWSError(err)
			errorMessage := err.Error()
			if awsErr != nil {
				errorMessage = awsErr.Error()
			}
			recoverableCondition.Message = &errorMessage
		} else if recoverableCondition != nil {
			recoverableCondition.Status = corev1.ConditionFalse
			recoverableCondition.Message = nil
		}
	}
	// Required to avoid the "declared but not used" error in the default case
	_ = syncCondition
	if terminalCondition != nil || recoverableCondition != nil || syncCondition != nil {
		return &resource{ko}, true // updated
	}
	return nil, false // not updated
}

// terminalAWSError returns awserr, true; if the supplied error is an aws Error type
// and if the exception indicates that it is a Terminal exception
// 'Terminal' exception are specified in generator configuration
func (rm *resourceManager) terminalAWSError(err error) bool {
	// No terminal_errors specified for this resource in generator config
	return false
}
