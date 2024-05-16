package companies

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    if815b544aed335c15fd039b4dcb81545e8ef32c8003c12e1dc36ea8df9bfeb2f "github.com/xifanyan/tddc/pkg/models"
)

// ItemCustodiansTraitsRequestBuilder builds and executes requests for operations under \companies\{company_id-id}\custodians\traits
type ItemCustodiansTraitsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemCustodiansTraitsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemCustodiansTraitsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ByTrait_name gets an item from the github.com/xifanyan/tddc/pkg.companies.item.custodians.traits.item collection
// returns a *ItemCustodiansTraitsWithTrait_nameItemRequestBuilder when successful
func (m *ItemCustodiansTraitsRequestBuilder) ByTrait_name(trait_name string)(*ItemCustodiansTraitsWithTrait_nameItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if trait_name != "" {
        urlTplParams["trait_name"] = trait_name
    }
    return NewItemCustodiansTraitsWithTrait_nameItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewItemCustodiansTraitsRequestBuilderInternal instantiates a new ItemCustodiansTraitsRequestBuilder and sets the default values.
func NewItemCustodiansTraitsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemCustodiansTraitsRequestBuilder) {
    m := &ItemCustodiansTraitsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/companies/{company_id%2Did}/custodians/traits", pathParameters),
    }
    return m
}
// NewItemCustodiansTraitsRequestBuilder instantiates a new ItemCustodiansTraitsRequestBuilder and sets the default values.
func NewItemCustodiansTraitsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemCustodiansTraitsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemCustodiansTraitsRequestBuilderInternal(urlParams, requestAdapter)
}
// Get list all unique trait names for all custodians in a company.
// returns a TraitNamesResponseable when successful
func (m *ItemCustodiansTraitsRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemCustodiansTraitsRequestBuilderGetRequestConfiguration)(if815b544aed335c15fd039b4dcb81545e8ef32c8003c12e1dc36ea8df9bfeb2f.TraitNamesResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, if815b544aed335c15fd039b4dcb81545e8ef32c8003c12e1dc36ea8df9bfeb2f.CreateTraitNamesResponseFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(if815b544aed335c15fd039b4dcb81545e8ef32c8003c12e1dc36ea8df9bfeb2f.TraitNamesResponseable), nil
}
// ToGetRequestInformation list all unique trait names for all custodians in a company.
// returns a *RequestInformation when successful
func (m *ItemCustodiansTraitsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemCustodiansTraitsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ItemCustodiansTraitsRequestBuilder when successful
func (m *ItemCustodiansTraitsRequestBuilder) WithUrl(rawUrl string)(*ItemCustodiansTraitsRequestBuilder) {
    return NewItemCustodiansTraitsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
