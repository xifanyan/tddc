package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type Matter struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The business_unit property
    business_unit *string
    // The caption property
    caption *string
    // The company_id property
    company_id *int32
    // The date_created property
    date_created *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The date_end property
    date_end *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The date_modified property
    date_modified *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The date_start property
    date_start *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The department property
    department *string
    // The id property
    id *int32
    // The matter_contacts property
    matter_contacts []Matter_contactable
    // The name property
    name *string
    // The notes property
    notes *string
    // The number property
    number *string
}
// NewMatter instantiates a new Matter and sets the default values.
func NewMatter()(*Matter) {
    m := &Matter{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateMatterFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateMatterFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewMatter(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *Matter) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetBusinessUnit gets the business_unit property value. The business_unit property
// returns a *string when successful
func (m *Matter) GetBusinessUnit()(*string) {
    return m.business_unit
}
// GetCaption gets the caption property value. The caption property
// returns a *string when successful
func (m *Matter) GetCaption()(*string) {
    return m.caption
}
// GetCompanyId gets the company_id property value. The company_id property
// returns a *int32 when successful
func (m *Matter) GetCompanyId()(*int32) {
    return m.company_id
}
// GetDateCreated gets the date_created property value. The date_created property
// returns a *Time when successful
func (m *Matter) GetDateCreated()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.date_created
}
// GetDateEnd gets the date_end property value. The date_end property
// returns a *Time when successful
func (m *Matter) GetDateEnd()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.date_end
}
// GetDateModified gets the date_modified property value. The date_modified property
// returns a *Time when successful
func (m *Matter) GetDateModified()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.date_modified
}
// GetDateStart gets the date_start property value. The date_start property
// returns a *Time when successful
func (m *Matter) GetDateStart()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.date_start
}
// GetDepartment gets the department property value. The department property
// returns a *string when successful
func (m *Matter) GetDepartment()(*string) {
    return m.department
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *Matter) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["business_unit"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBusinessUnit(val)
        }
        return nil
    }
    res["caption"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCaption(val)
        }
        return nil
    }
    res["company_id"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCompanyId(val)
        }
        return nil
    }
    res["date_created"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDateCreated(val)
        }
        return nil
    }
    res["date_end"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDateEnd(val)
        }
        return nil
    }
    res["date_modified"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDateModified(val)
        }
        return nil
    }
    res["date_start"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDateStart(val)
        }
        return nil
    }
    res["department"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDepartment(val)
        }
        return nil
    }
    res["id"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetId(val)
        }
        return nil
    }
    res["matter_contacts"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateMatter_contactFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]Matter_contactable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(Matter_contactable)
                }
            }
            m.SetMatterContacts(res)
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
    res["notes"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNotes(val)
        }
        return nil
    }
    res["number"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNumber(val)
        }
        return nil
    }
    return res
}
// GetId gets the id property value. The id property
// returns a *int32 when successful
func (m *Matter) GetId()(*int32) {
    return m.id
}
// GetMatterContacts gets the matter_contacts property value. The matter_contacts property
// returns a []Matter_contactable when successful
func (m *Matter) GetMatterContacts()([]Matter_contactable) {
    return m.matter_contacts
}
// GetName gets the name property value. The name property
// returns a *string when successful
func (m *Matter) GetName()(*string) {
    return m.name
}
// GetNotes gets the notes property value. The notes property
// returns a *string when successful
func (m *Matter) GetNotes()(*string) {
    return m.notes
}
// GetNumber gets the number property value. The number property
// returns a *string when successful
func (m *Matter) GetNumber()(*string) {
    return m.number
}
// Serialize serializes information the current object
func (m *Matter) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("business_unit", m.GetBusinessUnit())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("caption", m.GetCaption())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("company_id", m.GetCompanyId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteTimeValue("date_created", m.GetDateCreated())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteTimeValue("date_end", m.GetDateEnd())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteTimeValue("date_modified", m.GetDateModified())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteTimeValue("date_start", m.GetDateStart())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("department", m.GetDepartment())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("id", m.GetId())
        if err != nil {
            return err
        }
    }
    if m.GetMatterContacts() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetMatterContacts()))
        for i, v := range m.GetMatterContacts() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("matter_contacts", cast)
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
        err := writer.WriteStringValue("notes", m.GetNotes())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("number", m.GetNumber())
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
func (m *Matter) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetBusinessUnit sets the business_unit property value. The business_unit property
func (m *Matter) SetBusinessUnit(value *string)() {
    m.business_unit = value
}
// SetCaption sets the caption property value. The caption property
func (m *Matter) SetCaption(value *string)() {
    m.caption = value
}
// SetCompanyId sets the company_id property value. The company_id property
func (m *Matter) SetCompanyId(value *int32)() {
    m.company_id = value
}
// SetDateCreated sets the date_created property value. The date_created property
func (m *Matter) SetDateCreated(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.date_created = value
}
// SetDateEnd sets the date_end property value. The date_end property
func (m *Matter) SetDateEnd(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.date_end = value
}
// SetDateModified sets the date_modified property value. The date_modified property
func (m *Matter) SetDateModified(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.date_modified = value
}
// SetDateStart sets the date_start property value. The date_start property
func (m *Matter) SetDateStart(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.date_start = value
}
// SetDepartment sets the department property value. The department property
func (m *Matter) SetDepartment(value *string)() {
    m.department = value
}
// SetId sets the id property value. The id property
func (m *Matter) SetId(value *int32)() {
    m.id = value
}
// SetMatterContacts sets the matter_contacts property value. The matter_contacts property
func (m *Matter) SetMatterContacts(value []Matter_contactable)() {
    m.matter_contacts = value
}
// SetName sets the name property value. The name property
func (m *Matter) SetName(value *string)() {
    m.name = value
}
// SetNotes sets the notes property value. The notes property
func (m *Matter) SetNotes(value *string)() {
    m.notes = value
}
// SetNumber sets the number property value. The number property
func (m *Matter) SetNumber(value *string)() {
    m.number = value
}
type Matterable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetBusinessUnit()(*string)
    GetCaption()(*string)
    GetCompanyId()(*int32)
    GetDateCreated()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetDateEnd()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetDateModified()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetDateStart()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetDepartment()(*string)
    GetId()(*int32)
    GetMatterContacts()([]Matter_contactable)
    GetName()(*string)
    GetNotes()(*string)
    GetNumber()(*string)
    SetBusinessUnit(value *string)()
    SetCaption(value *string)()
    SetCompanyId(value *int32)()
    SetDateCreated(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetDateEnd(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetDateModified(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetDateStart(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetDepartment(value *string)()
    SetId(value *int32)()
    SetMatterContacts(value []Matter_contactable)()
    SetName(value *string)()
    SetNotes(value *string)()
    SetNumber(value *string)()
}
