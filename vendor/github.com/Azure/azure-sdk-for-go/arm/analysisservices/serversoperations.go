package analysisservices

// Copyright (c) Microsoft and contributors.  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by Microsoft (R) AutoRest Code Generator 0.17.0.0
// Changes may cause incorrect behavior and will be lost if the code is
// regenerated.

import (
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/autorest/validation"
	"net/http"
)

// ServersOperationsClient is the the Azure Analysis Services Web API provides
// a RESTful set of web services that enables users to create, retrieve,
// update, and delete Analysis Services servers
type ServersOperationsClient struct {
	ManagementClient
}

// NewServersOperationsClient creates an instance of the
// ServersOperationsClient client.
func NewServersOperationsClient(subscriptionID string) ServersOperationsClient {
	return NewServersOperationsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewServersOperationsClientWithBaseURI creates an instance of the
// ServersOperationsClient client.
func NewServersOperationsClientWithBaseURI(baseURI string, subscriptionID string) ServersOperationsClient {
	return ServersOperationsClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// Create provisions the specified Analysis Services server based on the
// configuration specified in the request. This method may poll for
// completion. Polling can be canceled by passing the cancel channel
// argument. The channel will be used to cancel polling and any outstanding
// HTTP requests.
//
// resourceGroupName is the name of the Azure Resource group of which a given
// Analysis Services server is part. This name must be at least 1 character
// in length, and no more than 90. serverName is the name of the Analysis
// Services server. It must be a minimum of 3 characters, and a maximum of
// 63. serverParameters is contains the information used to provision the
// Analysis Services server.
func (client ServersOperationsClient) Create(resourceGroupName string, serverName string, serverParameters Server, cancel <-chan struct{}) (result autorest.Response, err error) {
	if err := validation.Validate([]validation.Validation{
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._\(\)]+$`, Chain: nil}}},
		{TargetValue: serverName,
			Constraints: []validation.Constraint{{Target: "serverName", Name: validation.MaxLength, Rule: 63, Chain: nil},
				{Target: "serverName", Name: validation.MinLength, Rule: 3, Chain: nil},
				{Target: "serverName", Name: validation.Pattern, Rule: `^[a-z][a-z0-9]*$`, Chain: nil}}},
		{TargetValue: serverParameters,
			Constraints: []validation.Constraint{{Target: "serverParameters.ServerProperties", Name: validation.Null, Rule: false,
				Chain: []validation.Constraint{{Target: "serverParameters.ServerProperties.ProvisioningState", Name: validation.ReadOnly, Rule: true, Chain: nil},
					{Target: "serverParameters.ServerProperties.ServerFullName", Name: validation.ReadOnly, Rule: true, Chain: nil},
				}}}}}); err != nil {
		return result, validation.NewErrorWithValidationError(err, "analysisservices.ServersOperationsClient", "Create")
	}

	req, err := client.CreatePreparer(resourceGroupName, serverName, serverParameters, cancel)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "analysisservices.ServersOperationsClient", "Create", nil, "Failure preparing request")
	}

	resp, err := client.CreateSender(req)
	if err != nil {
		result.Response = resp
		return result, autorest.NewErrorWithError(err, "analysisservices.ServersOperationsClient", "Create", resp, "Failure sending request")
	}

	result, err = client.CreateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "analysisservices.ServersOperationsClient", "Create", resp, "Failure responding to request")
	}

	return
}

// CreatePreparer prepares the Create request.
func (client ServersOperationsClient) CreatePreparer(resourceGroupName string, serverName string, serverParameters Server, cancel <-chan struct{}) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"serverName":        autorest.Encode("path", serverName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	queryParameters := map[string]interface{}{
		"api-version": client.APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsJSON(),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AnalysisServices/servers/{serverName}", pathParameters),
		autorest.WithJSON(serverParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{Cancel: cancel})
}

// CreateSender sends the Create request. The method will close the
// http.Response Body if it receives an error.
func (client ServersOperationsClient) CreateSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client,
		req,
		azure.DoPollForAsynchronous(client.PollingDelay))
}

// CreateResponder handles the response to the Create request. The method always
// closes the http.Response Body.
func (client ServersOperationsClient) CreateResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Delete deletes the specified Analysis Services server. This method may poll
// for completion. Polling can be canceled by passing the cancel channel
// argument. The channel will be used to cancel polling and any outstanding
// HTTP requests.
//
// resourceGroupName is the name of the Azure Resource group of which a given
// Analysis Services server is part. This name must be at least 1 character
// in length, and no more than 90. serverName is the name of the Analysis
// Services server. It must be at least 3 characters in length, and no more
// than 63.
func (client ServersOperationsClient) Delete(resourceGroupName string, serverName string, cancel <-chan struct{}) (result autorest.Response, err error) {
	if err := validation.Validate([]validation.Validation{
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._\(\)]+$`, Chain: nil}}},
		{TargetValue: serverName,
			Constraints: []validation.Constraint{{Target: "serverName", Name: validation.MaxLength, Rule: 63, Chain: nil},
				{Target: "serverName", Name: validation.MinLength, Rule: 3, Chain: nil},
				{Target: "serverName", Name: validation.Pattern, Rule: `^[a-z][a-z0-9]*$`, Chain: nil}}}}); err != nil {
		return result, validation.NewErrorWithValidationError(err, "analysisservices.ServersOperationsClient", "Delete")
	}

	req, err := client.DeletePreparer(resourceGroupName, serverName, cancel)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "analysisservices.ServersOperationsClient", "Delete", nil, "Failure preparing request")
	}

	resp, err := client.DeleteSender(req)
	if err != nil {
		result.Response = resp
		return result, autorest.NewErrorWithError(err, "analysisservices.ServersOperationsClient", "Delete", resp, "Failure sending request")
	}

	result, err = client.DeleteResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "analysisservices.ServersOperationsClient", "Delete", resp, "Failure responding to request")
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client ServersOperationsClient) DeletePreparer(resourceGroupName string, serverName string, cancel <-chan struct{}) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"serverName":        autorest.Encode("path", serverName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	queryParameters := map[string]interface{}{
		"api-version": client.APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AnalysisServices/servers/{serverName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{Cancel: cancel})
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client ServersOperationsClient) DeleteSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client,
		req,
		azure.DoPollForAsynchronous(client.PollingDelay))
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client ServersOperationsClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusNoContent, http.StatusAccepted),
		autorest.ByClosing())
	result.Response = resp
	return
}

