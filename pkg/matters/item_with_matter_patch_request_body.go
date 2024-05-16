package matters

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type ItemWithMatter_PatchRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The business_unit property
    business_unit *string
    // The caption property
    caption *string
    // The company_id property
    company_id *int32
    // The date_end property
    date_end *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The date_start property
    date_start *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The department property
    department *string
    // The name property
    name *string
    // The notes property
    notes *string
    // The number property
    number *string
}
// NewItemWithMatter_PatchRequestBody instantiates a new ItemWithMatter_PatchRequestBody and sets the default values.
func NewItemWithMatter_PatchRequestBody()(*ItemWithMatter_PatchRequestBody) {
    m := &ItemWithMatter_PatchRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateItemWithMatter_PatchRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateItemWithMatter_PatchRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemWithMatter_PatchRequestBody(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *ItemWithMatter_PatchRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetBusinessUnit gets the business_unit property value. The business_unit property
// returns a *string when successful
func (m *ItemWithMatter_PatchRequestBody) GetBusinessUnit()(*string) {
    return m.business_unit
}
// GetCaption gets the caption property value. The caption property
// returns a *string when successful
func (m *ItemWithMatter_PatchRequestBody) GetCaption()(*string) {
    return m.caption
}
// GetCompanyId gets the company_id property value. The company_id property
// returns a *int32 when successful
func (m *ItemWithMatter_PatchRequestBody) GetCompanyId()(*int32) {
    return m.company_id
}
// GetDateEnd gets the date_end property value. The date_end property
// returns a *Time when successful
func (m *ItemWithMatter_PatchRequestBody) GetDateEnd()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.date_end
}
// GetDateStart gets the date_start property value. The date_start property
// returns a *Time when successful
func (m *ItemWithMatter_PatchRequestBody) GetDateStart()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.date_start
}
// GetDepartment gets the department property value. The department property
// returns a *string when successful
func (m *ItemWithMatter_PatchRequestBody) GetDepartment()(*string) {
    return m.department
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ItemWithMatter_PatchRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
// GetName gets the name property value. The name property
// returns a *string when successful
func (m *ItemWithMatter_PatchRequestBody) GetName()(*string) {
    return m.name
}
// GetNotes gets the notes property value. The notes property
// returns a *string when successful
func (m *ItemWithMatter_PatchRequestBody) GetNotes()(*string) {
    return m.notes
}
// GetNumber gets the number property value. The number property
// returns a *string when successful
func (m *ItemWithMatter_PatchRequestBody) GetNumber()(*string) {
    return m.number
}
// Serialize serializes information the current object
func (m *ItemWithMatter_PatchRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
        err := writer.WriteTimeValue("date_end", m.GetDateEnd())
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
func (m *ItemWithMatter_PatchRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetBusinessUnit sets the business_unit property value. The business_unit property
func (m *ItemWithMatter_PatchRequestBody) SetBusinessUnit(value *string)() {
    m.business_unit = value
}
// SetCaption sets the caption property value. The caption property
func (m *ItemWithMatter_PatchRequestBody) SetCaption(value *string)() {
    m.caption = value
}
// SetCompanyId sets the company_id property value. The company_id property
func (m *ItemWithMatter_PatchRequestBody) SetCompanyId(value *int32)() {
    m.company_id = value
}
// SetDateEnd sets the date_end property value. The date_end property
func (m *ItemWithMatter_PatchRequestBody) SetDateEnd(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.date_end = value
}
// SetDateStart sets the date_start property value. The date_start property
func (m *ItemWithMatter_PatchRequestBody) SetDateStart(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.date_start = value
}
// SetDepartment sets the department property value. The department property
func (m *ItemWithMatter_PatchRequestBody) SetDepartment(value *string)() {
    m.department = value
}
// SetName sets the name property value. The name property
func (m *ItemWithMatter_PatchRequestBody) SetName(value *string)() {
    m.name = value
}
// SetNotes sets the notes property value. The notes property
func (m *ItemWithMatter_PatchRequestBody) SetNotes(value *string)() {
    m.notes = value
}
// SetNumber sets the number property value. The number property
func (m *ItemWithMatter_PatchRequestBody) SetNumber(value *string)() {
    m.number = value
}
type ItemWithMatter_PatchRequestBodyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetBusinessUnit()(*string)
    GetCaption()(*string)
    GetCompanyId()(*int32)
    GetDateEnd()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetDateStart()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetDepartment()(*string)
    GetName()(*string)
    GetNotes()(*string)
    GetNumber()(*string)
    SetBusinessUnit(value *string)()
    SetCaption(value *string)()
    SetCompanyId(value *int32)()
    SetDateEnd(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetDateStart(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetDepartment(value *string)()
    SetName(value *string)()
    SetNotes(value *string)()
    SetNumber(value *string)()
}
