/*
Copyright 2021 The Crossplane Authors.

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

// Code generated by ack-generate. DO NOT EDIT.

package integrationresponse

import (
	"github.com/aws/aws-sdk-go/aws/awserr"
	svcsdk "github.com/aws/aws-sdk-go/service/apigateway"

	svcapitypes "github.com/crossplane-contrib/provider-aws/apis/apigateway/v1alpha1"
)

// NOTE(muvaf): We return pointers in case the function needs to start with an
// empty object, hence need to return a new pointer.

// GenerateGetIntegrationResponseInput returns input for read
// operation.
func GenerateGetIntegrationResponseInput(cr *svcapitypes.IntegrationResponse) *svcsdk.GetIntegrationResponseInput {
	res := &svcsdk.GetIntegrationResponseInput{}

	if cr.Spec.ForProvider.HTTPMethod != nil {
		res.SetHttpMethod(*cr.Spec.ForProvider.HTTPMethod)
	}
	if cr.Spec.ForProvider.StatusCode != nil {
		res.SetStatusCode(*cr.Spec.ForProvider.StatusCode)
	}

	return res
}

// GenerateIntegrationResponse returns the current state in the form of *svcapitypes.IntegrationResponse.
func GenerateIntegrationResponse(resp *svcsdk.IntegrationResponse) *svcapitypes.IntegrationResponse {
	cr := &svcapitypes.IntegrationResponse{}

	if resp.ContentHandling != nil {
		cr.Spec.ForProvider.ContentHandling = resp.ContentHandling
	} else {
		cr.Spec.ForProvider.ContentHandling = nil
	}
	if resp.ResponseParameters != nil {
		f1 := map[string]*string{}
		for f1key, f1valiter := range resp.ResponseParameters {
			var f1val string
			f1val = *f1valiter
			f1[f1key] = &f1val
		}
		cr.Spec.ForProvider.ResponseParameters = f1
	} else {
		cr.Spec.ForProvider.ResponseParameters = nil
	}
	if resp.ResponseTemplates != nil {
		f2 := map[string]*string{}
		for f2key, f2valiter := range resp.ResponseTemplates {
			var f2val string
			f2val = *f2valiter
			f2[f2key] = &f2val
		}
		cr.Spec.ForProvider.ResponseTemplates = f2
	} else {
		cr.Spec.ForProvider.ResponseTemplates = nil
	}
	if resp.SelectionPattern != nil {
		cr.Spec.ForProvider.SelectionPattern = resp.SelectionPattern
	} else {
		cr.Spec.ForProvider.SelectionPattern = nil
	}
	if resp.StatusCode != nil {
		cr.Spec.ForProvider.StatusCode = resp.StatusCode
	} else {
		cr.Spec.ForProvider.StatusCode = nil
	}

	return cr
}

// GeneratePutIntegrationResponseInput returns a create input.
func GeneratePutIntegrationResponseInput(cr *svcapitypes.IntegrationResponse) *svcsdk.PutIntegrationResponseInput {
	res := &svcsdk.PutIntegrationResponseInput{}

	if cr.Spec.ForProvider.ContentHandling != nil {
		res.SetContentHandling(*cr.Spec.ForProvider.ContentHandling)
	}
	if cr.Spec.ForProvider.HTTPMethod != nil {
		res.SetHttpMethod(*cr.Spec.ForProvider.HTTPMethod)
	}
	if cr.Spec.ForProvider.ResponseParameters != nil {
		f2 := map[string]*string{}
		for f2key, f2valiter := range cr.Spec.ForProvider.ResponseParameters {
			var f2val string
			f2val = *f2valiter
			f2[f2key] = &f2val
		}
		res.SetResponseParameters(f2)
	}
	if cr.Spec.ForProvider.ResponseTemplates != nil {
		f3 := map[string]*string{}
		for f3key, f3valiter := range cr.Spec.ForProvider.ResponseTemplates {
			var f3val string
			f3val = *f3valiter
			f3[f3key] = &f3val
		}
		res.SetResponseTemplates(f3)
	}
	if cr.Spec.ForProvider.SelectionPattern != nil {
		res.SetSelectionPattern(*cr.Spec.ForProvider.SelectionPattern)
	}
	if cr.Spec.ForProvider.StatusCode != nil {
		res.SetStatusCode(*cr.Spec.ForProvider.StatusCode)
	}

	return res
}

// GenerateUpdateIntegrationResponseInput returns an update input.
func GenerateUpdateIntegrationResponseInput(cr *svcapitypes.IntegrationResponse) *svcsdk.UpdateIntegrationResponseInput {
	res := &svcsdk.UpdateIntegrationResponseInput{}

	if cr.Spec.ForProvider.HTTPMethod != nil {
		res.SetHttpMethod(*cr.Spec.ForProvider.HTTPMethod)
	}
	if cr.Spec.ForProvider.StatusCode != nil {
		res.SetStatusCode(*cr.Spec.ForProvider.StatusCode)
	}

	return res
}

// GenerateDeleteIntegrationResponseInput returns a deletion input.
func GenerateDeleteIntegrationResponseInput(cr *svcapitypes.IntegrationResponse) *svcsdk.DeleteIntegrationResponseInput {
	res := &svcsdk.DeleteIntegrationResponseInput{}

	if cr.Spec.ForProvider.HTTPMethod != nil {
		res.SetHttpMethod(*cr.Spec.ForProvider.HTTPMethod)
	}
	if cr.Spec.ForProvider.StatusCode != nil {
		res.SetStatusCode(*cr.Spec.ForProvider.StatusCode)
	}

	return res
}

// IsNotFound returns whether the given error is of type NotFound or not.
func IsNotFound(err error) bool {
	awsErr, ok := err.(awserr.Error)
	return ok && awsErr.Code() == "NotFoundException"
}