package shipment

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"github.com/tealeg/xlsx"
	"time"
)

type Invoice struct {
	ID primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	InvoiceNumber string `json:"invoice_no" bson:"invoice_no"`
	Date string `json:"date" bson:"date"`
	ShipTo primitive.ObjectID `json:"ship_to" bson:"ship_to"` // oid of address
	BillTo primitive.ObjectID `json:"bill_to" bson:"bill_to"` // oid of address
	InvoiceItems []primitive.ObjectID `json:"invoice_items" bson:"invoice_items"`

	Version  int64 `json:"version" bson:"version"`
	CreateDate     string `json:"create_date" bson:"create_date"`
	CreateBy	primitive.ObjectID `json:"create_by" bson:"create_by"`
	ModifyDate string `json:"modify_date" bson:"modify_date"`
	ModifyBy primitive.ObjectID `json:"modify_by" bson:"modify_by"`
}

func (i *Invoice) SetDefault() {
	i.Version = 1
	t := time.Now()
	i.Date = t.Format("2006-01-02 15:04:05")
	i.CreateDate = t.Format("2006-01-02 15:04:05")
}

func (i *Invoice) SetDefaultExcel(create_by primitive.ObjectID, cells []*xlsx.Cell) error {
	i.Version = 1
	t := time.Now()
	i.CreateDate = t.Format("2006-01-02 15:04:05")
	i.CreateBy = create_by

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