// +build go1.13

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armmonitor

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/armcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"net/http"
)

// OperationsClient contains the methods for the Operations group.
// Don't use this type directly, use NewOperationsClient() instead.
type OperationsClient struct {
	con *armcore.Connection
}

// NewOperationsClient creates a new instance of OperationsClient with the specified values.
func NewOperationsClient(con *armcore.Connection) *OperationsClient {
	return &OperationsClient{con: con}
}

// List - Lists all of the available operations from Microsoft.Insights provider.
// If the operation fails it returns a generic error.
func (client *OperationsClient) List(ctx context.Context, options *OperationsListOptions) (OperationListResultResponse, error) {
	req, err := client.listCreateRequest(ctx, options)
	if err != nil {
		return OperationListResultResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return OperationListResultResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return OperationListResultResponse{}, client.listHandleError(resp)
	}
	return client.listHandleResponse(resp)
}

// listCreateRequest creates the List request.
func (client *OperationsClient) listCreateRequest(ctx context.Context, options *OperationsListOptions) (*azcore.Request, error) {
	urlPath := "/providers/Microsoft.Insights/operations"
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	reqQP.Set("api-version", "2015-04-01")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client *OperationsClient) listHandleResponse(resp *azcore.Response) (OperationListResultResponse, error) {
	var val *OperationListResult
	if err := resp.UnmarshalAsJSON(&val); err != nil {
		return OperationListResultResponse{}, err
	}
	return OperationListResultResponse{RawResponse: resp.Response, OperationListResult: val}, nil
}

// listHandleError handles the List error response.
func (client *OperationsClient) listHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
	if len(body) == 0 {
		return azcore.NewResponseError(errors.New(resp.Status), resp.Response)
	}
	return azcore.NewResponseError(errors.New(string(body)), resp.Response)
}
