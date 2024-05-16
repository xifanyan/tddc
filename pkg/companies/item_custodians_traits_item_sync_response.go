package companies

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use ItemCustodiansTraitsItemSyncPostResponseable instead.
type ItemCustodiansTraitsItemSyncResponse struct {
    ItemCustodiansTraitsItemSyncPostResponse
}
// NewItemCustodiansTraitsItemSyncResponse instantiates a new ItemCustodiansTraitsItemSyncResponse and sets the default values.
func NewItemCustodiansTraitsItemSyncResponse()(*ItemCustodiansTraitsItemSyncResponse) {
    m := &ItemCustodiansTraitsItemSyncResponse{
        ItemCustodiansTraitsItemSyncPostResponse: *NewItemCustodiansTraitsItemSyncPostResponse(),
    }
    return m
}
// CreateItemCustodiansTraitsItemSyncResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateItemCustodiansTraitsItemSyncResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemCustodiansTraitsItemSyncResponse(), nil
}
// Deprecated: This class is obsolete. Use ItemCustodiansTraitsItemSyncPostResponseable instead.
type ItemCustodiansTraitsItemSyncResponseable interface {
    ItemCustodiansTraitsItemSyncPostResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
