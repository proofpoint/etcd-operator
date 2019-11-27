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
	"github.com/satori/go.uuid"
	"net/http"
)

// AADObjectType enumerates the values for aad object type.
type AADObjectType string

const (
	// Group specifies the group state for aad object type.
	Group AADObjectType = "Group"
	// ServicePrincipal specifies the service principal state for aad object type.
	ServicePrincipal AADObjectType = "ServicePrincipal"
	// User specifies the user state for aad object type.
	User AADObjectType = "User"
)

// DataLakeAnalyticsAccountState enumerates the values for data lake analytics account state.
type DataLakeAnalyticsAccountState string

const (
	// Active specifies the active state for data lake analytics account state.
	Active DataLakeAnalyticsAccountState = "Active"
	// Suspended specifies the suspended state for data lake analytics account state.
	Suspended DataLakeAnalyticsAccountState = "Suspended"
)

// DataLakeAnalyticsAccountStatus enumerates the values for data lake analytics account status.
type DataLakeAnalyticsAccountStatus string

const (
	// Canceled specifies the canceled state for data lake analytics account status.
	Canceled DataLakeAnalyticsAccountStatus = "Canceled"
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
	// Undeleting specifies the undeleting state for data lake analytics account status.
	Undeleting DataLakeAnalyticsAccountStatus = "Undeleting"
)

// FirewallAllowAzureIpsState enumerates the values for firewall allow azure ips state.
type FirewallAllowAzureIpsState string

const (
	// Disabled specifies the disabled state for firewall allow azure ips state.
	Disabled FirewallAllowAzureIpsState = "Disabled"
	// Enabled specifies the enabled state for firewall allow azure ips state.
	Enabled FirewallAllowAzureIpsState = "Enabled"
)

// FirewallState enumerates the values for firewall state.
type FirewallState string

const (
	// FirewallStateDisabled specifies the firewall state disabled state for firewall state.
	FirewallStateDisabled FirewallState = "Disabled"
	// FirewallStateEnabled specifies the firewall state enabled state for firewall state.
	FirewallStateEnabled FirewallState = "Enabled"
)

// OperationOrigin enumerates the values for operation origin.
type OperationOrigin string

const (
	// OperationOriginSystem specifies the operation origin system state for operation origin.
	OperationOriginSystem OperationOrigin = "system"
	// OperationOriginUser specifies the operation origin user state for operation origin.
	OperationOriginUser OperationOrigin = "user"
	// OperationOriginUsersystem specifies the operation origin usersystem state for operation origin.
	OperationOriginUsersystem OperationOrigin = "user,system"
)

// SubscriptionState enumerates the values for subscription state.
type SubscriptionState string

const (
	// SubscriptionStateDeleted specifies the subscription state deleted state for subscription state.
	SubscriptionStateDeleted SubscriptionState = "Deleted"
	// SubscriptionStateRegistered specifies the subscription state registered state for subscription state.
	SubscriptionStateRegistered SubscriptionState = "Registered"
	// SubscriptionStateSuspended specifies the subscription state suspended state for subscription state.
	SubscriptionStateSuspended SubscriptionState = "Suspended"
	// SubscriptionStateUnregistered specifies the subscription state unregistered state for subscription state.
	SubscriptionStateUnregistered SubscriptionState = "Unregistered"
	// SubscriptionStateWarned specifies the subscription state warned state for subscription state.
	SubscriptionStateWarned SubscriptionState = "Warned"
)

// TierType enumerates the values for tier type.
type TierType string

const (
	// Commitment100000AUHours specifies the commitment 100000au hours state for tier type.
	Commitment100000AUHours TierType = "Commitment_100000AUHours"
	// Commitment10000AUHours specifies the commitment 10000au hours state for tier type.
	Commitment10000AUHours TierType = "Commitment_10000AUHours"
	// Commitment1000AUHours specifies the commitment 1000au hours state for tier type.
	Commitment1000AUHours TierType = "Commitment_1000AUHours"
	// Commitment100AUHours specifies the commitment 100au hours state for tier type.
	Commitment100AUHours TierType = "Commitment_100AUHours"
	// Commitment500000AUHours specifies the commitment 500000au hours state for tier type.
	Commitment500000AUHours TierType = "Commitment_500000AUHours"
	// Commitment50000AUHours specifies the commitment 50000au hours state for tier type.
	Commitment50000AUHours TierType = "Commitment_50000AUHours"
	// Commitment5000AUHours specifies the commitment 5000au hours state for tier type.
	Commitment5000AUHours TierType = "Commitment_5000AUHours"
	// Commitment500AUHours specifies the commitment 500au hours state for tier type.
	Commitment500AUHours TierType = "Commitment_500AUHours"
	// Consumption specifies the consumption state for tier type.
	Consumption TierType = "Consumption"
)

