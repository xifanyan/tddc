package searches

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// SearchesRequestBuilder builds and executes requests for operations under \searches
type SearchesRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// BySearch_id gets an item from the github.com/xifanyan/tddc/pkg.searches.item collection
// returns a *WithSearch_ItemRequestBuilder when successful
func (m *SearchesRequestBuilder) BySearch_id(search_id string)(*WithSearch_ItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if search_id != "" {
        urlTplParams["search_id"] = search_id
    }
    return NewWithSearch_ItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewSearchesRequestBuilderInternal instantiates a new SearchesRequestBuilder and sets the default values.
func NewSearchesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*SearchesRequestBuilder) {
    m := &SearchesRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/searches", pathParameters),
    }
    return m
}
// NewSearchesRequestBuilder instantiates a new SearchesRequestBuilder and sets the default values.
func NewSearchesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*SearchesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewSearchesRequestBuilderInternal(urlParams, requestAdapter)
}
