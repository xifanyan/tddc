package companies

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    if815b544aed335c15fd039b4dcb81545e8ef32c8003c12e1dc36ea8df9bfeb2f "github.com/xifanyan/tddc/pkg/models"
)

// ItemCustodiansRequestBuilder builds and executes requests for operations under \companies\{company_id-id}\custodians
type ItemCustodiansRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemCustodiansRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemCustodiansRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ItemCustodiansRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemCustodiansRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemCustodiansRequestBuilderInternal instantiates a new ItemCustodiansRequestBuilder and sets the default values.
func NewItemCustodiansRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemCustodiansRequestBuilder) {
    m := &ItemCustodiansRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/companies/{company_id%2Did}/custodians", pathParameters),
    }
    return m
}
// NewItemCustodiansRequestBuilder instantiates a new ItemCustodiansRequestBuilder and sets the default values.
func NewItemCustodiansRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemCustodiansRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemCustodiansRequestBuilderInternal(urlParams, requestAdapter)
}
// Get list all custodians in a company
// returns a []Custodianable when successful
func (m *ItemCustodiansRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemCustodiansRequestBuilderGetRequestConfiguration)([]if815b544aed335c15fd039b4dcb81545e8ef32c8003c12e1dc36ea8df9bfeb2f.Custodianable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.SendCollection(ctx, requestInfo, if815b544aed335c15fd039b4dcb81545e8ef32c8003c12e1dc36ea8df9bfeb2f.CreateCustodianFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    val := make([]if815b544aed335c15fd039b4dcb81545e8ef32c8003c12e1dc36ea8df9bfeb2f.Custodianable, len(res))
    for i, v := range res {
        if v != nil {
            val[i] = v.(if815b544aed335c15fd039b4dcb81545e8ef32c8003c12e1dc36ea8df9bfeb2f.Custodianable)
        }
    }
    return val, nil
}
// Locate the locate property
// returns a *ItemCustodiansLocateRequestBuilder when successful
func (m *ItemCustodiansRequestBuilder) Locate()(*ItemCustodiansLocateRequestBuilder) {
    return NewItemCustodiansLocateRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Post synchronizes Custodians to the Directory associated with the Company.Custodians are uniquely identified by the `unique_id_type` of the associated Directory. If no `unique_id_type` is set on the Directory, the sync operation defaults to the Custodian's `email` as its unique id. (Currently, this data can only be retrieved from the TotalDiscovery Web UI.)When the sync operation occurs, TotalDiscovery tries to match each supplied Custodian by its unique id to an existing Custodian in the requisite Directory. If a Custodian is found, all submitted attributes for that Custodian are updated to the supplied values, leaving other attributes unchanged. If no Custodian is found, then a new Custodian is created with the submitted attributes.
// returns a []byte when successful
// returns a CustodianSyncError error when the service returns a 422 status code
func (m *ItemCustodiansRequestBuilder) Post(ctx context.Context, body ItemCustodiansPostRequestBodyable, requestConfiguration *ItemCustodiansRequestBuilderPostRequestConfiguration)([]byte, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "422": if815b544aed335c15fd039b4dcb81545e8ef32c8003c12e1dc36ea8df9bfeb2f.CreateCustodianSyncErrorFromDiscriminatorValue,
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
// ToGetRequestInformation list all custodians in a company
// returns a *RequestInformation when successful
func (m *ItemCustodiansRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemCustodiansRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToPostRequestInformation synchronizes Custodians to the Directory associated with the Company.Custodians are uniquely identified by the `unique_id_type` of the associated Directory. If no `unique_id_type` is set on the Directory, the sync operation defaults to the Custodian's `email` as its unique id. (Currently, this data can only be retrieved from the TotalDiscovery Web UI.)When the sync operation occurs, TotalDiscovery tries to match each supplied Custodian by its unique id to an existing Custodian in the requisite Directory. If a Custodian is found, all submitted attributes for that Custodian are updated to the supplied values, leaving other attributes unchanged. If no Custodian is found, then a new Custodian is created with the submitted attributes.
// returns a *RequestInformation when successful
func (m *ItemCustodiansRequestBuilder) ToPostRequestInformation(ctx context.Context, body ItemCustodiansPostRequestBodyable, requestConfiguration *ItemCustodiansRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Traits the traits property
// returns a *ItemCustodiansTraitsRequestBuilder when successful
func (m *ItemCustodiansRequestBuilder) Traits()(*ItemCustodiansTraitsRequestBuilder) {
    return NewItemCustodiansTraitsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ItemCustodiansRequestBuilder when successful
func (m *ItemCustodiansRequestBuilder) WithUrl(rawUrl string)(*ItemCustodiansRequestBuilder) {
    return NewItemCustodiansRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
