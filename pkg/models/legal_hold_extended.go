package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type Legal_hold_extended struct {
    Legal_hold_with_matter
    // The custodians_legal_holds property
    custodians_legal_holds []Legal_hold_extended_custodians_legal_holdsable
}
// NewLegal_hold_extended instantiates a new Legal_hold_extended and sets the default values.
func NewLegal_hold_extended()(*Legal_hold_extended) {
    m := &Legal_hold_extended{
        Legal_hold_with_matter: *NewLegal_hold_with_matter(),
    }
    return m
}
// CreateLegal_hold_extendedFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateLegal_hold_extendedFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewLegal_hold_extended(), nil
}
// GetCustodiansLegalHolds gets the custodians_legal_holds property value. The custodians_legal_holds property
// returns a []Legal_hold_extended_custodians_legal_holdsable when successful
func (m *Legal_hold_extended) GetCustodiansLegalHolds()([]Legal_hold_extended_custodians_legal_holdsable) {
    return m.custodians_legal_holds
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *Legal_hold_extended) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Legal_hold_with_matter.GetFieldDeserializers()
    res["custodians_legal_holds"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateLegal_hold_extended_custodians_legal_holdsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]Legal_hold_extended_custodians_legal_holdsable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(Legal_hold_extended_custodians_legal_holdsable)
                }
            }
            m.SetCustodiansLegalHolds(res)
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *Legal_hold_extended) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Legal_hold_with_matter.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetCustodiansLegalHolds() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetCustodiansLegalHolds()))
        for i, v := range m.GetCustodiansLegalHolds() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("custodians_legal_holds", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetCustodiansLegalHolds sets the custodians_legal_holds property value. The custodians_legal_holds property
func (m *Legal_hold_extended) SetCustodiansLegalHolds(value []Legal_hold_extended_custodians_legal_holdsable)() {
    m.custodians_legal_holds = value
}
type Legal_hold_extendedable interface {
    Legal_hold_with_matterable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCustodiansLegalHolds()([]Legal_hold_extended_custodians_legal_holdsable)
    SetCustodiansLegalHolds(value []Legal_hold_extended_custodians_legal_holdsable)()
}
