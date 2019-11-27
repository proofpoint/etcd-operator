package account

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
	"github.com/Azure/go-autorest/autorest/date"
	"github.com/Azure/go-autorest/autorest/to"
	"net/http"
)

// DataLakeAnalyticsAccountState enumerates the values for data lake analytics account state.
type DataLakeAnalyticsAccountState string

const (
	// Active specifies the active state for data lake analytics account state.
	Active DataLakeAnalyticsAccountState = "active"
	// Suspended specifies the suspended state for data lake analytics account state.
	Suspended DataLakeAnalyticsAccountState = "suspended"
)

// DataLakeAnalyticsAccountStatus enumerates the values for data lake analytics account status.
type DataLakeAnalyticsAccountStatus string

const (
	// Creating specifies the creating state for data lake analytics account status.
	Creating DataLakeAnalyticsAccountStatus = "Creating"
	// Deleted specifies the deleted state for data lake analytics account status.
	Deleted DataLakeAnalyticsAccountStatus = "Deleted"
	// Deleting specifies the deleting state for data lake analytics account status.
	Deleting DataLakeAnalyticsAccountStatus = "Deleting"
	// Failed specifies the failed state for data lake analytics account status.
	Failed DataLakeAnalyticsAccountStatus = "Failed"
	// Patching specifies the patching state for data lake analytics account status.
	Patching DataLakeAnalyticsAccountStatus = "Patching"
	// Resuming specifies the resuming state for data lake analytics account status.
	Resuming DataLakeAnalyticsAccountStatus = "Resuming"
	// Running specifies the running state for data lake analytics account status.
	Running DataLakeAnalyticsAccountStatus = "Running"
	// Succeeded specifies the succeeded state for data lake analytics account status.
	Succeeded DataLakeAnalyticsAccountStatus = "Succeeded"
	// Suspending specifies the suspending state for data lake analytics account status.
	Suspending DataLakeAnalyticsAccountStatus = "Suspending"
)

// OperationStatus enumerates the values for operation status.
type OperationStatus string

const (
	// OperationStatusFailed specifies the operation status failed state for operation status.
	OperationStatusFailed OperationStatus = "Failed"
	// OperationStatusInProgress specifies the operation status in progress state for operation status.
	OperationStatusInProgress OperationStatus = "InProgress"
	// OperationStatusSucceeded specifies the operation status succeeded state for operation status.
	OperationStatusSucceeded OperationStatus = "Succeeded"
)

// AddDataLakeStoreParameters is additional Data Lake Store parameters.
type AddDataLakeStoreParameters struct {
	Properties *DataLakeStoreAccountInfoProperties `json:"properties,omitempty"`
}

// AddStorageAccountParameters is additional Azure Storage account parameters.
type AddStorageAccountParameters struct {
	Properties *StorageAccountProperties `json:"properties,omitempty"`
}

// AzureAsyncOperationResult is the response body contains the status of the specified asynchronous operation,
// indicating whether it has succeeded, is inprogress, or has failed. Note that this status is distinct from the HTTP
// status code returned for the Get Operation Status operation itself. If the asynchronous operation succeeded, the
// response body includes the HTTP status code for the successful request. If the asynchronous operation failed, the
// response body includes the HTTP status code for the failed request and error information regarding the failure.
type AzureAsyncOperationResult struct {
	Status OperationStatus `json:"status,omitempty"`
	Error  *Error          `json:"error,omitempty"`
}

// BlobContainer is azure Storage blob container information.
type BlobContainer struct {
	autorest.Response `json:"-"`
	Name              *string                  `json:"name,omitempty"`
	ID                *string                  `json:"id,omitempty"`
	Type              *string                  `json:"type,omitempty"`
	Properties        *BlobContainerProperties `json:"properties,omitempty"`
}

// BlobContainerProperties is azure Storage blob container properties information.
type BlobContainerProperties struct {
	LastModifiedTime *date.Time `json:"lastModifiedTime,omitempty"`
}

