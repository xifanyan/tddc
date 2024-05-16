package customers

import (
    "context"
    i53ac87e8cb3cc9276228f74d38694a208cacb99bb8ceb705eeae99fb88d4d274 "strconv"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    if815b544aed335c15fd039b4dcb81545e8ef32c8003c12e1dc36ea8df9bfeb2f "github.com/xifanyan/tddc/pkg/models"
)

// ItemDirectoriesRequestBuilder builds and executes requests for operations under \customers\{customer_id}\directories
type ItemDirectoriesRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemDirectoriesRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemDirectoriesRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ItemDirectoriesRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemDirectoriesRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ByDirectory_id gets an item from the github.com/xifanyan/tddc/pkg.customers.item.directories.item collection
// Deprecated: This indexer is deprecated and will be removed in the next major version. Use the one with the typed parameter instead.
// returns a *ItemDirectoriesWithDirectory_ItemRequestBuilder when successful
func (m *ItemDirectoriesRequestBuilder) ByDirectory_id(directory_id string)(*ItemDirectoriesWithDirectory_ItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if directory_id != "" {
        urlTplParams["directory_id"] = directory_id
    }
    return NewItemDirectoriesWithDirectory_ItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// ByDirectory_idInteger gets an item from the github.com/xifanyan/tddc/pkg.customers.item.directories.item collection
// returns a *ItemDirectoriesWithDirectory_ItemRequestBuilder when successful
func (m *ItemDirectoriesRequestBuilder) ByDirectory_idInteger(directory_id int32)(*ItemDirectoriesWithDirectory_ItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    urlTplParams["directory_id"] = i53ac87e8cb3cc9276228f74d38694a208cacb99bb8ceb705eeae99fb88d4d274.FormatInt(int64(directory_id), 10)
    return NewItemDirectoriesWithDirectory_ItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewItemDirectoriesRequestBuilderInternal instantiates a new ItemDirectoriesRequestBuilder and sets the default values.
func NewItemDirectoriesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemDirectoriesRequestBuilder) {
    m := &ItemDirectoriesRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/customers/{customer_id}/directories", pathParameters),
    }
    return m
}
// NewItemDirectoriesRequestBuilder instantiates a new ItemDirectoriesRequestBuilder and sets the default values.
func NewItemDirectoriesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemDirectoriesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemDirectoriesRequestBuilderInternal(urlParams, requestAdapter)
}
// Get list all directories under a customer.
// returns a []Directoryable when successful
func (m *ItemDirectoriesRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemDirectoriesRequestBuilderGetRequestConfiguration)([]if815b544aed335c15fd039b4dcb81545e8ef32c8003c12e1dc36ea8df9bfeb2f.Directoryable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.SendCollection(ctx, requestInfo, if815b544aed335c15fd039b4dcb81545e8ef32c8003c12e1dc36ea8df9bfeb2f.CreateDirectoryFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    val := make([]if815b544aed335c15fd039b4dcb81545e8ef32c8003c12e1dc36ea8df9bfeb2f.Directoryable, len(res))
    for i, v := range res {
        if v != nil {
            val[i] = v.(if815b544aed335c15fd039b4dcb81545e8ef32c8003c12e1dc36ea8df9bfeb2f.Directoryable)
        }
    }
    return val, nil
}
// Post creates a directory under the customer.Requires admin rights.Directories created in this way are "freestanding" and do not belong to a specific company, nor are they set asthe default directory used for the customer-level CLO.To attach this directory to a company, pass the returned id to the company as its directory_id.To make this directory the default directory for customer delegation, pass the returned id to the customer as itsdefault_directory_id.
// returns a Directoryable when successful
// returns a ApiError error when the service returns a 422 status code
func (m *ItemDirectoriesRequestBuilder) Post(ctx context.Context, body i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.MultipartBody, requestConfiguration *ItemDirectoriesRequestBuilderPostRequestConfiguration)(if815b544aed335c15fd039b4dcb81545e8ef32c8003c12e1dc36ea8df9bfeb2f.Directoryable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "422": if815b544aed335c15fd039b4dcb81545e8ef32c8003c12e1dc36ea8df9bfeb2f.CreateApiErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, if815b544aed335c15fd039b4dcb81545e8ef32c8003c12e1dc36ea8df9bfeb2f.CreateDirectoryFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(if815b544aed335c15fd039b4dcb81545e8ef32c8003c12e1dc36ea8df9bfeb2f.Directoryable), nil
}
// ToGetRequestInformation list all directories under a customer.
// returns a *RequestInformation when successful
func (m *ItemDirectoriesRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemDirectoriesRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToPostRequestInformation creates a directory under the customer.Requires admin rights.Directories created in this way are "freestanding" and do not belong to a specific company, nor are they set asthe default directory used for the customer-level CLO.To attach this directory to a company, pass the returned id to the company as its directory_id.To make this directory the default directory for customer delegation, pass the returned id to the customer as itsdefault_directory_id.
// returns a *RequestInformation when successful
func (m *ItemDirectoriesRequestBuilder) ToPostRequestInformation(ctx context.Context, body i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.MultipartBody, requestConfiguration *ItemDirectoriesRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "multipart/form-data", body)
    if err != nil {
        return nil, err
    }
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ItemDirectoriesRequestBuilder when successful
func (m *ItemDirectoriesRequestBuilder) WithUrl(rawUrl string)(*ItemDirectoriesRequestBuilder) {
    return NewItemDirectoriesRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
