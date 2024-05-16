package matters

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    if815b544aed335c15fd039b4dcb81545e8ef32c8003c12e1dc36ea8df9bfeb2f "github.com/xifanyan/tddc/pkg/models"
)

// ItemCustodiansRequestBuilder builds and executes requests for operations under \matters\{matter_id}\custodians
type ItemCustodiansRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemCustodiansRequestBuilderGetQueryParameters list all custodians in a matter
type ItemCustodiansRequestBuilderGetQueryParameters struct {
    Limit *float64 `uriparametername:"limit"`
    Offset *float64 `uriparametername:"offset"`
}
// ItemCustodiansRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemCustodiansRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemCustodiansRequestBuilderGetQueryParameters
}
// NewItemCustodiansRequestBuilderInternal instantiates a new ItemCustodiansRequestBuilder and sets the default values.
func NewItemCustodiansRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemCustodiansRequestBuilder) {
    m := &ItemCustodiansRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/matters/{matter_id}/custodians{?limit*,offset*}", pathParameters),
    }
    return m
}
// NewItemCustodiansRequestBuilder instantiates a new ItemCustodiansRequestBuilder and sets the default values.
func NewItemCustodiansRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemCustodiansRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemCustodiansRequestBuilderInternal(urlParams, requestAdapter)
}
// Get list all custodians in a matter
// returns a []Custodianable when successful
// returns a ApiError error when the service returns a 500 status code
func (m *ItemCustodiansRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemCustodiansRequestBuilderGetRequestConfiguration)([]if815b544aed335c15fd039b4dcb81545e8ef32c8003c12e1dc36ea8df9bfeb2f.Custodianable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "500": if815b544aed335c15fd039b4dcb81545e8ef32c8003c12e1dc36ea8df9bfeb2f.CreateApiErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.SendCollection(ctx, requestInfo, if815b544aed335c15fd039b4dcb81545e8ef32c8003c12e1dc36ea8df9bfeb2f.CreateCustodianFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    val := make([]if815b544aed335c15fd039b4dcb81545e8ef32c8003c12e1dc36ea8df9bfeb2f.Custodianable, len(res))
    for i, v := range res {
        if v != nil {
            val[i] = v.(if815b544aed335c15fd039b4dcb81545e8ef32c8003c12e1dc36ea8df9bfeb2f.Custodianable)
        }
    }
    return val, nil
}
// Locate the locate property
// returns a *ItemCustodiansLocateRequestBuilder when successful
func (m *ItemCustodiansRequestBuilder) Locate()(*ItemCustodiansLocateRequestBuilder) {
    return NewItemCustodiansLocateRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToGetRequestInformation list all custodians in a matter
// returns a *RequestInformation when successful
func (m *ItemCustodiansRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemCustodiansRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ItemCustodiansRequestBuilder when successful
func (m *ItemCustodiansRequestBuilder) WithUrl(rawUrl string)(*ItemCustodiansRequestBuilder) {
    return NewItemCustodiansRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
