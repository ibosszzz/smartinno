package masterData

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"github.com/tealeg/xlsx"
	"smartinno/controller"
	"time"
)

type Fleet struct {
	ID primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	FleetCode  string `json:"fleet_code" bson:"fleet_code" validate:"required"`
	FleetName string `json:"fleet_name" bson:"fleet_name"`
	Active  string `json:"active" bson:"active"`
	Version  int64 `json:"version" bson:"version"`
	CreateDate     string `json:"create_date" bson:"create_date"`
	CreateBy	primitive.ObjectID `json:"create_by" bson:"create_by"`
	ModifyDate string `json:"modify_date" bson:"modify_date"`
	ModifyBy primitive.ObjectID `json:"modify_by" bson:"modify_by"`
	HubID primitive.ObjectID `json:"hub_id" bson:"hub_id"`

}

func (f *Fleet) SetDefault() {
	f.Active = "true"
	f.Version = 1
	t := time.Now()
	f.CreateDate = t.Format("2006-01-02 15:04:05")
}

func (f *Fleet) SetDefaultExcel(create_by primitive.ObjectID, cells []*xlsx.Cell) error{
	f.Active = "true"
	f.Version = 1
	t := time.Now()
	f.CreateDate = t.Format("2006-01-02 15:04:05")
	f.CreateBy = create_by

	f.FleetCode = cells[0].String()
	f.FleetName = cells[1].String()
	var err error
	f.HubID, err = controller.GetOID("hub", "hub_name", cells[2].String())
	if err != nil {
		return err
	}
	return nil
}