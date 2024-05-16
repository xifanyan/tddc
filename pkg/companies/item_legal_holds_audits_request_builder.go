package companies

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    if815b544aed335c15fd039b4dcb81545e8ef32c8003c12e1dc36ea8df9bfeb2f "github.com/xifanyan/tddc/pkg/models"
)

// ItemLegal_holdsAuditsRequestBuilder builds and executes requests for operations under \companies\{company_id-id}\legal_holds\audits
type ItemLegal_holdsAuditsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemLegal_holdsAuditsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemLegal_holdsAuditsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemLegal_holdsAuditsRequestBuilderInternal instantiates a new ItemLegal_holdsAuditsRequestBuilder and sets the default values.
func NewItemLegal_holdsAuditsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemLegal_holdsAuditsRequestBuilder) {
    m := &ItemLegal_holdsAuditsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/companies/{company_id%2Did}/legal_holds/audits", pathParameters),
    }
    return m
}
// NewItemLegal_holdsAuditsRequestBuilder instantiates a new ItemLegal_holdsAuditsRequestBuilder and sets the default values.
func NewItemLegal_holdsAuditsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemLegal_holdsAuditsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemLegal_holdsAuditsRequestBuilderInternal(urlParams, requestAdapter)
}
// Get list legal hold status audits in a company
// returns a []Legal_hold_status_auditable when successful
func (m *ItemLegal_holdsAuditsRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemLegal_holdsAuditsRequestBuilderGetRequestConfiguration)([]if815b544aed335c15fd039b4dcb81545e8ef32c8003c12e1dc36ea8df9bfeb2f.Legal_hold_status_auditable, error) {
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
// ToGetRequestInformation list legal hold status audits in a company
// returns a *RequestInformation when successful
func (m *ItemLegal_holdsAuditsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemLegal_holdsAuditsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ItemLegal_holdsAuditsRequestBuilder when successful
func (m *ItemLegal_holdsAuditsRequestBuilder) WithUrl(rawUrl string)(*ItemLegal_holdsAuditsRequestBuilder) {
    return NewItemLegal_holdsAuditsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