// AddDataLakeStoreParameters is additional Data Lake Store parameters.
type AddDataLakeStoreParameters struct {
	*DataLakeStoreAccountInfoProperties `json:"properties,omitempty"`
}

// AddStorageAccountParameters is storage account parameters for a storage account being added to a Data Lake Analytics
// account.
type AddStorageAccountParameters struct {
	*StorageAccountProperties `json:"properties,omitempty"`
}

// CapabilityInformation is subscription-level properties and limits for Data Lake Analytics
type CapabilityInformation struct {
	autorest.Response `json:"-"`
	SubscriptionID    *uuid.UUID        `json:"subscriptionId,omitempty"`
	State             SubscriptionState `json:"state,omitempty"`
	MaxAccountCount   *int32            `json:"maxAccountCount,omitempty"`
	AccountCount      *int32            `json:"accountCount,omitempty"`
	MigrationState    *bool             `json:"migrationState,omitempty"`
}

// CheckNameAvailabilityParameters is data Lake Analytics account name availability check parameters
type CheckNameAvailabilityParameters struct {
	Name *string `json:"name,omitempty"`
	Type *string `json:"type,omitempty"`
}

// ComputePolicy is the parameters used to create a new compute policy.
type ComputePolicy struct {
	autorest.Response        `json:"-"`
	Name                     *string `json:"name,omitempty"`
	*ComputePolicyProperties `json:"properties,omitempty"`
}

// ComputePolicyAccountCreateParameters is the parameters used to create a new compute policy.
type ComputePolicyAccountCreateParameters struct {
	Name                                     *string `json:"name,omitempty"`
	*ComputePolicyPropertiesCreateParameters `json:"properties,omitempty"`
}

// ComputePolicyCreateOrUpdateParameters is the parameters used to create a new compute policy.
type ComputePolicyCreateOrUpdateParameters struct {
	*ComputePolicyPropertiesCreateParameters `json:"properties,omitempty"`
}

// ComputePolicyListResult is the list of compute policies in the account.
type ComputePolicyListResult struct {
	autorest.Response `json:"-"`
	Value             *[]ComputePolicy `json:"value,omitempty"`
	NextLink          *string          `json:"nextLink,omitempty"`
}

// ComputePolicyListResultPreparer prepares a request to retrieve the next set of results. It returns
// nil if no more results exist.
func (client ComputePolicyListResult) ComputePolicyListResultPreparer() (*http.Request, error) {
	if client.NextLink == nil || len(to.String(client.NextLink)) <= 0 {
		return nil, nil
	}
	return autorest.Prepare(&http.Request{},
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(to.String(client.NextLink)))
}

// ComputePolicyProperties is the compute policy properties to use when creating a new compute policy
type ComputePolicyProperties struct {
	ObjectID                     *uuid.UUID    `json:"objectId,omitempty"`
	ObjectType                   AADObjectType `json:"objectType,omitempty"`
	MaxDegreeOfParallelismPerJob *int32        `json:"maxDegreeOfParallelismPerJob,omitempty"`
	MinPriorityPerJob            *int32        `json:"minPriorityPerJob,omitempty"`
}

// ComputePolicyPropertiesCreateParameters is the compute policy properties to use when creating a new compute policy
type ComputePolicyPropertiesCreateParameters struct {
	ObjectID                     *uuid.UUID    `json:"objectId,omitempty"`
	ObjectType                   AADObjectType `json:"objectType,omitempty"`
	MaxDegreeOfParallelismPerJob *int32        `json:"maxDegreeOfParallelismPerJob,omitempty"`
	MinPriorityPerJob            *int32        `json:"minPriorityPerJob,omitempty"`
}

