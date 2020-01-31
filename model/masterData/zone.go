package masterData

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"github.com/tealeg/xlsx"
	"time"
)

type Zone struct {
	ID primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	ZoneCode  string `json:"zone_code" bson:"zone_code" validate:"required"`
	ZoneName string `json:"zone_name" bson:"zone_name" validate:"required"`
	Active  string `json:"active" bson:"active"`
	Version  int64 `json:"version" bson:"version"`
	CreateDate  string  `json:"create_date" bson:"create_date"`
	CreateBy	primitive.ObjectID `json:"create_by" bson:"create_by"`
	ModifyDate string `json:"modify_date" bson:"modify_date"`
	ModifyBy primitive.ObjectID `json:"modify_by" bson:"modify_by"`

}

func (z *Zone) SetDefault() {
	z.Active = "true"
	z.Version = 1
	t := time.Now()
	z.CreateDate = t.Format("2006-01-02 15:04:05")
}

func (z *Zone) SetDefaultExcel(create_by primitive.ObjectID, cells []*xlsx.Cell) {
	z.Active = "true"
	z.Version = 1
	t := time.Now()
	z.CreateDate = t.Format("2006-01-02 15:04:05")
	z.CreateBy = create_by

	z.ZoneCode = cells[0].String()
	z.ZoneName = cells[1].String()

}