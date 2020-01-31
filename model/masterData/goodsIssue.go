package masterData

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"github.com/tealeg/xlsx"
	"time"
)

type GoodsIssue struct {
	ID primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	GoodsIssueName string `json:"goods_issue_name" bson:"goods_issue_name"`
	Active  string `json:"active" bson:"active"`
	Version int64 `json:"version" bson:"version"`
	CreateDate     string `json:"create_date" bson:"create_date"`
	CreateBy	primitive.ObjectID `json:"create_by" bson:"create_by"`
	ModifyDate string `json:"modify_date" bson:"modify_date"`
	ModifyBy primitive.ObjectID `json:"modify_by" bson:"modify_by"`

}

func (gi *GoodsIssue) SetDefault() {
	gi.Active = "true"
	gi.Version = 1
	t := time.Now()
	gi.CreateDate = t.Format("2006-01-02 15:04:05")
}

func (gi *GoodsIssue) SetDefaultExcel(create_by primitive.ObjectID, cells []*xlsx.Cell) {
	gi.Active = "true"
	gi.Version = 1
	t := time.Now()
	gi.CreateDate = t.Format("2006-01-02 15:04:05")
	gi.CreateBy = create_by

	gi.GoodsIssueName = cells[0].String()
}