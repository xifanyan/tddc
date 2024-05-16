package legal_holds

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    if815b544aed335c15fd039b4dcb81545e8ef32c8003c12e1dc36ea8df9bfeb2f "github.com/xifanyan/tddc/pkg/models"
)

// ItemAuditsRequestBuilder builds and executes requests for operations under \legal_holds\{id}\audits
type ItemAuditsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemAuditsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemAuditsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemAuditsRequestBuilderInternal instantiates a new ItemAuditsRequestBuilder and sets the default values.
func NewItemAuditsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemAuditsRequestBuilder) {
    m := &ItemAuditsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/legal_holds/{id}/audits", pathParameters),
    }
    return m
}
// NewItemAuditsRequestBuilder instantiates a new ItemAuditsRequestBuilder and sets the default values.
func NewItemAuditsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemAuditsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemAuditsRequestBuilderInternal(urlParams, requestAdapter)
}
// Get list legal hold status audits for a legal hold
// returns a []Legal_hold_status_auditable when successful
func (m *ItemAuditsRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemAuditsRequestBuilderGetRequestConfiguration)([]if815b544aed335c15fd039b4dcb81545e8ef32c8003c12e1dc36ea8df9bfeb2f.Legal_hold_status_auditable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.SendCollection(ctx, requestInfo, if815b544aed335c15fd039b4dcb81545e8ef32c8003c12e1dc36ea8df9bfeb2f.CreateLegal_hold_status_auditFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    val := make([]if815b544aed335c15fd039b4dcb81545e8ef32c8003c12e1dc36ea8df9bfeb2f.Legal_hold_status_auditable, len(res))
    for i, v := range res {
        if v != nil {
            val[i] = v.(if815b544aed335c15fd039b4dcb81545e8ef32c8003c12e1dc36ea8df9bfeb2f.Legal_hold_status_auditable)
        }
    }
    return val, nil
}
// ToGetRequestInformation list legal hold status audits for a legal hold
// returns a *RequestInformation when successful
func (m *ItemAuditsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemAuditsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ItemAuditsRequestBuilder when successful
func (m *ItemAuditsRequestBuilder) WithUrl(rawUrl string)(*ItemAuditsRequestBuilder) {
    return NewItemAuditsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
