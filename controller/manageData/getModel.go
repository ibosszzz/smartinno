package manageData

import (
	master_data_model "smartinno/model/masterData"
	shipment_model "smartinno/model/shipment"
	"github.com/tealeg/xlsx"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetModelAndSetDefaultValue(collection_name string) interface{}{
	if collection_name == "zone" {
		m:= master_data_model.Zone{}
		m.SetDefault()
		return &m
	} else if collection_name == "hub"{
		m:= master_data_model.Hub{}
		m.SetDefault()
		return &m
	} else if collection_name == "fleet"{
		m:= master_data_model.Fleet{}
		m.SetDefault()
		return &m
	} else if collection_name == "carrier"{
		m:= master_data_model.Carrier{}
		m.SetDefault()
		return &m
	} else if collection_name == "vehicle"{
		m:= master_data_model.Vehicle{}
		m.SetDefault()
		return &m
	} else if collection_name == "vehicle_type"{
		m:= master_data_model.VehicleType{}
		m.SetDefault()
		return &m
	} else if collection_name == "driver"{
		m:= master_data_model.Driver{}
		m.SetDefault()
		return &m
	} else if collection_name == "helper"{
		m:= master_data_model.Helper{}
		m.SetDefault()
		return &m
	} else if collection_name == "address"{
		m:= master_data_model.Address{}
		m.SetDefault()
		return &m
	} else if collection_name == "goods_issue"{
		m:= master_data_model.GoodsIssue{}
		m.SetDefault()
		return &m
	} else if collection_name == "delivery_issue" {
		m:= master_data_model.DeliveryIssue{}
		m.SetDefault()
		return &m
	} else if collection_name == "shipment" {
		m:= shipment_model.Shipment{}
		m.SetDefault()
		return &m
	} else if collection_name == "shipment_item" {
		m:= shipment_model.ShipmentItem{}
		m.SetDefault()
		return &m
	} else if collection_name == "invoice" {
		m:= shipment_model.Invoice{}
		m.SetDefault()
		return &m
	} else if collection_name == "invoice_item" {
		m:= shipment_model.InvoiceItem{}
		m.SetDefault()
		return &m
	} else {
		return nil
	}
}

func GetModelAndSetDefaultValueExcel(collection_name string, create_by primitive.ObjectID, cells []*xlsx.Cell) (interface{}, error){
	if collection_name == "zone" {
		m:= master_data_model.Zone{}
		m.SetDefaultExcel(create_by, cells)
		return &m, nil
	} else if collection_name == "hub"{
		m:= master_data_model.Hub{}
		err := m.SetDefaultExcel(create_by, cells)
		return &m, err
	} else if collection_name == "fleet"{
		m:= master_data_model.Fleet{}
		err := m.SetDefaultExcel(create_by, cells)
		return &m, err
	} else if collection_name == "carrier"{
		m:= master_data_model.Carrier{}
		err := m.SetDefaultExcel(create_by, cells)
		return &m, err
	} else if collection_name == "vehicle"{
		m:= master_data_model.Vehicle{}
		err := m.SetDefaultExcel(create_by, cells)
		return &m, err
	} else if collection_name == "vehicle_type"{
		m:= master_data_model.VehicleType{}
		m.SetDefaultExcel(create_by, cells)
		return &m, nil
	} else if collection_name == "driver"{
		m:= master_data_model.Driver{}
		err := m.SetDefaultExcel(create_by, cells)
		return &m, err
	} else if collection_name == "helper"{
		m:= master_data_model.Helper{}
		err := m.SetDefaultExcel(create_by, cells)
		return &m, err
	} else if collection_name == "address"{
		m:= master_data_model.Address{}
		m.SetDefaultExcel(create_by, cells)
		return &m, nil
	} else if collection_name == "goods_issue"{
		m:= master_data_model.GoodsIssue{}
		m.SetDefaultExcel(create_by, cells)
		return &m, nil
	} else if collection_name == "delivery_issue" {
		m:= master_data_model.DeliveryIssue{}
		m.SetDefaultExcel(create_by, cells)
		return &m, nil
	} else if collection_name == "shipment" {
		m:= shipment_model.Shipment{}
		m.SetDefaultExcel(create_by, cells)
		return &m, nil
	} else if collection_name == "shipment_item" {
		m:= shipment_model.ShipmentItem{}
		m.SetDefaultExcel(create_by, cells)
		return &m, nil
	} else if collection_name == "invoice" {
		m:= shipment_model.Invoice{}
		m.SetDefaultExcel(create_by, cells)
		return &m, nil
	} else if collection_name == "invoice_item" {
		m:= shipment_model.InvoiceItem{}
		m.SetDefaultExcel(create_by, cells)
		return &m, nil
	} else {
		return nil, nil
	}
}
