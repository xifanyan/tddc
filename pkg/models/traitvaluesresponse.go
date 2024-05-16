package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type Traitvaluesresponse struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The trait_name property
    trait_name *string
    // The trait_values property
    trait_values []string
}
// NewTraitvaluesresponse instantiates a new Traitvaluesresponse and sets the default values.
func NewTraitvaluesresponse()(*Traitvaluesresponse) {
    m := &Traitvaluesresponse{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateTraitvaluesresponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateTraitvaluesresponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewTraitvaluesresponse(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *Traitvaluesresponse) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *Traitvaluesresponse) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["trait_name"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTraitName(val)
        }
        return nil
    }
    res["trait_values"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
            m.SetTraitValues(res)
        }
        return nil
    }
    return res
}
// GetTraitName gets the trait_name property value. The trait_name property
// returns a *string when successful
func (m *Traitvaluesresponse) GetTraitName()(*string) {
    return m.trait_name
}
// GetTraitValues gets the trait_values property value. The trait_values property
// returns a []string when successful
func (m *Traitvaluesresponse) GetTraitValues()([]string) {
    return m.trait_values
}
// Serialize serializes information the current object
func (m *Traitvaluesresponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("trait_name", m.GetTraitName())
        if err != nil {
            return err
        }
    }
    if m.GetTraitValues() != nil {
        err := writer.WriteCollectionOfStringValues("trait_values", m.GetTraitValues())
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
func (m *Traitvaluesresponse) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetTraitName sets the trait_name property value. The trait_name property
func (m *Traitvaluesresponse) SetTraitName(value *string)() {
    m.trait_name = value
}
// SetTraitValues sets the trait_values property value. The trait_values property
func (m *Traitvaluesresponse) SetTraitValues(value []string)() {
    m.trait_values = value
}
type Traitvaluesresponseable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetTraitName()(*string)
    GetTraitValues()([]string)
    SetTraitName(value *string)()
    SetTraitValues(value []string)()
}
