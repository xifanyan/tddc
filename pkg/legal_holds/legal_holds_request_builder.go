package legal_holds

import (
    i53ac87e8cb3cc9276228f74d38694a208cacb99bb8ceb705eeae99fb88d4d274 "strconv"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// Legal_holdsRequestBuilder builds and executes requests for operations under \legal_holds
type Legal_holdsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ById gets an item from the github.com/xifanyan/tddc/pkg.legal_holds.item collection
// Deprecated: This indexer is deprecated and will be removed in the next major version. Use the one with the typed parameter instead.
// returns a *Legal_holdsItemRequestBuilder when successful
func (m *Legal_holdsRequestBuilder) ById(id string)(*Legal_holdsItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["id"] = id
    }
    return NewLegal_holdsItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// ByIdInteger gets an item from the github.com/xifanyan/tddc/pkg.legal_holds.item collection
// returns a *Legal_holdsItemRequestBuilder when successful
func (m *Legal_holdsRequestBuilder) ByIdInteger(id int32)(*Legal_holdsItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    urlTplParams["id"] = i53ac87e8cb3cc9276228f74d38694a208cacb99bb8ceb705eeae99fb88d4d274.FormatInt(int64(id), 10)
    return NewLegal_holdsItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewLegal_holdsRequestBuilderInternal instantiates a new Legal_holdsRequestBuilder and sets the default values.
func NewLegal_holdsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*Legal_holdsRequestBuilder) {
    m := &Legal_holdsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/legal_holds", pathParameters),
    }
    return m
}
// NewLegal_holdsRequestBuilder instantiates a new Legal_holdsRequestBuilder and sets the default values.
func NewLegal_holdsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*Legal_holdsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewLegal_holdsRequestBuilderInternal(urlParams, requestAdapter)
}
