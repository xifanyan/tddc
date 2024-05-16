package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type Search_diag struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The full_query property
    full_query Search_diag_full_queryable
}
// NewSearch_diag instantiates a new Search_diag and sets the default values.
func NewSearch_diag()(*Search_diag) {
    m := &Search_diag{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateSearch_diagFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateSearch_diagFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewSearch_diag(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *Search_diag) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *Search_diag) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["full_query"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateSearch_diag_full_queryFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFullQuery(val.(Search_diag_full_queryable))
        }
        return nil
    }
    return res
}
// GetFullQuery gets the full_query property value. The full_query property
// returns a Search_diag_full_queryable when successful
func (m *Search_diag) GetFullQuery()(Search_diag_full_queryable) {
    return m.full_query
}
// Serialize serializes information the current object
func (m *Search_diag) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("full_query", m.GetFullQuery())
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
func (m *Search_diag) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetFullQuery sets the full_query property value. The full_query property
func (m *Search_diag) SetFullQuery(value Search_diag_full_queryable)() {
    m.full_query = value
}
type Search_diagable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetFullQuery()(Search_diag_full_queryable)
    SetFullQuery(value Search_diag_full_queryable)()
}
