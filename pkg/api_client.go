package pkg

import (
    i25911dc319edd61cbac496af7eab5ef20b6069a42515e22ec6a9bc97bf598488 "github.com/microsoft/kiota-serialization-json-go"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i4bcdc892e61ac17e2afc10b5e2b536b29f4fd6c1ad30f4a5a68df47495db3347 "github.com/microsoft/kiota-serialization-form-go"
    i56887720f41ac882814261620b1c8459c4a992a0207af547c4453dd39fabc426 "github.com/microsoft/kiota-serialization-multipart-go"
    i7294a22093d408fdca300f11b81a887d89c47b764af06c8b803e2323973fdb83 "github.com/microsoft/kiota-serialization-text-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    i193388f9f459f47d2c8274c8ac796c8f50986cfd778a4c8d70f6cdf03b70cf69 "github.com/xifanyan/tddc/pkg/collections"
    i1de8afbc6843ca76cd39c931c5b733e39dbefccfeecbbe14bda2f40fad45ecb4 "github.com/xifanyan/tddc/pkg/matters"
    i3d76b5e22f38add1010d104e0b80c877b70818b18473d854a923d9019345c26b "github.com/xifanyan/tddc/pkg/legal_holds"
    i45c651b865c1f94388852d91754503c3df6c397790455fddbfcc199084e7560d "github.com/xifanyan/tddc/pkg/custodians"
    i4f8be26726d1e5cf99e42978603d7d3ca9281228732ec1fed238bee11dcf4f70 "github.com/xifanyan/tddc/pkg/customers"
    ic0845eadd364ce60c0c26d7e82aed4f535c37906cb39c2f2466b69dad5422b34 "github.com/xifanyan/tddc/pkg/companies"
    ic407fcfda44cf6e2c0a0901ae8bf6763dd7099a6e827bc350481077dd34134c2 "github.com/xifanyan/tddc/pkg/searches"
    ie90620e879a034e2f4301f9e71412ff912d0b514e22a1f26b5968328c9d23ce4 "github.com/xifanyan/tddc/pkg/directories"
)

// ApiClient the main entry point of the SDK, exposes the configuration and the fluent API.
type ApiClient struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// Collections the collections property
// returns a *CollectionsRequestBuilder when successful
func (m *ApiClient) Collections()(*i193388f9f459f47d2c8274c8ac796c8f50986cfd778a4c8d70f6cdf03b70cf69.CollectionsRequestBuilder) {
    return i193388f9f459f47d2c8274c8ac796c8f50986cfd778a4c8d70f6cdf03b70cf69.NewCollectionsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Companies the companies property
// returns a *CompaniesRequestBuilder when successful
func (m *ApiClient) Companies()(*ic0845eadd364ce60c0c26d7e82aed4f535c37906cb39c2f2466b69dad5422b34.CompaniesRequestBuilder) {
    return ic0845eadd364ce60c0c26d7e82aed4f535c37906cb39c2f2466b69dad5422b34.NewCompaniesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewApiClient instantiates a new ApiClient and sets the default values.
func NewApiClient(requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ApiClient) {
    m := &ApiClient{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}", map[string]string{}),
    }
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RegisterDefaultSerializer(func() i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriterFactory { return i25911dc319edd61cbac496af7eab5ef20b6069a42515e22ec6a9bc97bf598488.NewJsonSerializationWriterFactory() })
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RegisterDefaultSerializer(func() i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriterFactory { return i7294a22093d408fdca300f11b81a887d89c47b764af06c8b803e2323973fdb83.NewTextSerializationWriterFactory() })
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RegisterDefaultSerializer(func() i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriterFactory { return i4bcdc892e61ac17e2afc10b5e2b536b29f4fd6c1ad30f4a5a68df47495db3347.NewFormSerializationWriterFactory() })
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RegisterDefaultSerializer(func() i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriterFactory { return i56887720f41ac882814261620b1c8459c4a992a0207af547c4453dd39fabc426.NewMultipartSerializationWriterFactory() })
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RegisterDefaultDeserializer(func() i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNodeFactory { return i25911dc319edd61cbac496af7eab5ef20b6069a42515e22ec6a9bc97bf598488.NewJsonParseNodeFactory() })
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RegisterDefaultDeserializer(func() i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNodeFactory { return i7294a22093d408fdca300f11b81a887d89c47b764af06c8b803e2323973fdb83.NewTextParseNodeFactory() })
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RegisterDefaultDeserializer(func() i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNodeFactory { return i4bcdc892e61ac17e2afc10b5e2b536b29f4fd6c1ad30f4a5a68df47495db3347.NewFormParseNodeFactory() })
    if m.BaseRequestBuilder.RequestAdapter.GetBaseUrl() == "" {
        m.BaseRequestBuilder.RequestAdapter.SetBaseUrl("https://app.totaldiscovery.com/api/v1")
    }
    m.BaseRequestBuilder.PathParameters["baseurl"] = m.BaseRequestBuilder.RequestAdapter.GetBaseUrl()
    return m
}
// Custodians the custodians property
// returns a *CustodiansRequestBuilder when successful
func (m *ApiClient) Custodians()(*i45c651b865c1f94388852d91754503c3df6c397790455fddbfcc199084e7560d.CustodiansRequestBuilder) {
    return i45c651b865c1f94388852d91754503c3df6c397790455fddbfcc199084e7560d.NewCustodiansRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Customers the customers property
// returns a *CustomersRequestBuilder when successful
func (m *ApiClient) Customers()(*i4f8be26726d1e5cf99e42978603d7d3ca9281228732ec1fed238bee11dcf4f70.CustomersRequestBuilder) {
    return i4f8be26726d1e5cf99e42978603d7d3ca9281228732ec1fed238bee11dcf4f70.NewCustomersRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Directories the directories property
// returns a *DirectoriesRequestBuilder when successful
func (m *ApiClient) Directories()(*ie90620e879a034e2f4301f9e71412ff912d0b514e22a1f26b5968328c9d23ce4.DirectoriesRequestBuilder) {
    return ie90620e879a034e2f4301f9e71412ff912d0b514e22a1f26b5968328c9d23ce4.NewDirectoriesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Legal_holds the legal_holds property
// returns a *Legal_holdsRequestBuilder when successful
func (m *ApiClient) Legal_holds()(*i3d76b5e22f38add1010d104e0b80c877b70818b18473d854a923d9019345c26b.Legal_holdsRequestBuilder) {
    return i3d76b5e22f38add1010d104e0b80c877b70818b18473d854a923d9019345c26b.NewLegal_holdsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Matters the matters property
// returns a *MattersRequestBuilder when successful
func (m *ApiClient) Matters()(*i1de8afbc6843ca76cd39c931c5b733e39dbefccfeecbbe14bda2f40fad45ecb4.MattersRequestBuilder) {
    return i1de8afbc6843ca76cd39c931c5b733e39dbefccfeecbbe14bda2f40fad45ecb4.NewMattersRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Searches the searches property
// returns a *SearchesRequestBuilder when successful
func (m *ApiClient) Searches()(*ic407fcfda44cf6e2c0a0901ae8bf6763dd7099a6e827bc350481077dd34134c2.SearchesRequestBuilder) {
    return ic407fcfda44cf6e2c0a0901ae8bf6763dd7099a6e827bc350481077dd34134c2.NewSearchesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
