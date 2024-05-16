package collections

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    if815b544aed335c15fd039b4dcb81545e8ef32c8003c12e1dc36ea8df9bfeb2f "github.com/xifanyan/tddc/pkg/models"
)

// WithCollection_ItemRequestBuilder builds and executes requests for operations under \collections\{collection_id}
type WithCollection_ItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// WithCollection_ItemRequestBuilderGetQueryParameters get Collection
type WithCollection_ItemRequestBuilderGetQueryParameters struct {
    // If appended, will show the Collection Server ID and Pocoa Tag. These are internal IDs, but knowing them can be helpful in getting support. Note: Please assume these values will change or be removed in the future. 
    Show_internal_ids *bool `uriparametername:"show_internal_ids"`
}
// WithCollection_ItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type WithCollection_ItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *WithCollection_ItemRequestBuilderGetQueryParameters
}
// NewWithCollection_ItemRequestBuilderInternal instantiates a new WithCollection_ItemRequestBuilder and sets the default values.
func NewWithCollection_ItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*WithCollection_ItemRequestBuilder) {
    m := &WithCollection_ItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/collections/{collection_id}{?show_internal_ids*}", pathParameters),
    }
    return m
}
// NewWithCollection_ItemRequestBuilder instantiates a new WithCollection_ItemRequestBuilder and sets the default values.
func NewWithCollection_ItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*WithCollection_ItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewWithCollection_ItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Get get Collection
// returns a CollectionOutputable when successful
func (m *WithCollection_ItemRequestBuilder) Get(ctx context.Context, requestConfiguration *WithCollection_ItemRequestBuilderGetRequestConfiguration)(if815b544aed335c15fd039b4dcb81545e8ef32c8003c12e1dc36ea8df9bfeb2f.CollectionOutputable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, if815b544aed335c15fd039b4dcb81545e8ef32c8003c12e1dc36ea8df9bfeb2f.CreateCollectionOutputFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(if815b544aed335c15fd039b4dcb81545e8ef32c8003c12e1dc36ea8df9bfeb2f.CollectionOutputable), nil
}
// ToGetRequestInformation get Collection
// returns a *RequestInformation when successful
func (m *WithCollection_ItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *WithCollection_ItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *WithCollection_ItemRequestBuilder when successful
func (m *WithCollection_ItemRequestBuilder) WithUrl(rawUrl string)(*WithCollection_ItemRequestBuilder) {
    return NewWithCollection_ItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
