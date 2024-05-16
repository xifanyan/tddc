package searches

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    if815b544aed335c15fd039b4dcb81545e8ef32c8003c12e1dc36ea8df9bfeb2f "github.com/xifanyan/tddc/pkg/models"
)

// WithSearch_ItemRequestBuilder builds and executes requests for operations under \searches\{search_id}
type WithSearch_ItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// WithSearch_ItemRequestBuilderGetQueryParameters get a single search
type WithSearch_ItemRequestBuilderGetQueryParameters struct {
    Diag *bool `uriparametername:"diag"`
}
// WithSearch_ItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type WithSearch_ItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *WithSearch_ItemRequestBuilderGetQueryParameters
}
// NewWithSearch_ItemRequestBuilderInternal instantiates a new WithSearch_ItemRequestBuilder and sets the default values.
func NewWithSearch_ItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*WithSearch_ItemRequestBuilder) {
    m := &WithSearch_ItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/searches/{search_id}{?diag*}", pathParameters),
    }
    return m
}
// NewWithSearch_ItemRequestBuilder instantiates a new WithSearch_ItemRequestBuilder and sets the default values.
func NewWithSearch_ItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*WithSearch_ItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewWithSearch_ItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Get get a single search
// returns a Searchable when successful
func (m *WithSearch_ItemRequestBuilder) Get(ctx context.Context, requestConfiguration *WithSearch_ItemRequestBuilderGetRequestConfiguration)(if815b544aed335c15fd039b4dcb81545e8ef32c8003c12e1dc36ea8df9bfeb2f.Searchable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, if815b544aed335c15fd039b4dcb81545e8ef32c8003c12e1dc36ea8df9bfeb2f.CreateSearchFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(if815b544aed335c15fd039b4dcb81545e8ef32c8003c12e1dc36ea8df9bfeb2f.Searchable), nil
}
// ToGetRequestInformation get a single search
// returns a *RequestInformation when successful
func (m *WithSearch_ItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *WithSearch_ItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *WithSearch_ItemRequestBuilder when successful
func (m *WithSearch_ItemRequestBuilder) WithUrl(rawUrl string)(*WithSearch_ItemRequestBuilder) {
    return NewWithSearch_ItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
