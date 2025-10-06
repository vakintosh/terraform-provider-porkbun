package apiclient

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/oapi-codegen/runtime"
)

// DnssecCreateRecordJSONBody defines the parameters for DnssecCreateRecord
type DnssecCreateRecordJSONBody struct {
	Apikey          string `json:"apikey"`
	Secretapikey    string `json:"secretapikey"`
	KeyTag          string `json:"keyTag"`
	Alg             string `json:"alg"`
	DigestType      string `json:"digestType"`
	Digest          string `json:"digest"`
	MaxSigLife      string `json:"maxSigLife"`
	KeyDataFlags    string `json:"keyFlags"`
	KeyDataProtocol string `json:"keyProtocol"`
	KeyDataAlgo     string `json:"keyAlgo"`
	KeyDataPubKey   string `json:"keyPubKey"`
}

// DnssecCreateRecordJSONRequestBody defines body for DnssecCreateRecord for application/json ContentType.
type DnssecCreateRecordJSONRequestBody DnssecCreateRecordJSONBody

// // The interface specification for the client above.
// type ClientInterface interface {

// 	// DnssecCreateRecordWithBody request with any body
// 	DnssecCreateRecordWithBody(ctx context.Context, domain DomainPath, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

// 	DnssecCreateRecord(ctx context.Context, domain DomainPath, body DnssecCreateRecordJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

// }

func (c *Client) DnssecCreateRecordWithBody(ctx context.Context, domain DomainPath, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewDnssecCreateRecordRequestWithBody(c.Server, domain, contentType, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) DnssecCreateRecord(ctx context.Context, domain DomainPath, body DnssecCreateRecordJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewDnssecCreateRecordRequest(c.Server, domain, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

// NewDnssecCreateRecordRequest calls the generic DnssecCreateRecord builder with application/json body
func NewDnssecCreateRecordRequest(server string, domain DomainPath, body DnssecCreateRecordJSONRequestBody) (*http.Request, error) {
	var bodyReader io.Reader
	buf, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	bodyReader = bytes.NewReader(buf)
	return NewDnssecCreateRecordRequestWithBody(server, domain, "application/json", bodyReader)
}

// NewDnssecCreateRecordRequestWithBody generates requests for DnssecCreateRecord with any type of body
func NewDnssecCreateRecordRequestWithBody(server string, domain DomainPath, contentType string, body io.Reader) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "domain", runtime.ParamLocationPath, domain)
	if err != nil {
		return nil, err
	}

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/v3/dns/createDnssecRecord/%s", pathParam0)
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", queryURL.String(), body)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", contentType)

	return req, nil
}

// type ClientWithResponsesInterface interface {

// 	// DnssecCreateRecordWithBodyWithResponse request with any body
// 	DnssecCreateRecordWithBodyWithResponse(ctx context.Context, domain DomainPath, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*DnssecCreateRecordResp, error)

// 	DnssecCreateRecordWithResponse(ctx context.Context, domain DomainPath, body DnssecCreateRecordJSONRequestBody, reqEditors ...RequestEditorFn) (*DnssecCreateRecordResp, error)
// }

type DnssecCreateRecordResp struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *struct {
		Status string `json:"status"`
	}
}

// Status returns HTTPResponse.Status
func (r DnssecCreateRecordResp) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r DnssecCreateRecordResp) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

