package shipment

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"github.com/tealeg/xlsx"
	"time"

)

type Shipment struct {
	ID primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	ShipmentNumber string `json:"shipment_no" bson:"shipment_no"`
	ShipmentItems []primitive.ObjectID `json:"shipment_items" bson:"shipment_items"`
	Vehicle primitive.ObjectID `json:"vehicle" bson:"vehicle"`
	Helper []primitive.ObjectID `json:"helper" bson:"helper"`

	Version  int64 `json:"version" bson:"version"`
	CreateDate     string `json:"create_date" bson:"create_date"`
	CreateBy	primitive.ObjectID `json:"create_by" bson:"create_by"`
	ModifyDate string `json:"modify_date" bson:"modify_date"`
	ModifyBy primitive.ObjectID `json:"modify_by" bson:"modify_by"`
}

func (s *Shipment) SetDefault() {
	s.Version = 1
	t := time.Now()
	s.CreateDate = t.Format("2006-01-02 15:04:05")
}

func (s *Shipment) SetDefaultExcel(create_by primitive.ObjectID, cells []*xlsx.Cell) error {
	s.Version = 1
	t := time.Now()
	s.CreateDate = t.Format("2006-01-02 15:04:05")
	s.CreateBy = create_by

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