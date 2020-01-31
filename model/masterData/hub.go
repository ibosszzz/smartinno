package masterData

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"github.com/tealeg/xlsx"
	"smartinno/controller"
	"time"
)

type Hub struct {
	ID primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	HubCode  string `json:"hub_code" bson:"hub_code" validate:"required"`
	HubName string `json:"hub_name" bson:"hub_name"`
	Active  string `json:"active" bson:"active"`
	Address1 string `json:"address1" bson:"address1"`
	Address2 string `json:"address2" bson:"address2"`
	ZipCode string `json:"zipcode" bson:"zipcode"`
	Version  int64 `json:"version" bson:"version"`
	CreateDate     string `json:"create_date" bson:"create_date"`
	CreateBy	primitive.ObjectID `json:"create_by" bson:"create_by"`
	ModifyDate string `json:"modify_date" bson:"modify_date"`
	ModifyBy primitive.ObjectID `json:"modify_by" bson:"modify_by"`
	ZoneID primitive.ObjectID `json:"zone_id" bson:"zone_id"`

}

func (h *Hub) SetDefault() {
	h.Active = "true"
	h.Version = 1
	t := time.Now()
	h.CreateDate = t.Format("2006-01-02 15:04:05")
}

func (h *Hub) SetDefaultExcel(create_by primitive.ObjectID, cells []*xlsx.Cell) error{
	h.Active = "true"
	h.Version = 1
	t := time.Now()
	h.CreateDate = t.Format("2006-01-02 15:04:05")
	h.CreateBy = create_by

	h.HubCode = cells[0].String()
	h.HubName = cells[1].String()
	h.Address1 = cells[2].String()
	h.Address2 = cells[3].String()
	h.ZipCode = cells[4].String()
	var err error
	h.ZoneID, err = controller.GetOID("zone", "zone_name", cells[5].String())
	if err != nil {
		return err
	}
	return nil
}