// GetDetails gets details about the specified Analysis Services server.
//
// resourceGroupName is the name of the Azure Resource group of which a given
// Analysis Services server is part. This name must be at least 1 character
// in length, and no more than 90. serverName is the name of the Analysis
// Services server. It must be a minimum of 3 characters, and a maximum of
// 63.
func (client ServersOperationsClient) GetDetails(resourceGroupName string, serverName string) (result Server, err error) {
	if err := validation.Validate([]validation.Validation{
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._\(\)]+$`, Chain: nil}}},
		{TargetValue: serverName,
			Constraints: []validation.Constraint{{Target: "serverName", Name: validation.MaxLength, Rule: 63, Chain: nil},
				{Target: "serverName", Name: validation.MinLength, Rule: 3, Chain: nil},
				{Target: "serverName", Name: validation.Pattern, Rule: `^[a-z][a-z0-9]*$`, Chain: nil}}}}); err != nil {
		return result, validation.NewErrorWithValidationError(err, "analysisservices.ServersOperationsClient", "GetDetails")
	}

	req, err := client.GetDetailsPreparer(resourceGroupName, serverName)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "analysisservices.ServersOperationsClient", "GetDetails", nil, "Failure preparing request")
	}

	resp, err := client.GetDetailsSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "analysisservices.ServersOperationsClient", "GetDetails", resp, "Failure sending request")
	}

	result, err = client.GetDetailsResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "analysisservices.ServersOperationsClient", "GetDetails", resp, "Failure responding to request")
	}

	return
}

// GetDetailsPreparer prepares the GetDetails request.
func (client ServersOperationsClient) GetDetailsPreparer(resourceGroupName string, serverName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"serverName":        autorest.Encode("path", serverName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	queryParameters := map[string]interface{}{
		"api-version": client.APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AnalysisServices/servers/{serverName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// GetDetailsSender sends the GetDetails request. The method will close the
// http.Response Body if it receives an error.
func (client ServersOperationsClient) GetDetailsSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// GetDetailsResponder handles the response to the GetDetails request. The method always
// closes the http.Response Body.
func (client ServersOperationsClient) GetDetailsResponder(resp *http.Response) (result Server, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// List lists all the Analysis Services servers for the given subscription.
func (client ServersOperationsClient) List() (result Servers, err error) {
	req, err := client.ListPreparer()
	if err != nil {
		return result, autorest.NewErrorWithError(err, "analysisservices.ServersOperationsClient", "List", nil, "Failure preparing request")
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "analysisservices.ServersOperationsClient", "List", resp, "Failure sending request")
	}

	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "analysisservices.ServersOperationsClient", "List", resp, "Failure responding to request")
	}

	return
}

// ListPreparer prepares the List request.
func (client ServersOperationsClient) ListPreparer() (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"subscriptionId": autorest.Encode("path", client.SubscriptionID),
	}

	queryParameters := map[string]interface{}{
		"api-version": client.APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.AnalysisServices/servers", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client ServersOperationsClient) ListSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client ServersOperationsClient) ListResponder(resp *http.Response) (result Servers, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListByResourceGroup gets all the Analysis Services servers for the given
// resource group.
//
// resourceGroupName is the name of the Azure Resource group of which a given
// Analysis Services server is part. This name must be at least 1 character
// in length, and no more than 90.
func (client ServersOperationsClient) ListByResourceGroup(resourceGroupName string) (result Servers, err error) {
	if err := validation.Validate([]validation.Validation{
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._\(\)]+$`, Chain: nil}}}}); err != nil {
		return result, validation.NewErrorWithValidationError(err, "analysisservices.ServersOperationsClient", "ListByResourceGroup")
	}

	req, err := client.ListByResourceGroupPreparer(resourceGroupName)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "analysisservices.ServersOperationsClient", "ListByResourceGroup", nil, "Failure preparing request")
	}

	resp, err := client.ListByResourceGroupSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "analysisservices.ServersOperationsClient", "ListByResourceGroup", resp, "Failure sending request")
	}

	result, err = client.ListByResourceGroupResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "analysisservices.ServersOperationsClient", "ListByResourceGroup", resp, "Failure responding to request")
	}

	return
}

