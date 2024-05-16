package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type Companies struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The address_1 property
    address_1 *string
    // The address_2 property
    address_2 *string
    // The city property
    city *string
    // The contact_email property
    contact_email *string
    // The contact_name property
    contact_name *string
    // The contact_phone property
    contact_phone *string
    // The customer_id property
    customer_id *int32
    // The directory_id property
    directory_id *int32
    // The id property
    id *int32
    // The name property
    name *string
    // The notes property
    notes *string
    // The state property
    state *string
    // The zip property
    zip *string
}
// NewCompanies instantiates a new Companies and sets the default values.
func NewCompanies()(*Companies) {
    m := &Companies{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateCompaniesFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateCompaniesFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewCompanies(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *Companies) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetAddress1 gets the address_1 property value. The address_1 property
// returns a *string when successful
func (m *Companies) GetAddress1()(*string) {
    return m.address_1
}
// GetAddress2 gets the address_2 property value. The address_2 property
// returns a *string when successful
func (m *Companies) GetAddress2()(*string) {
    return m.address_2
}
// GetCity gets the city property value. The city property
// returns a *string when successful
func (m *Companies) GetCity()(*string) {
    return m.city
}
// GetContactEmail gets the contact_email property value. The contact_email property
// returns a *string when successful
func (m *Companies) GetContactEmail()(*string) {
    return m.contact_email
}
// GetContactName gets the contact_name property value. The contact_name property
// returns a *string when successful
func (m *Companies) GetContactName()(*string) {
    return m.contact_name
}
// GetContactPhone gets the contact_phone property value. The contact_phone property
// returns a *string when successful
func (m *Companies) GetContactPhone()(*string) {
    return m.contact_phone
}
// GetCustomerId gets the customer_id property value. The customer_id property
// returns a *int32 when successful
func (m *Companies) GetCustomerId()(*int32) {
    return m.customer_id
}
// GetDirectoryId gets the directory_id property value. The directory_id property
// returns a *int32 when successful
func (m *Companies) GetDirectoryId()(*int32) {
    return m.directory_id
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *Companies) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
    res["directory_id"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDirectoryId(val)
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
// GetId gets the id property value. The id property
// returns a *int32 when successful
func (m *Companies) GetId()(*int32) {
    return m.id
}
// GetName gets the name property value. The name property
// returns a *string when successful
func (m *Companies) GetName()(*string) {
    return m.name
}
// GetNotes gets the notes property value. The notes property
// returns a *string when successful
func (m *Companies) GetNotes()(*string) {
    return m.notes
}
// GetState gets the state property value. The state property
// returns a *string when successful
func (m *Companies) GetState()(*string) {
    return m.state
}
// GetZip gets the zip property value. The zip property
// returns a *string when successful
func (m *Companies) GetZip()(*string) {
    return m.zip
}
// Serialize serializes information the current object
func (m *Companies) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
        err := writer.WriteInt32Value("directory_id", m.GetDirectoryId())
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
func (m *Companies) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetAddress1 sets the address_1 property value. The address_1 property
func (m *Companies) SetAddress1(value *string)() {
    m.address_1 = value
}
// SetAddress2 sets the address_2 property value. The address_2 property
func (m *Companies) SetAddress2(value *string)() {
    m.address_2 = value
}
// SetCity sets the city property value. The city property
func (m *Companies) SetCity(value *string)() {
    m.city = value
}
// SetContactEmail sets the contact_email property value. The contact_email property
func (m *Companies) SetContactEmail(value *string)() {
    m.contact_email = value
}
// SetContactName sets the contact_name property value. The contact_name property
func (m *Companies) SetContactName(value *string)() {
    m.contact_name = value
}
// SetContactPhone sets the contact_phone property value. The contact_phone property
func (m *Companies) SetContactPhone(value *string)() {
    m.contact_phone = value
}
// SetCustomerId sets the customer_id property value. The customer_id property
func (m *Companies) SetCustomerId(value *int32)() {
    m.customer_id = value
}
// SetDirectoryId sets the directory_id property value. The directory_id property
func (m *Companies) SetDirectoryId(value *int32)() {
    m.directory_id = value
}
// SetId sets the id property value. The id property
func (m *Companies) SetId(value *int32)() {
    m.id = value
}
// SetName sets the name property value. The name property
func (m *Companies) SetName(value *string)() {
    m.name = value
}
// SetNotes sets the notes property value. The notes property
func (m *Companies) SetNotes(value *string)() {
    m.notes = value
}
// SetState sets the state property value. The state property
func (m *Companies) SetState(value *string)() {
    m.state = value
}
// SetZip sets the zip property value. The zip property
func (m *Companies) SetZip(value *string)() {
    m.zip = value
}
type Companiesable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAddress1()(*string)
    GetAddress2()(*string)
    GetCity()(*string)
    GetContactEmail()(*string)
    GetContactName()(*string)
    GetContactPhone()(*string)
    GetCustomerId()(*int32)
    GetDirectoryId()(*int32)
    GetId()(*int32)
    GetName()(*string)
    GetNotes()(*string)
    GetState()(*string)
    GetZip()(*string)
    SetAddress1(value *string)()
    SetAddress2(value *string)()
    SetCity(value *string)()
    SetContactEmail(value *string)()
    SetContactName(value *string)()
    SetContactPhone(value *string)()
    SetCustomerId(value *int32)()
    SetDirectoryId(value *int32)()
    SetId(value *int32)()
    SetName(value *string)()
    SetNotes(value *string)()
    SetState(value *string)()
    SetZip(value *string)()
}
