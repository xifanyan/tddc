package companies

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ItemUnique_idDry_runRequestBuilder builds and executes requests for operations under \companies\{company_id-id}\unique_id\dry_run
type ItemUnique_idDry_runRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ByKey_type gets an item from the github.com/xifanyan/tddc/pkg.companies.item.unique_id.dry_run.item collection
// returns a *ItemUnique_idDry_runWithKey_typeItemRequestBuilder when successful
func (m *ItemUnique_idDry_runRequestBuilder) ByKey_type(key_type string)(*ItemUnique_idDry_runWithKey_typeItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if key_type != "" {
        urlTplParams["key_type"] = key_type
    }
    return NewItemUnique_idDry_runWithKey_typeItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewItemUnique_idDry_runRequestBuilderInternal instantiates a new ItemUnique_idDry_runRequestBuilder and sets the default values.
func NewItemUnique_idDry_runRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemUnique_idDry_runRequestBuilder) {
    m := &ItemUnique_idDry_runRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/companies/{company_id%2Did}/unique_id/dry_run", pathParameters),
    }
    return m
}
// NewItemUnique_idDry_runRequestBuilder instantiates a new ItemUnique_idDry_runRequestBuilder and sets the default values.
func NewItemUnique_idDry_runRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemUnique_idDry_runRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemUnique_idDry_runRequestBuilderInternal(urlParams, requestAdapter)
}