// DataLakeAnalyticsAccount is a Data Lake Analytics account object, containing all information associated with the
// named Data Lake Analytics account.
type DataLakeAnalyticsAccount struct {
	autorest.Response `json:"-"`
	Location          *string                             `json:"location,omitempty"`
	Name              *string                             `json:"name,omitempty"`
	Type              *string                             `json:"type,omitempty"`
	ID                *string                             `json:"id,omitempty"`
	Tags              *map[string]*string                 `json:"tags,omitempty"`
	Properties        *DataLakeAnalyticsAccountProperties `json:"properties,omitempty"`
}

// DataLakeAnalyticsAccountListDataLakeStoreResult is data Lake Account list information.
type DataLakeAnalyticsAccountListDataLakeStoreResult struct {
	autorest.Response `json:"-"`
	Value             *[]DataLakeStoreAccountInfo `json:"value,omitempty"`
	Count             *int32                      `json:"count,omitempty"`
	NextLink          *string                     `json:"nextLink,omitempty"`
}

// DataLakeAnalyticsAccountListDataLakeStoreResultPreparer prepares a request to retrieve the next set of results. It returns
// nil if no more results exist.
func (client DataLakeAnalyticsAccountListDataLakeStoreResult) DataLakeAnalyticsAccountListDataLakeStoreResultPreparer() (*http.Request, error) {
	if client.NextLink == nil || len(to.String(client.NextLink)) <= 0 {
		return nil, nil
	}
	return autorest.Prepare(&http.Request{},
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(to.String(client.NextLink)))
}

// DataLakeAnalyticsAccountListResult is dataLakeAnalytics Account list information.
type DataLakeAnalyticsAccountListResult struct {
	autorest.Response `json:"-"`
	Value             *[]DataLakeAnalyticsAccount `json:"value,omitempty"`
	NextLink          *string                     `json:"nextLink,omitempty"`
}

// DataLakeAnalyticsAccountListResultPreparer prepares a request to retrieve the next set of results. It returns
// nil if no more results exist.
func (client DataLakeAnalyticsAccountListResult) DataLakeAnalyticsAccountListResultPreparer() (*http.Request, error) {
	if client.NextLink == nil || len(to.String(client.NextLink)) <= 0 {
		return nil, nil
	}
	return autorest.Prepare(&http.Request{},
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(to.String(client.NextLink)))
}

// DataLakeAnalyticsAccountListStorageAccountsResult is azure Storage Account list information.
type DataLakeAnalyticsAccountListStorageAccountsResult struct {
	autorest.Response `json:"-"`
	Value             *[]StorageAccountInfo `json:"value,omitempty"`
	Count             *int32                `json:"count,omitempty"`
	NextLink          *string               `json:"nextLink,omitempty"`
}

// DataLakeAnalyticsAccountListStorageAccountsResultPreparer prepares a request to retrieve the next set of results. It returns
// nil if no more results exist.
func (client DataLakeAnalyticsAccountListStorageAccountsResult) DataLakeAnalyticsAccountListStorageAccountsResultPreparer() (*http.Request, error) {
	if client.NextLink == nil || len(to.String(client.NextLink)) <= 0 {
		return nil, nil
	}
	return autorest.Prepare(&http.Request{},
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(to.String(client.NextLink)))
}

