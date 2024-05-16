package customers

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type ItemWithCustomer_PatchRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The billing_start_date property
    billing_start_date *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The combine_hold_reminders property
    combine_hold_reminders *bool
    // The custom_company_label property
    custom_company_label *string
    // The directory assigned to new companies when not delegating CLO.
    default_directory_id *int32
    // The delegate_clo property
    delegate_clo *bool
    // The email_display_name property
    email_display_name *string
    // The email_reply_to property
    email_reply_to *string
    // The email_send_as_domains property
    email_send_as_domains *string
    // The include_active_holds_in_release_notice property
    include_active_holds_in_release_notice *int32
    // The inline_images_as_attachments property
    inline_images_as_attachments *bool
    // The license_code property
    license_code *string
    // The name property
    name *string
    // The salesperson property
    salesperson *string
    // The use_email_send_as property
    use_email_send_as *bool
}
// NewItemWithCustomer_PatchRequestBody instantiates a new ItemWithCustomer_PatchRequestBody and sets the default values.
func NewItemWithCustomer_PatchRequestBody()(*ItemWithCustomer_PatchRequestBody) {
    m := &ItemWithCustomer_PatchRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateItemWithCustomer_PatchRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateItemWithCustomer_PatchRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemWithCustomer_PatchRequestBody(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *ItemWithCustomer_PatchRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetBillingStartDate gets the billing_start_date property value. The billing_start_date property
// returns a *Time when successful
func (m *ItemWithCustomer_PatchRequestBody) GetBillingStartDate()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.billing_start_date
}
// GetCombineHoldReminders gets the combine_hold_reminders property value. The combine_hold_reminders property
// returns a *bool when successful
func (m *ItemWithCustomer_PatchRequestBody) GetCombineHoldReminders()(*bool) {
    return m.combine_hold_reminders
}
// GetCustomCompanyLabel gets the custom_company_label property value. The custom_company_label property
// returns a *string when successful
func (m *ItemWithCustomer_PatchRequestBody) GetCustomCompanyLabel()(*string) {
    return m.custom_company_label
}
// GetDefaultDirectoryId gets the default_directory_id property value. The directory assigned to new companies when not delegating CLO.
// returns a *int32 when successful
func (m *ItemWithCustomer_PatchRequestBody) GetDefaultDirectoryId()(*int32) {
    return m.default_directory_id
}
// GetDelegateClo gets the delegate_clo property value. The delegate_clo property
// returns a *bool when successful
func (m *ItemWithCustomer_PatchRequestBody) GetDelegateClo()(*bool) {
    return m.delegate_clo
}
// GetEmailDisplayName gets the email_display_name property value. The email_display_name property
// returns a *string when successful
func (m *ItemWithCustomer_PatchRequestBody) GetEmailDisplayName()(*string) {
    return m.email_display_name
}
// GetEmailReplyTo gets the email_reply_to property value. The email_reply_to property
// returns a *string when successful
func (m *ItemWithCustomer_PatchRequestBody) GetEmailReplyTo()(*string) {
    return m.email_reply_to
}
// GetEmailSendAsDomains gets the email_send_as_domains property value. The email_send_as_domains property
// returns a *string when successful
func (m *ItemWithCustomer_PatchRequestBody) GetEmailSendAsDomains()(*string) {
    return m.email_send_as_domains
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ItemWithCustomer_PatchRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["billing_start_date"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBillingStartDate(val)
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
    res["custom_company_label"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCustomCompanyLabel(val)
        }
        return nil
    }
    res["default_directory_id"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDefaultDirectoryId(val)
        }
        return nil
    }
    res["delegate_clo"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDelegateClo(val)
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
    res["email_send_as_domains"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEmailSendAsDomains(val)
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
    res["inline_images_as_attachments"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetInlineImagesAsAttachments(val)
        }
        return nil
    }
    res["license_code"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLicenseCode(val)
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
    res["salesperson"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSalesperson(val)
        }
        return nil
    }
    res["use_email_send_as"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUseEmailSendAs(val)
        }
        return nil
    }
    return res
}
// GetIncludeActiveHoldsInReleaseNotice gets the include_active_holds_in_release_notice property value. The include_active_holds_in_release_notice property
// returns a *int32 when successful
func (m *ItemWithCustomer_PatchRequestBody) GetIncludeActiveHoldsInReleaseNotice()(*int32) {
    return m.include_active_holds_in_release_notice
}
// GetInlineImagesAsAttachments gets the inline_images_as_attachments property value. The inline_images_as_attachments property
// returns a *bool when successful
func (m *ItemWithCustomer_PatchRequestBody) GetInlineImagesAsAttachments()(*bool) {
    return m.inline_images_as_attachments
}
// GetLicenseCode gets the license_code property value. The license_code property
// returns a *string when successful
func (m *ItemWithCustomer_PatchRequestBody) GetLicenseCode()(*string) {
    return m.license_code
}
// GetName gets the name property value. The name property
// returns a *string when successful
func (m *ItemWithCustomer_PatchRequestBody) GetName()(*string) {
    return m.name
}
// GetSalesperson gets the salesperson property value. The salesperson property
// returns a *string when successful
func (m *ItemWithCustomer_PatchRequestBody) GetSalesperson()(*string) {
    return m.salesperson
}
// GetUseEmailSendAs gets the use_email_send_as property value. The use_email_send_as property
// returns a *bool when successful
func (m *ItemWithCustomer_PatchRequestBody) GetUseEmailSendAs()(*bool) {
    return m.use_email_send_as
}
// Serialize serializes information the current object
func (m *ItemWithCustomer_PatchRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteTimeValue("billing_start_date", m.GetBillingStartDate())
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
        err := writer.WriteStringValue("custom_company_label", m.GetCustomCompanyLabel())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("default_directory_id", m.GetDefaultDirectoryId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("delegate_clo", m.GetDelegateClo())
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
        err := writer.WriteStringValue("email_send_as_domains", m.GetEmailSendAsDomains())
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
        err := writer.WriteBoolValue("inline_images_as_attachments", m.GetInlineImagesAsAttachments())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("license_code", m.GetLicenseCode())
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
        err := writer.WriteStringValue("salesperson", m.GetSalesperson())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("use_email_send_as", m.GetUseEmailSendAs())
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
func (m *ItemWithCustomer_PatchRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetBillingStartDate sets the billing_start_date property value. The billing_start_date property
func (m *ItemWithCustomer_PatchRequestBody) SetBillingStartDate(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.billing_start_date = value
}
// SetCombineHoldReminders sets the combine_hold_reminders property value. The combine_hold_reminders property
func (m *ItemWithCustomer_PatchRequestBody) SetCombineHoldReminders(value *bool)() {
    m.combine_hold_reminders = value
}
// SetCustomCompanyLabel sets the custom_company_label property value. The custom_company_label property
func (m *ItemWithCustomer_PatchRequestBody) SetCustomCompanyLabel(value *string)() {
    m.custom_company_label = value
}
// SetDefaultDirectoryId sets the default_directory_id property value. The directory assigned to new companies when not delegating CLO.
func (m *ItemWithCustomer_PatchRequestBody) SetDefaultDirectoryId(value *int32)() {
    m.default_directory_id = value
}
// SetDelegateClo sets the delegate_clo property value. The delegate_clo property
func (m *ItemWithCustomer_PatchRequestBody) SetDelegateClo(value *bool)() {
    m.delegate_clo = value
}
// SetEmailDisplayName sets the email_display_name property value. The email_display_name property
func (m *ItemWithCustomer_PatchRequestBody) SetEmailDisplayName(value *string)() {
    m.email_display_name = value
}
// SetEmailReplyTo sets the email_reply_to property value. The email_reply_to property
func (m *ItemWithCustomer_PatchRequestBody) SetEmailReplyTo(value *string)() {
    m.email_reply_to = value
}
// SetEmailSendAsDomains sets the email_send_as_domains property value. The email_send_as_domains property
func (m *ItemWithCustomer_PatchRequestBody) SetEmailSendAsDomains(value *string)() {
    m.email_send_as_domains = value
}
// SetIncludeActiveHoldsInReleaseNotice sets the include_active_holds_in_release_notice property value. The include_active_holds_in_release_notice property
func (m *ItemWithCustomer_PatchRequestBody) SetIncludeActiveHoldsInReleaseNotice(value *int32)() {
    m.include_active_holds_in_release_notice = value
}
// SetInlineImagesAsAttachments sets the inline_images_as_attachments property value. The inline_images_as_attachments property
func (m *ItemWithCustomer_PatchRequestBody) SetInlineImagesAsAttachments(value *bool)() {
    m.inline_images_as_attachments = value
}
// SetLicenseCode sets the license_code property value. The license_code property
func (m *ItemWithCustomer_PatchRequestBody) SetLicenseCode(value *string)() {
    m.license_code = value
}
// SetName sets the name property value. The name property
func (m *ItemWithCustomer_PatchRequestBody) SetName(value *string)() {
    m.name = value
}
// SetSalesperson sets the salesperson property value. The salesperson property
func (m *ItemWithCustomer_PatchRequestBody) SetSalesperson(value *string)() {
    m.salesperson = value
}
// SetUseEmailSendAs sets the use_email_send_as property value. The use_email_send_as property
func (m *ItemWithCustomer_PatchRequestBody) SetUseEmailSendAs(value *bool)() {
    m.use_email_send_as = value
}
type ItemWithCustomer_PatchRequestBodyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetBillingStartDate()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetCombineHoldReminders()(*bool)
    GetCustomCompanyLabel()(*string)
    GetDefaultDirectoryId()(*int32)
    GetDelegateClo()(*bool)
    GetEmailDisplayName()(*string)
    GetEmailReplyTo()(*string)
    GetEmailSendAsDomains()(*string)
    GetIncludeActiveHoldsInReleaseNotice()(*int32)
    GetInlineImagesAsAttachments()(*bool)
    GetLicenseCode()(*string)
    GetName()(*string)
    GetSalesperson()(*string)
    GetUseEmailSendAs()(*bool)
    SetBillingStartDate(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetCombineHoldReminders(value *bool)()
    SetCustomCompanyLabel(value *string)()
    SetDefaultDirectoryId(value *int32)()
    SetDelegateClo(value *bool)()
    SetEmailDisplayName(value *string)()
    SetEmailReplyTo(value *string)()
    SetEmailSendAsDomains(value *string)()
    SetIncludeActiveHoldsInReleaseNotice(value *int32)()
    SetInlineImagesAsAttachments(value *bool)()
    SetLicenseCode(value *string)()
    SetName(value *string)()
    SetSalesperson(value *string)()
    SetUseEmailSendAs(value *bool)()
}