// DnssecCreateRecordWithBodyWithResponse request with arbitrary body returning *DnssecCreateRecordResp
func (c *ClientWithResponses) DnssecCreateRecordWithBodyWithResponse(ctx context.Context, domain DomainPath, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*DnssecCreateRecordResp, error) {
	rsp, err := c.DnssecCreateRecordWithBody(ctx, domain, contentType, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseDnssecCreateRecordResp(rsp)
}

func (c *ClientWithResponses) DnssecCreateRecordWithResponse(ctx context.Context, domain DomainPath, body DnssecCreateRecordJSONRequestBody, reqEditors ...RequestEditorFn) (*DnssecCreateRecordResp, error) {
	rsp, err := c.DnssecCreateRecord(ctx, domain, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseDnssecCreateRecordResp(rsp)
}

// ParseDnssecCreateRecordResp parses an HTTP response from a DnssecCreateRecordWithResponse call
func ParseDnssecCreateRecordResp(rsp *http.Response) (*DnssecCreateRecordResp, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &DnssecCreateRecordResp{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest struct {
			Status string `json:"status"`
		}
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest
	}

	return response, nil
}

//
// GET RECORDS
//

// DnssecGetRecord defines the model for a DNSSEC record.
type DnssecGetRecord struct {
	Id           string `json:"id"`
	Apikey       string `json:"apikey"`
	Secretapikey string `json:"secretapikey"`
	// Domain          *string `json:"domain,omitempty"`
	KeyTag          string  `json:"keyTag"`
	Alg             string  `json:"alg"`
	DigestType      string  `json:"digestType"`
	Digest          string  `json:"digest"`
	MaxSigLife      *string `json:"maxSigLife,omitempty"`
	KeyDataFlags    *string `json:"keyFlags,omitempty"`
	KeyDataProtocol *string `json:"keyProtocol,omitempty"`
	KeyDataAlgo     *string `json:"keyAlgo,omitempty"`
	KeyDataPubKey   *string `json:"keyPubKey,omitempty"`
}

// DnssecGetRecordsResponse defines the model for the DNSSEC Get Records response.
type DnssecGetRecordsResponse struct {
	Status  string                     `json:"status"`
	Records map[string]DnssecGetRecord `json:"records"`
}

// DnssecGetRecordsJSONRequestBody defines the body for DnssecGetRecords for application/json ContentType.
type DnssecGetRecordsJSONRequestBody = ApiKeyAndSecretKey

// // The interface specification for the client above.
// type ClientInterface interface {
//
// 	// DnssecGetRecordsWithBody request with any body
// 	DnssecGetRecordsWithBody(ctx context.Context, domain DomainPath, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

// 	// DnssecGetRecords request with the standard JSON body
// 	DnssecGetRecords(ctx context.Context, domain DomainPath, body DnssecGetRecordsJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)
// }

func (c *Client) DnssecGetRecordsWithBody(ctx context.Context, domain DomainPath, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewDnssecGetRecordsRequestWithBody(c.Server, domain, contentType, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) DnssecGetRecords(ctx context.Context, domain DomainPath, body DnssecGetRecordsJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewDnssecGetRecordsRequest(c.Server, domain, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

// NewDnssecGetRecordsRequest calls the generic DnssecGetRecords builder with application/json body
func NewDnssecGetRecordsRequest(server string, domain DomainPath, body DnssecGetRecordsJSONRequestBody) (*http.Request, error) {
	var bodyReader io.Reader
	buf, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	bodyReader = bytes.NewReader(buf)
	return NewDnssecGetRecordsRequestWithBody(server, domain, "application/json", bodyReader)
}

// NewDnssecGetRecordsRequestWithBody generates requests for DnssecGetRecords with any type of body
func NewDnssecGetRecordsRequestWithBody(server string, domain DomainPath, contentType string, body io.Reader) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "domain", runtime.ParamLocationPath, domain)
	if err != nil {
		return nil, err
	}

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/v3/dns/getDnssecRecords/%s", pathParam0)
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", queryURL.String(), body)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", contentType)

	return req, nil
}

// type ClientWithResponsesInterface interface {

// 	// DnssecGetRecordsWithBodyWithResponse request with any body
// 	DnssecGetRecordsWithBodyWithResponse(ctx context.Context, domain DomainPath, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*DnssecGetRecordsResp, error)

// 	DnssecGetRecordsWithResponse(ctx context.Context, domain DomainPath, body DnssecGetRecordsJSONRequestBody, reqEditors ...RequestEditorFn) (*DnssecGetRecordsResp, error)
// }

type DnssecGetRecordsResp struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *DnssecGetRecordsResponse
}

// Status returns HTTPResponse.Status
func (r DnssecGetRecordsResp) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r DnssecGetRecordsResp) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

