package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type Legal_hold_extended_custodians_legal_holds struct {
    // The acknowledged_at property
    acknowledged_at *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The custodian property
    custodian Custodianable
    // The error property
    error *bool
    // The escalation_error property
    escalation_error *bool
    // The last_acknowledgement_reminder_sent_at property
    last_acknowledgement_reminder_sent_at *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The last_hold_reminder_sent_at property
    last_hold_reminder_sent_at *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The legal_hold_status property
    legal_hold_status Legal_hold_extended_custodians_legal_holds_legal_hold_statusable
    // The released_at property
    released_at *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The sent_at property
    sent_at *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
}
// NewLegal_hold_extended_custodians_legal_holds instantiates a new Legal_hold_extended_custodians_legal_holds and sets the default values.
func NewLegal_hold_extended_custodians_legal_holds()(*Legal_hold_extended_custodians_legal_holds) {
    m := &Legal_hold_extended_custodians_legal_holds{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateLegal_hold_extended_custodians_legal_holdsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateLegal_hold_extended_custodians_legal_holdsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewLegal_hold_extended_custodians_legal_holds(), nil
}
// GetAcknowledgedAt gets the acknowledged_at property value. The acknowledged_at property
// returns a *Time when successful
func (m *Legal_hold_extended_custodians_legal_holds) GetAcknowledgedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.acknowledged_at
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *Legal_hold_extended_custodians_legal_holds) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetCustodian gets the custodian property value. The custodian property
// returns a Custodianable when successful
func (m *Legal_hold_extended_custodians_legal_holds) GetCustodian()(Custodianable) {
    return m.custodian
}
// GetError gets the error property value. The error property
// returns a *bool when successful
func (m *Legal_hold_extended_custodians_legal_holds) GetError()(*bool) {
    return m.error
}
// GetEscalationError gets the escalation_error property value. The escalation_error property
// returns a *bool when successful
func (m *Legal_hold_extended_custodians_legal_holds) GetEscalationError()(*bool) {
    return m.escalation_error
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *Legal_hold_extended_custodians_legal_holds) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["acknowledged_at"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAcknowledgedAt(val)
        }
        return nil
    }
    res["custodian"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateCustodianFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCustodian(val.(Custodianable))
        }
        return nil
    }
    res["error"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetError(val)
        }
        return nil
    }
    res["escalation_error"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEscalationError(val)
        }
        return nil
    }
    res["last_acknowledgement_reminder_sent_at"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastAcknowledgementReminderSentAt(val)
        }
        return nil
    }
    res["last_hold_reminder_sent_at"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastHoldReminderSentAt(val)
        }
        return nil
    }
    res["legal_hold_status"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateLegal_hold_extended_custodians_legal_holds_legal_hold_statusFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLegalHoldStatus(val.(Legal_hold_extended_custodians_legal_holds_legal_hold_statusable))
        }
        return nil
    }
    res["released_at"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetReleasedAt(val)
        }
        return nil
    }
    res["sent_at"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSentAt(val)
        }
        return nil
    }
    return res
}
// GetLastAcknowledgementReminderSentAt gets the last_acknowledgement_reminder_sent_at property value. The last_acknowledgement_reminder_sent_at property
// returns a *Time when successful
func (m *Legal_hold_extended_custodians_legal_holds) GetLastAcknowledgementReminderSentAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.last_acknowledgement_reminder_sent_at
}
// GetLastHoldReminderSentAt gets the last_hold_reminder_sent_at property value. The last_hold_reminder_sent_at property
// returns a *Time when successful
func (m *Legal_hold_extended_custodians_legal_holds) GetLastHoldReminderSentAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.last_hold_reminder_sent_at
}
// GetLegalHoldStatus gets the legal_hold_status property value. The legal_hold_status property
// returns a Legal_hold_extended_custodians_legal_holds_legal_hold_statusable when successful
func (m *Legal_hold_extended_custodians_legal_holds) GetLegalHoldStatus()(Legal_hold_extended_custodians_legal_holds_legal_hold_statusable) {
    return m.legal_hold_status
}
// GetReleasedAt gets the released_at property value. The released_at property
// returns a *Time when successful
func (m *Legal_hold_extended_custodians_legal_holds) GetReleasedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.released_at
}
// GetSentAt gets the sent_at property value. The sent_at property
// returns a *Time when successful
func (m *Legal_hold_extended_custodians_legal_holds) GetSentAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.sent_at
}
// Serialize serializes information the current object
func (m *Legal_hold_extended_custodians_legal_holds) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteTimeValue("acknowledged_at", m.GetAcknowledgedAt())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("custodian", m.GetCustodian())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("error", m.GetError())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("escalation_error", m.GetEscalationError())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteTimeValue("last_acknowledgement_reminder_sent_at", m.GetLastAcknowledgementReminderSentAt())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteTimeValue("last_hold_reminder_sent_at", m.GetLastHoldReminderSentAt())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("legal_hold_status", m.GetLegalHoldStatus())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteTimeValue("released_at", m.GetReleasedAt())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteTimeValue("sent_at", m.GetSentAt())
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
// SetAcknowledgedAt sets the acknowledged_at property value. The acknowledged_at property
func (m *Legal_hold_extended_custodians_legal_holds) SetAcknowledgedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.acknowledged_at = value
}
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *Legal_hold_extended_custodians_legal_holds) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetCustodian sets the custodian property value. The custodian property
func (m *Legal_hold_extended_custodians_legal_holds) SetCustodian(value Custodianable)() {
    m.custodian = value
}
// SetError sets the error property value. The error property
func (m *Legal_hold_extended_custodians_legal_holds) SetError(value *bool)() {
    m.error = value
}
// SetEscalationError sets the escalation_error property value. The escalation_error property
func (m *Legal_hold_extended_custodians_legal_holds) SetEscalationError(value *bool)() {
    m.escalation_error = value
}
// SetLastAcknowledgementReminderSentAt sets the last_acknowledgement_reminder_sent_at property value. The last_acknowledgement_reminder_sent_at property
func (m *Legal_hold_extended_custodians_legal_holds) SetLastAcknowledgementReminderSentAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.last_acknowledgement_reminder_sent_at = value
}
// SetLastHoldReminderSentAt sets the last_hold_reminder_sent_at property value. The last_hold_reminder_sent_at property
func (m *Legal_hold_extended_custodians_legal_holds) SetLastHoldReminderSentAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.last_hold_reminder_sent_at = value
}
// SetLegalHoldStatus sets the legal_hold_status property value. The legal_hold_status property
func (m *Legal_hold_extended_custodians_legal_holds) SetLegalHoldStatus(value Legal_hold_extended_custodians_legal_holds_legal_hold_statusable)() {
    m.legal_hold_status = value
}
// SetReleasedAt sets the released_at property value. The released_at property
func (m *Legal_hold_extended_custodians_legal_holds) SetReleasedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.released_at = value
}
// SetSentAt sets the sent_at property value. The sent_at property
func (m *Legal_hold_extended_custodians_legal_holds) SetSentAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.sent_at = value
}
type Legal_hold_extended_custodians_legal_holdsable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAcknowledgedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetCustodian()(Custodianable)
    GetError()(*bool)
    GetEscalationError()(*bool)
    GetLastAcknowledgementReminderSentAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetLastHoldReminderSentAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetLegalHoldStatus()(Legal_hold_extended_custodians_legal_holds_legal_hold_statusable)
    GetReleasedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetSentAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    SetAcknowledgedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetCustodian(value Custodianable)()
    SetError(value *bool)()
    SetEscalationError(value *bool)()
    SetLastAcknowledgementReminderSentAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetLastHoldReminderSentAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetLegalHoldStatus(value Legal_hold_extended_custodians_legal_holds_legal_hold_statusable)()
    SetReleasedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetSentAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
}