// DataLakeAnalyticsAccount is a Data Lake Analytics account object, containing all information associated with the
// named Data Lake Analytics account.
type DataLakeAnalyticsAccount struct {
	autorest.Response                   `json:"-"`
	ID                                  *string             `json:"id,omitempty"`
	Name                                *string             `json:"name,omitempty"`
	Type                                *string             `json:"type,omitempty"`
	Location                            *string             `json:"location,omitempty"`
	Tags                                *map[string]*string `json:"tags,omitempty"`
	*DataLakeAnalyticsAccountProperties `json:"properties,omitempty"`
}

// DataLakeAnalyticsAccountBasic is a Data Lake Analytics account object, containing all information associated with
// the named Data Lake Analytics account.
type DataLakeAnalyticsAccountBasic struct {
	ID                                       *string             `json:"id,omitempty"`
	Name                                     *string             `json:"name,omitempty"`
	Type                                     *string             `json:"type,omitempty"`
	Location                                 *string             `json:"location,omitempty"`
	Tags                                     *map[string]*string `json:"tags,omitempty"`
	*DataLakeAnalyticsAccountPropertiesBasic `json:"properties,omitempty"`
}

// DataLakeAnalyticsAccountListDataLakeStoreResult is data Lake Account list information.
type DataLakeAnalyticsAccountListDataLakeStoreResult struct {
	autorest.Response `json:"-"`
	Value             *[]DataLakeStoreAccountInfo `json:"value,omitempty"`
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
	Value             *[]DataLakeAnalyticsAccountBasic `json:"value,omitempty"`
	NextLink          *string                          `json:"nextLink,omitempty"`
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
// Lake Analytics account. Returned only when retrieving a specific account.
type DataLakeAnalyticsAccountProperties struct {
	ProvisioningState            DataLakeAnalyticsAccountStatus          `json:"provisioningState,omitempty"`
	State                        DataLakeAnalyticsAccountState           `json:"state,omitempty"`
	CreationTime                 *date.Time                              `json:"creationTime,omitempty"`
	LastModifiedTime             *date.Time                              `json:"lastModifiedTime,omitempty"`
	Endpoint                     *string                                 `json:"endpoint,omitempty"`
	AccountID                    *uuid.UUID                              `json:"accountId,omitempty"`
	DefaultDataLakeStoreAccount  *string                                 `json:"defaultDataLakeStoreAccount,omitempty"`
	MaxDegreeOfParallelism       *int32                                  `json:"maxDegreeOfParallelism,omitempty"`
	QueryStoreRetention          *int32                                  `json:"queryStoreRetention,omitempty"`
	MaxJobCount                  *int32                                  `json:"maxJobCount,omitempty"`
	SystemMaxDegreeOfParallelism *int32                                  `json:"systemMaxDegreeOfParallelism,omitempty"`
	SystemMaxJobCount            *int32                                  `json:"systemMaxJobCount,omitempty"`
	DataLakeStoreAccounts        *[]DataLakeStoreAccountInfo             `json:"dataLakeStoreAccounts,omitempty"`
	StorageAccounts              *[]StorageAccountInfo                   `json:"storageAccounts,omitempty"`
	NewTier                      TierType                                `json:"newTier,omitempty"`
	CurrentTier                  TierType                                `json:"currentTier,omitempty"`
	FirewallState                FirewallState                           `json:"firewallState,omitempty"`
	FirewallAllowAzureIps        FirewallAllowAzureIpsState              `json:"firewallAllowAzureIps,omitempty"`
	FirewallRules                *[]FirewallRule                         `json:"firewallRules,omitempty"`
	MaxDegreeOfParallelismPerJob *int32                                  `json:"maxDegreeOfParallelismPerJob,omitempty"`
	MinPriorityPerJob            *int32                                  `json:"minPriorityPerJob,omitempty"`
	ComputePolicies              *[]ComputePolicyAccountCreateParameters `json:"computePolicies,omitempty"`
}

// DataLakeAnalyticsAccountPropertiesBasic is the basic account specific properties that are associated with an
// underlying Data Lake Analytics account.
type DataLakeAnalyticsAccountPropertiesBasic struct {
	ProvisioningState DataLakeAnalyticsAccountStatus `json:"provisioningState,omitempty"`
	State             DataLakeAnalyticsAccountState  `json:"state,omitempty"`
	CreationTime      *date.Time                     `json:"creationTime,omitempty"`
	LastModifiedTime  *date.Time                     `json:"lastModifiedTime,omitempty"`
	Endpoint          *string                        `json:"endpoint,omitempty"`
	AccountID         *uuid.UUID                     `json:"accountId,omitempty"`
}

// DataLakeAnalyticsAccountUpdateParameters is the parameters that can be used to update an existing Data Lake
// Analytics account.
type DataLakeAnalyticsAccountUpdateParameters struct {
	Tags                                      *map[string]*string `json:"tags,omitempty"`
	*UpdateDataLakeAnalyticsAccountProperties `json:"properties,omitempty"`
}

// DataLakeAnalyticsFirewallRuleListResult is data Lake Analytics firewall rule list information.
type DataLakeAnalyticsFirewallRuleListResult struct {
	autorest.Response `json:"-"`
	Value             *[]FirewallRule `json:"value,omitempty"`
	NextLink          *string         `json:"nextLink,omitempty"`
}

// DataLakeAnalyticsFirewallRuleListResultPreparer prepares a request to retrieve the next set of results. It returns
// nil if no more results exist.
func (client DataLakeAnalyticsFirewallRuleListResult) DataLakeAnalyticsFirewallRuleListResultPreparer() (*http.Request, error) {
	if client.NextLink == nil || len(to.String(client.NextLink)) <= 0 {
		return nil, nil
	}
	return autorest.Prepare(&http.Request{},
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(to.String(client.NextLink)))
}

// DataLakeStoreAccountInfo is data Lake Store account information.
type DataLakeStoreAccountInfo struct {
	autorest.Response                   `json:"-"`
	ID                                  *string `json:"id,omitempty"`
	Name                                *string `json:"name,omitempty"`
	Type                                *string `json:"type,omitempty"`
	*DataLakeStoreAccountInfoProperties `json:"properties,omitempty"`
}

// DataLakeStoreAccountInfoProperties is data Lake Store account properties information.
type DataLakeStoreAccountInfoProperties struct {
	Suffix *string `json:"suffix,omitempty"`
}

// FirewallRule is data Lake Analytics firewall rule information
type FirewallRule struct {
	autorest.Response       `json:"-"`
	ID                      *string `json:"id,omitempty"`
	Name                    *string `json:"name,omitempty"`
	Type                    *string `json:"type,omitempty"`
	*FirewallRuleProperties `json:"properties,omitempty"`
}

// FirewallRuleProperties is data Lake Analytics firewall rule properties information
type FirewallRuleProperties struct {
	StartIPAddress *string `json:"startIpAddress,omitempty"`
	EndIPAddress   *string `json:"endIpAddress,omitempty"`
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

// ListStorageContainersResult is the list of blob containers associated with the storage account attached to the Data
// Lake Analytics account.
type ListStorageContainersResult struct {
	autorest.Response `json:"-"`
	Value             *[]StorageContainer `json:"value,omitempty"`
	NextLink          *string             `json:"nextLink,omitempty"`
}

// ListStorageContainersResultPreparer prepares a request to retrieve the next set of results. It returns
// nil if no more results exist.
func (client ListStorageContainersResult) ListStorageContainersResultPreparer() (*http.Request, error) {
	if client.NextLink == nil || len(to.String(client.NextLink)) <= 0 {
		return nil, nil
	}
	return autorest.Prepare(&http.Request{},
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(to.String(client.NextLink)))
}

// NameAvailabilityInformation is data Lake Analytics account name availability result information
type NameAvailabilityInformation struct {
	autorest.Response `json:"-"`
	NameAvailable     *bool   `json:"nameAvailable,omitempty"`
	Reason            *string `json:"reason,omitempty"`
	Message           *string `json:"message,omitempty"`
}

// Operation is an available operation for Data Lake Analytics
type Operation struct {
	Name    *string           `json:"name,omitempty"`
	Display *OperationDisplay `json:"display,omitempty"`
	Origin  OperationOrigin   `json:"origin,omitempty"`
}

// OperationDisplay is the display information for a particular operation
type OperationDisplay struct {
	Provider    *string `json:"provider,omitempty"`
	Resource    *string `json:"resource,omitempty"`
	Operation   *string `json:"operation,omitempty"`
	Description *string `json:"description,omitempty"`
}

// OperationListResult is the list of available operations for Data Lake Analytics
type OperationListResult struct {
	autorest.Response `json:"-"`
	Value             *[]Operation `json:"value,omitempty"`
	NextLink          *string      `json:"nextLink,omitempty"`
}

// OptionalSubResource is the Resource model definition for a nested resource with no required properties.
type OptionalSubResource struct {
	ID   *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
	Type *string `json:"type,omitempty"`
}

// Resource is the Resource model definition.
type Resource struct {
	ID       *string             `json:"id,omitempty"`
	Name     *string             `json:"name,omitempty"`
	Type     *string             `json:"type,omitempty"`
	Location *string             `json:"location,omitempty"`
	Tags     *map[string]*string `json:"tags,omitempty"`
}

// SasTokenInfo is SAS token information.
type SasTokenInfo struct {
	AccessToken *string `json:"accessToken,omitempty"`
}

// StorageAccountInfo is azure Storage account information.
type StorageAccountInfo struct {
	autorest.Response         `json:"-"`
	ID                        *string `json:"id,omitempty"`
	Name                      *string `json:"name,omitempty"`
	Type                      *string `json:"type,omitempty"`
	*StorageAccountProperties `json:"properties,omitempty"`
}

// StorageAccountProperties is azure Storage account properties information.
type StorageAccountProperties struct {
	AccessKey *string `json:"accessKey,omitempty"`
	Suffix    *string `json:"suffix,omitempty"`
}

// StorageContainer is azure Storage blob container information.
type StorageContainer struct {
	autorest.Response           `json:"-"`
	ID                          *string `json:"id,omitempty"`
	Name                        *string `json:"name,omitempty"`
	Type                        *string `json:"type,omitempty"`
	*StorageContainerProperties `json:"properties,omitempty"`
}

// StorageContainerProperties is azure Storage blob container properties information.
type StorageContainerProperties struct {
	LastModifiedTime *date.Time `json:"lastModifiedTime,omitempty"`
}

// SubResource is the Sub Resource model definition.
type SubResource struct {
	ID   *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
	Type *string `json:"type,omitempty"`
}

// UpdateDataLakeAnalyticsAccountProperties is the properties to update that are associated with an underlying Data
// Lake Analytics account to.
type UpdateDataLakeAnalyticsAccountProperties struct {
	MaxDegreeOfParallelism       *int32                     `json:"maxDegreeOfParallelism,omitempty"`
	QueryStoreRetention          *int32                     `json:"queryStoreRetention,omitempty"`
	MaxJobCount                  *int32                     `json:"maxJobCount,omitempty"`
	NewTier                      TierType                   `json:"newTier,omitempty"`
	FirewallState                FirewallState              `json:"firewallState,omitempty"`
	FirewallAllowAzureIps        FirewallAllowAzureIpsState `json:"firewallAllowAzureIps,omitempty"`
	FirewallRules                *[]FirewallRule            `json:"firewallRules,omitempty"`
	MaxDegreeOfParallelismPerJob *int32                     `json:"maxDegreeOfParallelismPerJob,omitempty"`
	MinPriorityPerJob            *int32                     `json:"minPriorityPerJob,omitempty"`
	ComputePolicies              *[]ComputePolicy           `json:"computePolicies,omitempty"`
}

// UpdateFirewallRuleParameters is data Lake Analytics firewall rule update parameters
type UpdateFirewallRuleParameters struct {
	*UpdateFirewallRuleProperties `json:"properties,omitempty"`
}

// UpdateFirewallRuleProperties is data Lake Analytics firewall rule properties information
type UpdateFirewallRuleProperties struct {
	StartIPAddress *string `json:"startIpAddress,omitempty"`
	EndIPAddress   *string `json:"endIpAddress,omitempty"`
}

// UpdateStorageAccountParameters is storage account parameters for a storage account being updated in a Data Lake
// Analytics account.
type UpdateStorageAccountParameters struct {
	*UpdateStorageAccountProperties `json:"properties,omitempty"`
}

// UpdateStorageAccountProperties is azure Storage account properties information to update.
type UpdateStorageAccountProperties struct {
	AccessKey *string `json:"accessKey,omitempty"`
	Suffix    *string `json:"suffix,omitempty"`
}
