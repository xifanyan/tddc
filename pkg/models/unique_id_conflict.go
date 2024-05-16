package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type UniqueIdConflict struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The conflict property
    conflict *string
    // The custodian_ids property
    custodian_ids []int32
}
// NewUniqueIdConflict instantiates a new UniqueIdConflict and sets the default values.
func NewUniqueIdConflict()(*UniqueIdConflict) {
    m := &UniqueIdConflict{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateUniqueIdConflictFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateUniqueIdConflictFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewUniqueIdConflict(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *UniqueIdConflict) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetConflict gets the conflict property value. The conflict property
// returns a *string when successful
func (m *UniqueIdConflict) GetConflict()(*string) {
    return m.conflict
}
// GetCustodianIds gets the custodian_ids property value. The custodian_ids property
// returns a []int32 when successful
func (m *UniqueIdConflict) GetCustodianIds()([]int32) {
    return m.custodian_ids
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *UniqueIdConflict) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["conflict"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetConflict(val)
        }
        return nil
    }
    res["custodian_ids"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("int32")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]int32, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = *(v.(*int32))
                }
            }
            m.SetCustodianIds(res)
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *UniqueIdConflict) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("conflict", m.GetConflict())
        if err != nil {
            return err
        }
    }
    if m.GetCustodianIds() != nil {
        err := writer.WriteCollectionOfInt32Values("custodian_ids", m.GetCustodianIds())
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
func (m *UniqueIdConflict) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetConflict sets the conflict property value. The conflict property
func (m *UniqueIdConflict) SetConflict(value *string)() {
    m.conflict = value
}
// SetCustodianIds sets the custodian_ids property value. The custodian_ids property
func (m *UniqueIdConflict) SetCustodianIds(value []int32)() {
    m.custodian_ids = value
}
type UniqueIdConflictable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetConflict()(*string)
    GetCustodianIds()([]int32)
    SetConflict(value *string)()
    SetCustodianIds(value []int32)()
}
