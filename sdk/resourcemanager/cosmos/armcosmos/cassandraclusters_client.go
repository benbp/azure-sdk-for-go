//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armcosmos

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

// CassandraClustersClient contains the methods for the CassandraClusters group.
// Don't use this type directly, use NewCassandraClustersClient() instead.
type CassandraClustersClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewCassandraClustersClient creates a new instance of CassandraClustersClient with the specified values.
//   - subscriptionID - The ID of the target subscription.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewCassandraClustersClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*CassandraClustersClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &CassandraClustersClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// BeginCreateUpdate - Create or update a managed Cassandra cluster. When updating, you must specify all writable properties.
// To update only some properties, use PATCH.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-11-15
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - clusterName - Managed Cassandra cluster name.
//   - body - The properties specifying the desired state of the managed Cassandra cluster.
//   - options - CassandraClustersClientBeginCreateUpdateOptions contains the optional parameters for the CassandraClustersClient.BeginCreateUpdate
//     method.
func (client *CassandraClustersClient) BeginCreateUpdate(ctx context.Context, resourceGroupName string, clusterName string, body ClusterResource, options *CassandraClustersClientBeginCreateUpdateOptions) (*runtime.Poller[CassandraClustersClientCreateUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createUpdate(ctx, resourceGroupName, clusterName, body, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[CassandraClustersClientCreateUpdateResponse]{
			Tracer: client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[CassandraClustersClientCreateUpdateResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// CreateUpdate - Create or update a managed Cassandra cluster. When updating, you must specify all writable properties. To
// update only some properties, use PATCH.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-11-15
func (client *CassandraClustersClient) createUpdate(ctx context.Context, resourceGroupName string, clusterName string, body ClusterResource, options *CassandraClustersClientBeginCreateUpdateOptions) (*http.Response, error) {
	var err error
	const operationName = "CassandraClustersClient.BeginCreateUpdate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.createUpdateCreateRequest(ctx, resourceGroupName, clusterName, body, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusCreated) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// createUpdateCreateRequest creates the CreateUpdate request.
func (client *CassandraClustersClient) createUpdateCreateRequest(ctx context.Context, resourceGroupName string, clusterName string, body ClusterResource, options *CassandraClustersClientBeginCreateUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/cassandraClusters/{clusterName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if clusterName == "" {
		return nil, errors.New("parameter clusterName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{clusterName}", url.PathEscape(clusterName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-11-15")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, body); err != nil {
		return nil, err
	}
	return req, nil
}

// BeginDeallocate - Deallocate the Managed Cassandra Cluster and Associated Data Centers. Deallocation will deallocate the
// host virtual machine of this cluster, and reserved the data disk. This won't do anything on an
// already deallocated cluster. Use Start to restart the cluster.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-11-15
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - clusterName - Managed Cassandra cluster name.
//   - options - CassandraClustersClientBeginDeallocateOptions contains the optional parameters for the CassandraClustersClient.BeginDeallocate
//     method.
func (client *CassandraClustersClient) BeginDeallocate(ctx context.Context, resourceGroupName string, clusterName string, options *CassandraClustersClientBeginDeallocateOptions) (*runtime.Poller[CassandraClustersClientDeallocateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deallocate(ctx, resourceGroupName, clusterName, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[CassandraClustersClientDeallocateResponse]{
			Tracer: client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[CassandraClustersClientDeallocateResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// Deallocate - Deallocate the Managed Cassandra Cluster and Associated Data Centers. Deallocation will deallocate the host
// virtual machine of this cluster, and reserved the data disk. This won't do anything on an
// already deallocated cluster. Use Start to restart the cluster.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-11-15
func (client *CassandraClustersClient) deallocate(ctx context.Context, resourceGroupName string, clusterName string, options *CassandraClustersClientBeginDeallocateOptions) (*http.Response, error) {
	var err error
	const operationName = "CassandraClustersClient.BeginDeallocate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.deallocateCreateRequest(ctx, resourceGroupName, clusterName, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusAccepted) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// deallocateCreateRequest creates the Deallocate request.
func (client *CassandraClustersClient) deallocateCreateRequest(ctx context.Context, resourceGroupName string, clusterName string, options *CassandraClustersClientBeginDeallocateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/cassandraClusters/{clusterName}/deallocate"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if clusterName == "" {
		return nil, errors.New("parameter clusterName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{clusterName}", url.PathEscape(clusterName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-11-15")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// BeginDelete - Deletes a managed Cassandra cluster.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-11-15
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - clusterName - Managed Cassandra cluster name.
//   - options - CassandraClustersClientBeginDeleteOptions contains the optional parameters for the CassandraClustersClient.BeginDelete
//     method.
func (client *CassandraClustersClient) BeginDelete(ctx context.Context, resourceGroupName string, clusterName string, options *CassandraClustersClientBeginDeleteOptions) (*runtime.Poller[CassandraClustersClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, clusterName, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[CassandraClustersClientDeleteResponse]{
			Tracer: client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[CassandraClustersClientDeleteResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// Delete - Deletes a managed Cassandra cluster.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-11-15
func (client *CassandraClustersClient) deleteOperation(ctx context.Context, resourceGroupName string, clusterName string, options *CassandraClustersClientBeginDeleteOptions) (*http.Response, error) {
	var err error
	const operationName = "CassandraClustersClient.BeginDelete"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, clusterName, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusAccepted, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// deleteCreateRequest creates the Delete request.
func (client *CassandraClustersClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, clusterName string, options *CassandraClustersClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/cassandraClusters/{clusterName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if clusterName == "" {
		return nil, errors.New("parameter clusterName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{clusterName}", url.PathEscape(clusterName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-11-15")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Get the properties of a managed Cassandra cluster.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-11-15
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - clusterName - Managed Cassandra cluster name.
//   - options - CassandraClustersClientGetOptions contains the optional parameters for the CassandraClustersClient.Get method.
func (client *CassandraClustersClient) Get(ctx context.Context, resourceGroupName string, clusterName string, options *CassandraClustersClientGetOptions) (CassandraClustersClientGetResponse, error) {
	var err error
	const operationName = "CassandraClustersClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, resourceGroupName, clusterName, options)
	if err != nil {
		return CassandraClustersClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return CassandraClustersClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return CassandraClustersClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *CassandraClustersClient) getCreateRequest(ctx context.Context, resourceGroupName string, clusterName string, options *CassandraClustersClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/cassandraClusters/{clusterName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if clusterName == "" {
		return nil, errors.New("parameter clusterName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{clusterName}", url.PathEscape(clusterName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-11-15")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *CassandraClustersClient) getHandleResponse(resp *http.Response) (CassandraClustersClientGetResponse, error) {
	result := CassandraClustersClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ClusterResource); err != nil {
		return CassandraClustersClientGetResponse{}, err
	}
	return result, nil
}

// BeginInvokeCommand - Invoke a command like nodetool for cassandra maintenance
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-11-15
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - clusterName - Managed Cassandra cluster name.
//   - body - Specification which command to run where
//   - options - CassandraClustersClientBeginInvokeCommandOptions contains the optional parameters for the CassandraClustersClient.BeginInvokeCommand
//     method.
func (client *CassandraClustersClient) BeginInvokeCommand(ctx context.Context, resourceGroupName string, clusterName string, body CommandPostBody, options *CassandraClustersClientBeginInvokeCommandOptions) (*runtime.Poller[CassandraClustersClientInvokeCommandResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.invokeCommand(ctx, resourceGroupName, clusterName, body, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[CassandraClustersClientInvokeCommandResponse]{
			Tracer: client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[CassandraClustersClientInvokeCommandResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// InvokeCommand - Invoke a command like nodetool for cassandra maintenance
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-11-15
func (client *CassandraClustersClient) invokeCommand(ctx context.Context, resourceGroupName string, clusterName string, body CommandPostBody, options *CassandraClustersClientBeginInvokeCommandOptions) (*http.Response, error) {
	var err error
	const operationName = "CassandraClustersClient.BeginInvokeCommand"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.invokeCommandCreateRequest(ctx, resourceGroupName, clusterName, body, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusAccepted) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// invokeCommandCreateRequest creates the InvokeCommand request.
func (client *CassandraClustersClient) invokeCommandCreateRequest(ctx context.Context, resourceGroupName string, clusterName string, body CommandPostBody, options *CassandraClustersClientBeginInvokeCommandOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/cassandraClusters/{clusterName}/invokeCommand"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if clusterName == "" {
		return nil, errors.New("parameter clusterName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{clusterName}", url.PathEscape(clusterName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-11-15")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, body); err != nil {
		return nil, err
	}
	return req, nil
}

// NewListByResourceGroupPager - List all managed Cassandra clusters in this resource group.
//
// Generated from API version 2024-11-15
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - options - CassandraClustersClientListByResourceGroupOptions contains the optional parameters for the CassandraClustersClient.NewListByResourceGroupPager
//     method.
func (client *CassandraClustersClient) NewListByResourceGroupPager(resourceGroupName string, options *CassandraClustersClientListByResourceGroupOptions) *runtime.Pager[CassandraClustersClientListByResourceGroupResponse] {
	return runtime.NewPager(runtime.PagingHandler[CassandraClustersClientListByResourceGroupResponse]{
		More: func(page CassandraClustersClientListByResourceGroupResponse) bool {
			return false
		},
		Fetcher: func(ctx context.Context, page *CassandraClustersClientListByResourceGroupResponse) (CassandraClustersClientListByResourceGroupResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "CassandraClustersClient.NewListByResourceGroupPager")
			req, err := client.listByResourceGroupCreateRequest(ctx, resourceGroupName, options)
			if err != nil {
				return CassandraClustersClientListByResourceGroupResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return CassandraClustersClientListByResourceGroupResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return CassandraClustersClientListByResourceGroupResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByResourceGroupHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listByResourceGroupCreateRequest creates the ListByResourceGroup request.
func (client *CassandraClustersClient) listByResourceGroupCreateRequest(ctx context.Context, resourceGroupName string, options *CassandraClustersClientListByResourceGroupOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/cassandraClusters"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-11-15")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByResourceGroupHandleResponse handles the ListByResourceGroup response.
func (client *CassandraClustersClient) listByResourceGroupHandleResponse(resp *http.Response) (CassandraClustersClientListByResourceGroupResponse, error) {
	result := CassandraClustersClientListByResourceGroupResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ListClusters); err != nil {
		return CassandraClustersClientListByResourceGroupResponse{}, err
	}
	return result, nil
}

// NewListBySubscriptionPager - List all managed Cassandra clusters in this subscription.
//
// Generated from API version 2024-11-15
//   - options - CassandraClustersClientListBySubscriptionOptions contains the optional parameters for the CassandraClustersClient.NewListBySubscriptionPager
//     method.
func (client *CassandraClustersClient) NewListBySubscriptionPager(options *CassandraClustersClientListBySubscriptionOptions) *runtime.Pager[CassandraClustersClientListBySubscriptionResponse] {
	return runtime.NewPager(runtime.PagingHandler[CassandraClustersClientListBySubscriptionResponse]{
		More: func(page CassandraClustersClientListBySubscriptionResponse) bool {
			return false
		},
		Fetcher: func(ctx context.Context, page *CassandraClustersClientListBySubscriptionResponse) (CassandraClustersClientListBySubscriptionResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "CassandraClustersClient.NewListBySubscriptionPager")
			req, err := client.listBySubscriptionCreateRequest(ctx, options)
			if err != nil {
				return CassandraClustersClientListBySubscriptionResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return CassandraClustersClientListBySubscriptionResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return CassandraClustersClientListBySubscriptionResponse{}, runtime.NewResponseError(resp)
			}
			return client.listBySubscriptionHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listBySubscriptionCreateRequest creates the ListBySubscription request.
func (client *CassandraClustersClient) listBySubscriptionCreateRequest(ctx context.Context, options *CassandraClustersClientListBySubscriptionOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.DocumentDB/cassandraClusters"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-11-15")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listBySubscriptionHandleResponse handles the ListBySubscription response.
func (client *CassandraClustersClient) listBySubscriptionHandleResponse(resp *http.Response) (CassandraClustersClientListBySubscriptionResponse, error) {
	result := CassandraClustersClientListBySubscriptionResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ListClusters); err != nil {
		return CassandraClustersClientListBySubscriptionResponse{}, err
	}
	return result, nil
}

// BeginStart - Start the Managed Cassandra Cluster and Associated Data Centers. Start will start the host virtual machine
// of this cluster with reserved data disk. This won't do anything on an already running
// cluster. Use Deallocate to deallocate the cluster.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-11-15
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - clusterName - Managed Cassandra cluster name.
//   - options - CassandraClustersClientBeginStartOptions contains the optional parameters for the CassandraClustersClient.BeginStart
//     method.
func (client *CassandraClustersClient) BeginStart(ctx context.Context, resourceGroupName string, clusterName string, options *CassandraClustersClientBeginStartOptions) (*runtime.Poller[CassandraClustersClientStartResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.start(ctx, resourceGroupName, clusterName, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[CassandraClustersClientStartResponse]{
			Tracer: client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[CassandraClustersClientStartResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// Start - Start the Managed Cassandra Cluster and Associated Data Centers. Start will start the host virtual machine of this
// cluster with reserved data disk. This won't do anything on an already running
// cluster. Use Deallocate to deallocate the cluster.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-11-15
func (client *CassandraClustersClient) start(ctx context.Context, resourceGroupName string, clusterName string, options *CassandraClustersClientBeginStartOptions) (*http.Response, error) {
	var err error
	const operationName = "CassandraClustersClient.BeginStart"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.startCreateRequest(ctx, resourceGroupName, clusterName, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusAccepted) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// startCreateRequest creates the Start request.
func (client *CassandraClustersClient) startCreateRequest(ctx context.Context, resourceGroupName string, clusterName string, options *CassandraClustersClientBeginStartOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/cassandraClusters/{clusterName}/start"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if clusterName == "" {
		return nil, errors.New("parameter clusterName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{clusterName}", url.PathEscape(clusterName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-11-15")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Status - Gets the CPU, memory, and disk usage statistics for each Cassandra node in a cluster.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-11-15
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - clusterName - Managed Cassandra cluster name.
//   - options - CassandraClustersClientStatusOptions contains the optional parameters for the CassandraClustersClient.Status
//     method.
func (client *CassandraClustersClient) Status(ctx context.Context, resourceGroupName string, clusterName string, options *CassandraClustersClientStatusOptions) (CassandraClustersClientStatusResponse, error) {
	var err error
	const operationName = "CassandraClustersClient.Status"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.statusCreateRequest(ctx, resourceGroupName, clusterName, options)
	if err != nil {
		return CassandraClustersClientStatusResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return CassandraClustersClientStatusResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return CassandraClustersClientStatusResponse{}, err
	}
	resp, err := client.statusHandleResponse(httpResp)
	return resp, err
}

// statusCreateRequest creates the Status request.
func (client *CassandraClustersClient) statusCreateRequest(ctx context.Context, resourceGroupName string, clusterName string, options *CassandraClustersClientStatusOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/cassandraClusters/{clusterName}/status"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if clusterName == "" {
		return nil, errors.New("parameter clusterName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{clusterName}", url.PathEscape(clusterName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-11-15")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// statusHandleResponse handles the Status response.
func (client *CassandraClustersClient) statusHandleResponse(resp *http.Response) (CassandraClustersClientStatusResponse, error) {
	result := CassandraClustersClientStatusResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.CassandraClusterPublicStatus); err != nil {
		return CassandraClustersClientStatusResponse{}, err
	}
	return result, nil
}

// BeginUpdate - Updates some of the properties of a managed Cassandra cluster.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-11-15
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - clusterName - Managed Cassandra cluster name.
//   - body - Parameters to provide for specifying the managed Cassandra cluster.
//   - options - CassandraClustersClientBeginUpdateOptions contains the optional parameters for the CassandraClustersClient.BeginUpdate
//     method.
func (client *CassandraClustersClient) BeginUpdate(ctx context.Context, resourceGroupName string, clusterName string, body ClusterResource, options *CassandraClustersClientBeginUpdateOptions) (*runtime.Poller[CassandraClustersClientUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.update(ctx, resourceGroupName, clusterName, body, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[CassandraClustersClientUpdateResponse]{
			Tracer: client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[CassandraClustersClientUpdateResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// Update - Updates some of the properties of a managed Cassandra cluster.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-11-15
func (client *CassandraClustersClient) update(ctx context.Context, resourceGroupName string, clusterName string, body ClusterResource, options *CassandraClustersClientBeginUpdateOptions) (*http.Response, error) {
	var err error
	const operationName = "CassandraClustersClient.BeginUpdate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.updateCreateRequest(ctx, resourceGroupName, clusterName, body, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusAccepted) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// updateCreateRequest creates the Update request.
func (client *CassandraClustersClient) updateCreateRequest(ctx context.Context, resourceGroupName string, clusterName string, body ClusterResource, options *CassandraClustersClientBeginUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/cassandraClusters/{clusterName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if clusterName == "" {
		return nil, errors.New("parameter clusterName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{clusterName}", url.PathEscape(clusterName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-11-15")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, body); err != nil {
		return nil, err
	}
	return req, nil
}
