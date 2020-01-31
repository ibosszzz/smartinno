package masterData

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"github.com/tealeg/xlsx"
	"time"
	"fmt"
)

type VehicleType struct {
	ID primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	VehicleTypeCode  string `json:"vehicle_type_code" bson:"vehicle_type_code"`
	Description string `json:"description" bson:"description"`
	Wheels int64 `json:"wheels" bson:"wheels"`
	Width int64 `json:"width" bson:"width"`
	Length int64 `json:"length" bson:"length"`
	Height int64 `json:"height" bson:"height"`
	Payload int64 `json:"payload" bson:"payload"`
	Type string `json:"type" bson:"type"`
	Active  string `json:"active" bson:"active"`
	Version  int64 `json:"version" bson:"version"`
	CreateDate     string `json:"create_date" bson:"create_date"`
	CreateBy	primitive.ObjectID `json:"create_by" bson:"create_by"`
	ModifyDate string `json:"modify_date" bson:"modify_date"`
	ModifyBy primitive.ObjectID `json:"modify_by" bson:"modify_by"`

}

func (vt *VehicleType) SetDefault() {
	vt.Active = "true"
	vt.Version = 1
	t := time.Now()
	vt.CreateDate = t.Format("2006-01-02 15:04:05")
}

func (vt *VehicleType) SetDefaultExcel(create_by primitive.ObjectID, cells []*xlsx.Cell) {
	vt.Active = "true"
	vt.Version = 1
	t := time.Now()
	vt.CreateDate = t.Format("2006-01-02 15:04:05")
	vt.CreateBy = create_by

	vt.VehicleTypeCode = cells[0].String()
	vt.Description = cells[1].String()
	// check error
	var err error
	var f float64
	f, err = cells[2].Float()
	if err != nil{
		fmt.Println(err.Error())
	}
	vt.Wheels = int64(f)
	f, err = cells[3].Float()
	if err != nil{
		fmt.Println(err.Error())
	}
	vt.Width = int64(f)
	f, err = cells[4].Float()
	if err != nil{
		fmt.Println(err.Error())
	}
	vt.Length = int64(f)
	f, err = cells[5].Float()
	if err != nil{
		fmt.Println(err.Error())
	}
	vt.Height = int64(f)
	f, err = cells[6].Float()
	if err != nil{
		fmt.Println(err.Error())
	}
	vt.Payload = int64(f)
	vt.Type = cells[7].String()
}