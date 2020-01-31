package masterData

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"github.com/tealeg/xlsx"
	"smartinno/controller"
	"time"
)

type Driver struct {
	ID primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	DriverName string `json:"driver_name" bson:"driver_name"`
	PersonalID string `json:"personal_id" bson:"personal_id"`
	Mobile string `json:"mobile" bson:"mobile"`
	DateOfBirth string `json:"date_of_birth" bson:"date_of_birth"`
	DriverLicenseID string `json:"driver_license_id" bson:"driver_license_id"`
	DriverLicenseExpired string `json:"driver_license_expired" bson:"driver_license_expired"`

	Username string `json:"username" bson:"username"`
	Password string `json:"password" bson:"password"`

	ReadyToMatch string `json:"ready_to_match" bson:"ready_to_match"`
	Active  string `json:"active" bson:"active"`
	Version  int64 `json:"version" bson:"version"`
	CreateDate    string  `json:"create_date" bson:"create_date"`
	CreateBy	primitive.ObjectID `json:"create_by" bson:"create_by"`
	ModifyDate string `json:"modify_date" bson:"modify_date"`
	ModifyBy primitive.ObjectID `json:"modify_by" bson:"modify_by"`
	CarrierID primitive.ObjectID `json:"carrier_id" bson:"carrier_id"`

}

func (d *Driver) SetDefault() {
	d.Active = "true"
	d.ReadyToMatch = "true"
	d.Version = 1
	t := time.Now()
	d.CreateDate = t.Format("2006-01-02 15:04:05")
}

func (d *Driver) SetDefaultExcel(create_by primitive.ObjectID, cells []*xlsx.Cell) error {
	d.Active = "true"
	d.ReadyToMatch = "true"
	d.Version = 1
	t := time.Now()
	d.CreateDate = t.Format("2006-01-02 15:04:05")
	d.CreateBy = create_by

	d.DriverName = cells[0].String()
	d.PersonalID = cells[1].String()
	d.Mobile = cells[2].String()
	d.DateOfBirth = cells[3].String()
	d.DriverLicenseID = cells[4].String()
	d.DriverLicenseExpired = cells[5].String()
	d.Username = cells[6].String()
	d.Password = cells[7].String()
	var err error
	d.CarrierID, err = controller.GetOID("carrier", "carrier_name", cells[8].String())
	if err != nil {
		return err
	}
	return nil
}