// DnssecGetRecordsWithBodyWithResponse request with arbitrary body returning *DnssecGetRecordsResp
func (c *ClientWithResponses) DnssecGetRecordsWithBodyWithResponse(ctx context.Context, domain DomainPath, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*DnssecGetRecordsResp, error) {
	rsp, err := c.DnssecGetRecordsWithBody(ctx, domain, contentType, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseDnssecGetRecordsResp(rsp)
}

func (c *ClientWithResponses) DnssecGetRecordsWithResponse(ctx context.Context, domain DomainPath, body DnssecGetRecordsJSONRequestBody, reqEditors ...RequestEditorFn) (*DnssecGetRecordsResp, error) {
	rsp, err := c.DnssecGetRecords(ctx, domain, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseDnssecGetRecordsResp(rsp)
}

// ParseDnssecGetRecordsResp parses an HTTP response from a DnssecGetRecordsWithResponse call
func ParseDnssecGetRecordsResp(rsp *http.Response) (*DnssecGetRecordsResp, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &DnssecGetRecordsResp{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest DnssecGetRecordsResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest
	}

	return response, nil
}

//
// Delete Record
//

// DnssecDeleteRecordJSONRequestBody defines the body for DnssecDeleteRecord for application/json ContentType.
type DnssecDeleteRecordJSONRequestBody = ApiKeyAndSecretKey

// // The interface specification for the client above.
// type ClientInterface interface {
//
// 	// DnssecDeleteRecordWithBody request with any body
// 	DnssecDeleteRecordWithBody(ctx context.Context, domain DomainPath, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

// 	DnssecDeleteRecord(ctx context.Context, domain DomainPath, body DnssecDeleteRecordJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)
// }

func (c *Client) DnssecDeleteRecordWithBody(ctx context.Context, domain DomainPath, keytag string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewDnssecDeleteRecordRequestWithBody(c.Server, domain, keytag, contentType, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) DnssecDeleteRecord(ctx context.Context, domain DomainPath, keytag string, body DnssecDeleteRecordJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewDnssecDeleteRecordRequest(c.Server, domain, keytag, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

// NewDnssecDeleteRecordRequest calls the generic DnssecDeleteRecord builder with application/json body
func NewDnssecDeleteRecordRequest(server string, domain DomainPath, keytag string, body DnssecDeleteRecordJSONRequestBody) (*http.Request, error) {
	var bodyReader io.Reader
	buf, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	bodyReader = bytes.NewReader(buf)
	return NewDnssecDeleteRecordRequestWithBody(server, domain, keytag, "application/json", bodyReader)
}

// NewDnssecDeleteRecordRequestWithBody generates requests for DnssecDeleteRecord with any type of body
func NewDnssecDeleteRecordRequestWithBody(server string, domain DomainPath, keytag string, contentType string, body io.Reader) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "domain", runtime.ParamLocationPath, domain)
	if err != nil {
		return nil, err
	}

	var pathParam1 string

	pathParam1, err = runtime.StyleParamWithLocation("simple", false, "keytag", runtime.ParamLocationPath, keytag)
	if err != nil {
		return nil, err
	}

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/v3/dns/deleteDnssecRecord/%s/%s", pathParam0, pathParam1)
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", queryURL.String(), body)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", contentType)

	return req, nil
}

// type ClientWithResponsesInterface interface {

// 	// DnssecDeleteRecordWithBodyWithResponse request with any body
// 	DnssecDeleteRecordWithBodyWithResponse(ctx context.Context, domain DomainPath, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*DnssecDeleteRecordResp, error)

// 	DnssecDeleteRecordByKeyTagWithResponse(ctx context.Context, domain DomainPath, body DnssecDeleteRecordJSONRequestBody, reqEditors ...RequestEditorFn) (*DnssecDeleteRecordResp, error)
// }

type DnssecDeleteRecordResp struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *struct {
		Status string `json:"status"`
	}
}

// Status returns HTTPResponse.Status
func (r DnssecDeleteRecordResp) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r DnssecDeleteRecordResp) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

// DnssecDeleteRecordWithBodyWithResponse request with arbitrary body returning *DnssecDeleteRecordResp
func (c *ClientWithResponses) DnssecDeleteRecordWithBodyWithResponse(ctx context.Context, domain DomainPath, keytag string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*DnssecDeleteRecordResp, error) {
	rsp, err := c.DnssecDeleteRecordWithBody(ctx, domain, keytag, contentType, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseDnssecDeleteRecordResp(rsp)
}

func (c *ClientWithResponses) DnssecDeleteRecordByKeyTagWithResponse(ctx context.Context, domain DomainPath, keytag string, body DnssecDeleteRecordJSONRequestBody, reqEditors ...RequestEditorFn) (*DnssecDeleteRecordResp, error) {
	rsp, err := c.DnssecDeleteRecord(ctx, domain, keytag, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseDnssecDeleteRecordResp(rsp)
}

// ParseDnssecDeleteRecordResp parses an HTTP response from a DnssecDeleteRecordByKeyTagWithResponse call
func ParseDnssecDeleteRecordResp(rsp *http.Response) (*DnssecDeleteRecordResp, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &DnssecDeleteRecordResp{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest struct {
			Status string `json:"status"`
		}
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest
	}

	return response, nil
}
