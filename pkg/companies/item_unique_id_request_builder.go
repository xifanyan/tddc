package companies

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ItemUnique_idRequestBuilder builds and executes requests for operations under \companies\{company_id-id}\unique_id
type ItemUnique_idRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// NewItemUnique_idRequestBuilderInternal instantiates a new ItemUnique_idRequestBuilder and sets the default values.
func NewItemUnique_idRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemUnique_idRequestBuilder) {
    m := &ItemUnique_idRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/companies/{company_id%2Did}/unique_id", pathParameters),
    }
    return m
}
// NewItemUnique_idRequestBuilder instantiates a new ItemUnique_idRequestBuilder and sets the default values.
func NewItemUnique_idRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemUnique_idRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemUnique_idRequestBuilderInternal(urlParams, requestAdapter)
}
// Dry_run the dry_run property
// returns a *ItemUnique_idDry_runRequestBuilder when successful
func (m *ItemUnique_idRequestBuilder) Dry_run()(*ItemUnique_idDry_runRequestBuilder) {
    return NewItemUnique_idDry_runRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
