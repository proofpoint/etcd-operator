package recoveryservicessiterecovery

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
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"net/http"
)

// ReplicationPoliciesClient is the client for the ReplicationPolicies methods of the Recoveryservicessiterecovery
// service.
type ReplicationPoliciesClient struct {
	ManagementClient
}

// NewReplicationPoliciesClient creates an instance of the ReplicationPoliciesClient client.
func NewReplicationPoliciesClient(subscriptionID string, resourceGroupName string, resourceName string) ReplicationPoliciesClient {
	return NewReplicationPoliciesClientWithBaseURI(DefaultBaseURI, subscriptionID, resourceGroupName, resourceName)
}

// NewReplicationPoliciesClientWithBaseURI creates an instance of the ReplicationPoliciesClient client.
func NewReplicationPoliciesClientWithBaseURI(baseURI string, subscriptionID string, resourceGroupName string, resourceName string) ReplicationPoliciesClient {
	return ReplicationPoliciesClient{NewWithBaseURI(baseURI, subscriptionID, resourceGroupName, resourceName)}
}

// Create the operation to create a replication policy This method may poll for completion. Polling can be canceled by
// passing the cancel channel argument. The channel will be used to cancel polling and any outstanding HTTP requests.
//
// policyName is replication policy name input is create policy input
func (client ReplicationPoliciesClient) Create(policyName string, input CreatePolicyInput, cancel <-chan struct{}) (<-chan Policy, <-chan error) {
	resultChan := make(chan Policy, 1)
	errChan := make(chan error, 1)
	go func() {
		var err error
		var result Policy
		defer func() {
			if err != nil {
				errChan <- err
			}
			resultChan <- result
			close(resultChan)
			close(errChan)
		}()
		req, err := client.CreatePreparer(policyName, input, cancel)
		if err != nil {
			err = autorest.NewErrorWithError(err, "recoveryservicessiterecovery.ReplicationPoliciesClient", "Create", nil, "Failure preparing request")
			return
		}

		resp, err := client.CreateSender(req)
		if err != nil {
			result.Response = autorest.Response{Response: resp}
			err = autorest.NewErrorWithError(err, "recoveryservicessiterecovery.ReplicationPoliciesClient", "Create", resp, "Failure sending request")
			return
		}

		result, err = client.CreateResponder(resp)
		if err != nil {
			err = autorest.NewErrorWithError(err, "recoveryservicessiterecovery.ReplicationPoliciesClient", "Create", resp, "Failure responding to request")
		}
	}()
	return resultChan, errChan
}

// CreatePreparer prepares the Create request.
func (client ReplicationPoliciesClient) CreatePreparer(policyName string, input CreatePolicyInput, cancel <-chan struct{}) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"policyName":        autorest.Encode("path", policyName),
		"resourceGroupName": autorest.Encode("path", client.ResourceGroupName),
		"resourceName":      autorest.Encode("path", client.ResourceName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2016-08-10"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsJSON(),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/Subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{resourceName}/replicationPolicies/{policyName}", pathParameters),
		autorest.WithJSON(input),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{Cancel: cancel})
}

// CreateSender sends the Create request. The method will close the
// http.Response Body if it receives an error.
func (client ReplicationPoliciesClient) CreateSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client,
		req,
		azure.DoRetryWithRegistration(client.Client),
		azure.DoPollForAsynchronous(client.PollingDelay))
}

