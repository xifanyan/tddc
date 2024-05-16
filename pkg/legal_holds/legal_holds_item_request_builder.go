package legal_holds

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    if815b544aed335c15fd039b4dcb81545e8ef32c8003c12e1dc36ea8df9bfeb2f "github.com/xifanyan/tddc/pkg/models"
)

// Legal_holdsItemRequestBuilder builds and executes requests for operations under \legal_holds\{id}
type Legal_holdsItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// Legal_holdsItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type Legal_holdsItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// Audits the audits property
// returns a *ItemAuditsRequestBuilder when successful
func (m *Legal_holdsItemRequestBuilder) Audits()(*ItemAuditsRequestBuilder) {
    return NewItemAuditsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewLegal_holdsItemRequestBuilderInternal instantiates a new Legal_holdsItemRequestBuilder and sets the default values.
func NewLegal_holdsItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*Legal_holdsItemRequestBuilder) {
    m := &Legal_holdsItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/legal_holds/{id}", pathParameters),
    }
    return m
}
// NewLegal_holdsItemRequestBuilder instantiates a new Legal_holdsItemRequestBuilder and sets the default values.
func NewLegal_holdsItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*Legal_holdsItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewLegal_holdsItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Get display legal hold
// returns a Legal_hold_extendedable when successful
func (m *Legal_holdsItemRequestBuilder) Get(ctx context.Context, requestConfiguration *Legal_holdsItemRequestBuilderGetRequestConfiguration)(if815b544aed335c15fd039b4dcb81545e8ef32c8003c12e1dc36ea8df9bfeb2f.Legal_hold_extendedable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, if815b544aed335c15fd039b4dcb81545e8ef32c8003c12e1dc36ea8df9bfeb2f.CreateLegal_hold_extendedFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(if815b544aed335c15fd039b4dcb81545e8ef32c8003c12e1dc36ea8df9bfeb2f.Legal_hold_extendedable), nil
}
// ToGetRequestInformation display legal hold
// returns a *RequestInformation when successful
func (m *Legal_holdsItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *Legal_holdsItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *Legal_holdsItemRequestBuilder when successful
func (m *Legal_holdsItemRequestBuilder) WithUrl(rawUrl string)(*Legal_holdsItemRequestBuilder) {
    return NewLegal_holdsItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
