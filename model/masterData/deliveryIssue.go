package masterData

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"github.com/tealeg/xlsx"
	"time"
)

type DeliveryIssue struct {
	ID primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	DeliveryIssueName string `json:"delivery_issue_name" bson:"delivery_issue_name"`
	Active  string `json:"active" bson:"active"`
	Version int64 `json:"version" bson:"version"`
	CreateDate     string `json:"create_date" bson:"create_date"`
	CreateBy	primitive.ObjectID `json:"create_by" bson:"create_by"`
	ModifyDate string `json:"modify_date" bson:"modify_date"`
	ModifyBy primitive.ObjectID `json:"modify_by" bson:"modify_by"`

}

func (di *DeliveryIssue) SetDefault() {
	di.Active = "true"
	di.Version = 1
	t := time.Now()
	di.CreateDate = t.Format("2006-01-02 15:04:05")
}

func (di *DeliveryIssue) SetDefaultExcel(create_by primitive.ObjectID, cells []*xlsx.Cell) {
	di.Active = "true"
	di.Version = 1
	t := time.Now()
	di.CreateDate = t.Format("2006-01-02 15:04:05")
	di.CreateBy = create_by

	di.DeliveryIssueName = cells[0].String()
}