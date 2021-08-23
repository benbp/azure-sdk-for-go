// +build go1.13

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armkeyvault

import (
	"context"
	"errors"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/armcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

// MHSMPrivateEndpointConnectionsClient contains the methods for the MHSMPrivateEndpointConnections group.
// Don't use this type directly, use NewMHSMPrivateEndpointConnectionsClient() instead.
type MHSMPrivateEndpointConnectionsClient struct {
	con            *armcore.Connection
	subscriptionID string
}

// NewMHSMPrivateEndpointConnectionsClient creates a new instance of MHSMPrivateEndpointConnectionsClient with the specified values.
func NewMHSMPrivateEndpointConnectionsClient(con *armcore.Connection, subscriptionID string) *MHSMPrivateEndpointConnectionsClient {
	return &MHSMPrivateEndpointConnectionsClient{con: con, subscriptionID: subscriptionID}
}

// BeginDelete - Deletes the specified private endpoint connection associated with the managed hsm pool.
// If the operation fails it returns the *CloudError error type.
func (client *MHSMPrivateEndpointConnectionsClient) BeginDelete(ctx context.Context, resourceGroupName string, name string, privateEndpointConnectionName string, options *MHSMPrivateEndpointConnectionsBeginDeleteOptions) (MHSMPrivateEndpointConnectionPollerResponse, error) {
	resp, err := client.deleteOperation(ctx, resourceGroupName, name, privateEndpointConnectionName, options)
	if err != nil {
		return MHSMPrivateEndpointConnectionPollerResponse{}, err
	}
	result := MHSMPrivateEndpointConnectionPollerResponse{
		RawResponse: resp.Response,
	}
	pt, err := armcore.NewLROPoller("MHSMPrivateEndpointConnectionsClient.Delete", "", resp, client.con.Pipeline(), client.deleteHandleError)
	if err != nil {
		return MHSMPrivateEndpointConnectionPollerResponse{}, err
	}
	poller := &mhsmPrivateEndpointConnectionPoller{
		pt: pt,
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (MHSMPrivateEndpointConnectionResponse, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

// ResumeDelete creates a new MHSMPrivateEndpointConnectionPoller from the specified resume token.
// token - The value must come from a previous call to MHSMPrivateEndpointConnectionPoller.ResumeToken().
func (client *MHSMPrivateEndpointConnectionsClient) ResumeDelete(ctx context.Context, token string) (MHSMPrivateEndpointConnectionPollerResponse, error) {
	pt, err := armcore.NewLROPollerFromResumeToken("MHSMPrivateEndpointConnectionsClient.Delete", token, client.con.Pipeline(), client.deleteHandleError)
	if err != nil {
		return MHSMPrivateEndpointConnectionPollerResponse{}, err
	}
	poller := &mhsmPrivateEndpointConnectionPoller{
		pt: pt,
	}
	resp, err := poller.Poll(ctx)
	if err != nil {
		return MHSMPrivateEndpointConnectionPollerResponse{}, err
	}
	result := MHSMPrivateEndpointConnectionPollerResponse{
		RawResponse: resp,
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (MHSMPrivateEndpointConnectionResponse, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

// Delete - Deletes the specified private endpoint connection associated with the managed hsm pool.
// If the operation fails it returns the *CloudError error type.
func (client *MHSMPrivateEndpointConnectionsClient) deleteOperation(ctx context.Context, resourceGroupName string, name string, privateEndpointConnectionName string, options *MHSMPrivateEndpointConnectionsBeginDeleteOptions) (*azcore.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, name, privateEndpointConnectionName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent) {
		return nil, client.deleteHandleError(resp)
	}
	return resp, nil
}

// deleteCreateRequest creates the Delete request.
func (client *MHSMPrivateEndpointConnectionsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, name string, privateEndpointConnectionName string, options *MHSMPrivateEndpointConnectionsBeginDeleteOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.KeyVault/managedHSMs/{name}/privateEndpointConnections/{privateEndpointConnectionName}"
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
	if privateEndpointConnectionName == "" {
		return nil, errors.New("parameter privateEndpointConnectionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{privateEndpointConnectionName}", url.PathEscape(privateEndpointConnectionName))
	req, err := azcore.NewRequest(ctx, http.MethodDelete, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	reqQP.Set("api-version", "2021-04-01-preview")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// deleteHandleError handles the Delete error response.
func (client *MHSMPrivateEndpointConnectionsClient) deleteHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
	errType := CloudError{raw: string(body)}
	if err := resp.UnmarshalAsJSON(&errType); err != nil {
		return azcore.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp.Response)
	}
	return azcore.NewResponseError(&errType, resp.Response)
}

// Get - Gets the specified private endpoint connection associated with the managed HSM Pool.
// If the operation fails it returns the *ManagedHsmError error type.
func (client *MHSMPrivateEndpointConnectionsClient) Get(ctx context.Context, resourceGroupName string, name string, privateEndpointConnectionName string, options *MHSMPrivateEndpointConnectionsGetOptions) (MHSMPrivateEndpointConnectionResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, name, privateEndpointConnectionName, options)
	if err != nil {
		return MHSMPrivateEndpointConnectionResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return MHSMPrivateEndpointConnectionResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return MHSMPrivateEndpointConnectionResponse{}, client.getHandleError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *MHSMPrivateEndpointConnectionsClient) getCreateRequest(ctx context.Context, resourceGroupName string, name string, privateEndpointConnectionName string, options *MHSMPrivateEndpointConnectionsGetOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.KeyVault/managedHSMs/{name}/privateEndpointConnections/{privateEndpointConnectionName}"
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
	if privateEndpointConnectionName == "" {
		return nil, errors.New("parameter privateEndpointConnectionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{privateEndpointConnectionName}", url.PathEscape(privateEndpointConnectionName))
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	reqQP.Set("api-version", "2021-04-01-preview")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *MHSMPrivateEndpointConnectionsClient) getHandleResponse(resp *azcore.Response) (MHSMPrivateEndpointConnectionResponse, error) {
	var val *MHSMPrivateEndpointConnection
	if err := resp.UnmarshalAsJSON(&val); err != nil {
		return MHSMPrivateEndpointConnectionResponse{}, err
	}
	result := MHSMPrivateEndpointConnectionResponse{RawResponse: resp.Response, MHSMPrivateEndpointConnection: val}
	if val := resp.Header.Get("Retry-After"); val != "" {
		retryAfter32, err := strconv.ParseInt(val, 10, 32)
		retryAfter := int32(retryAfter32)
		if err != nil {
			return MHSMPrivateEndpointConnectionResponse{}, err
		}
		result.RetryAfter = &retryAfter
	}
	if val := resp.Header.Get("Azure-AsyncOperation"); val != "" {
		result.AzureAsyncOperation = &val
	}
	return result, nil
}

// getHandleError handles the Get error response.
func (client *MHSMPrivateEndpointConnectionsClient) getHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
	errType := ManagedHsmError{raw: string(body)}
	if err := resp.UnmarshalAsJSON(&errType); err != nil {
		return azcore.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp.Response)
	}
	return azcore.NewResponseError(&errType, resp.Response)
}

// ListByResource - The List operation gets information about the private endpoint connections associated with the managed HSM Pool.
// If the operation fails it returns the *ManagedHsmError error type.
func (client *MHSMPrivateEndpointConnectionsClient) ListByResource(resourceGroupName string, name string, options *MHSMPrivateEndpointConnectionsListByResourceOptions) MHSMPrivateEndpointConnectionsListResultPager {
	return &mhsmPrivateEndpointConnectionsListResultPager{
		pipeline: client.con.Pipeline(),
		requester: func(ctx context.Context) (*azcore.Request, error) {
			return client.listByResourceCreateRequest(ctx, resourceGroupName, name, options)
		},
		responder: client.listByResourceHandleResponse,
		errorer:   client.listByResourceHandleError,
		advancer: func(ctx context.Context, resp MHSMPrivateEndpointConnectionsListResultResponse) (*azcore.Request, error) {
			return azcore.NewRequest(ctx, http.MethodGet, *resp.MHSMPrivateEndpointConnectionsListResult.NextLink)
		},
		statusCodes: []int{http.StatusOK},
	}
}

// listByResourceCreateRequest creates the ListByResource request.
func (client *MHSMPrivateEndpointConnectionsClient) listByResourceCreateRequest(ctx context.Context, resourceGroupName string, name string, options *MHSMPrivateEndpointConnectionsListByResourceOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.KeyVault/managedHSMs/{name}/privateEndpointConnections"
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
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	reqQP.Set("api-version", "2021-04-01-preview")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// listByResourceHandleResponse handles the ListByResource response.
func (client *MHSMPrivateEndpointConnectionsClient) listByResourceHandleResponse(resp *azcore.Response) (MHSMPrivateEndpointConnectionsListResultResponse, error) {
	var val *MHSMPrivateEndpointConnectionsListResult
	if err := resp.UnmarshalAsJSON(&val); err != nil {
		return MHSMPrivateEndpointConnectionsListResultResponse{}, err
	}
	return MHSMPrivateEndpointConnectionsListResultResponse{RawResponse: resp.Response, MHSMPrivateEndpointConnectionsListResult: val}, nil
}

// listByResourceHandleError handles the ListByResource error response.
func (client *MHSMPrivateEndpointConnectionsClient) listByResourceHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
	errType := ManagedHsmError{raw: string(body)}
	if err := resp.UnmarshalAsJSON(&errType); err != nil {
		return azcore.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp.Response)
	}
	return azcore.NewResponseError(&errType, resp.Response)
}

// Put - Updates the specified private endpoint connection associated with the managed hsm pool.
// If the operation fails it returns the *CloudError error type.
func (client *MHSMPrivateEndpointConnectionsClient) Put(ctx context.Context, resourceGroupName string, name string, privateEndpointConnectionName string, properties MHSMPrivateEndpointConnection, options *MHSMPrivateEndpointConnectionsPutOptions) (MHSMPrivateEndpointConnectionResponse, error) {
	req, err := client.putCreateRequest(ctx, resourceGroupName, name, privateEndpointConnectionName, properties, options)
	if err != nil {
		return MHSMPrivateEndpointConnectionResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return MHSMPrivateEndpointConnectionResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return MHSMPrivateEndpointConnectionResponse{}, client.putHandleError(resp)
	}
	return client.putHandleResponse(resp)
}

// putCreateRequest creates the Put request.
func (client *MHSMPrivateEndpointConnectionsClient) putCreateRequest(ctx context.Context, resourceGroupName string, name string, privateEndpointConnectionName string, properties MHSMPrivateEndpointConnection, options *MHSMPrivateEndpointConnectionsPutOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.KeyVault/managedHSMs/{name}/privateEndpointConnections/{privateEndpointConnectionName}"
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
	if privateEndpointConnectionName == "" {
		return nil, errors.New("parameter privateEndpointConnectionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{privateEndpointConnectionName}", url.PathEscape(privateEndpointConnectionName))
	req, err := azcore.NewRequest(ctx, http.MethodPut, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	reqQP.Set("api-version", "2021-04-01-preview")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, req.MarshalAsJSON(properties)
}

// putHandleResponse handles the Put response.
func (client *MHSMPrivateEndpointConnectionsClient) putHandleResponse(resp *azcore.Response) (MHSMPrivateEndpointConnectionResponse, error) {
	var val *MHSMPrivateEndpointConnection
	if err := resp.UnmarshalAsJSON(&val); err != nil {
		return MHSMPrivateEndpointConnectionResponse{}, err
	}
	result := MHSMPrivateEndpointConnectionResponse{RawResponse: resp.Response, MHSMPrivateEndpointConnection: val}
	if val := resp.Header.Get("Retry-After"); val != "" {
		retryAfter32, err := strconv.ParseInt(val, 10, 32)
		retryAfter := int32(retryAfter32)
		if err != nil {
			return MHSMPrivateEndpointConnectionResponse{}, err
		}
		result.RetryAfter = &retryAfter
	}
	if val := resp.Header.Get("Azure-AsyncOperation"); val != "" {
		result.AzureAsyncOperation = &val
	}
	return result, nil
}

// putHandleError handles the Put error response.
func (client *MHSMPrivateEndpointConnectionsClient) putHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
	errType := CloudError{raw: string(body)}
	if err := resp.UnmarshalAsJSON(&errType); err != nil {
		return azcore.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp.Response)
	}
	return azcore.NewResponseError(&errType, resp.Response)
}
