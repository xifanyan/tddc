package custodians

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    if815b544aed335c15fd039b4dcb81545e8ef32c8003c12e1dc36ea8df9bfeb2f "github.com/xifanyan/tddc/pkg/models"
)

// ItemTraitsRequestBuilder builds and executes requests for operations under \custodians\{custodian_id}\traits
type ItemTraitsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemTraitsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemTraitsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ByTrait_name gets an item from the github.com/xifanyan/tddc/pkg.custodians.item.traits.item collection
// returns a *ItemTraitsWithTrait_nameItemRequestBuilder when successful
func (m *ItemTraitsRequestBuilder) ByTrait_name(trait_name string)(*ItemTraitsWithTrait_nameItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if trait_name != "" {
        urlTplParams["trait_name"] = trait_name
    }
    return NewItemTraitsWithTrait_nameItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewItemTraitsRequestBuilderInternal instantiates a new ItemTraitsRequestBuilder and sets the default values.
func NewItemTraitsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemTraitsRequestBuilder) {
    m := &ItemTraitsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/custodians/{custodian_id}/traits", pathParameters),
    }
    return m
}
// NewItemTraitsRequestBuilder instantiates a new ItemTraitsRequestBuilder and sets the default values.
func NewItemTraitsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemTraitsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemTraitsRequestBuilderInternal(urlParams, requestAdapter)
}
// Get list all unique trait names for a custodian.
// returns a TraitNamesResponseable when successful
func (m *ItemTraitsRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemTraitsRequestBuilderGetRequestConfiguration)(if815b544aed335c15fd039b4dcb81545e8ef32c8003c12e1dc36ea8df9bfeb2f.TraitNamesResponseable, error) {
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
// ToGetRequestInformation list all unique trait names for a custodian.
// returns a *RequestInformation when successful
func (m *ItemTraitsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemTraitsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ItemTraitsRequestBuilder when successful
func (m *ItemTraitsRequestBuilder) WithUrl(rawUrl string)(*ItemTraitsRequestBuilder) {
    return NewItemTraitsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
