package shipment

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"github.com/tealeg/xlsx"
	"time"
)

type InvoiceItem struct {
	ID primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Item string `json:"item" bson:"item"`
	Quantity int `json:"quantity" bson:"quantity"`
	Remark string `json:"remark" bson:"remark"`
	Version  int64 `json:"version" bson:"version"`
	CreateDate     string `json:"create_date" bson:"create_date"`
	CreateBy	primitive.ObjectID `json:"create_by" bson:"create_by"`
	ModifyDate string `json:"modify_date" bson:"modify_date"`
	ModifyBy primitive.ObjectID `json:"modify_by" bson:"modify_by"`
}

func (invItem *InvoiceItem) SetDefault() {
	invItem.Version = 1
	t := time.Now()
	invItem.CreateDate = t.Format("2006-01-02 15:04:05")
}

func (invItem *InvoiceItem) SetDefaultExcel(create_by primitive.ObjectID, cells []*xlsx.Cell) error {
	invItem.Version = 1
	t := time.Now()
	invItem.CreateDate = t.Format("2006-01-02 15:04:05")
	invItem.CreateBy = create_by

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