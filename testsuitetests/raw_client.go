package testsuitetests

import (
	context "context"
	server "github.com/VapiAI/server-sdk-go"
	core "github.com/VapiAI/server-sdk-go/core"
	internal "github.com/VapiAI/server-sdk-go/internal"
	option "github.com/VapiAI/server-sdk-go/option"
	http "net/http"
)

type RawClient struct {
	baseURL string
	caller  *internal.Caller
	header  http.Header
}

func NewRawClient(options *core.RequestOptions) *RawClient {
	return &RawClient{
		baseURL: options.BaseURL,
		caller: internal.NewCaller(
			&internal.CallerParams{
				Client:      options.HTTPClient,
				MaxAttempts: options.MaxAttempts,
			},
		),
		header: options.ToHeader(),
	}
}

func (r *RawClient) TestSuiteTestControllerFindAllPaginated(
	ctx context.Context,
	testSuiteId string,
	request *server.TestSuiteTestControllerFindAllPaginatedRequest,
	opts ...option.RequestOption,
) (*core.Response[*server.TestSuiteTestsPaginatedResponse], error) {
	options := core.NewRequestOptions(opts...)
	baseURL := internal.ResolveBaseURL(
		options.BaseURL,
		r.baseURL,
		"https://api.vapi.ai",
	)
	endpointURL := internal.EncodeURL(
		baseURL+"/test-suite/%v/test",
		testSuiteId,
	)
	queryParams, err := internal.QueryValues(request)
	if err != nil {
		return nil, err
	}
	if len(queryParams) > 0 {
		endpointURL += "?" + queryParams.Encode()
	}
	headers := internal.MergeHeaders(
		r.header.Clone(),
		options.ToHeader(),
	)
	var response *server.TestSuiteTestsPaginatedResponse
	raw, err := r.caller.Call(
		ctx,
		&internal.CallParams{
			URL:             endpointURL,
			Method:          http.MethodGet,
			Headers:         headers,
			MaxAttempts:     options.MaxAttempts,
			BodyProperties:  options.BodyProperties,
			QueryParameters: options.QueryParameters,
			Client:          options.HTTPClient,
			Request:         request,
			Response:        &response,
		},
	)
	if err != nil {
		return nil, err
	}
	return &core.Response[*server.TestSuiteTestsPaginatedResponse]{
		StatusCode: raw.StatusCode,
		Header:     raw.Header,
		Body:       response,
	}, nil
}

func (r *RawClient) TestSuiteTestControllerCreate(
	ctx context.Context,
	testSuiteId string,
	request *server.TestSuiteTestControllerCreateRequest,
	opts ...option.RequestOption,
) (*core.Response[*server.TestSuiteTestControllerCreateResponse], error) {
	options := core.NewRequestOptions(opts...)
	baseURL := internal.ResolveBaseURL(
		options.BaseURL,
		r.baseURL,
		"https://api.vapi.ai",
	)
	endpointURL := internal.EncodeURL(
		baseURL+"/test-suite/%v/test",
		testSuiteId,
	)
	headers := internal.MergeHeaders(
		r.header.Clone(),
		options.ToHeader(),
	)
	var response *server.TestSuiteTestControllerCreateResponse
	raw, err := r.caller.Call(
		ctx,
		&internal.CallParams{
			URL:             endpointURL,
			Method:          http.MethodPost,
			Headers:         headers,
			MaxAttempts:     options.MaxAttempts,
			BodyProperties:  options.BodyProperties,
			QueryParameters: options.QueryParameters,
			Client:          options.HTTPClient,
			Request:         request,
			Response:        &response,
		},
	)
	if err != nil {
		return nil, err
	}
	return &core.Response[*server.TestSuiteTestControllerCreateResponse]{
		StatusCode: raw.StatusCode,
		Header:     raw.Header,
		Body:       response,
	}, nil
}

