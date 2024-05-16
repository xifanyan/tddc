package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type Custodian struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The business property
    business *string
    // The country property
    country *string
    // The department property
    department *string
    // The directory_id property
    directory_id *int32
    // The email property
    email *string
    // Formerly employee_number.
    employee_identifier *string
    // The employee_status property
    employee_status *string
    // The employee_type property
    employee_type *string
    // The external_sync property
    external_sync *bool
    // The first_name property
    first_name *string
    // The function property
    function *string
    // The id property
    id *int32
    // The last_name property
    last_name *string
    // The location property
    location *string
    // The name property
    name *string
    // The notes property
    notes *string
    // The phone property
    phone *string
    // The supervisor_email property
    supervisor_email *string
    // The supervisor_name property
    supervisor_name *string
    // The title property
    title *string
}
// NewCustodian instantiates a new Custodian and sets the default values.
func NewCustodian()(*Custodian) {
    m := &Custodian{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateCustodianFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateCustodianFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewCustodian(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *Custodian) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetBusiness gets the business property value. The business property
// returns a *string when successful
func (m *Custodian) GetBusiness()(*string) {
    return m.business
}
// GetCountry gets the country property value. The country property
// returns a *string when successful
func (m *Custodian) GetCountry()(*string) {
    return m.country
}
// GetDepartment gets the department property value. The department property
// returns a *string when successful
func (m *Custodian) GetDepartment()(*string) {
    return m.department
}
// GetDirectoryId gets the directory_id property value. The directory_id property
// returns a *int32 when successful
func (m *Custodian) GetDirectoryId()(*int32) {
    return m.directory_id
}
// GetEmail gets the email property value. The email property
// returns a *string when successful
func (m *Custodian) GetEmail()(*string) {
    return m.email
}
// GetEmployeeIdentifier gets the employee_identifier property value. Formerly employee_number.
// returns a *string when successful
func (m *Custodian) GetEmployeeIdentifier()(*string) {
    return m.employee_identifier
}
// GetEmployeeStatus gets the employee_status property value. The employee_status property
// returns a *string when successful
func (m *Custodian) GetEmployeeStatus()(*string) {
    return m.employee_status
}
// GetEmployeeType gets the employee_type property value. The employee_type property
// returns a *string when successful
func (m *Custodian) GetEmployeeType()(*string) {
    return m.employee_type
}
// GetExternalSync gets the external_sync property value. The external_sync property
// returns a *bool when successful
func (m *Custodian) GetExternalSync()(*bool) {
    return m.external_sync
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *Custodian) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["business"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBusiness(val)
        }
        return nil
    }
    res["country"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCountry(val)
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
    res["email"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEmail(val)
        }
        return nil
    }
    res["employee_identifier"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEmployeeIdentifier(val)
        }
        return nil
    }
    res["employee_status"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEmployeeStatus(val)
        }
        return nil
    }
    res["employee_type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEmployeeType(val)
        }
        return nil
    }
    res["external_sync"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetExternalSync(val)
        }
        return nil
    }
    res["first_name"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFirstName(val)
        }
        return nil
    }
    res["function"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFunction(val)
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
    res["last_name"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastName(val)
        }
        return nil
    }
    res["location"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLocation(val)
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
    res["phone"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPhone(val)
        }
        return nil
    }
    res["supervisor_email"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSupervisorEmail(val)
        }
        return nil
    }
    res["supervisor_name"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSupervisorName(val)
        }
        return nil
    }
    res["title"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTitle(val)
        }
        return nil
    }
    return res
}
// GetFirstName gets the first_name property value. The first_name property
// returns a *string when successful
func (m *Custodian) GetFirstName()(*string) {
    return m.first_name
}
// GetFunction gets the function property value. The function property
// returns a *string when successful
func (m *Custodian) GetFunction()(*string) {
    return m.function
}
// GetId gets the id property value. The id property
// returns a *int32 when successful
func (m *Custodian) GetId()(*int32) {
    return m.id
}
// GetLastName gets the last_name property value. The last_name property
// returns a *string when successful
func (m *Custodian) GetLastName()(*string) {
    return m.last_name
}
// GetLocation gets the location property value. The location property
// returns a *string when successful
func (m *Custodian) GetLocation()(*string) {
    return m.location
}
// GetName gets the name property value. The name property
// returns a *string when successful
func (m *Custodian) GetName()(*string) {
    return m.name
}
// GetNotes gets the notes property value. The notes property
// returns a *string when successful
func (m *Custodian) GetNotes()(*string) {
    return m.notes
}
// GetPhone gets the phone property value. The phone property
// returns a *string when successful
func (m *Custodian) GetPhone()(*string) {
    return m.phone
}
// GetSupervisorEmail gets the supervisor_email property value. The supervisor_email property
// returns a *string when successful
func (m *Custodian) GetSupervisorEmail()(*string) {
    return m.supervisor_email
}
// GetSupervisorName gets the supervisor_name property value. The supervisor_name property
// returns a *string when successful
func (m *Custodian) GetSupervisorName()(*string) {
    return m.supervisor_name
}
// GetTitle gets the title property value. The title property
// returns a *string when successful
func (m *Custodian) GetTitle()(*string) {
    return m.title
}
// Serialize serializes information the current object
func (m *Custodian) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("business", m.GetBusiness())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("country", m.GetCountry())
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
        err := writer.WriteInt32Value("directory_id", m.GetDirectoryId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("email", m.GetEmail())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("employee_identifier", m.GetEmployeeIdentifier())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("employee_status", m.GetEmployeeStatus())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("employee_type", m.GetEmployeeType())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("external_sync", m.GetExternalSync())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("first_name", m.GetFirstName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("function", m.GetFunction())
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
        err := writer.WriteStringValue("last_name", m.GetLastName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("location", m.GetLocation())
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
        err := writer.WriteStringValue("phone", m.GetPhone())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("supervisor_email", m.GetSupervisorEmail())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("supervisor_name", m.GetSupervisorName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("title", m.GetTitle())
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
func (m *Custodian) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetBusiness sets the business property value. The business property
func (m *Custodian) SetBusiness(value *string)() {
    m.business = value
}
// SetCountry sets the country property value. The country property
func (m *Custodian) SetCountry(value *string)() {
    m.country = value
}
// SetDepartment sets the department property value. The department property
func (m *Custodian) SetDepartment(value *string)() {
    m.department = value
}
// SetDirectoryId sets the directory_id property value. The directory_id property
func (m *Custodian) SetDirectoryId(value *int32)() {
    m.directory_id = value
}
// SetEmail sets the email property value. The email property
func (m *Custodian) SetEmail(value *string)() {
    m.email = value
}
// SetEmployeeIdentifier sets the employee_identifier property value. Formerly employee_number.
func (m *Custodian) SetEmployeeIdentifier(value *string)() {
    m.employee_identifier = value
}
// SetEmployeeStatus sets the employee_status property value. The employee_status property
func (m *Custodian) SetEmployeeStatus(value *string)() {
    m.employee_status = value
}
// SetEmployeeType sets the employee_type property value. The employee_type property
func (m *Custodian) SetEmployeeType(value *string)() {
    m.employee_type = value
}
// SetExternalSync sets the external_sync property value. The external_sync property
func (m *Custodian) SetExternalSync(value *bool)() {
    m.external_sync = value
}
// SetFirstName sets the first_name property value. The first_name property
func (m *Custodian) SetFirstName(value *string)() {
    m.first_name = value
}
// SetFunction sets the function property value. The function property
func (m *Custodian) SetFunction(value *string)() {
    m.function = value
}
// SetId sets the id property value. The id property
func (m *Custodian) SetId(value *int32)() {
    m.id = value
}
// SetLastName sets the last_name property value. The last_name property
func (m *Custodian) SetLastName(value *string)() {
    m.last_name = value
}
// SetLocation sets the location property value. The location property
func (m *Custodian) SetLocation(value *string)() {
    m.location = value
}
// SetName sets the name property value. The name property
func (m *Custodian) SetName(value *string)() {
    m.name = value
}
// SetNotes sets the notes property value. The notes property
func (m *Custodian) SetNotes(value *string)() {
    m.notes = value
}
// SetPhone sets the phone property value. The phone property
func (m *Custodian) SetPhone(value *string)() {
    m.phone = value
}
// SetSupervisorEmail sets the supervisor_email property value. The supervisor_email property
func (m *Custodian) SetSupervisorEmail(value *string)() {
    m.supervisor_email = value
}
// SetSupervisorName sets the supervisor_name property value. The supervisor_name property
func (m *Custodian) SetSupervisorName(value *string)() {
    m.supervisor_name = value
}
// SetTitle sets the title property value. The title property
func (m *Custodian) SetTitle(value *string)() {
    m.title = value
}
type Custodianable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetBusiness()(*string)
    GetCountry()(*string)
    GetDepartment()(*string)
    GetDirectoryId()(*int32)
    GetEmail()(*string)
    GetEmployeeIdentifier()(*string)
    GetEmployeeStatus()(*string)
    GetEmployeeType()(*string)
    GetExternalSync()(*bool)
    GetFirstName()(*string)
    GetFunction()(*string)
    GetId()(*int32)
    GetLastName()(*string)
    GetLocation()(*string)
    GetName()(*string)
    GetNotes()(*string)
    GetPhone()(*string)
    GetSupervisorEmail()(*string)
    GetSupervisorName()(*string)
    GetTitle()(*string)
    SetBusiness(value *string)()
    SetCountry(value *string)()
    SetDepartment(value *string)()
    SetDirectoryId(value *int32)()
    SetEmail(value *string)()
    SetEmployeeIdentifier(value *string)()
    SetEmployeeStatus(value *string)()
    SetEmployeeType(value *string)()
    SetExternalSync(value *bool)()
    SetFirstName(value *string)()
    SetFunction(value *string)()
    SetId(value *int32)()
    SetLastName(value *string)()
    SetLocation(value *string)()
    SetName(value *string)()
    SetNotes(value *string)()
    SetPhone(value *string)()
    SetSupervisorEmail(value *string)()
    SetSupervisorName(value *string)()
    SetTitle(value *string)()
}
