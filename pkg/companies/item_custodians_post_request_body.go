package companies

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    if815b544aed335c15fd039b4dcb81545e8ef32c8003c12e1dc36ea8df9bfeb2f "github.com/xifanyan/tddc/pkg/models"
)

type ItemCustodiansPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The custodians property
    custodians []if815b544aed335c15fd039b4dcb81545e8ef32c8003c12e1dc36ea8df9bfeb2f.Custodianable
}
// NewItemCustodiansPostRequestBody instantiates a new ItemCustodiansPostRequestBody and sets the default values.
func NewItemCustodiansPostRequestBody()(*ItemCustodiansPostRequestBody) {
    m := &ItemCustodiansPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateItemCustodiansPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateItemCustodiansPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemCustodiansPostRequestBody(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *ItemCustodiansPostRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetCustodians gets the custodians property value. The custodians property
// returns a []Custodianable when successful
func (m *ItemCustodiansPostRequestBody) GetCustodians()([]if815b544aed335c15fd039b4dcb81545e8ef32c8003c12e1dc36ea8df9bfeb2f.Custodianable) {
    return m.custodians
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ItemCustodiansPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["custodians"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(if815b544aed335c15fd039b4dcb81545e8ef32c8003c12e1dc36ea8df9bfeb2f.CreateCustodianFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]if815b544aed335c15fd039b4dcb81545e8ef32c8003c12e1dc36ea8df9bfeb2f.Custodianable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(if815b544aed335c15fd039b4dcb81545e8ef32c8003c12e1dc36ea8df9bfeb2f.Custodianable)
                }
            }
            m.SetCustodians(res)
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *ItemCustodiansPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetCustodians() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetCustodians()))
        for i, v := range m.GetCustodians() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("custodians", cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ItemCustodiansPostRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetCustodians sets the custodians property value. The custodians property
func (m *ItemCustodiansPostRequestBody) SetCustodians(value []if815b544aed335c15fd039b4dcb81545e8ef32c8003c12e1dc36ea8df9bfeb2f.Custodianable)() {
    m.custodians = value
}
type ItemCustodiansPostRequestBodyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCustodians()([]if815b544aed335c15fd039b4dcb81545e8ef32c8003c12e1dc36ea8df9bfeb2f.Custodianable)
    SetCustodians(value []if815b544aed335c15fd039b4dcb81545e8ef32c8003c12e1dc36ea8df9bfeb2f.Custodianable)()
}
