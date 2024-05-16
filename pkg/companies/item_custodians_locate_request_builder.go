package companies

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    if815b544aed335c15fd039b4dcb81545e8ef32c8003c12e1dc36ea8df9bfeb2f "github.com/xifanyan/tddc/pkg/models"
)

// ItemCustodiansLocateRequestBuilder builds and executes requests for operations under \companies\{company_id-id}\custodians\locate
type ItemCustodiansLocateRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemCustodiansLocateRequestBuilderGetQueryParameters finds a custodian by their Unique ID
type ItemCustodiansLocateRequestBuilderGetQueryParameters struct {
    // ID of the custodian
    Id *string `uriparametername:"id"`
}
// ItemCustodiansLocateRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemCustodiansLocateRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemCustodiansLocateRequestBuilderGetQueryParameters
}
// NewItemCustodiansLocateRequestBuilderInternal instantiates a new ItemCustodiansLocateRequestBuilder and sets the default values.
func NewItemCustodiansLocateRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemCustodiansLocateRequestBuilder) {
    m := &ItemCustodiansLocateRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/companies/{company_id%2Did}/custodians/locate?id={id}", pathParameters),
    }
    return m
}
// NewItemCustodiansLocateRequestBuilder instantiates a new ItemCustodiansLocateRequestBuilder and sets the default values.
func NewItemCustodiansLocateRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemCustodiansLocateRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemCustodiansLocateRequestBuilderInternal(urlParams, requestAdapter)
}
// Get finds a custodian by their Unique ID
// returns a Custodianable when successful
func (m *ItemCustodiansLocateRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemCustodiansLocateRequestBuilderGetRequestConfiguration)(if815b544aed335c15fd039b4dcb81545e8ef32c8003c12e1dc36ea8df9bfeb2f.Custodianable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, if815b544aed335c15fd039b4dcb81545e8ef32c8003c12e1dc36ea8df9bfeb2f.CreateCustodianFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(if815b544aed335c15fd039b4dcb81545e8ef32c8003c12e1dc36ea8df9bfeb2f.Custodianable), nil
}
// ToGetRequestInformation finds a custodian by their Unique ID
// returns a *RequestInformation when successful
func (m *ItemCustodiansLocateRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemCustodiansLocateRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ItemCustodiansLocateRequestBuilder when successful
func (m *ItemCustodiansLocateRequestBuilder) WithUrl(rawUrl string)(*ItemCustodiansLocateRequestBuilder) {
    return NewItemCustodiansLocateRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
