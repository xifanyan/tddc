package custodians

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    if815b544aed335c15fd039b4dcb81545e8ef32c8003c12e1dc36ea8df9bfeb2f "github.com/xifanyan/tddc/pkg/models"
)

// ItemTraitsWithTrait_nameItemRequestBuilder builds and executes requests for operations under \custodians\{custodian_id}\traits\{trait_name}
type ItemTraitsWithTrait_nameItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemTraitsWithTrait_nameItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemTraitsWithTrait_nameItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemTraitsWithTrait_nameItemRequestBuilderInternal instantiates a new ItemTraitsWithTrait_nameItemRequestBuilder and sets the default values.
func NewItemTraitsWithTrait_nameItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemTraitsWithTrait_nameItemRequestBuilder) {
    m := &ItemTraitsWithTrait_nameItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/custodians/{custodian_id}/traits/{trait_name}", pathParameters),
    }
    return m
}
// NewItemTraitsWithTrait_nameItemRequestBuilder instantiates a new ItemTraitsWithTrait_nameItemRequestBuilder and sets the default values.
func NewItemTraitsWithTrait_nameItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemTraitsWithTrait_nameItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemTraitsWithTrait_nameItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Get list all unique trait values for a trait name for a custodian.
// returns a Traitvaluesresponseable when successful
func (m *ItemTraitsWithTrait_nameItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemTraitsWithTrait_nameItemRequestBuilderGetRequestConfiguration)(if815b544aed335c15fd039b4dcb81545e8ef32c8003c12e1dc36ea8df9bfeb2f.Traitvaluesresponseable, error) {
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
// ToGetRequestInformation list all unique trait values for a trait name for a custodian.
// returns a *RequestInformation when successful
func (m *ItemTraitsWithTrait_nameItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemTraitsWithTrait_nameItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ItemTraitsWithTrait_nameItemRequestBuilder when successful
func (m *ItemTraitsWithTrait_nameItemRequestBuilder) WithUrl(rawUrl string)(*ItemTraitsWithTrait_nameItemRequestBuilder) {
    return NewItemTraitsWithTrait_nameItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
