package matters

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    if815b544aed335c15fd039b4dcb81545e8ef32c8003c12e1dc36ea8df9bfeb2f "github.com/xifanyan/tddc/pkg/models"
)

// ItemLegal_holdsStatsRequestBuilder builds and executes requests for operations under \matters\{matter_id}\legal_holds\stats
type ItemLegal_holdsStatsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemLegal_holdsStatsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemLegal_holdsStatsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemLegal_holdsStatsRequestBuilderInternal instantiates a new ItemLegal_holdsStatsRequestBuilder and sets the default values.
func NewItemLegal_holdsStatsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemLegal_holdsStatsRequestBuilder) {
    m := &ItemLegal_holdsStatsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/matters/{matter_id}/legal_holds/stats", pathParameters),
    }
    return m
}
// NewItemLegal_holdsStatsRequestBuilder instantiates a new ItemLegal_holdsStatsRequestBuilder and sets the default values.
func NewItemLegal_holdsStatsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemLegal_holdsStatsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemLegal_holdsStatsRequestBuilderInternal(urlParams, requestAdapter)
}
// Get show legal hold stats for a matter
// returns a Legal_hold_statsable when successful
func (m *ItemLegal_holdsStatsRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemLegal_holdsStatsRequestBuilderGetRequestConfiguration)(if815b544aed335c15fd039b4dcb81545e8ef32c8003c12e1dc36ea8df9bfeb2f.Legal_hold_statsable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, if815b544aed335c15fd039b4dcb81545e8ef32c8003c12e1dc36ea8df9bfeb2f.CreateLegal_hold_statsFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(if815b544aed335c15fd039b4dcb81545e8ef32c8003c12e1dc36ea8df9bfeb2f.Legal_hold_statsable), nil
}
// ToGetRequestInformation show legal hold stats for a matter
// returns a *RequestInformation when successful
func (m *ItemLegal_holdsStatsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemLegal_holdsStatsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ItemLegal_holdsStatsRequestBuilder when successful
func (m *ItemLegal_holdsStatsRequestBuilder) WithUrl(rawUrl string)(*ItemLegal_holdsStatsRequestBuilder) {
    return NewItemLegal_holdsStatsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
