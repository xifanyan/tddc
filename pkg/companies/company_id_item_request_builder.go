package companies

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    if815b544aed335c15fd039b4dcb81545e8ef32c8003c12e1dc36ea8df9bfeb2f "github.com/xifanyan/tddc/pkg/models"
)

// Company_idItemRequestBuilder builds and executes requests for operations under \companies\{company_id-id}
type Company_idItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// Company_idItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type Company_idItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// Company_idItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type Company_idItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewCompany_idItemRequestBuilderInternal instantiates a new Company_idItemRequestBuilder and sets the default values.
func NewCompany_idItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*Company_idItemRequestBuilder) {
    m := &Company_idItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/companies/{company_id%2Did}", pathParameters),
    }
    return m
}
// NewCompany_idItemRequestBuilder instantiates a new Company_idItemRequestBuilder and sets the default values.
func NewCompany_idItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*Company_idItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewCompany_idItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Custodians the custodians property
// returns a *ItemCustodiansRequestBuilder when successful
func (m *Company_idItemRequestBuilder) Custodians()(*ItemCustodiansRequestBuilder) {
    return NewItemCustodiansRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Directory the directory property
// returns a *ItemDirectoryRequestBuilder when successful
func (m *Company_idItemRequestBuilder) Directory()(*ItemDirectoryRequestBuilder) {
    return NewItemDirectoryRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get get a single company
// returns a Companiesable when successful
func (m *Company_idItemRequestBuilder) Get(ctx context.Context, requestConfiguration *Company_idItemRequestBuilderGetRequestConfiguration)(if815b544aed335c15fd039b4dcb81545e8ef32c8003c12e1dc36ea8df9bfeb2f.Companiesable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, if815b544aed335c15fd039b4dcb81545e8ef32c8003c12e1dc36ea8df9bfeb2f.CreateCompaniesFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(if815b544aed335c15fd039b4dcb81545e8ef32c8003c12e1dc36ea8df9bfeb2f.Companiesable), nil
}
// Legal_holds the legal_holds property
// returns a *ItemLegal_holdsRequestBuilder when successful
func (m *Company_idItemRequestBuilder) Legal_holds()(*ItemLegal_holdsRequestBuilder) {
    return NewItemLegal_holdsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Matters the matters property
// returns a *ItemMattersRequestBuilder when successful
func (m *Company_idItemRequestBuilder) Matters()(*ItemMattersRequestBuilder) {
    return NewItemMattersRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Patch update Companies
// returns a Companiesable when successful
func (m *Company_idItemRequestBuilder) Patch(ctx context.Context, body ItemCompany_idPatchRequestBodyable, requestConfiguration *Company_idItemRequestBuilderPatchRequestConfiguration)(if815b544aed335c15fd039b4dcb81545e8ef32c8003c12e1dc36ea8df9bfeb2f.Companiesable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, if815b544aed335c15fd039b4dcb81545e8ef32c8003c12e1dc36ea8df9bfeb2f.CreateCompaniesFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(if815b544aed335c15fd039b4dcb81545e8ef32c8003c12e1dc36ea8df9bfeb2f.Companiesable), nil
}
// ToGetRequestInformation get a single company
// returns a *RequestInformation when successful
func (m *Company_idItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *Company_idItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToPatchRequestInformation update Companies
// returns a *RequestInformation when successful
func (m *Company_idItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ItemCompany_idPatchRequestBodyable, requestConfiguration *Company_idItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Unique_id the unique_id property
// returns a *ItemUnique_idRequestBuilder when successful
func (m *Company_idItemRequestBuilder) Unique_id()(*ItemUnique_idRequestBuilder) {
    return NewItemUnique_idRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *Company_idItemRequestBuilder when successful
func (m *Company_idItemRequestBuilder) WithUrl(rawUrl string)(*Company_idItemRequestBuilder) {
    return NewCompany_idItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
