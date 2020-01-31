package masterData

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"github.com/tealeg/xlsx"
	"smartinno/controller"
	"time"
)

type Helper struct {
	ID primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	HelperName string `json:"helper_name" bson:"helper_name"`
	HelperNickName string `json:"helper_nickname" bson:"helper_nickname"`
	PersonalID string `json:"personal_id" bson:"personal_id"`
	Mobile string `json:"mobile" bson:"mobile"`
	DateOfBirth string `json:"date_of_birth" bson:"date_of_birth"`
	Active  string `json:"active" bson:"active"`
	Version  int64 `json:"version" bson:"version"`
	CreateDate     string `json:"create_date" bson:"create_date"`
	CreateBy	 primitive.ObjectID `json:"create_by" bson:"create_by"`
	ModifyDate string `json:"modify_date" bson:"modify_date"`
	ModifyBy  primitive.ObjectID `json:"modify_by" bson:"modify_by"`
	CarrierID primitive.ObjectID `json:"carrier_id" bson:"carrier_id"`

}

func (h *Helper) SetDefault() {
	h.Active = "true"
	h.Version = 1
	t := time.Now()
	h.CreateDate = t.Format("2006-01-02 15:04:05")
}

func (h *Helper) SetDefaultExcel(create_by primitive.ObjectID, cells []*xlsx.Cell) error {
	h.Active = "true"
	h.Version = 1
	t := time.Now()
	h.CreateDate = t.Format("2006-01-02 15:04:05")
	h.CreateBy = create_by

	h.HelperName = cells[0].String()
	h.HelperNickName = cells[1].String()
	h.PersonalID = cells[2].String()
	h.Mobile = cells[3].String()
	h.DateOfBirth = cells[4].String()
	var err error
	h.CarrierID, err = controller.GetOID("carrier", "carrier_name", cells[5].String())
	if err != nil {
		return err
	}
	return nil
}