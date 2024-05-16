package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type Legal_hold_status_audit struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The created_at property
    created_at *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The custodian property
    custodian Custodianable
    // The legal_hold property
    legal_hold Legal_hold_with_matterable
    // The status_type property
    status_type *string
}
// NewLegal_hold_status_audit instantiates a new Legal_hold_status_audit and sets the default values.
func NewLegal_hold_status_audit()(*Legal_hold_status_audit) {
    m := &Legal_hold_status_audit{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateLegal_hold_status_auditFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateLegal_hold_status_auditFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewLegal_hold_status_audit(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *Legal_hold_status_audit) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetCreatedAt gets the created_at property value. The created_at property
// returns a *Time when successful
func (m *Legal_hold_status_audit) GetCreatedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.created_at
}
// GetCustodian gets the custodian property value. The custodian property
// returns a Custodianable when successful
func (m *Legal_hold_status_audit) GetCustodian()(Custodianable) {
    return m.custodian
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *Legal_hold_status_audit) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["created_at"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCreatedAt(val)
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
    res["legal_hold"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateLegal_hold_with_matterFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLegalHold(val.(Legal_hold_with_matterable))
        }
        return nil
    }
    res["status_type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStatusType(val)
        }
        return nil
    }
    return res
}
// GetLegalHold gets the legal_hold property value. The legal_hold property
// returns a Legal_hold_with_matterable when successful
func (m *Legal_hold_status_audit) GetLegalHold()(Legal_hold_with_matterable) {
    return m.legal_hold
}
// GetStatusType gets the status_type property value. The status_type property
// returns a *string when successful
func (m *Legal_hold_status_audit) GetStatusType()(*string) {
    return m.status_type
}
// Serialize serializes information the current object
func (m *Legal_hold_status_audit) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteTimeValue("created_at", m.GetCreatedAt())
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
        err := writer.WriteObjectValue("legal_hold", m.GetLegalHold())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("status_type", m.GetStatusType())
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
func (m *Legal_hold_status_audit) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetCreatedAt sets the created_at property value. The created_at property
func (m *Legal_hold_status_audit) SetCreatedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.created_at = value
}
// SetCustodian sets the custodian property value. The custodian property
func (m *Legal_hold_status_audit) SetCustodian(value Custodianable)() {
    m.custodian = value
}
// SetLegalHold sets the legal_hold property value. The legal_hold property
func (m *Legal_hold_status_audit) SetLegalHold(value Legal_hold_with_matterable)() {
    m.legal_hold = value
}
// SetStatusType sets the status_type property value. The status_type property
func (m *Legal_hold_status_audit) SetStatusType(value *string)() {
    m.status_type = value
}
type Legal_hold_status_auditable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCreatedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetCustodian()(Custodianable)
    GetLegalHold()(Legal_hold_with_matterable)
    GetStatusType()(*string)
    SetCreatedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetCustodian(value Custodianable)()
    SetLegalHold(value Legal_hold_with_matterable)()
    SetStatusType(value *string)()
}
