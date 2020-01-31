package masterData

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"github.com/tealeg/xlsx"
	"smartinno/controller"
	"time"
)

type Carrier struct {
	ID primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	CarrierCode  string `json:"carrier_code" bson:"carrier_code" validate:"required"`
	CarrierName string `json:"carrier_name" bson:"carrier_name"`
	ContactName string `json:"contact_name" bson:"contact_name"`
	Email string `json:"email" bson:"email"`
	Mobile string `json:"mobile" bson:"mobile"`
	Active  string `json:"active" bson:"active"`
	Version  int64 `json:"version" bson:"version"`
	CreateDate     string `json:"create_date" bson:"create_date"`
	CreateBy	primitive.ObjectID `json:"create_by" bson:"create_by"`
	ModifyDate string `json:"modify_date" bson:"modify_date"`
	ModifyBy primitive.ObjectID `json:"modify_by" bson:"modify_by"`
	FleetID primitive.ObjectID `json:"fleet_id" bson:"fleet_id"`

}

func (c *Carrier) SetDefault() {
	c.Active = "true"
	c.Version = 1
	t := time.Now()
	c.CreateDate = t.Format("2006-01-02 15:04:05")
}


func (c *Carrier) SetDefaultExcel(create_by primitive.ObjectID, cells []*xlsx.Cell) error {
	c.Active = "true"
	c.Version = 1
	t := time.Now()
	c.CreateDate = t.Format("2006-01-02 15:04:05")
	c.CreateBy = create_by

	c.CarrierCode = cells[0].String()
	c.CarrierName = cells[1].String()
	c.ContactName = cells[2].String()
	c.Email = cells[3].String()
	c.Mobile = cells[4].String()

	var err error
	c.FleetID, err = controller.GetOID("fleet", "fleet_name", cells[5].String())
	if err != nil {
		return err
	}
	return nil
}