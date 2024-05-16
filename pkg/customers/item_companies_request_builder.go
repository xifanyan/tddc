package customers

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    if815b544aed335c15fd039b4dcb81545e8ef32c8003c12e1dc36ea8df9bfeb2f "github.com/xifanyan/tddc/pkg/models"
)

// ItemCompaniesRequestBuilder builds and executes requests for operations under \customers\{customer_id}\companies
type ItemCompaniesRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemCompaniesRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemCompaniesRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemCompaniesRequestBuilderInternal instantiates a new ItemCompaniesRequestBuilder and sets the default values.
func NewItemCompaniesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemCompaniesRequestBuilder) {
    m := &ItemCompaniesRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/customers/{customer_id}/companies", pathParameters),
    }
    return m
}
// NewItemCompaniesRequestBuilder instantiates a new ItemCompaniesRequestBuilder and sets the default values.
func NewItemCompaniesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemCompaniesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemCompaniesRequestBuilderInternal(urlParams, requestAdapter)
}
// Get list all companies in a customer
// returns a []Companiesable when successful
func (m *ItemCompaniesRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemCompaniesRequestBuilderGetRequestConfiguration)([]if815b544aed335c15fd039b4dcb81545e8ef32c8003c12e1dc36ea8df9bfeb2f.Companiesable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.SendCollection(ctx, requestInfo, if815b544aed335c15fd039b4dcb81545e8ef32c8003c12e1dc36ea8df9bfeb2f.CreateCompaniesFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    val := make([]if815b544aed335c15fd039b4dcb81545e8ef32c8003c12e1dc36ea8df9bfeb2f.Companiesable, len(res))
    for i, v := range res {
        if v != nil {
            val[i] = v.(if815b544aed335c15fd039b4dcb81545e8ef32c8003c12e1dc36ea8df9bfeb2f.Companiesable)
        }
    }
    return val, nil
}
// ToGetRequestInformation list all companies in a customer
// returns a *RequestInformation when successful
func (m *ItemCompaniesRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemCompaniesRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ItemCompaniesRequestBuilder when successful
func (m *ItemCompaniesRequestBuilder) WithUrl(rawUrl string)(*ItemCompaniesRequestBuilder) {
    return NewItemCompaniesRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