func (r *RawClient) TestSuiteTestControllerFindOne(
	ctx context.Context,
	testSuiteId string,
	id string,
	opts ...option.RequestOption,
) (*core.Response[*server.TestSuiteTestControllerFindOneResponse], error) {
	options := core.NewRequestOptions(opts...)
	baseURL := internal.ResolveBaseURL(
		options.BaseURL,
		r.baseURL,
		"https://api.vapi.ai",
	)
	endpointURL := internal.EncodeURL(
		baseURL+"/test-suite/%v/test/%v",
		testSuiteId,
		id,
	)
	headers := internal.MergeHeaders(
		r.header.Clone(),
		options.ToHeader(),
	)
	var response *server.TestSuiteTestControllerFindOneResponse
	raw, err := r.caller.Call(
		ctx,
		&internal.CallParams{
			URL:             endpointURL,
			Method:          http.MethodGet,
			Headers:         headers,
			MaxAttempts:     options.MaxAttempts,
			BodyProperties:  options.BodyProperties,
			QueryParameters: options.QueryParameters,
			Client:          options.HTTPClient,
			Response:        &response,
		},
	)
	if err != nil {
		return nil, err
	}
	return &core.Response[*server.TestSuiteTestControllerFindOneResponse]{
		StatusCode: raw.StatusCode,
		Header:     raw.Header,
		Body:       response,
	}, nil
}

func (r *RawClient) TestSuiteTestControllerRemove(
	ctx context.Context,
	testSuiteId string,
	id string,
	opts ...option.RequestOption,
) (*core.Response[*server.TestSuiteTestControllerRemoveResponse], error) {
	options := core.NewRequestOptions(opts...)
	baseURL := internal.ResolveBaseURL(
		options.BaseURL,
		r.baseURL,
		"https://api.vapi.ai",
	)
	endpointURL := internal.EncodeURL(
		baseURL+"/test-suite/%v/test/%v",
		testSuiteId,
		id,
	)
	headers := internal.MergeHeaders(
		r.header.Clone(),
		options.ToHeader(),
	)
	var response *server.TestSuiteTestControllerRemoveResponse
	raw, err := r.caller.Call(
		ctx,
		&internal.CallParams{
			URL:             endpointURL,
			Method:          http.MethodDelete,
			Headers:         headers,
			MaxAttempts:     options.MaxAttempts,
			BodyProperties:  options.BodyProperties,
			QueryParameters: options.QueryParameters,
			Client:          options.HTTPClient,
			Response:        &response,
		},
	)
	if err != nil {
		return nil, err
	}
	return &core.Response[*server.TestSuiteTestControllerRemoveResponse]{
		StatusCode: raw.StatusCode,
		Header:     raw.Header,
		Body:       response,
	}, nil
}

func (r *RawClient) TestSuiteTestControllerUpdate(
	ctx context.Context,
	testSuiteId string,
	id string,
	request *server.TestSuiteTestControllerUpdateRequest,
	opts ...option.RequestOption,
) (*core.Response[*server.TestSuiteTestControllerUpdateResponse], error) {
	options := core.NewRequestOptions(opts...)
	baseURL := internal.ResolveBaseURL(
		options.BaseURL,
		r.baseURL,
		"https://api.vapi.ai",
	)
	endpointURL := internal.EncodeURL(
		baseURL+"/test-suite/%v/test/%v",
		testSuiteId,
		id,
	)
	headers := internal.MergeHeaders(
		r.header.Clone(),
		options.ToHeader(),
	)
	var response *server.TestSuiteTestControllerUpdateResponse
	raw, err := r.caller.Call(
		ctx,
		&internal.CallParams{
			URL:             endpointURL,
			Method:          http.MethodPatch,
			Headers:         headers,
			MaxAttempts:     options.MaxAttempts,
			BodyProperties:  options.BodyProperties,
			QueryParameters: options.QueryParameters,
			Client:          options.HTTPClient,
			Request:         request,
			Response:        &response,
		},
	)
	if err != nil {
		return nil, err
	}
	return &core.Response[*server.TestSuiteTestControllerUpdateResponse]{
		StatusCode: raw.StatusCode,
		Header:     raw.Header,
		Body:       response,
	}, nil
}
