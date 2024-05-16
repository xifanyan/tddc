package companies

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    if815b544aed335c15fd039b4dcb81545e8ef32c8003c12e1dc36ea8df9bfeb2f "github.com/xifanyan/tddc/pkg/models"
)

// ItemCustodiansTraitsWithTrait_nameItemRequestBuilder builds and executes requests for operations under \companies\{company_id-id}\custodians\traits\{trait_name}
type ItemCustodiansTraitsWithTrait_nameItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemCustodiansTraitsWithTrait_nameItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemCustodiansTraitsWithTrait_nameItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemCustodiansTraitsWithTrait_nameItemRequestBuilderInternal instantiates a new ItemCustodiansTraitsWithTrait_nameItemRequestBuilder and sets the default values.
func NewItemCustodiansTraitsWithTrait_nameItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemCustodiansTraitsWithTrait_nameItemRequestBuilder) {
    m := &ItemCustodiansTraitsWithTrait_nameItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/companies/{company_id%2Did}/custodians/traits/{trait_name}", pathParameters),
    }
    return m
}
// NewItemCustodiansTraitsWithTrait_nameItemRequestBuilder instantiates a new ItemCustodiansTraitsWithTrait_nameItemRequestBuilder and sets the default values.
func NewItemCustodiansTraitsWithTrait_nameItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemCustodiansTraitsWithTrait_nameItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemCustodiansTraitsWithTrait_nameItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Get list all unique trait values for a trait name for all custodians in a company
// returns a Traitvaluesresponseable when successful
func (m *ItemCustodiansTraitsWithTrait_nameItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemCustodiansTraitsWithTrait_nameItemRequestBuilderGetRequestConfiguration)(if815b544aed335c15fd039b4dcb81545e8ef32c8003c12e1dc36ea8df9bfeb2f.Traitvaluesresponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, if815b544aed335c15fd039b4dcb81545e8ef32c8003c12e1dc36ea8df9bfeb2f.CreateTraitvaluesresponseFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(if815b544aed335c15fd039b4dcb81545e8ef32c8003c12e1dc36ea8df9bfeb2f.Traitvaluesresponseable), nil
}
// Sync the sync property
// returns a *ItemCustodiansTraitsItemSyncRequestBuilder when successful
func (m *ItemCustodiansTraitsWithTrait_nameItemRequestBuilder) Sync()(*ItemCustodiansTraitsItemSyncRequestBuilder) {
    return NewItemCustodiansTraitsItemSyncRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToGetRequestInformation list all unique trait values for a trait name for all custodians in a company
// returns a *RequestInformation when successful
func (m *ItemCustodiansTraitsWithTrait_nameItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemCustodiansTraitsWithTrait_nameItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ItemCustodiansTraitsWithTrait_nameItemRequestBuilder when successful
func (m *ItemCustodiansTraitsWithTrait_nameItemRequestBuilder) WithUrl(rawUrl string)(*ItemCustodiansTraitsWithTrait_nameItemRequestBuilder) {
    return NewItemCustodiansTraitsWithTrait_nameItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