// ListByResourceGroupPreparer prepares the ListByResourceGroup request.
func (client ServersOperationsClient) ListByResourceGroupPreparer(resourceGroupName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	queryParameters := map[string]interface{}{
		"api-version": client.APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AnalysisServices/servers", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// ListByResourceGroupSender sends the ListByResourceGroup request. The method will close the
// http.Response Body if it receives an error.
func (client ServersOperationsClient) ListByResourceGroupSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// ListByResourceGroupResponder handles the response to the ListByResourceGroup request. The method always
// closes the http.Response Body.
func (client ServersOperationsClient) ListByResourceGroupResponder(resp *http.Response) (result Servers, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Resume resumes operation of the specified Analysis Services server
// instance. This method may poll for completion. Polling can be canceled by
// passing the cancel channel argument. The channel will be used to cancel
// polling and any outstanding HTTP requests.
//
// resourceGroupName is the name of the Azure Resource group of which a given
// Analysis Services server is part. This name must be at least 1 character
// in length, and no more than 90. serverName is the name of the Analysis
// Services server. It must be at least 3 characters in length, and no more
// than 63.
func (client ServersOperationsClient) Resume(resourceGroupName string, serverName string, cancel <-chan struct{}) (result autorest.Response, err error) {
	if err := validation.Validate([]validation.Validation{
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._\(\)]+$`, Chain: nil}}},
		{TargetValue: serverName,
			Constraints: []validation.Constraint{{Target: "serverName", Name: validation.MaxLength, Rule: 63, Chain: nil},
				{Target: "serverName", Name: validation.MinLength, Rule: 3, Chain: nil},
				{Target: "serverName", Name: validation.Pattern, Rule: `^[a-z][a-z0-9]*$`, Chain: nil}}}}); err != nil {
		return result, validation.NewErrorWithValidationError(err, "analysisservices.ServersOperationsClient", "Resume")
	}

	req, err := client.ResumePreparer(resourceGroupName, serverName, cancel)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "analysisservices.ServersOperationsClient", "Resume", nil, "Failure preparing request")
	}

	resp, err := client.ResumeSender(req)
	if err != nil {
		result.Response = resp
		return result, autorest.NewErrorWithError(err, "analysisservices.ServersOperationsClient", "Resume", resp, "Failure sending request")
	}

	result, err = client.ResumeResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "analysisservices.ServersOperationsClient", "Resume", resp, "Failure responding to request")
	}

	return
}

// ResumePreparer prepares the Resume request.
func (client ServersOperationsClient) ResumePreparer(resourceGroupName string, serverName string, cancel <-chan struct{}) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"serverName":        autorest.Encode("path", serverName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	queryParameters := map[string]interface{}{
		"api-version": client.APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AnalysisServices/servers/{serverName}/resume", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{Cancel: cancel})
}

// ResumeSender sends the Resume request. The method will close the
// http.Response Body if it receives an error.
func (client ServersOperationsClient) ResumeSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client,
		req,
		azure.DoPollForAsynchronous(client.PollingDelay))
}

// ResumeResponder handles the response to the Resume request. The method always
// closes the http.Response Body.
func (client ServersOperationsClient) ResumeResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Suspend supends operation of the specified Analysis Services server
// instance. This method may poll for completion. Polling can be canceled by
// passing the cancel channel argument. The channel will be used to cancel
// polling and any outstanding HTTP requests.
//
// resourceGroupName is the name of the Azure Resource group of which a given
// Analysis Services server is part. This name must be at least 1 character
// in length, and no more than 90. serverName is the name of the Analysis
// Services server. It must be at least 3 characters in length, and no more
// than 63.
func (client ServersOperationsClient) Suspend(resourceGroupName string, serverName string, cancel <-chan struct{}) (result autorest.Response, err error) {
	if err := validation.Validate([]validation.Validation{
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._\(\)]+$`, Chain: nil}}},
		{TargetValue: serverName,
			Constraints: []validation.Constraint{{Target: "serverName", Name: validation.MaxLength, Rule: 63, Chain: nil},
				{Target: "serverName", Name: validation.MinLength, Rule: 3, Chain: nil},
				{Target: "serverName", Name: validation.Pattern, Rule: `^[a-z][a-z0-9]*$`, Chain: nil}}}}); err != nil {
		return result, validation.NewErrorWithValidationError(err, "analysisservices.ServersOperationsClient", "Suspend")
	}

	req, err := client.SuspendPreparer(resourceGroupName, serverName, cancel)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "analysisservices.ServersOperationsClient", "Suspend", nil, "Failure preparing request")
	}

	resp, err := client.SuspendSender(req)
	if err != nil {
		result.Response = resp
		return result, autorest.NewErrorWithError(err, "analysisservices.ServersOperationsClient", "Suspend", resp, "Failure sending request")
	}

	result, err = client.SuspendResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "analysisservices.ServersOperationsClient", "Suspend", resp, "Failure responding to request")
	}

	return
}

