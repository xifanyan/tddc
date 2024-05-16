package custodians

import (
    i53ac87e8cb3cc9276228f74d38694a208cacb99bb8ceb705eeae99fb88d4d274 "strconv"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// CustodiansRequestBuilder builds and executes requests for operations under \custodians
type CustodiansRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ByCustodian_id gets an item from the github.com/xifanyan/tddc/pkg.custodians.item collection
// Deprecated: This indexer is deprecated and will be removed in the next major version. Use the one with the typed parameter instead.
// returns a *WithCustodian_ItemRequestBuilder when successful
func (m *CustodiansRequestBuilder) ByCustodian_id(custodian_id string)(*WithCustodian_ItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if custodian_id != "" {
        urlTplParams["custodian_id"] = custodian_id
    }
    return NewWithCustodian_ItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// ByCustodian_idInteger gets an item from the github.com/xifanyan/tddc/pkg.custodians.item collection
// returns a *WithCustodian_ItemRequestBuilder when successful
func (m *CustodiansRequestBuilder) ByCustodian_idInteger(custodian_id int32)(*WithCustodian_ItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    urlTplParams["custodian_id"] = i53ac87e8cb3cc9276228f74d38694a208cacb99bb8ceb705eeae99fb88d4d274.FormatInt(int64(custodian_id), 10)
    return NewWithCustodian_ItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewCustodiansRequestBuilderInternal instantiates a new CustodiansRequestBuilder and sets the default values.
func NewCustodiansRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CustodiansRequestBuilder) {
    m := &CustodiansRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/custodians", pathParameters),
    }
    return m
}
// NewCustodiansRequestBuilder instantiates a new CustodiansRequestBuilder and sets the default values.
func NewCustodiansRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CustodiansRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewCustodiansRequestBuilderInternal(urlParams, requestAdapter)
}
