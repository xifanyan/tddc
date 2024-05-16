package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type Legal_hold_stats struct {
    // The count of holds that have at least one active custodian
    active_legal_holds *int32
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The total *unique* count of custodians on hold
    custodians_on_hold *int32
    // The count of holds that have all custodians released.
    released_legal_holds *int32
    // How many legal holds have sent notification out to custodians. Note: this count will include legal holds that have been released, as they had sent a notice in their past
    sent_legal_holds *int32
    // Total *unique* count of custodians that have at least one hold they are, and they have not acknowledged
    total_not_acknowledged *int32
}
// NewLegal_hold_stats instantiates a new Legal_hold_stats and sets the default values.
func NewLegal_hold_stats()(*Legal_hold_stats) {
    m := &Legal_hold_stats{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateLegal_hold_statsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateLegal_hold_statsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewLegal_hold_stats(), nil
}
// GetActiveLegalHolds gets the active_legal_holds property value. The count of holds that have at least one active custodian
// returns a *int32 when successful
func (m *Legal_hold_stats) GetActiveLegalHolds()(*int32) {
    return m.active_legal_holds
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *Legal_hold_stats) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetCustodiansOnHold gets the custodians_on_hold property value. The total *unique* count of custodians on hold
// returns a *int32 when successful
func (m *Legal_hold_stats) GetCustodiansOnHold()(*int32) {
    return m.custodians_on_hold
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *Legal_hold_stats) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["active_legal_holds"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetActiveLegalHolds(val)
        }
        return nil
    }
    res["custodians_on_hold"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCustodiansOnHold(val)
        }
        return nil
    }
    res["released_legal_holds"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetReleasedLegalHolds(val)
        }
        return nil
    }
    res["sent_legal_holds"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSentLegalHolds(val)
        }
        return nil
    }
    res["total_not_acknowledged"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTotalNotAcknowledged(val)
        }
        return nil
    }
    return res
}
// GetReleasedLegalHolds gets the released_legal_holds property value. The count of holds that have all custodians released.
// returns a *int32 when successful
func (m *Legal_hold_stats) GetReleasedLegalHolds()(*int32) {
    return m.released_legal_holds
}
// GetSentLegalHolds gets the sent_legal_holds property value. How many legal holds have sent notification out to custodians. Note: this count will include legal holds that have been released, as they had sent a notice in their past
// returns a *int32 when successful
func (m *Legal_hold_stats) GetSentLegalHolds()(*int32) {
    return m.sent_legal_holds
}
// GetTotalNotAcknowledged gets the total_not_acknowledged property value. Total *unique* count of custodians that have at least one hold they are, and they have not acknowledged
// returns a *int32 when successful
func (m *Legal_hold_stats) GetTotalNotAcknowledged()(*int32) {
    return m.total_not_acknowledged
}
// Serialize serializes information the current object
func (m *Legal_hold_stats) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteInt32Value("active_legal_holds", m.GetActiveLegalHolds())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("custodians_on_hold", m.GetCustodiansOnHold())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("released_legal_holds", m.GetReleasedLegalHolds())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("sent_legal_holds", m.GetSentLegalHolds())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("total_not_acknowledged", m.GetTotalNotAcknowledged())
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
// SetActiveLegalHolds sets the active_legal_holds property value. The count of holds that have at least one active custodian
func (m *Legal_hold_stats) SetActiveLegalHolds(value *int32)() {
    m.active_legal_holds = value
}
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *Legal_hold_stats) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetCustodiansOnHold sets the custodians_on_hold property value. The total *unique* count of custodians on hold
func (m *Legal_hold_stats) SetCustodiansOnHold(value *int32)() {
    m.custodians_on_hold = value
}
// SetReleasedLegalHolds sets the released_legal_holds property value. The count of holds that have all custodians released.
func (m *Legal_hold_stats) SetReleasedLegalHolds(value *int32)() {
    m.released_legal_holds = value
}
// SetSentLegalHolds sets the sent_legal_holds property value. How many legal holds have sent notification out to custodians. Note: this count will include legal holds that have been released, as they had sent a notice in their past
func (m *Legal_hold_stats) SetSentLegalHolds(value *int32)() {
    m.sent_legal_holds = value
}
// SetTotalNotAcknowledged sets the total_not_acknowledged property value. Total *unique* count of custodians that have at least one hold they are, and they have not acknowledged
func (m *Legal_hold_stats) SetTotalNotAcknowledged(value *int32)() {
    m.total_not_acknowledged = value
}
type Legal_hold_statsable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetActiveLegalHolds()(*int32)
    GetCustodiansOnHold()(*int32)
    GetReleasedLegalHolds()(*int32)
    GetSentLegalHolds()(*int32)
    GetTotalNotAcknowledged()(*int32)
    SetActiveLegalHolds(value *int32)()
    SetCustodiansOnHold(value *int32)()
    SetReleasedLegalHolds(value *int32)()
    SetSentLegalHolds(value *int32)()
    SetTotalNotAcknowledged(value *int32)()
}
