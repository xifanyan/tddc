package collections

import (
    i53ac87e8cb3cc9276228f74d38694a208cacb99bb8ceb705eeae99fb88d4d274 "strconv"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// CollectionsRequestBuilder builds and executes requests for operations under \collections
type CollectionsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ByCollection_id gets an item from the github.com/xifanyan/tddc/pkg.collections.item collection
// Deprecated: This indexer is deprecated and will be removed in the next major version. Use the one with the typed parameter instead.
// returns a *WithCollection_ItemRequestBuilder when successful
func (m *CollectionsRequestBuilder) ByCollection_id(collection_id string)(*WithCollection_ItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if collection_id != "" {
        urlTplParams["collection_id"] = collection_id
    }
    return NewWithCollection_ItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// ByCollection_idInteger gets an item from the github.com/xifanyan/tddc/pkg.collections.item collection
// returns a *WithCollection_ItemRequestBuilder when successful
func (m *CollectionsRequestBuilder) ByCollection_idInteger(collection_id int32)(*WithCollection_ItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    urlTplParams["collection_id"] = i53ac87e8cb3cc9276228f74d38694a208cacb99bb8ceb705eeae99fb88d4d274.FormatInt(int64(collection_id), 10)
    return NewWithCollection_ItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewCollectionsRequestBuilderInternal instantiates a new CollectionsRequestBuilder and sets the default values.
func NewCollectionsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CollectionsRequestBuilder) {
    m := &CollectionsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/collections", pathParameters),
    }
    return m
}
// NewCollectionsRequestBuilder instantiates a new CollectionsRequestBuilder and sets the default values.
func NewCollectionsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CollectionsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewCollectionsRequestBuilderInternal(urlParams, requestAdapter)
}
