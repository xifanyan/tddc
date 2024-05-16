package matters

import (
    i53ac87e8cb3cc9276228f74d38694a208cacb99bb8ceb705eeae99fb88d4d274 "strconv"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// MattersRequestBuilder builds and executes requests for operations under \matters
type MattersRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ByMatter_id gets an item from the github.com/xifanyan/tddc/pkg.matters.item collection
// Deprecated: This indexer is deprecated and will be removed in the next major version. Use the one with the typed parameter instead.
// returns a *WithMatter_ItemRequestBuilder when successful
func (m *MattersRequestBuilder) ByMatter_id(matter_id string)(*WithMatter_ItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if matter_id != "" {
        urlTplParams["matter_id"] = matter_id
    }
    return NewWithMatter_ItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// ByMatter_idInteger gets an item from the github.com/xifanyan/tddc/pkg.matters.item collection
// returns a *WithMatter_ItemRequestBuilder when successful
func (m *MattersRequestBuilder) ByMatter_idInteger(matter_id int32)(*WithMatter_ItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    urlTplParams["matter_id"] = i53ac87e8cb3cc9276228f74d38694a208cacb99bb8ceb705eeae99fb88d4d274.FormatInt(int64(matter_id), 10)
    return NewWithMatter_ItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewMattersRequestBuilderInternal instantiates a new MattersRequestBuilder and sets the default values.
func NewMattersRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MattersRequestBuilder) {
    m := &MattersRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/matters", pathParameters),
    }
    return m
}
// NewMattersRequestBuilder instantiates a new MattersRequestBuilder and sets the default values.
func NewMattersRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MattersRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewMattersRequestBuilderInternal(urlParams, requestAdapter)
}
