package shipment

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"github.com/tealeg/xlsx"
	"time"
)

type ShipmentItem struct {
	ID primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	ShipmentItemNumber string `json:"shipment_item_no" bson:"shipment_item_no"` // DO Number
	Invoices []primitive.ObjectID `json:"invoices" bson:"invoices"`
	Type string `json:"type" bson:"type"` // pickup, delivery
	PlanCheckIn string `json:"plan_check_in" bson:"plan_check_in"`
	PlanCheckOut string `json:"plan_check_out" bson:"plan_check_out"`
	DeliveryCondition string `json:"delivery_condition" bson:"delivery_condition"` // ex time
	DeliveryReamrk string `json:"delivery_remark" bson:"delivery_remark"` // text

	Version  int64 `json:"version" bson:"version"`
	CreateDate     string `json:"create_date" bson:"create_date"`
	CreateBy	primitive.ObjectID `json:"create_by" bson:"create_by"`
	ModifyDate string `json:"modify_date" bson:"modify_date"`
	ModifyBy primitive.ObjectID `json:"modify_by" bson:"modify_by"`
}

func (si *ShipmentItem) SetDefault() {
	si.Version = 1
	t := time.Now()
	si.CreateDate = t.Format("2006-01-02 15:04:05")
}

func (si *ShipmentItem) SetDefaultExcel(create_by primitive.ObjectID, cells []*xlsx.Cell) error {
	si.Version = 1
	t := time.Now()
	si.CreateDate = t.Format("2006-01-02 15:04:05")
	si.CreateBy = create_by

	// i.CarrierCode = cells[0].String()
	// i.CarrierName = cells[1].String()
	// i.ContactName = cells[2].String()
	// i.Email = cells[3].String()
	// c.Mobile = cells[4].String()

	// var err error
	// c.FleetID, err = controller.GetOID("fleet", "fleet_name", cells[5].String())
	// if err != nil {
	// 	return err
	// }
	return nil
}