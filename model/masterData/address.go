package masterData

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"github.com/tealeg/xlsx"
	"time"
	"fmt"
)

type Address struct {
	ID primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	AddressCode  string `json:"address_code" bson:"address_code"`
	AddressName string `json:"address_name" bson:"address_name"`
	PickupOrDeliveryPoint string `json:"pickup_or_delivery_point" bson:"pickup_or_delivery_point"`
	Building string `json:"building" bson:"building"`
	Address1 string `json:"address1" bson:"address1"`
	Address2 string `json:"address2" bson:"address2"`
	ZipCode string `json:"zipcode" bson:"zipcode"`
	City string `json:"city" bson:"city"`
	Country string `json:"country" bson:"country"`

	ContactMobile string `json:"contact_mobile" bson:"contact_mobile"`
	ContactName string `json:"contact_name" bson:"contact_name"`
	Latitude float64 `json:"latitude" bson:"latitude"`
	Longitude float64 `json:"longitude" bson:"longitude"`
	Radius int64 `json:"radius" bson:"radius"`
	Suggested_Latitude float64 `json:"suggested_latitude" bson:"suggested_latitude"`
	Suggested_Longitude float64 `json:"suggested_longitude" bson:"suggested_longitude"`
	Remark string `json:"remark" bson:"remark"`

	Active  string `json:"active" bson:"active"`
	Version  int64 `json:"version" bson:"version"`
	CreateDate     string `json:"create_date" bson:"create_date"`
	CreateBy	primitive.ObjectID `json:"create_by" bson:"create_by"`
	ModifyDate string `json:"modify_date" bson:"modify_date"`
	ModifyBy primitive.ObjectID `json:"modify_by" bson:"modify_by"`

}

func (a *Address) SetDefault() {
	a.Active = "true"
	a.Version = 1
	a.Radius = 300
	t := time.Now()
	a.CreateDate = t.Format("2006-01-02 15:04:05")
}

func (a *Address) SetDefaultExcel(create_by primitive.ObjectID, cells []*xlsx.Cell) {
	a.Active = "true"
	a.Version = 1
	a.Radius = 300
	t := time.Now()
	a.CreateDate = t.Format("2006-01-02 15:04:05")
	a.CreateBy = create_by

	a.AddressCode = cells[0].String()
	a.AddressName = cells[1].String()
	a.PickupOrDeliveryPoint = cells[2].String()
	a.Building = cells[3].String()
	a.Address1 = cells[4].String()
	a.Address2 = cells[5].String()
	a.ZipCode = cells[6].String()
	a.City = cells[7].String()
	a.Country = cells[8].String()
	a.ContactMobile = cells[9].String()
	a.ContactName = cells[10].String()

	// check error
	var err error
	a.Latitude, err = cells[11].Float()
	if err != nil{
		fmt.Println(err.Error())
	}
	a.Longitude, err = cells[12].Float()
	if err != nil{
		fmt.Println(err.Error())
	}
	var f float64
	f, err = cells[13].Float()
	if err != nil{
		fmt.Println(err.Error())
	}
	a.Radius= int64(f)
	
	a.Remark = cells[14].String()
}
