package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type Legal_hold_with_matter struct {
    Legal_hold
    // The matter property
    matter Matterable
}
// NewLegal_hold_with_matter instantiates a new Legal_hold_with_matter and sets the default values.
func NewLegal_hold_with_matter()(*Legal_hold_with_matter) {
    m := &Legal_hold_with_matter{
        Legal_hold: *NewLegal_hold(),
    }
    return m
}
// CreateLegal_hold_with_matterFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateLegal_hold_with_matterFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewLegal_hold_with_matter(), nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *Legal_hold_with_matter) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Legal_hold.GetFieldDeserializers()
    res["matter"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateMatterFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMatter(val.(Matterable))
        }
        return nil
    }
    return res
}
// GetMatter gets the matter property value. The matter property
// returns a Matterable when successful
func (m *Legal_hold_with_matter) GetMatter()(Matterable) {
    return m.matter
}
// Serialize serializes information the current object
func (m *Legal_hold_with_matter) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Legal_hold.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("matter", m.GetMatter())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetMatter sets the matter property value. The matter property
func (m *Legal_hold_with_matter) SetMatter(value Matterable)() {
    m.matter = value
}
type Legal_hold_with_matterable interface {
    Legal_holdable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetMatter()(Matterable)
    SetMatter(value Matterable)()
}
