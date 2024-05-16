package companies

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    if815b544aed335c15fd039b4dcb81545e8ef32c8003c12e1dc36ea8df9bfeb2f "github.com/xifanyan/tddc/pkg/models"
)

// ItemUnique_idDry_runWithKey_typeItemRequestBuilder builds and executes requests for operations under \companies\{company_id-id}\unique_id\dry_run\{key_type}
type ItemUnique_idDry_runWithKey_typeItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemUnique_idDry_runWithKey_typeItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemUnique_idDry_runWithKey_typeItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemUnique_idDry_runWithKey_typeItemRequestBuilderInternal instantiates a new ItemUnique_idDry_runWithKey_typeItemRequestBuilder and sets the default values.
func NewItemUnique_idDry_runWithKey_typeItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemUnique_idDry_runWithKey_typeItemRequestBuilder) {
    m := &ItemUnique_idDry_runWithKey_typeItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/companies/{company_id%2Did}/unique_id/dry_run/{key_type}", pathParameters),
    }
    return m
}
// NewItemUnique_idDry_runWithKey_typeItemRequestBuilder instantiates a new ItemUnique_idDry_runWithKey_typeItemRequestBuilder and sets the default values.
func NewItemUnique_idDry_runWithKey_typeItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemUnique_idDry_runWithKey_typeItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemUnique_idDry_runWithKey_typeItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Get validate unique id conflicts
// returns a []UniqueIdConflictable when successful
func (m *ItemUnique_idDry_runWithKey_typeItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemUnique_idDry_runWithKey_typeItemRequestBuilderGetRequestConfiguration)([]if815b544aed335c15fd039b4dcb81545e8ef32c8003c12e1dc36ea8df9bfeb2f.UniqueIdConflictable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.SendCollection(ctx, requestInfo, if815b544aed335c15fd039b4dcb81545e8ef32c8003c12e1dc36ea8df9bfeb2f.CreateUniqueIdConflictFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    val := make([]if815b544aed335c15fd039b4dcb81545e8ef32c8003c12e1dc36ea8df9bfeb2f.UniqueIdConflictable, len(res))
    for i, v := range res {
        if v != nil {
            val[i] = v.(if815b544aed335c15fd039b4dcb81545e8ef32c8003c12e1dc36ea8df9bfeb2f.UniqueIdConflictable)
        }
    }
    return val, nil
}
// ToGetRequestInformation validate unique id conflicts
// returns a *RequestInformation when successful
func (m *ItemUnique_idDry_runWithKey_typeItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemUnique_idDry_runWithKey_typeItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ItemUnique_idDry_runWithKey_typeItemRequestBuilder when successful
func (m *ItemUnique_idDry_runWithKey_typeItemRequestBuilder) WithUrl(rawUrl string)(*ItemUnique_idDry_runWithKey_typeItemRequestBuilder) {
    return NewItemUnique_idDry_runWithKey_typeItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
