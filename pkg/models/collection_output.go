package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type CollectionOutput struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The collection_end_date property
    collection_end_date *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The collection_start_date property
    collection_start_date *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The created_at property
    created_at *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The custodian_id property
    custodian_id *int32
    // The custodian_name property
    custodian_name *string
    // The id property
    id *int32
    // The item_count property
    item_count *int32
    // The item_size_in_bytes property
    item_size_in_bytes *int32
    // The matter_id property
    matter_id *int32
    // The name property
    name *string
    // The status_type property
    status_type *string
}
// NewCollectionOutput instantiates a new CollectionOutput and sets the default values.
func NewCollectionOutput()(*CollectionOutput) {
    m := &CollectionOutput{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateCollectionOutputFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateCollectionOutputFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewCollectionOutput(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *CollectionOutput) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetCollectionEndDate gets the collection_end_date property value. The collection_end_date property
// returns a *Time when successful
func (m *CollectionOutput) GetCollectionEndDate()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.collection_end_date
}
// GetCollectionStartDate gets the collection_start_date property value. The collection_start_date property
// returns a *Time when successful
func (m *CollectionOutput) GetCollectionStartDate()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.collection_start_date
}
// GetCreatedAt gets the created_at property value. The created_at property
// returns a *Time when successful
func (m *CollectionOutput) GetCreatedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.created_at
}
// GetCustodianId gets the custodian_id property value. The custodian_id property
// returns a *int32 when successful
func (m *CollectionOutput) GetCustodianId()(*int32) {
    return m.custodian_id
}
// GetCustodianName gets the custodian_name property value. The custodian_name property
// returns a *string when successful
func (m *CollectionOutput) GetCustodianName()(*string) {
    return m.custodian_name
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *CollectionOutput) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["collection_end_date"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCollectionEndDate(val)
        }
        return nil
    }
    res["collection_start_date"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCollectionStartDate(val)
        }
        return nil
    }
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
    res["custodian_id"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCustodianId(val)
        }
        return nil
    }
    res["custodian_name"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCustodianName(val)
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
    res["item_count"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetItemCount(val)
        }
        return nil
    }
    res["item_size_in_bytes"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetItemSizeInBytes(val)
        }
        return nil
    }
    res["matter_id"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMatterId(val)
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
// GetId gets the id property value. The id property
// returns a *int32 when successful
func (m *CollectionOutput) GetId()(*int32) {
    return m.id
}
// GetItemCount gets the item_count property value. The item_count property
// returns a *int32 when successful
func (m *CollectionOutput) GetItemCount()(*int32) {
    return m.item_count
}
// GetItemSizeInBytes gets the item_size_in_bytes property value. The item_size_in_bytes property
// returns a *int32 when successful
func (m *CollectionOutput) GetItemSizeInBytes()(*int32) {
    return m.item_size_in_bytes
}
// GetMatterId gets the matter_id property value. The matter_id property
// returns a *int32 when successful
func (m *CollectionOutput) GetMatterId()(*int32) {
    return m.matter_id
}
// GetName gets the name property value. The name property
// returns a *string when successful
func (m *CollectionOutput) GetName()(*string) {
    return m.name
}
// GetStatusType gets the status_type property value. The status_type property
// returns a *string when successful
func (m *CollectionOutput) GetStatusType()(*string) {
    return m.status_type
}
// Serialize serializes information the current object
func (m *CollectionOutput) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteTimeValue("collection_end_date", m.GetCollectionEndDate())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteTimeValue("collection_start_date", m.GetCollectionStartDate())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteTimeValue("created_at", m.GetCreatedAt())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("custodian_id", m.GetCustodianId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("custodian_name", m.GetCustodianName())
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
        err := writer.WriteInt32Value("item_count", m.GetItemCount())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("item_size_in_bytes", m.GetItemSizeInBytes())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("matter_id", m.GetMatterId())
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
func (m *CollectionOutput) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetCollectionEndDate sets the collection_end_date property value. The collection_end_date property
func (m *CollectionOutput) SetCollectionEndDate(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.collection_end_date = value
}
// SetCollectionStartDate sets the collection_start_date property value. The collection_start_date property
func (m *CollectionOutput) SetCollectionStartDate(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.collection_start_date = value
}
// SetCreatedAt sets the created_at property value. The created_at property
func (m *CollectionOutput) SetCreatedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.created_at = value
}
// SetCustodianId sets the custodian_id property value. The custodian_id property
func (m *CollectionOutput) SetCustodianId(value *int32)() {
    m.custodian_id = value
}
// SetCustodianName sets the custodian_name property value. The custodian_name property
func (m *CollectionOutput) SetCustodianName(value *string)() {
    m.custodian_name = value
}
// SetId sets the id property value. The id property
func (m *CollectionOutput) SetId(value *int32)() {
    m.id = value
}
// SetItemCount sets the item_count property value. The item_count property
func (m *CollectionOutput) SetItemCount(value *int32)() {
    m.item_count = value
}
// SetItemSizeInBytes sets the item_size_in_bytes property value. The item_size_in_bytes property
func (m *CollectionOutput) SetItemSizeInBytes(value *int32)() {
    m.item_size_in_bytes = value
}
// SetMatterId sets the matter_id property value. The matter_id property
func (m *CollectionOutput) SetMatterId(value *int32)() {
    m.matter_id = value
}
// SetName sets the name property value. The name property
func (m *CollectionOutput) SetName(value *string)() {
    m.name = value
}
// SetStatusType sets the status_type property value. The status_type property
func (m *CollectionOutput) SetStatusType(value *string)() {
    m.status_type = value
}
type CollectionOutputable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCollectionEndDate()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetCollectionStartDate()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetCreatedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetCustodianId()(*int32)
    GetCustodianName()(*string)
    GetId()(*int32)
    GetItemCount()(*int32)
    GetItemSizeInBytes()(*int32)
    GetMatterId()(*int32)
    GetName()(*string)
    GetStatusType()(*string)
    SetCollectionEndDate(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetCollectionStartDate(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetCreatedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetCustodianId(value *int32)()
    SetCustodianName(value *string)()
    SetId(value *int32)()
    SetItemCount(value *int32)()
    SetItemSizeInBytes(value *int32)()
    SetMatterId(value *int32)()
    SetName(value *string)()
    SetStatusType(value *string)()
}
