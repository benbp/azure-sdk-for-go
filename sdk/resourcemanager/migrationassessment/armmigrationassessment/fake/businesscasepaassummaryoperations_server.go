// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package fake

import (
	"context"
	"errors"
	"fmt"
	azfake "github.com/Azure/azure-sdk-for-go/sdk/azcore/fake"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/fake/server"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/migrationassessment/armmigrationassessment"
	"net/http"
	"net/url"
	"regexp"
)

// BusinessCasePaasSummaryOperationsServer is a fake server for instances of the armmigrationassessment.BusinessCasePaasSummaryOperationsClient type.
type BusinessCasePaasSummaryOperationsServer struct {
	// Get is the fake for method BusinessCasePaasSummaryOperationsClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, projectName string, businessCaseName string, paasSummaryName string, options *armmigrationassessment.BusinessCasePaasSummaryOperationsClientGetOptions) (resp azfake.Responder[armmigrationassessment.BusinessCasePaasSummaryOperationsClientGetResponse], errResp azfake.ErrorResponder)

	// NewListByBusinessCasePager is the fake for method BusinessCasePaasSummaryOperationsClient.NewListByBusinessCasePager
	// HTTP status codes to indicate success: http.StatusOK
	NewListByBusinessCasePager func(resourceGroupName string, projectName string, businessCaseName string, options *armmigrationassessment.BusinessCasePaasSummaryOperationsClientListByBusinessCaseOptions) (resp azfake.PagerResponder[armmigrationassessment.BusinessCasePaasSummaryOperationsClientListByBusinessCaseResponse])
}

// NewBusinessCasePaasSummaryOperationsServerTransport creates a new instance of BusinessCasePaasSummaryOperationsServerTransport with the provided implementation.
// The returned BusinessCasePaasSummaryOperationsServerTransport instance is connected to an instance of armmigrationassessment.BusinessCasePaasSummaryOperationsClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewBusinessCasePaasSummaryOperationsServerTransport(srv *BusinessCasePaasSummaryOperationsServer) *BusinessCasePaasSummaryOperationsServerTransport {
	return &BusinessCasePaasSummaryOperationsServerTransport{
		srv:                        srv,
		newListByBusinessCasePager: newTracker[azfake.PagerResponder[armmigrationassessment.BusinessCasePaasSummaryOperationsClientListByBusinessCaseResponse]](),
	}
}

// BusinessCasePaasSummaryOperationsServerTransport connects instances of armmigrationassessment.BusinessCasePaasSummaryOperationsClient to instances of BusinessCasePaasSummaryOperationsServer.
// Don't use this type directly, use NewBusinessCasePaasSummaryOperationsServerTransport instead.
type BusinessCasePaasSummaryOperationsServerTransport struct {
	srv                        *BusinessCasePaasSummaryOperationsServer
	newListByBusinessCasePager *tracker[azfake.PagerResponder[armmigrationassessment.BusinessCasePaasSummaryOperationsClientListByBusinessCaseResponse]]
}

// Do implements the policy.Transporter interface for BusinessCasePaasSummaryOperationsServerTransport.
func (b *BusinessCasePaasSummaryOperationsServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	return b.dispatchToMethodFake(req, method)
}

func (b *BusinessCasePaasSummaryOperationsServerTransport) dispatchToMethodFake(req *http.Request, method string) (*http.Response, error) {
	resultChan := make(chan result)
	defer close(resultChan)

	go func() {
		var intercepted bool
		var res result
		if businessCasePaasSummaryOperationsServerTransportInterceptor != nil {
			res.resp, res.err, intercepted = businessCasePaasSummaryOperationsServerTransportInterceptor.Do(req)
		}
		if !intercepted {
			switch method {
			case "BusinessCasePaasSummaryOperationsClient.Get":
				res.resp, res.err = b.dispatchGet(req)
			case "BusinessCasePaasSummaryOperationsClient.NewListByBusinessCasePager":
				res.resp, res.err = b.dispatchNewListByBusinessCasePager(req)
			default:
				res.err = fmt.Errorf("unhandled API %s", method)
			}

		}
		select {
		case resultChan <- res:
		case <-req.Context().Done():
		}
	}()

	select {
	case <-req.Context().Done():
		return nil, req.Context().Err()
	case res := <-resultChan:
		return res.resp, res.err
	}
}

func (b *BusinessCasePaasSummaryOperationsServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if b.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Migrate/assessmentProjects/(?P<projectName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/businessCases/(?P<businessCaseName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/paasSummaries/(?P<paasSummaryName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 5 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	projectNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("projectName")])
	if err != nil {
		return nil, err
	}
	businessCaseNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("businessCaseName")])
	if err != nil {
		return nil, err
	}
	paasSummaryNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("paasSummaryName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := b.srv.Get(req.Context(), resourceGroupNameParam, projectNameParam, businessCaseNameParam, paasSummaryNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).PaasSummary, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (b *BusinessCasePaasSummaryOperationsServerTransport) dispatchNewListByBusinessCasePager(req *http.Request) (*http.Response, error) {
	if b.srv.NewListByBusinessCasePager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListByBusinessCasePager not implemented")}
	}
	newListByBusinessCasePager := b.newListByBusinessCasePager.get(req)
	if newListByBusinessCasePager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Migrate/assessmentProjects/(?P<projectName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/businessCases/(?P<businessCaseName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/paasSummaries`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		projectNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("projectName")])
		if err != nil {
			return nil, err
		}
		businessCaseNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("businessCaseName")])
		if err != nil {
			return nil, err
		}
		resp := b.srv.NewListByBusinessCasePager(resourceGroupNameParam, projectNameParam, businessCaseNameParam, nil)
		newListByBusinessCasePager = &resp
		b.newListByBusinessCasePager.add(req, newListByBusinessCasePager)
		server.PagerResponderInjectNextLinks(newListByBusinessCasePager, req, func(page *armmigrationassessment.BusinessCasePaasSummaryOperationsClientListByBusinessCaseResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListByBusinessCasePager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		b.newListByBusinessCasePager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListByBusinessCasePager) {
		b.newListByBusinessCasePager.remove(req)
	}
	return resp, nil
}

// set this to conditionally intercept incoming requests to BusinessCasePaasSummaryOperationsServerTransport
var businessCasePaasSummaryOperationsServerTransportInterceptor interface {
	// Do returns true if the server transport should use the returned response/error
	Do(*http.Request) (*http.Response, error, bool)
}
