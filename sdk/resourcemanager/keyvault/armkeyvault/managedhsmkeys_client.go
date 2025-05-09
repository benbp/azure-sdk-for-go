// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armkeyvault

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// ManagedHsmKeysClient contains the methods for the ManagedHsmKeys group.
// Don't use this type directly, use NewManagedHsmKeysClient() instead.
type ManagedHsmKeysClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewManagedHsmKeysClient creates a new instance of ManagedHsmKeysClient with the specified values.
//   - subscriptionID - Subscription credentials which uniquely identify Microsoft Azure subscription. The subscription ID forms
//     part of the URI for every service call.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewManagedHsmKeysClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*ManagedHsmKeysClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &ManagedHsmKeysClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// CreateIfNotExist - Creates the first version of a new key if it does not exist. If it already exists, then the existing
// key is returned without any write operations being performed. This API does not create subsequent
// versions, and does not update existing keys.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-11-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - name - The name of the Managed HSM Pool within the specified resource group.
//   - keyName - The name of the key to be created. The value you provide may be copied globally for the purpose of running the
//     service. The value provided should not include personally identifiable or sensitive
//     information.
//   - parameters - The parameters used to create the specified key.
//   - options - ManagedHsmKeysClientCreateIfNotExistOptions contains the optional parameters for the ManagedHsmKeysClient.CreateIfNotExist
//     method.
func (client *ManagedHsmKeysClient) CreateIfNotExist(ctx context.Context, resourceGroupName string, name string, keyName string, parameters ManagedHsmKeyCreateParameters, options *ManagedHsmKeysClientCreateIfNotExistOptions) (ManagedHsmKeysClientCreateIfNotExistResponse, error) {
	var err error
	const operationName = "ManagedHsmKeysClient.CreateIfNotExist"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.createIfNotExistCreateRequest(ctx, resourceGroupName, name, keyName, parameters, options)
	if err != nil {
		return ManagedHsmKeysClientCreateIfNotExistResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ManagedHsmKeysClientCreateIfNotExistResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return ManagedHsmKeysClientCreateIfNotExistResponse{}, err
	}
	resp, err := client.createIfNotExistHandleResponse(httpResp)
	return resp, err
}

