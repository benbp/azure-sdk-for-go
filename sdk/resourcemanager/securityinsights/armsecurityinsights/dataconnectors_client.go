//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armsecurityinsights

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

// DataConnectorsClient contains the methods for the DataConnectors group.
// Don't use this type directly, use NewDataConnectorsClient() instead.
type DataConnectorsClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewDataConnectorsClient creates a new instance of DataConnectorsClient with the specified values.
//   - subscriptionID - The ID of the target subscription.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewDataConnectorsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*DataConnectorsClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &DataConnectorsClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// Connect - Connects a data connector.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-09-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - workspaceName - The name of the workspace.
//   - dataConnectorID - Connector ID
//   - connectBody - The data connector
//   - options - DataConnectorsClientConnectOptions contains the optional parameters for the DataConnectorsClient.Connect method.
func (client *DataConnectorsClient) Connect(ctx context.Context, resourceGroupName string, workspaceName string, dataConnectorID string, connectBody DataConnectorConnectBody, options *DataConnectorsClientConnectOptions) (DataConnectorsClientConnectResponse, error) {
	var err error
	const operationName = "DataConnectorsClient.Connect"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.connectCreateRequest(ctx, resourceGroupName, workspaceName, dataConnectorID, connectBody, options)
	if err != nil {
		return DataConnectorsClientConnectResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return DataConnectorsClientConnectResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return DataConnectorsClientConnectResponse{}, err
	}
	return DataConnectorsClientConnectResponse{}, nil
}

// connectCreateRequest creates the Connect request.
func (client *DataConnectorsClient) connectCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, dataConnectorID string, connectBody DataConnectorConnectBody, options *DataConnectorsClientConnectOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.OperationalInsights/workspaces/{workspaceName}/providers/Microsoft.SecurityInsights/dataConnectors/{dataConnectorId}/connect"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if workspaceName == "" {
		return nil, errors.New("parameter workspaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceName}", url.PathEscape(workspaceName))
	if dataConnectorID == "" {
		return nil, errors.New("parameter dataConnectorID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{dataConnectorId}", url.PathEscape(dataConnectorID))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-09-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, connectBody); err != nil {
		return nil, err
	}
	return req, nil
}

// CreateOrUpdate - Creates or updates the data connector.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-09-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - workspaceName - The name of the workspace.
//   - dataConnectorID - Connector ID
//   - dataConnector - The data connector
//   - options - DataConnectorsClientCreateOrUpdateOptions contains the optional parameters for the DataConnectorsClient.CreateOrUpdate
//     method.
func (client *DataConnectorsClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, workspaceName string, dataConnectorID string, dataConnector DataConnectorClassification, options *DataConnectorsClientCreateOrUpdateOptions) (DataConnectorsClientCreateOrUpdateResponse, error) {
	var err error
	const operationName = "DataConnectorsClient.CreateOrUpdate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, workspaceName, dataConnectorID, dataConnector, options)
	if err != nil {
		return DataConnectorsClientCreateOrUpdateResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return DataConnectorsClientCreateOrUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusCreated) {
		err = runtime.NewResponseError(httpResp)
		return DataConnectorsClientCreateOrUpdateResponse{}, err
	}
	resp, err := client.createOrUpdateHandleResponse(httpResp)
	return resp, err
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *DataConnectorsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, dataConnectorID string, dataConnector DataConnectorClassification, options *DataConnectorsClientCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.OperationalInsights/workspaces/{workspaceName}/providers/Microsoft.SecurityInsights/dataConnectors/{dataConnectorId}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if workspaceName == "" {
		return nil, errors.New("parameter workspaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceName}", url.PathEscape(workspaceName))
	if dataConnectorID == "" {
		return nil, errors.New("parameter dataConnectorID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{dataConnectorId}", url.PathEscape(dataConnectorID))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-09-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, dataConnector); err != nil {
		return nil, err
	}
	return req, nil
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *DataConnectorsClient) createOrUpdateHandleResponse(resp *http.Response) (DataConnectorsClientCreateOrUpdateResponse, error) {
	result := DataConnectorsClientCreateOrUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result); err != nil {
		return DataConnectorsClientCreateOrUpdateResponse{}, err
	}
	return result, nil
}

