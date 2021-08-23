// +build go1.13

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armcontainerregistry

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/armcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"net/http"
	"net/url"
	"strings"
	"time"
)

// PipelineRunsClient contains the methods for the PipelineRuns group.
// Don't use this type directly, use NewPipelineRunsClient() instead.
type PipelineRunsClient struct {
	con            *armcore.Connection
	subscriptionID string
}

// NewPipelineRunsClient creates a new instance of PipelineRunsClient with the specified values.
func NewPipelineRunsClient(con *armcore.Connection, subscriptionID string) *PipelineRunsClient {
	return &PipelineRunsClient{con: con, subscriptionID: subscriptionID}
}

// BeginCreate - Creates a pipeline run for a container registry with the specified parameters
// If the operation fails it returns a generic error.
func (client *PipelineRunsClient) BeginCreate(ctx context.Context, resourceGroupName string, registryName string, pipelineRunName string, pipelineRunCreateParameters PipelineRun, options *PipelineRunsBeginCreateOptions) (PipelineRunPollerResponse, error) {
	resp, err := client.create(ctx, resourceGroupName, registryName, pipelineRunName, pipelineRunCreateParameters, options)
	if err != nil {
		return PipelineRunPollerResponse{}, err
	}
	result := PipelineRunPollerResponse{
		RawResponse: resp.Response,
	}
	pt, err := armcore.NewLROPoller("PipelineRunsClient.Create", "", resp, client.con.Pipeline(), client.createHandleError)
	if err != nil {
		return PipelineRunPollerResponse{}, err
	}
	poller := &pipelineRunPoller{
		pt: pt,
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (PipelineRunResponseType, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

// ResumeCreate creates a new PipelineRunPoller from the specified resume token.
// token - The value must come from a previous call to PipelineRunPoller.ResumeToken().
func (client *PipelineRunsClient) ResumeCreate(ctx context.Context, token string) (PipelineRunPollerResponse, error) {
	pt, err := armcore.NewLROPollerFromResumeToken("PipelineRunsClient.Create", token, client.con.Pipeline(), client.createHandleError)
	if err != nil {
		return PipelineRunPollerResponse{}, err
	}
	poller := &pipelineRunPoller{
		pt: pt,
	}
	resp, err := poller.Poll(ctx)
	if err != nil {
		return PipelineRunPollerResponse{}, err
	}
	result := PipelineRunPollerResponse{
		RawResponse: resp,
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (PipelineRunResponseType, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

// Create - Creates a pipeline run for a container registry with the specified parameters
// If the operation fails it returns a generic error.
func (client *PipelineRunsClient) create(ctx context.Context, resourceGroupName string, registryName string, pipelineRunName string, pipelineRunCreateParameters PipelineRun, options *PipelineRunsBeginCreateOptions) (*azcore.Response, error) {
	req, err := client.createCreateRequest(ctx, resourceGroupName, registryName, pipelineRunName, pipelineRunCreateParameters, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK, http.StatusCreated) {
		return nil, client.createHandleError(resp)
	}
	return resp, nil
}

// createCreateRequest creates the Create request.
func (client *PipelineRunsClient) createCreateRequest(ctx context.Context, resourceGroupName string, registryName string, pipelineRunName string, pipelineRunCreateParameters PipelineRun, options *PipelineRunsBeginCreateOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerRegistry/registries/{registryName}/pipelineRuns/{pipelineRunName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if registryName == "" {
		return nil, errors.New("parameter registryName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{registryName}", url.PathEscape(registryName))
	if pipelineRunName == "" {
		return nil, errors.New("parameter pipelineRunName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{pipelineRunName}", url.PathEscape(pipelineRunName))
	req, err := azcore.NewRequest(ctx, http.MethodPut, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	reqQP.Set("api-version", "2021-06-01-preview")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, req.MarshalAsJSON(pipelineRunCreateParameters)
}

// createHandleError handles the Create error response.
func (client *PipelineRunsClient) createHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
	if len(body) == 0 {
		return azcore.NewResponseError(errors.New(resp.Status), resp.Response)
	}
	return azcore.NewResponseError(errors.New(string(body)), resp.Response)
}

// BeginDelete - Deletes a pipeline run from a container registry.
// If the operation fails it returns a generic error.
func (client *PipelineRunsClient) BeginDelete(ctx context.Context, resourceGroupName string, registryName string, pipelineRunName string, options *PipelineRunsBeginDeleteOptions) (HTTPPollerResponse, error) {
	resp, err := client.deleteOperation(ctx, resourceGroupName, registryName, pipelineRunName, options)
	if err != nil {
		return HTTPPollerResponse{}, err
	}
	result := HTTPPollerResponse{
		RawResponse: resp.Response,
	}
	pt, err := armcore.NewLROPoller("PipelineRunsClient.Delete", "", resp, client.con.Pipeline(), client.deleteHandleError)
	if err != nil {
		return HTTPPollerResponse{}, err
	}
	poller := &httpPoller{
		pt: pt,
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (*http.Response, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

// ResumeDelete creates a new HTTPPoller from the specified resume token.
// token - The value must come from a previous call to HTTPPoller.ResumeToken().
func (client *PipelineRunsClient) ResumeDelete(ctx context.Context, token string) (HTTPPollerResponse, error) {
	pt, err := armcore.NewLROPollerFromResumeToken("PipelineRunsClient.Delete", token, client.con.Pipeline(), client.deleteHandleError)
	if err != nil {
		return HTTPPollerResponse{}, err
	}
	poller := &httpPoller{
		pt: pt,
	}
	resp, err := poller.Poll(ctx)
	if err != nil {
		return HTTPPollerResponse{}, err
	}
	result := HTTPPollerResponse{
		RawResponse: resp,
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (*http.Response, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

// Delete - Deletes a pipeline run from a container registry.
// If the operation fails it returns a generic error.
func (client *PipelineRunsClient) deleteOperation(ctx context.Context, resourceGroupName string, registryName string, pipelineRunName string, options *PipelineRunsBeginDeleteOptions) (*azcore.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, registryName, pipelineRunName, options)
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
func (client *PipelineRunsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, registryName string, pipelineRunName string, options *PipelineRunsBeginDeleteOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerRegistry/registries/{registryName}/pipelineRuns/{pipelineRunName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if registryName == "" {
		return nil, errors.New("parameter registryName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{registryName}", url.PathEscape(registryName))
	if pipelineRunName == "" {
		return nil, errors.New("parameter pipelineRunName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{pipelineRunName}", url.PathEscape(pipelineRunName))
	req, err := azcore.NewRequest(ctx, http.MethodDelete, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	reqQP.Set("api-version", "2021-06-01-preview")
	req.URL.RawQuery = reqQP.Encode()
	return req, nil
}

// deleteHandleError handles the Delete error response.
func (client *PipelineRunsClient) deleteHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
	if len(body) == 0 {
		return azcore.NewResponseError(errors.New(resp.Status), resp.Response)
	}
	return azcore.NewResponseError(errors.New(string(body)), resp.Response)
}

// Get - Gets the detailed information for a given pipeline run.
// If the operation fails it returns a generic error.
func (client *PipelineRunsClient) Get(ctx context.Context, resourceGroupName string, registryName string, pipelineRunName string, options *PipelineRunsGetOptions) (PipelineRunResponseType, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, registryName, pipelineRunName, options)
	if err != nil {
		return PipelineRunResponseType{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return PipelineRunResponseType{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return PipelineRunResponseType{}, client.getHandleError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *PipelineRunsClient) getCreateRequest(ctx context.Context, resourceGroupName string, registryName string, pipelineRunName string, options *PipelineRunsGetOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerRegistry/registries/{registryName}/pipelineRuns/{pipelineRunName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if registryName == "" {
		return nil, errors.New("parameter registryName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{registryName}", url.PathEscape(registryName))
	if pipelineRunName == "" {
		return nil, errors.New("parameter pipelineRunName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{pipelineRunName}", url.PathEscape(pipelineRunName))
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	reqQP.Set("api-version", "2021-06-01-preview")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *PipelineRunsClient) getHandleResponse(resp *azcore.Response) (PipelineRunResponseType, error) {
	var val *PipelineRun
	if err := resp.UnmarshalAsJSON(&val); err != nil {
		return PipelineRunResponseType{}, err
	}
	return PipelineRunResponseType{RawResponse: resp.Response, PipelineRun: val}, nil
}

// getHandleError handles the Get error response.
func (client *PipelineRunsClient) getHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
	if len(body) == 0 {
		return azcore.NewResponseError(errors.New(resp.Status), resp.Response)
	}
	return azcore.NewResponseError(errors.New(string(body)), resp.Response)
}

// List - Lists all the pipeline runs for the specified container registry.
// If the operation fails it returns a generic error.
func (client *PipelineRunsClient) List(resourceGroupName string, registryName string, options *PipelineRunsListOptions) PipelineRunListResultPager {
	return &pipelineRunListResultPager{
		pipeline: client.con.Pipeline(),
		requester: func(ctx context.Context) (*azcore.Request, error) {
			return client.listCreateRequest(ctx, resourceGroupName, registryName, options)
		},
		responder: client.listHandleResponse,
		errorer:   client.listHandleError,
		advancer: func(ctx context.Context, resp PipelineRunListResultResponse) (*azcore.Request, error) {
			return azcore.NewRequest(ctx, http.MethodGet, *resp.PipelineRunListResult.NextLink)
		},
		statusCodes: []int{http.StatusOK},
	}
}

// listCreateRequest creates the List request.
func (client *PipelineRunsClient) listCreateRequest(ctx context.Context, resourceGroupName string, registryName string, options *PipelineRunsListOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerRegistry/registries/{registryName}/pipelineRuns"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if registryName == "" {
		return nil, errors.New("parameter registryName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{registryName}", url.PathEscape(registryName))
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	reqQP.Set("api-version", "2021-06-01-preview")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client *PipelineRunsClient) listHandleResponse(resp *azcore.Response) (PipelineRunListResultResponse, error) {
	var val *PipelineRunListResult
	if err := resp.UnmarshalAsJSON(&val); err != nil {
		return PipelineRunListResultResponse{}, err
	}
	return PipelineRunListResultResponse{RawResponse: resp.Response, PipelineRunListResult: val}, nil
}

// listHandleError handles the List error response.
func (client *PipelineRunsClient) listHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
	if len(body) == 0 {
		return azcore.NewResponseError(errors.New(resp.Status), resp.Response)
	}
	return azcore.NewResponseError(errors.New(string(body)), resp.Response)
}