// createIfNotExistCreateRequest creates the CreateIfNotExist request.
func (client *ManagedHsmKeysClient) createIfNotExistCreateRequest(ctx context.Context, resourceGroupName string, name string, keyName string, parameters ManagedHsmKeyCreateParameters, _ *ManagedHsmKeysClientCreateIfNotExistOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.KeyVault/managedHSMs/{name}/keys/{keyName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if name == "" {
		return nil, errors.New("parameter name cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{name}", url.PathEscape(name))
	if keyName == "" {
		return nil, errors.New("parameter keyName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{keyName}", url.PathEscape(keyName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-11-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, parameters); err != nil {
		return nil, err
	}
	return req, nil
}

// createIfNotExistHandleResponse handles the CreateIfNotExist response.
func (client *ManagedHsmKeysClient) createIfNotExistHandleResponse(resp *http.Response) (ManagedHsmKeysClientCreateIfNotExistResponse, error) {
	result := ManagedHsmKeysClientCreateIfNotExistResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ManagedHsmKey); err != nil {
		return ManagedHsmKeysClientCreateIfNotExistResponse{}, err
	}
	return result, nil
}

// Get - Gets the current version of the specified key from the specified managed HSM.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-11-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - name - The name of the Managed HSM Pool within the specified resource group.
//   - keyName - The name of the key to be created. The value you provide may be copied globally for the purpose of running the
//     service. The value provided should not include personally identifiable or sensitive
//     information.
//   - options - ManagedHsmKeysClientGetOptions contains the optional parameters for the ManagedHsmKeysClient.Get method.
func (client *ManagedHsmKeysClient) Get(ctx context.Context, resourceGroupName string, name string, keyName string, options *ManagedHsmKeysClientGetOptions) (ManagedHsmKeysClientGetResponse, error) {
	var err error
	const operationName = "ManagedHsmKeysClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, resourceGroupName, name, keyName, options)
	if err != nil {
		return ManagedHsmKeysClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ManagedHsmKeysClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return ManagedHsmKeysClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *ManagedHsmKeysClient) getCreateRequest(ctx context.Context, resourceGroupName string, name string, keyName string, _ *ManagedHsmKeysClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.KeyVault/managedHSMs/{name}/keys/{keyName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if name == "" {
		return nil, errors.New("parameter name cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{name}", url.PathEscape(name))
	if keyName == "" {
		return nil, errors.New("parameter keyName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{keyName}", url.PathEscape(keyName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-11-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *ManagedHsmKeysClient) getHandleResponse(resp *http.Response) (ManagedHsmKeysClientGetResponse, error) {
	result := ManagedHsmKeysClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ManagedHsmKey); err != nil {
		return ManagedHsmKeysClientGetResponse{}, err
	}
	return result, nil
}

// GetVersion - Gets the specified version of the specified key in the specified managed HSM.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-11-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - name - The name of the Managed HSM Pool within the specified resource group.
//   - keyName - The name of the key to be created. The value you provide may be copied globally for the purpose of running the
//     service. The value provided should not include personally identifiable or sensitive
//     information.
//   - keyVersion - The version of the key to be retrieved.
//   - options - ManagedHsmKeysClientGetVersionOptions contains the optional parameters for the ManagedHsmKeysClient.GetVersion
//     method.
func (client *ManagedHsmKeysClient) GetVersion(ctx context.Context, resourceGroupName string, name string, keyName string, keyVersion string, options *ManagedHsmKeysClientGetVersionOptions) (ManagedHsmKeysClientGetVersionResponse, error) {
	var err error
	const operationName = "ManagedHsmKeysClient.GetVersion"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getVersionCreateRequest(ctx, resourceGroupName, name, keyName, keyVersion, options)
	if err != nil {
		return ManagedHsmKeysClientGetVersionResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ManagedHsmKeysClientGetVersionResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return ManagedHsmKeysClientGetVersionResponse{}, err
	}
	resp, err := client.getVersionHandleResponse(httpResp)
	return resp, err
}

// getVersionCreateRequest creates the GetVersion request.
func (client *ManagedHsmKeysClient) getVersionCreateRequest(ctx context.Context, resourceGroupName string, name string, keyName string, keyVersion string, _ *ManagedHsmKeysClientGetVersionOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.KeyVault/managedHSMs/{name}/keys/{keyName}/versions/{keyVersion}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if name == "" {
		return nil, errors.New("parameter name cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{name}", url.PathEscape(name))
	if keyName == "" {
		return nil, errors.New("parameter keyName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{keyName}", url.PathEscape(keyName))
	if keyVersion == "" {
		return nil, errors.New("parameter keyVersion cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{keyVersion}", url.PathEscape(keyVersion))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-11-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getVersionHandleResponse handles the GetVersion response.
func (client *ManagedHsmKeysClient) getVersionHandleResponse(resp *http.Response) (ManagedHsmKeysClientGetVersionResponse, error) {
	result := ManagedHsmKeysClientGetVersionResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ManagedHsmKey); err != nil {
		return ManagedHsmKeysClientGetVersionResponse{}, err
	}
	return result, nil
}

// NewListPager - Lists the keys in the specified managed HSM.
//
// Generated from API version 2024-11-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - name - The name of the Managed HSM Pool within the specified resource group.
//   - options - ManagedHsmKeysClientListOptions contains the optional parameters for the ManagedHsmKeysClient.NewListPager method.
func (client *ManagedHsmKeysClient) NewListPager(resourceGroupName string, name string, options *ManagedHsmKeysClientListOptions) *runtime.Pager[ManagedHsmKeysClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[ManagedHsmKeysClientListResponse]{
		More: func(page ManagedHsmKeysClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *ManagedHsmKeysClientListResponse) (ManagedHsmKeysClientListResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "ManagedHsmKeysClient.NewListPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listCreateRequest(ctx, resourceGroupName, name, options)
			}, nil)
			if err != nil {
				return ManagedHsmKeysClientListResponse{}, err
			}
			return client.listHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listCreateRequest creates the List request.
func (client *ManagedHsmKeysClient) listCreateRequest(ctx context.Context, resourceGroupName string, name string, _ *ManagedHsmKeysClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.KeyVault/managedHSMs/{name}/keys"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if name == "" {
		return nil, errors.New("parameter name cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{name}", url.PathEscape(name))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-11-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *ManagedHsmKeysClient) listHandleResponse(resp *http.Response) (ManagedHsmKeysClientListResponse, error) {
	result := ManagedHsmKeysClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ManagedHsmKeyListResult); err != nil {
		return ManagedHsmKeysClientListResponse{}, err
	}
	return result, nil
}

// NewListVersionsPager - Lists the versions of the specified key in the specified managed HSM.
//
// Generated from API version 2024-11-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - name - The name of the Managed HSM Pool within the specified resource group.
//   - keyName - The name of the key to be created. The value you provide may be copied globally for the purpose of running the
//     service. The value provided should not include personally identifiable or sensitive
//     information.
//   - options - ManagedHsmKeysClientListVersionsOptions contains the optional parameters for the ManagedHsmKeysClient.NewListVersionsPager
//     method.
func (client *ManagedHsmKeysClient) NewListVersionsPager(resourceGroupName string, name string, keyName string, options *ManagedHsmKeysClientListVersionsOptions) *runtime.Pager[ManagedHsmKeysClientListVersionsResponse] {
	return runtime.NewPager(runtime.PagingHandler[ManagedHsmKeysClientListVersionsResponse]{
		More: func(page ManagedHsmKeysClientListVersionsResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *ManagedHsmKeysClientListVersionsResponse) (ManagedHsmKeysClientListVersionsResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "ManagedHsmKeysClient.NewListVersionsPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listVersionsCreateRequest(ctx, resourceGroupName, name, keyName, options)
			}, nil)
			if err != nil {
				return ManagedHsmKeysClientListVersionsResponse{}, err
			}
			return client.listVersionsHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listVersionsCreateRequest creates the ListVersions request.
func (client *ManagedHsmKeysClient) listVersionsCreateRequest(ctx context.Context, resourceGroupName string, name string, keyName string, _ *ManagedHsmKeysClientListVersionsOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.KeyVault/managedHSMs/{name}/keys/{keyName}/versions"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if name == "" {
		return nil, errors.New("parameter name cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{name}", url.PathEscape(name))
	if keyName == "" {
		return nil, errors.New("parameter keyName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{keyName}", url.PathEscape(keyName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-11-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listVersionsHandleResponse handles the ListVersions response.
func (client *ManagedHsmKeysClient) listVersionsHandleResponse(resp *http.Response) (ManagedHsmKeysClientListVersionsResponse, error) {
	result := ManagedHsmKeysClientListVersionsResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ManagedHsmKeyListResult); err != nil {
		return ManagedHsmKeysClientListVersionsResponse{}, err
	}
	return result, nil
}