// DataLakeAnalyticsAccountProperties is the account specific properties that are associated with an underlying Data
// Lake Analytics account.
type DataLakeAnalyticsAccountProperties struct {
	ProvisioningState           DataLakeAnalyticsAccountStatus `json:"provisioningState,omitempty"`
	State                       DataLakeAnalyticsAccountState  `json:"state,omitempty"`
	DefaultDataLakeStoreAccount *string                        `json:"defaultDataLakeStoreAccount,omitempty"`
	MaxDegreeOfParallelism      *int32                         `json:"maxDegreeOfParallelism,omitempty"`
	MaxJobCount                 *int32                         `json:"maxJobCount,omitempty"`
	DataLakeStoreAccounts       *[]DataLakeStoreAccountInfo    `json:"dataLakeStoreAccounts,omitempty"`
	StorageAccounts             *[]StorageAccountInfo          `json:"storageAccounts,omitempty"`
	CreationTime                *date.Time                     `json:"creationTime,omitempty"`
	LastModifiedTime            *date.Time                     `json:"lastModifiedTime,omitempty"`
	Endpoint                    *string                        `json:"endpoint,omitempty"`
}

// DataLakeStoreAccountInfo is data Lake Store account information.
type DataLakeStoreAccountInfo struct {
	autorest.Response `json:"-"`
	Name              *string                             `json:"name,omitempty"`
	Properties        *DataLakeStoreAccountInfoProperties `json:"properties,omitempty"`
}

// DataLakeStoreAccountInfoProperties is data Lake Store account properties information.
type DataLakeStoreAccountInfoProperties struct {
	Suffix *string `json:"suffix,omitempty"`
}

// Error is generic resource error information.
type Error struct {
	Code       *string         `json:"code,omitempty"`
	Message    *string         `json:"message,omitempty"`
	Target     *string         `json:"target,omitempty"`
	Details    *[]ErrorDetails `json:"details,omitempty"`
	InnerError *InnerError     `json:"innerError,omitempty"`
}

// ErrorDetails is generic resource error details information.
type ErrorDetails struct {
	Code    *string `json:"code,omitempty"`
	Message *string `json:"message,omitempty"`
	Target  *string `json:"target,omitempty"`
}

// InnerError is generic resource inner error information.
type InnerError struct {
	Trace   *string `json:"trace,omitempty"`
	Context *string `json:"context,omitempty"`
}

// ListBlobContainersResult is the list of blob containers associated with the storage account attached to the Data
// Lake Analytics account.
type ListBlobContainersResult struct {
	autorest.Response `json:"-"`
	Value             *[]BlobContainer `json:"value,omitempty"`
	NextLink          *string          `json:"nextLink,omitempty"`
}

// ListBlobContainersResultPreparer prepares a request to retrieve the next set of results. It returns
// nil if no more results exist.
func (client ListBlobContainersResult) ListBlobContainersResultPreparer() (*http.Request, error) {
	if client.NextLink == nil || len(to.String(client.NextLink)) <= 0 {
		return nil, nil
	}
	return autorest.Prepare(&http.Request{},
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(to.String(client.NextLink)))
}

// ListSasTokensResult is the SAS response that contains the storage account, container and associated SAS token for
// connection use.
type ListSasTokensResult struct {
	autorest.Response `json:"-"`
	Value             *[]SasTokenInfo `json:"value,omitempty"`
	NextLink          *string         `json:"nextLink,omitempty"`
}

// ListSasTokensResultPreparer prepares a request to retrieve the next set of results. It returns
// nil if no more results exist.
func (client ListSasTokensResult) ListSasTokensResultPreparer() (*http.Request, error) {
	if client.NextLink == nil || len(to.String(client.NextLink)) <= 0 {
		return nil, nil
	}
	return autorest.Prepare(&http.Request{},
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(to.String(client.NextLink)))
}

// SasTokenInfo is SAS token information.
type SasTokenInfo struct {
	AccessToken *string `json:"accessToken,omitempty"`
}

// StorageAccountInfo is azure Storage account information.
type StorageAccountInfo struct {
	autorest.Response `json:"-"`
	Name              *string                   `json:"name,omitempty"`
	Properties        *StorageAccountProperties `json:"properties,omitempty"`
}

// StorageAccountProperties is azure Storage account properties information.
type StorageAccountProperties struct {
	AccessKey *string `json:"accessKey,omitempty"`
	Suffix    *string `json:"suffix,omitempty"`
}