// SuspendPreparer prepares the Suspend request.
func (client ServersOperationsClient) SuspendPreparer(resourceGroupName string, serverName string, cancel <-chan struct{}) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"serverName":        autorest.Encode("path", serverName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	queryParameters := map[string]interface{}{
		"api-version": client.APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AnalysisServices/servers/{serverName}/suspend", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{Cancel: cancel})
}

// SuspendSender sends the Suspend request. The method will close the
// http.Response Body if it receives an error.
func (client ServersOperationsClient) SuspendSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client,
		req,
		azure.DoPollForAsynchronous(client.PollingDelay))
}

// SuspendResponder handles the response to the Suspend request. The method always
// closes the http.Response Body.
func (client ServersOperationsClient) SuspendResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Update updates the current state of the specified Analysis Services server.
//
// resourceGroupName is the name of the Azure Resource group of which a given
// Analysis Services server is part. This name must be at least 1 character
// in length, and no more than 90. serverName is the name of the Analysis
// Services server. It must be at least 3 characters in length, and no more
// than 63. serverUpdateParameters is request object that contains the
// updated information for the server.
func (client ServersOperationsClient) Update(resourceGroupName string, serverName string, serverUpdateParameters ServerUpdateParameters) (result Server, err error) {
	if err := validation.Validate([]validation.Validation{
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._\(\)]+$`, Chain: nil}}},
		{TargetValue: serverName,
			Constraints: []validation.Constraint{{Target: "serverName", Name: validation.MaxLength, Rule: 63, Chain: nil},
				{Target: "serverName", Name: validation.MinLength, Rule: 3, Chain: nil},
				{Target: "serverName", Name: validation.Pattern, Rule: `^[a-z][a-z0-9]*$`, Chain: nil}}}}); err != nil {
		return result, validation.NewErrorWithValidationError(err, "analysisservices.ServersOperationsClient", "Update")
	}

	req, err := client.UpdatePreparer(resourceGroupName, serverName, serverUpdateParameters)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "analysisservices.ServersOperationsClient", "Update", nil, "Failure preparing request")
	}

	resp, err := client.UpdateSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "analysisservices.ServersOperationsClient", "Update", resp, "Failure sending request")
	}

	result, err = client.UpdateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "analysisservices.ServersOperationsClient", "Update", resp, "Failure responding to request")
	}

	return
}

// UpdatePreparer prepares the Update request.
func (client ServersOperationsClient) UpdatePreparer(resourceGroupName string, serverName string, serverUpdateParameters ServerUpdateParameters) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"serverName":        autorest.Encode("path", serverName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	queryParameters := map[string]interface{}{
		"api-version": client.APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsJSON(),
		autorest.AsPatch(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AnalysisServices/servers/{serverName}", pathParameters),
		autorest.WithJSON(serverUpdateParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// UpdateSender sends the Update request. The method will close the
// http.Response Body if it receives an error.
func (client ServersOperationsClient) UpdateSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// UpdateResponder handles the response to the Update request. The method always
// closes the http.Response Body.
func (client ServersOperationsClient) UpdateResponder(resp *http.Response) (result Server, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}