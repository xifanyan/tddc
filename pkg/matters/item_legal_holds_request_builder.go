package matters

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    if815b544aed335c15fd039b4dcb81545e8ef32c8003c12e1dc36ea8df9bfeb2f "github.com/xifanyan/tddc/pkg/models"
)

// ItemLegal_holdsRequestBuilder builds and executes requests for operations under \matters\{matter_id}\legal_holds
type ItemLegal_holdsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemLegal_holdsRequestBuilderGetQueryParameters list legal holds in a matter
type ItemLegal_holdsRequestBuilderGetQueryParameters struct {
    Limit *float64 `uriparametername:"limit"`
    Offset *float64 `uriparametername:"offset"`
}
// ItemLegal_holdsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemLegal_holdsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemLegal_holdsRequestBuilderGetQueryParameters
}
// Audits the audits property
// returns a *ItemLegal_holdsAuditsRequestBuilder when successful
func (m *ItemLegal_holdsRequestBuilder) Audits()(*ItemLegal_holdsAuditsRequestBuilder) {
    return NewItemLegal_holdsAuditsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewItemLegal_holdsRequestBuilderInternal instantiates a new ItemLegal_holdsRequestBuilder and sets the default values.
func NewItemLegal_holdsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemLegal_holdsRequestBuilder) {
    m := &ItemLegal_holdsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/matters/{matter_id}/legal_holds{?limit*,offset*}", pathParameters),
    }
    return m
}
// NewItemLegal_holdsRequestBuilder instantiates a new ItemLegal_holdsRequestBuilder and sets the default values.
func NewItemLegal_holdsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemLegal_holdsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemLegal_holdsRequestBuilderInternal(urlParams, requestAdapter)
}
// Get list legal holds in a matter
// returns a []Legal_hold_extendedable when successful
func (m *ItemLegal_holdsRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemLegal_holdsRequestBuilderGetRequestConfiguration)([]if815b544aed335c15fd039b4dcb81545e8ef32c8003c12e1dc36ea8df9bfeb2f.Legal_hold_extendedable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.SendCollection(ctx, requestInfo, if815b544aed335c15fd039b4dcb81545e8ef32c8003c12e1dc36ea8df9bfeb2f.CreateLegal_hold_extendedFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    val := make([]if815b544aed335c15fd039b4dcb81545e8ef32c8003c12e1dc36ea8df9bfeb2f.Legal_hold_extendedable, len(res))
    for i, v := range res {
        if v != nil {
            val[i] = v.(if815b544aed335c15fd039b4dcb81545e8ef32c8003c12e1dc36ea8df9bfeb2f.Legal_hold_extendedable)
        }
    }
    return val, nil
}
// Stats the stats property
// returns a *ItemLegal_holdsStatsRequestBuilder when successful
func (m *ItemLegal_holdsRequestBuilder) Stats()(*ItemLegal_holdsStatsRequestBuilder) {
    return NewItemLegal_holdsStatsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToGetRequestInformation list legal holds in a matter
// returns a *RequestInformation when successful
func (m *ItemLegal_holdsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemLegal_holdsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ItemLegal_holdsRequestBuilder when successful
func (m *ItemLegal_holdsRequestBuilder) WithUrl(rawUrl string)(*ItemLegal_holdsRequestBuilder) {
    return NewItemLegal_holdsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
