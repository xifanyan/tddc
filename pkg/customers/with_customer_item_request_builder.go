package customers

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    if815b544aed335c15fd039b4dcb81545e8ef32c8003c12e1dc36ea8df9bfeb2f "github.com/xifanyan/tddc/pkg/models"
)

// WithCustomer_ItemRequestBuilder builds and executes requests for operations under \customers\{customer_id}
type WithCustomer_ItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// WithCustomer_ItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type WithCustomer_ItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// WithCustomer_ItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type WithCustomer_ItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// Companies the companies property
// returns a *ItemCompaniesRequestBuilder when successful
func (m *WithCustomer_ItemRequestBuilder) Companies()(*ItemCompaniesRequestBuilder) {
    return NewItemCompaniesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewWithCustomer_ItemRequestBuilderInternal instantiates a new WithCustomer_ItemRequestBuilder and sets the default values.
func NewWithCustomer_ItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*WithCustomer_ItemRequestBuilder) {
    m := &WithCustomer_ItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/customers/{customer_id}", pathParameters),
    }
    return m
}
// NewWithCustomer_ItemRequestBuilder instantiates a new WithCustomer_ItemRequestBuilder and sets the default values.
func NewWithCustomer_ItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*WithCustomer_ItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewWithCustomer_ItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Directories the directories property
// returns a *ItemDirectoriesRequestBuilder when successful
func (m *WithCustomer_ItemRequestBuilder) Directories()(*ItemDirectoriesRequestBuilder) {
    return NewItemDirectoriesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get get a single customer
// returns a Customerable when successful
func (m *WithCustomer_ItemRequestBuilder) Get(ctx context.Context, requestConfiguration *WithCustomer_ItemRequestBuilderGetRequestConfiguration)(if815b544aed335c15fd039b4dcb81545e8ef32c8003c12e1dc36ea8df9bfeb2f.Customerable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, if815b544aed335c15fd039b4dcb81545e8ef32c8003c12e1dc36ea8df9bfeb2f.CreateCustomerFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(if815b544aed335c15fd039b4dcb81545e8ef32c8003c12e1dc36ea8df9bfeb2f.Customerable), nil
}
// Patch updates a customer.To change CLO to Customer level: Set `delegate_clo` to false and `default_directory_id` to a valid directory.To change CLO to Company level: Set `delegate_clo` to true and `default_directory_id` to nil.
// returns a Customerable when successful
// returns a ApiError error when the service returns a 422 status code
// returns a ApiError error when the service returns a 500 status code
func (m *WithCustomer_ItemRequestBuilder) Patch(ctx context.Context, body ItemWithCustomer_PatchRequestBodyable, requestConfiguration *WithCustomer_ItemRequestBuilderPatchRequestConfiguration)(if815b544aed335c15fd039b4dcb81545e8ef32c8003c12e1dc36ea8df9bfeb2f.Customerable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "422": if815b544aed335c15fd039b4dcb81545e8ef32c8003c12e1dc36ea8df9bfeb2f.CreateApiErrorFromDiscriminatorValue,
        "500": if815b544aed335c15fd039b4dcb81545e8ef32c8003c12e1dc36ea8df9bfeb2f.CreateApiErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, if815b544aed335c15fd039b4dcb81545e8ef32c8003c12e1dc36ea8df9bfeb2f.CreateCustomerFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(if815b544aed335c15fd039b4dcb81545e8ef32c8003c12e1dc36ea8df9bfeb2f.Customerable), nil
}
// ToGetRequestInformation get a single customer
// returns a *RequestInformation when successful
func (m *WithCustomer_ItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *WithCustomer_ItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToPatchRequestInformation updates a customer.To change CLO to Customer level: Set `delegate_clo` to false and `default_directory_id` to a valid directory.To change CLO to Company level: Set `delegate_clo` to true and `default_directory_id` to nil.
// returns a *RequestInformation when successful
func (m *WithCustomer_ItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ItemWithCustomer_PatchRequestBodyable, requestConfiguration *WithCustomer_ItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *WithCustomer_ItemRequestBuilder when successful
func (m *WithCustomer_ItemRequestBuilder) WithUrl(rawUrl string)(*WithCustomer_ItemRequestBuilder) {
    return NewWithCustomer_ItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