// CreateResponder handles the response to the Create request. The method always
// closes the http.Response Body.
func (client ReplicationPoliciesClient) CreateResponder(resp *http.Response) (result Policy, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete the operation to delete a replication policy. This method may poll for completion. Polling can be canceled by
// passing the cancel channel argument. The channel will be used to cancel polling and any outstanding HTTP requests.
//
// policyName is replication policy name.
func (client ReplicationPoliciesClient) Delete(policyName string, cancel <-chan struct{}) (<-chan autorest.Response, <-chan error) {
	resultChan := make(chan autorest.Response, 1)
	errChan := make(chan error, 1)
	go func() {
		var err error
		var result autorest.Response
		defer func() {
			if err != nil {
				errChan <- err
			}
			resultChan <- result
			close(resultChan)
			close(errChan)
		}()
		req, err := client.DeletePreparer(policyName, cancel)
		if err != nil {
			err = autorest.NewErrorWithError(err, "recoveryservicessiterecovery.ReplicationPoliciesClient", "Delete", nil, "Failure preparing request")
			return
		}

		resp, err := client.DeleteSender(req)
		if err != nil {
			result.Response = resp
			err = autorest.NewErrorWithError(err, "recoveryservicessiterecovery.ReplicationPoliciesClient", "Delete", resp, "Failure sending request")
			return
		}

		result, err = client.DeleteResponder(resp)
		if err != nil {
			err = autorest.NewErrorWithError(err, "recoveryservicessiterecovery.ReplicationPoliciesClient", "Delete", resp, "Failure responding to request")
		}
	}()
	return resultChan, errChan
}

// DeletePreparer prepares the Delete request.
func (client ReplicationPoliciesClient) DeletePreparer(policyName string, cancel <-chan struct{}) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"policyName":        autorest.Encode("path", policyName),
		"resourceGroupName": autorest.Encode("path", client.ResourceGroupName),
		"resourceName":      autorest.Encode("path", client.ResourceName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2016-08-10"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/Subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{resourceName}/replicationPolicies/{policyName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{Cancel: cancel})
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client ReplicationPoliciesClient) DeleteSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client,
		req,
		azure.DoRetryWithRegistration(client.Client),
		azure.DoPollForAsynchronous(client.PollingDelay))
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client ReplicationPoliciesClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Get gets the details of a replication policy.
//
// policyName is replication policy name.
func (client ReplicationPoliciesClient) Get(policyName string) (result Policy, err error) {
	req, err := client.GetPreparer(policyName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "recoveryservicessiterecovery.ReplicationPoliciesClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "recoveryservicessiterecovery.ReplicationPoliciesClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "recoveryservicessiterecovery.ReplicationPoliciesClient", "Get", resp, "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client ReplicationPoliciesClient) GetPreparer(policyName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"policyName":        autorest.Encode("path", policyName),
		"resourceGroupName": autorest.Encode("path", client.ResourceGroupName),
		"resourceName":      autorest.Encode("path", client.ResourceName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2016-08-10"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/Subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{resourceName}/replicationPolicies/{policyName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client ReplicationPoliciesClient) GetSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client,
		req,
		azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client ReplicationPoliciesClient) GetResponder(resp *http.Response) (result Policy, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// List lists the replication policies for a vault.
func (client ReplicationPoliciesClient) List() (result PolicyCollection, err error) {
	req, err := client.ListPreparer()
	if err != nil {
		err = autorest.NewErrorWithError(err, "recoveryservicessiterecovery.ReplicationPoliciesClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "recoveryservicessiterecovery.ReplicationPoliciesClient", "List", resp, "Failure sending request")
		return
	}

	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "recoveryservicessiterecovery.ReplicationPoliciesClient", "List", resp, "Failure responding to request")
	}

	return
}

// ListPreparer prepares the List request.
func (client ReplicationPoliciesClient) ListPreparer() (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", client.ResourceGroupName),
		"resourceName":      autorest.Encode("path", client.ResourceName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2016-08-10"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/Subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{resourceName}/replicationPolicies", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client ReplicationPoliciesClient) ListSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client,
		req,
		azure.DoRetryWithRegistration(client.Client))
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client ReplicationPoliciesClient) ListResponder(resp *http.Response) (result PolicyCollection, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListNextResults retrieves the next set of results, if any.
func (client ReplicationPoliciesClient) ListNextResults(lastResults PolicyCollection) (result PolicyCollection, err error) {
	req, err := lastResults.PolicyCollectionPreparer()
	if err != nil {
		return result, autorest.NewErrorWithError(err, "recoveryservicessiterecovery.ReplicationPoliciesClient", "List", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "recoveryservicessiterecovery.ReplicationPoliciesClient", "List", resp, "Failure sending next results request")
	}

	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "recoveryservicessiterecovery.ReplicationPoliciesClient", "List", resp, "Failure responding to next results request")
	}

	return
}

