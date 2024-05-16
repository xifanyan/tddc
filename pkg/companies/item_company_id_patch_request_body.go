package companies

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type ItemCompany_idPatchRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The address_1 property
    address_1 *string
    // The address_2 property
    address_2 *string
    // The city property
    city *string
    // The combine_hold_reminders property
    combine_hold_reminders *bool
    // The contact_email property
    contact_email *string
    // The contact_name property
    contact_name *string
    // The contact_phone property
    contact_phone *string
    // The customer_id property
    customer_id *int32
    // The email_display_name property
    email_display_name *string
    // The email_reply_to property
    email_reply_to *string
    // The include_active_holds_in_release_notice property
    include_active_holds_in_release_notice *int32
    // The name property
    name *string
    // The notes property
    notes *string
    // The number property
    number *string
    // The state property
    state *string
    // The zip property
    zip *string
}
// NewItemCompany_idPatchRequestBody instantiates a new ItemCompany_idPatchRequestBody and sets the default values.
func NewItemCompany_idPatchRequestBody()(*ItemCompany_idPatchRequestBody) {
    m := &ItemCompany_idPatchRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateItemCompany_idPatchRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateItemCompany_idPatchRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemCompany_idPatchRequestBody(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *ItemCompany_idPatchRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetAddress1 gets the address_1 property value. The address_1 property
// returns a *string when successful
func (m *ItemCompany_idPatchRequestBody) GetAddress1()(*string) {
    return m.address_1
}
// GetAddress2 gets the address_2 property value. The address_2 property
// returns a *string when successful
func (m *ItemCompany_idPatchRequestBody) GetAddress2()(*string) {
    return m.address_2
}
// GetCity gets the city property value. The city property
// returns a *string when successful
func (m *ItemCompany_idPatchRequestBody) GetCity()(*string) {
    return m.city
}
// GetCombineHoldReminders gets the combine_hold_reminders property value. The combine_hold_reminders property
// returns a *bool when successful
func (m *ItemCompany_idPatchRequestBody) GetCombineHoldReminders()(*bool) {
    return m.combine_hold_reminders
}
// GetContactEmail gets the contact_email property value. The contact_email property
// returns a *string when successful
func (m *ItemCompany_idPatchRequestBody) GetContactEmail()(*string) {
    return m.contact_email
}
// GetContactName gets the contact_name property value. The contact_name property
// returns a *string when successful
func (m *ItemCompany_idPatchRequestBody) GetContactName()(*string) {
    return m.contact_name
}
// GetContactPhone gets the contact_phone property value. The contact_phone property
// returns a *string when successful
func (m *ItemCompany_idPatchRequestBody) GetContactPhone()(*string) {
    return m.contact_phone
}
// GetCustomerId gets the customer_id property value. The customer_id property
// returns a *int32 when successful
func (m *ItemCompany_idPatchRequestBody) GetCustomerId()(*int32) {
    return m.customer_id
}
// GetEmailDisplayName gets the email_display_name property value. The email_display_name property
// returns a *string when successful
func (m *ItemCompany_idPatchRequestBody) GetEmailDisplayName()(*string) {
    return m.email_display_name
}
// GetEmailReplyTo gets the email_reply_to property value. The email_reply_to property
// returns a *string when successful
func (m *ItemCompany_idPatchRequestBody) GetEmailReplyTo()(*string) {
    return m.email_reply_to
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ItemCompany_idPatchRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["address_1"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAddress1(val)
        }
        return nil
    }
    res["address_2"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAddress2(val)
        }
        return nil
    }
    res["city"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCity(val)
        }
        return nil
    }
    res["combine_hold_reminders"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCombineHoldReminders(val)
        }
        return nil
    }
    res["contact_email"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetContactEmail(val)
        }
        return nil
    }
    res["contact_name"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetContactName(val)
        }
        return nil
    }
    res["contact_phone"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetContactPhone(val)
        }
        return nil
    }
    res["customer_id"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCustomerId(val)
        }
        return nil
    }
    res["email_display_name"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEmailDisplayName(val)
        }
        return nil
    }
    res["email_reply_to"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEmailReplyTo(val)
        }
        return nil
    }
    res["include_active_holds_in_release_notice"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIncludeActiveHoldsInReleaseNotice(val)
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
    res["state"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetState(val)
        }
        return nil
    }
    res["zip"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetZip(val)
        }
        return nil
    }
    return res
}
// GetIncludeActiveHoldsInReleaseNotice gets the include_active_holds_in_release_notice property value. The include_active_holds_in_release_notice property
// returns a *int32 when successful
func (m *ItemCompany_idPatchRequestBody) GetIncludeActiveHoldsInReleaseNotice()(*int32) {
    return m.include_active_holds_in_release_notice
}
// GetName gets the name property value. The name property
// returns a *string when successful
func (m *ItemCompany_idPatchRequestBody) GetName()(*string) {
    return m.name
}
// GetNotes gets the notes property value. The notes property
// returns a *string when successful
func (m *ItemCompany_idPatchRequestBody) GetNotes()(*string) {
    return m.notes
}
// GetNumber gets the number property value. The number property
// returns a *string when successful
func (m *ItemCompany_idPatchRequestBody) GetNumber()(*string) {
    return m.number
}
// GetState gets the state property value. The state property
// returns a *string when successful
func (m *ItemCompany_idPatchRequestBody) GetState()(*string) {
    return m.state
}
// GetZip gets the zip property value. The zip property
// returns a *string when successful
func (m *ItemCompany_idPatchRequestBody) GetZip()(*string) {
    return m.zip
}
// Serialize serializes information the current object
func (m *ItemCompany_idPatchRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("address_1", m.GetAddress1())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("address_2", m.GetAddress2())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("city", m.GetCity())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("combine_hold_reminders", m.GetCombineHoldReminders())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("contact_email", m.GetContactEmail())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("contact_name", m.GetContactName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("contact_phone", m.GetContactPhone())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("customer_id", m.GetCustomerId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("email_display_name", m.GetEmailDisplayName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("email_reply_to", m.GetEmailReplyTo())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("include_active_holds_in_release_notice", m.GetIncludeActiveHoldsInReleaseNotice())
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
        err := writer.WriteStringValue("state", m.GetState())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("zip", m.GetZip())
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
func (m *ItemCompany_idPatchRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetAddress1 sets the address_1 property value. The address_1 property
func (m *ItemCompany_idPatchRequestBody) SetAddress1(value *string)() {
    m.address_1 = value
}
// SetAddress2 sets the address_2 property value. The address_2 property
func (m *ItemCompany_idPatchRequestBody) SetAddress2(value *string)() {
    m.address_2 = value
}
// SetCity sets the city property value. The city property
func (m *ItemCompany_idPatchRequestBody) SetCity(value *string)() {
    m.city = value
}
// SetCombineHoldReminders sets the combine_hold_reminders property value. The combine_hold_reminders property
func (m *ItemCompany_idPatchRequestBody) SetCombineHoldReminders(value *bool)() {
    m.combine_hold_reminders = value
}
// SetContactEmail sets the contact_email property value. The contact_email property
func (m *ItemCompany_idPatchRequestBody) SetContactEmail(value *string)() {
    m.contact_email = value
}
// SetContactName sets the contact_name property value. The contact_name property
func (m *ItemCompany_idPatchRequestBody) SetContactName(value *string)() {
    m.contact_name = value
}
// SetContactPhone sets the contact_phone property value. The contact_phone property
func (m *ItemCompany_idPatchRequestBody) SetContactPhone(value *string)() {
    m.contact_phone = value
}
// SetCustomerId sets the customer_id property value. The customer_id property
func (m *ItemCompany_idPatchRequestBody) SetCustomerId(value *int32)() {
    m.customer_id = value
}
// SetEmailDisplayName sets the email_display_name property value. The email_display_name property
func (m *ItemCompany_idPatchRequestBody) SetEmailDisplayName(value *string)() {
    m.email_display_name = value
}
// SetEmailReplyTo sets the email_reply_to property value. The email_reply_to property
func (m *ItemCompany_idPatchRequestBody) SetEmailReplyTo(value *string)() {
    m.email_reply_to = value
}
// SetIncludeActiveHoldsInReleaseNotice sets the include_active_holds_in_release_notice property value. The include_active_holds_in_release_notice property
func (m *ItemCompany_idPatchRequestBody) SetIncludeActiveHoldsInReleaseNotice(value *int32)() {
    m.include_active_holds_in_release_notice = value
}
// SetName sets the name property value. The name property
func (m *ItemCompany_idPatchRequestBody) SetName(value *string)() {
    m.name = value
}
// SetNotes sets the notes property value. The notes property
func (m *ItemCompany_idPatchRequestBody) SetNotes(value *string)() {
    m.notes = value
}
// SetNumber sets the number property value. The number property
func (m *ItemCompany_idPatchRequestBody) SetNumber(value *string)() {
    m.number = value
}
// SetState sets the state property value. The state property
func (m *ItemCompany_idPatchRequestBody) SetState(value *string)() {
    m.state = value
}
// SetZip sets the zip property value. The zip property
func (m *ItemCompany_idPatchRequestBody) SetZip(value *string)() {
    m.zip = value
}
type ItemCompany_idPatchRequestBodyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAddress1()(*string)
    GetAddress2()(*string)
    GetCity()(*string)
    GetCombineHoldReminders()(*bool)
    GetContactEmail()(*string)
    GetContactName()(*string)
    GetContactPhone()(*string)
    GetCustomerId()(*int32)
    GetEmailDisplayName()(*string)
    GetEmailReplyTo()(*string)
    GetIncludeActiveHoldsInReleaseNotice()(*int32)
    GetName()(*string)
    GetNotes()(*string)
    GetNumber()(*string)
    GetState()(*string)
    GetZip()(*string)
    SetAddress1(value *string)()
    SetAddress2(value *string)()
    SetCity(value *string)()
    SetCombineHoldReminders(value *bool)()
    SetContactEmail(value *string)()
    SetContactName(value *string)()
    SetContactPhone(value *string)()
    SetCustomerId(value *int32)()
    SetEmailDisplayName(value *string)()
    SetEmailReplyTo(value *string)()
    SetIncludeActiveHoldsInReleaseNotice(value *int32)()
    SetName(value *string)()
    SetNotes(value *string)()
    SetNumber(value *string)()
    SetState(value *string)()
    SetZip(value *string)()
}