// Delete - Delete the data connector.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-09-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - workspaceName - The name of the workspace.
//   - dataConnectorID - Connector ID
//   - options - DataConnectorsClientDeleteOptions contains the optional parameters for the DataConnectorsClient.Delete method.
func (client *DataConnectorsClient) Delete(ctx context.Context, resourceGroupName string, workspaceName string, dataConnectorID string, options *DataConnectorsClientDeleteOptions) (DataConnectorsClientDeleteResponse, error) {
	var err error
	const operationName = "DataConnectorsClient.Delete"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, workspaceName, dataConnectorID, options)
	if err != nil {
		return DataConnectorsClientDeleteResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return DataConnectorsClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return DataConnectorsClientDeleteResponse{}, err
	}
	return DataConnectorsClientDeleteResponse{}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *DataConnectorsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, dataConnectorID string, options *DataConnectorsClientDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.OperationalInsights/workspaces/{workspaceName}/providers/Microsoft.SecurityInsights/dataConnectors/{dataConnectorId}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if workspaceName == "" {
		return nil, errors.New("parameter workspaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceName}", url.PathEscape(workspaceName))
	if dataConnectorID == "" {
		return nil, errors.New("parameter dataConnectorID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{dataConnectorId}", url.PathEscape(dataConnectorID))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-09-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Disconnect - Disconnect a data connector.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-09-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - workspaceName - The name of the workspace.
//   - dataConnectorID - Connector ID
//   - options - DataConnectorsClientDisconnectOptions contains the optional parameters for the DataConnectorsClient.Disconnect
//     method.
func (client *DataConnectorsClient) Disconnect(ctx context.Context, resourceGroupName string, workspaceName string, dataConnectorID string, options *DataConnectorsClientDisconnectOptions) (DataConnectorsClientDisconnectResponse, error) {
	var err error
	const operationName = "DataConnectorsClient.Disconnect"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.disconnectCreateRequest(ctx, resourceGroupName, workspaceName, dataConnectorID, options)
	if err != nil {
		return DataConnectorsClientDisconnectResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return DataConnectorsClientDisconnectResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return DataConnectorsClientDisconnectResponse{}, err
	}
	return DataConnectorsClientDisconnectResponse{}, nil
}

// disconnectCreateRequest creates the Disconnect request.
func (client *DataConnectorsClient) disconnectCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, dataConnectorID string, options *DataConnectorsClientDisconnectOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.OperationalInsights/workspaces/{workspaceName}/providers/Microsoft.SecurityInsights/dataConnectors/{dataConnectorId}/disconnect"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if workspaceName == "" {
		return nil, errors.New("parameter workspaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceName}", url.PathEscape(workspaceName))
	if dataConnectorID == "" {
		return nil, errors.New("parameter dataConnectorID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{dataConnectorId}", url.PathEscape(dataConnectorID))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-09-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Gets a data connector.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-09-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - workspaceName - The name of the workspace.
//   - dataConnectorID - Connector ID
//   - options - DataConnectorsClientGetOptions contains the optional parameters for the DataConnectorsClient.Get method.
func (client *DataConnectorsClient) Get(ctx context.Context, resourceGroupName string, workspaceName string, dataConnectorID string, options *DataConnectorsClientGetOptions) (DataConnectorsClientGetResponse, error) {
	var err error
	const operationName = "DataConnectorsClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, resourceGroupName, workspaceName, dataConnectorID, options)
	if err != nil {
		return DataConnectorsClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return DataConnectorsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return DataConnectorsClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *DataConnectorsClient) getCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, dataConnectorID string, options *DataConnectorsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.OperationalInsights/workspaces/{workspaceName}/providers/Microsoft.SecurityInsights/dataConnectors/{dataConnectorId}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if workspaceName == "" {
		return nil, errors.New("parameter workspaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceName}", url.PathEscape(workspaceName))
	if dataConnectorID == "" {
		return nil, errors.New("parameter dataConnectorID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{dataConnectorId}", url.PathEscape(dataConnectorID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-09-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *DataConnectorsClient) getHandleResponse(resp *http.Response) (DataConnectorsClientGetResponse, error) {
	result := DataConnectorsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result); err != nil {
		return DataConnectorsClientGetResponse{}, err
	}
	return result, nil
}

// NewListPager - Gets all data connectors.
//
// Generated from API version 2022-09-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - workspaceName - The name of the workspace.
//   - options - DataConnectorsClientListOptions contains the optional parameters for the DataConnectorsClient.NewListPager method.
func (client *DataConnectorsClient) NewListPager(resourceGroupName string, workspaceName string, options *DataConnectorsClientListOptions) *runtime.Pager[DataConnectorsClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[DataConnectorsClientListResponse]{
		More: func(page DataConnectorsClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *DataConnectorsClientListResponse) (DataConnectorsClientListResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "DataConnectorsClient.NewListPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listCreateRequest(ctx, resourceGroupName, workspaceName, options)
			}, nil)
			if err != nil {
				return DataConnectorsClientListResponse{}, err
			}
			return client.listHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listCreateRequest creates the List request.
func (client *DataConnectorsClient) listCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, options *DataConnectorsClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.OperationalInsights/workspaces/{workspaceName}/providers/Microsoft.SecurityInsights/dataConnectors"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if workspaceName == "" {
		return nil, errors.New("parameter workspaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceName}", url.PathEscape(workspaceName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-09-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *DataConnectorsClient) listHandleResponse(resp *http.Response) (DataConnectorsClientListResponse, error) {
	result := DataConnectorsClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.DataConnectorList); err != nil {
		return DataConnectorsClientListResponse{}, err
	}
	return result, nil
}