// ListComplete gets all elements from the list without paging.
func (client ReplicationPoliciesClient) ListComplete(cancel <-chan struct{}) (<-chan Policy, <-chan error) {
	resultChan := make(chan Policy)
	errChan := make(chan error, 1)
	go func() {
		defer func() {
			close(resultChan)
			close(errChan)
		}()
		list, err := client.List()
		if err != nil {
			errChan <- err
			return
		}
		if list.Value != nil {
			for _, item := range *list.Value {
				select {
				case <-cancel:
					return
				case resultChan <- item:
					// Intentionally left blank
				}
			}
		}
		for list.NextLink != nil {
			list, err = client.ListNextResults(list)
			if err != nil {
				errChan <- err
				return
			}
			if list.Value != nil {
				for _, item := range *list.Value {
					select {
					case <-cancel:
						return
					case resultChan <- item:
						// Intentionally left blank
					}
				}
			}
		}
	}()
	return resultChan, errChan
}

// Update the operation to update a replication policy. This method may poll for completion. Polling can be canceled by
// passing the cancel channel argument. The channel will be used to cancel polling and any outstanding HTTP requests.
//
// policyName is protection profile Id. input is update Protection Profile Input
func (client ReplicationPoliciesClient) Update(policyName string, input UpdatePolicyInput, cancel <-chan struct{}) (<-chan Policy, <-chan error) {
	resultChan := make(chan Policy, 1)
	errChan := make(chan error, 1)
	go func() {
		var err error
		var result Policy
		defer func() {
			if err != nil {
				errChan <- err
			}
			resultChan <- result
			close(resultChan)
			close(errChan)
		}()
		req, err := client.UpdatePreparer(policyName, input, cancel)
		if err != nil {
			err = autorest.NewErrorWithError(err, "recoveryservicessiterecovery.ReplicationPoliciesClient", "Update", nil, "Failure preparing request")
			return
		}

		resp, err := client.UpdateSender(req)
		if err != nil {
			result.Response = autorest.Response{Response: resp}
			err = autorest.NewErrorWithError(err, "recoveryservicessiterecovery.ReplicationPoliciesClient", "Update", resp, "Failure sending request")
			return
		}

		result, err = client.UpdateResponder(resp)
		if err != nil {
			err = autorest.NewErrorWithError(err, "recoveryservicessiterecovery.ReplicationPoliciesClient", "Update", resp, "Failure responding to request")
		}
	}()
	return resultChan, errChan
}

// UpdatePreparer prepares the Update request.
func (client ReplicationPoliciesClient) UpdatePreparer(policyName string, input UpdatePolicyInput, cancel <-chan struct{}) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"policyName":        autorest.Encode("path", policyName),
		"resourceGroupName": autorest.Encode("path", client.ResourceGroupName),
		"resourceName":      autorest.Encode("path", client.ResourceName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2016-08-10"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsJSON(),
		autorest.AsPatch(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/Subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{resourceName}/replicationPolicies/{policyName}", pathParameters),
		autorest.WithJSON(input),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{Cancel: cancel})
}

// UpdateSender sends the Update request. The method will close the
// http.Response Body if it receives an error.
func (client ReplicationPoliciesClient) UpdateSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client,
		req,
		azure.DoRetryWithRegistration(client.Client),
		azure.DoPollForAsynchronous(client.PollingDelay))
}

// UpdateResponder handles the response to the Update request. The method always
// closes the http.Response Body.
func (client ReplicationPoliciesClient) UpdateResponder(resp *http.Response) (result Policy, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
