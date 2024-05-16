package matters

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    if815b544aed335c15fd039b4dcb81545e8ef32c8003c12e1dc36ea8df9bfeb2f "github.com/xifanyan/tddc/pkg/models"
)

// WithMatter_ItemRequestBuilder builds and executes requests for operations under \matters\{matter_id}
type WithMatter_ItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// WithMatter_ItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type WithMatter_ItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// WithMatter_ItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type WithMatter_ItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// Collections the collections property
// returns a *ItemCollectionsRequestBuilder when successful
func (m *WithMatter_ItemRequestBuilder) Collections()(*ItemCollectionsRequestBuilder) {
    return NewItemCollectionsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewWithMatter_ItemRequestBuilderInternal instantiates a new WithMatter_ItemRequestBuilder and sets the default values.
func NewWithMatter_ItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*WithMatter_ItemRequestBuilder) {
    m := &WithMatter_ItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/matters/{matter_id}", pathParameters),
    }
    return m
}
// NewWithMatter_ItemRequestBuilder instantiates a new WithMatter_ItemRequestBuilder and sets the default values.
func NewWithMatter_ItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*WithMatter_ItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewWithMatter_ItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Custodians the custodians property
// returns a *ItemCustodiansRequestBuilder when successful
func (m *WithMatter_ItemRequestBuilder) Custodians()(*ItemCustodiansRequestBuilder) {
    return NewItemCustodiansRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get get a single matter
// returns a Matterable when successful
func (m *WithMatter_ItemRequestBuilder) Get(ctx context.Context, requestConfiguration *WithMatter_ItemRequestBuilderGetRequestConfiguration)(if815b544aed335c15fd039b4dcb81545e8ef32c8003c12e1dc36ea8df9bfeb2f.Matterable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, if815b544aed335c15fd039b4dcb81545e8ef32c8003c12e1dc36ea8df9bfeb2f.CreateMatterFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(if815b544aed335c15fd039b4dcb81545e8ef32c8003c12e1dc36ea8df9bfeb2f.Matterable), nil
}
// Legal_holds the legal_holds property
// returns a *ItemLegal_holdsRequestBuilder when successful
func (m *WithMatter_ItemRequestBuilder) Legal_holds()(*ItemLegal_holdsRequestBuilder) {
    return NewItemLegal_holdsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Patch update a matter
// returns a []byte when successful
// returns a ItemWithMatter_4XXError error when the service returns a 4XX status code
// returns a ItemWithMatter_5XXError error when the service returns a 5XX status code
func (m *WithMatter_ItemRequestBuilder) Patch(ctx context.Context, body ItemWithMatter_PatchRequestBodyable, requestConfiguration *WithMatter_ItemRequestBuilderPatchRequestConfiguration)([]byte, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": CreateItemWithMatter_4XXErrorFromDiscriminatorValue,
        "5XX": CreateItemWithMatter_5XXErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.SendPrimitive(ctx, requestInfo, "[]byte", errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.([]byte), nil
}
// ToGetRequestInformation get a single matter
// returns a *RequestInformation when successful
func (m *WithMatter_ItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *WithMatter_ItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToPatchRequestInformation update a matter
// returns a *RequestInformation when successful
func (m *WithMatter_ItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ItemWithMatter_PatchRequestBodyable, requestConfiguration *WithMatter_ItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *WithMatter_ItemRequestBuilder when successful
func (m *WithMatter_ItemRequestBuilder) WithUrl(rawUrl string)(*WithMatter_ItemRequestBuilder) {
    return NewWithMatter_ItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
