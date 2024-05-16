package companies

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    if815b544aed335c15fd039b4dcb81545e8ef32c8003c12e1dc36ea8df9bfeb2f "github.com/xifanyan/tddc/pkg/models"
)

// ItemCustodiansTraitsItemSyncRequestBuilder builds and executes requests for operations under \companies\{company_id-id}\custodians\traits\{trait_name}\sync
type ItemCustodiansTraitsItemSyncRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemCustodiansTraitsItemSyncRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemCustodiansTraitsItemSyncRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemCustodiansTraitsItemSyncRequestBuilderInternal instantiates a new ItemCustodiansTraitsItemSyncRequestBuilder and sets the default values.
func NewItemCustodiansTraitsItemSyncRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemCustodiansTraitsItemSyncRequestBuilder) {
    m := &ItemCustodiansTraitsItemSyncRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/companies/{company_id%2Did}/custodians/traits/{trait_name}/sync", pathParameters),
    }
    return m
}
// NewItemCustodiansTraitsItemSyncRequestBuilder instantiates a new ItemCustodiansTraitsItemSyncRequestBuilder and sets the default values.
func NewItemCustodiansTraitsItemSyncRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemCustodiansTraitsItemSyncRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemCustodiansTraitsItemSyncRequestBuilderInternal(urlParams, requestAdapter)
}
// Post sync traits for a company using custodian's unique_id and a given trait name.
// Deprecated: This method is obsolete. Use PostAsSyncPostResponse instead.
// returns a ItemCustodiansTraitsItemSyncResponseable when successful
// returns a TraitSyncError error when the service returns a 422 status code
func (m *ItemCustodiansTraitsItemSyncRequestBuilder) Post(ctx context.Context, body if815b544aed335c15fd039b4dcb81545e8ef32c8003c12e1dc36ea8df9bfeb2f.Synctraitrequestable, requestConfiguration *ItemCustodiansTraitsItemSyncRequestBuilderPostRequestConfiguration)(ItemCustodiansTraitsItemSyncResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "422": if815b544aed335c15fd039b4dcb81545e8ef32c8003c12e1dc36ea8df9bfeb2f.CreateTraitSyncErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateItemCustodiansTraitsItemSyncResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ItemCustodiansTraitsItemSyncResponseable), nil
}
// PostAsSyncPostResponse sync traits for a company using custodian's unique_id and a given trait name.
// returns a ItemCustodiansTraitsItemSyncPostResponseable when successful
// returns a TraitSyncError error when the service returns a 422 status code
func (m *ItemCustodiansTraitsItemSyncRequestBuilder) PostAsSyncPostResponse(ctx context.Context, body if815b544aed335c15fd039b4dcb81545e8ef32c8003c12e1dc36ea8df9bfeb2f.Synctraitrequestable, requestConfiguration *ItemCustodiansTraitsItemSyncRequestBuilderPostRequestConfiguration)(ItemCustodiansTraitsItemSyncPostResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "422": if815b544aed335c15fd039b4dcb81545e8ef32c8003c12e1dc36ea8df9bfeb2f.CreateTraitSyncErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateItemCustodiansTraitsItemSyncPostResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ItemCustodiansTraitsItemSyncPostResponseable), nil
}
// ToPostRequestInformation sync traits for a company using custodian's unique_id and a given trait name.
// returns a *RequestInformation when successful
func (m *ItemCustodiansTraitsItemSyncRequestBuilder) ToPostRequestInformation(ctx context.Context, body if815b544aed335c15fd039b4dcb81545e8ef32c8003c12e1dc36ea8df9bfeb2f.Synctraitrequestable, requestConfiguration *ItemCustodiansTraitsItemSyncRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
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
// returns a *ItemCustodiansTraitsItemSyncRequestBuilder when successful
func (m *ItemCustodiansTraitsItemSyncRequestBuilder) WithUrl(rawUrl string)(*ItemCustodiansTraitsItemSyncRequestBuilder) {
    return NewItemCustodiansTraitsItemSyncRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
