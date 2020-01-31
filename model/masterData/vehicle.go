package masterData

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"github.com/tealeg/xlsx"
	"smartinno/controller"
	"time"
	"fmt"
)

type Vehicle struct {
	ID primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	License  string `json:"license" bson:"license" xlsx:"0"`
	LicenseExpired  string `json:"license_expired" bson:"license_expired" xlsx:"1"`
	Mileage int64 `json:"mileage" bson:"mileage" xlsx:"2"`
	MileageStart int64 `json:"mileage_start" bson:"mileage_start" xlsx:"3"`
	MileageStartDate string `json:"mileage_start_date" bson:"mileage_start_date" xlsx:"4"`
	Remark string `json:"remark" bson:"remark" xlsx:"5"`
	Active  string `json:"active" bson:"active"`
	Version  int64 `json:"version" bson:"version"`
	CreateDate     string `json:"create_date" bson:"create_date"`
	CreateBy	primitive.ObjectID `json:"create_by" bson:"create_by"`
	ModifyDate string `json:"modify_date" bson:"modify_date"`
	ModifyBy primitive.ObjectID `json:"modify_by" bson:"modify_by"`
	CarrierID primitive.ObjectID `json:"carrier_id" bson:"carrier_id"`
	VehicleTypeID primitive.ObjectID `json:"vehicle_type_id" bson:"vehicle_type_id"`
	DriverID primitive.ObjectID `json:"driver_id" bson:"driver_id"`

}

func (v *Vehicle) SetDefault() {
	v.Active = "true"
	v.Version = 1
	t := time.Now()
	v.CreateDate = t.Format("2006-01-02 15:04:05")
}

func (v *Vehicle) SetDefaultExcel(create_by primitive.ObjectID, cells []*xlsx.Cell) error {
	v.Active = "true"
	v.Version = 1
	t := time.Now()
	v.CreateDate = t.Format("2006-01-02 15:04:05")
	v.CreateBy = create_by

	v.License = cells[0].String()
	v.LicenseExpired = cells[1].String()

	// check error
	var err error
	var f float64
	f, err = cells[2].Float()
	v.Mileage = int64(f)
	if err != nil {
		fmt.Println(err.Error())
	}
	f, err = cells[3].Float()
	if err != nil {
		fmt.Println(err.Error())
	}
	v.MileageStart = int64(f)
	v.MileageStartDate = cells[4].String() +" "+ cells[5].String()
	v.Remark = cells[6].String()
	v.CarrierID, err = controller.GetOID("carrier", "carrier_name", cells[7].String())
	if err != nil {
		return err
	}
	v.VehicleTypeID, err = controller.GetOID("vehicle_type", "description", cells[8].String())
	if err != nil {
		return err
	}
	return nil
}

