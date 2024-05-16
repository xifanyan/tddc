package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type Synctraitrequest struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The custodian_unique_ids property
    custodian_unique_ids []string
    // The trait_value property
    trait_value *string
}
// NewSynctraitrequest instantiates a new Synctraitrequest and sets the default values.
func NewSynctraitrequest()(*Synctraitrequest) {
    m := &Synctraitrequest{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateSynctraitrequestFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateSynctraitrequestFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewSynctraitrequest(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *Synctraitrequest) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetCustodianUniqueIds gets the custodian_unique_ids property value. The custodian_unique_ids property
// returns a []string when successful
func (m *Synctraitrequest) GetCustodianUniqueIds()([]string) {
    return m.custodian_unique_ids
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *Synctraitrequest) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["custodian_unique_ids"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = *(v.(*string))
                }
            }
            m.SetCustodianUniqueIds(res)
        }
        return nil
    }
    res["trait_value"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTraitValue(val)
        }
        return nil
    }
    return res
}
// GetTraitValue gets the trait_value property value. The trait_value property
// returns a *string when successful
func (m *Synctraitrequest) GetTraitValue()(*string) {
    return m.trait_value
}
// Serialize serializes information the current object
func (m *Synctraitrequest) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetCustodianUniqueIds() != nil {
        err := writer.WriteCollectionOfStringValues("custodian_unique_ids", m.GetCustodianUniqueIds())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("trait_value", m.GetTraitValue())
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
func (m *Synctraitrequest) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetCustodianUniqueIds sets the custodian_unique_ids property value. The custodian_unique_ids property
func (m *Synctraitrequest) SetCustodianUniqueIds(value []string)() {
    m.custodian_unique_ids = value
}
// SetTraitValue sets the trait_value property value. The trait_value property
func (m *Synctraitrequest) SetTraitValue(value *string)() {
    m.trait_value = value
}
type Synctraitrequestable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCustodianUniqueIds()([]string)
    GetTraitValue()(*string)
    SetCustodianUniqueIds(value []string)()
    SetTraitValue(value *string)()
}
