package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type Legal_hold_extended_custodians_legal_holds_legal_hold_status struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The is_on_hold property
    is_on_hold *string
    // The name property
    name *string
}
// NewLegal_hold_extended_custodians_legal_holds_legal_hold_status instantiates a new Legal_hold_extended_custodians_legal_holds_legal_hold_status and sets the default values.
func NewLegal_hold_extended_custodians_legal_holds_legal_hold_status()(*Legal_hold_extended_custodians_legal_holds_legal_hold_status) {
    m := &Legal_hold_extended_custodians_legal_holds_legal_hold_status{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateLegal_hold_extended_custodians_legal_holds_legal_hold_statusFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateLegal_hold_extended_custodians_legal_holds_legal_hold_statusFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewLegal_hold_extended_custodians_legal_holds_legal_hold_status(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *Legal_hold_extended_custodians_legal_holds_legal_hold_status) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *Legal_hold_extended_custodians_legal_holds_legal_hold_status) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["is_on_hold"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsOnHold(val)
        }
        return nil
    }
    res["name"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetName(val)
        }
        return nil
    }
    return res
}
// GetIsOnHold gets the is_on_hold property value. The is_on_hold property
// returns a *string when successful
func (m *Legal_hold_extended_custodians_legal_holds_legal_hold_status) GetIsOnHold()(*string) {
    return m.is_on_hold
}
// GetName gets the name property value. The name property
// returns a *string when successful
func (m *Legal_hold_extended_custodians_legal_holds_legal_hold_status) GetName()(*string) {
    return m.name
}
// Serialize serializes information the current object
func (m *Legal_hold_extended_custodians_legal_holds_legal_hold_status) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("is_on_hold", m.GetIsOnHold())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("name", m.GetName())
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
func (m *Legal_hold_extended_custodians_legal_holds_legal_hold_status) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetIsOnHold sets the is_on_hold property value. The is_on_hold property
func (m *Legal_hold_extended_custodians_legal_holds_legal_hold_status) SetIsOnHold(value *string)() {
    m.is_on_hold = value
}
// SetName sets the name property value. The name property
func (m *Legal_hold_extended_custodians_legal_holds_legal_hold_status) SetName(value *string)() {
    m.name = value
}
type Legal_hold_extended_custodians_legal_holds_legal_hold_statusable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetIsOnHold()(*string)
    GetName()(*string)
    SetIsOnHold(value *string)()
    SetName(value *string)()
}
