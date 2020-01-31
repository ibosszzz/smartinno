package main

import (
	"fmt"
	"log"
	"net/http"
	"smartinno/config/validation"
	"smartinno/controller"
	"smartinno/controller/manageData"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

const (
	STATIC_DIR = "/static/"
)

func main() {
	r := mux.NewRouter()

	// SetUniqueKey
	r.HandleFunc("/set_unique_key", validation.PreConfigSetUniqueKey).Methods("GET")
	r.HandleFunc("/test", controller.Test).Methods("GET")

	r.
		PathPrefix(STATIC_DIR).
		Handler(http.StripPrefix(STATIC_DIR, http.FileServer(http.Dir("."+STATIC_DIR))))

	//Zone
	zone := r.PathPrefix("/api/v1/zone").Subrouter()
	zone.HandleFunc("/add", manageData.Add).Methods("POST")
	zone.HandleFunc("/update/id/{id}", manageData.Update).Methods("POST")
	zone.HandleFunc("/delete/id/{id}", manageData.Delete).Methods("DELETE")
	zone.HandleFunc("", manageData.FindAll).Methods("GET")
	zone.HandleFunc("/id/{id}", manageData.FindByID).Methods("GET")
	zone.HandleFunc("/excel/upload/create_by/{create_by}", manageData.UploadExcel).Methods("POST")
	zone.HandleFunc("/excel/download", manageData.DownloadExcel).Methods("GET")
	zone.HandleFunc("/search_or_filter", manageData.Search).Methods("GET")

	//Hub
	hub := r.PathPrefix("/api/v1/hub").Subrouter()
	hub.HandleFunc("/add", manageData.Add).Methods("POST")
	hub.HandleFunc("/update/id/{id}", manageData.Update).Methods("POST")
	hub.HandleFunc("/delete/id/{id}", manageData.Delete).Methods("DELETE")
	hub.HandleFunc("", manageData.FindAll).Methods("GET")
	hub.HandleFunc("/id/{id}", manageData.FindByID).Methods("GET")
	hub.HandleFunc("/excel/upload/create_by/{create_by}", manageData.UploadExcel).Methods("POST")
	hub.HandleFunc("/excel/download", manageData.DownloadExcel).Methods("GET")
	hub.HandleFunc("/search_or_filter", manageData.Search).Methods("GET")

	//Fleet
	fleet := r.PathPrefix("/api/v1/fleet").Subrouter()
	fleet.HandleFunc("/add", manageData.Add).Methods("POST")
	fleet.HandleFunc("/update/id/{id}", manageData.Update).Methods("POST")
	fleet.HandleFunc("/delete/id/{id}", manageData.Delete).Methods("DELETE")
	fleet.HandleFunc("", manageData.FindAll).Methods("GET")
	fleet.HandleFunc("/id/{id}", manageData.FindByID).Methods("GET")
	fleet.HandleFunc("/excel/upload/create_by/{create_by}", manageData.UploadExcel).Methods("POST")
	fleet.HandleFunc("/excel/download", manageData.DownloadExcel).Methods("GET")
	fleet.HandleFunc("/search_or_filter", manageData.Search).Methods("GET")

	//Carrier
	carrier := r.PathPrefix("/api/v1/carrier").Subrouter()
	carrier.HandleFunc("/add", manageData.Add).Methods("POST")
	carrier.HandleFunc("/update/id/{id}", manageData.Update).Methods("POST")
	carrier.HandleFunc("/delete/id/{id}", manageData.Delete).Methods("DELETE")
	carrier.HandleFunc("", manageData.FindAll).Methods("GET")
	carrier.HandleFunc("/id/{id}", manageData.FindByID).Methods("GET")
	carrier.HandleFunc("/excel/upload/create_by/{create_by}", manageData.UploadExcel).Methods("POST")
	carrier.HandleFunc("/excel/download", manageData.DownloadExcel).Methods("GET")
	carrier.HandleFunc("/search_or_filter", manageData.Search).Methods("GET")

	//Vehicle
	vehicle := r.PathPrefix("/api/v1/vehicle").Subrouter()
	vehicle.HandleFunc("/add", manageData.Add).Methods("POST")
	vehicle.HandleFunc("/update/id/{id}", manageData.Update).Methods("POST")
	vehicle.HandleFunc("/delete/id/{id}", manageData.Delete).Methods("DELETE")
	vehicle.HandleFunc("", manageData.FindAll).Methods("GET")
	vehicle.HandleFunc("/id/{id}", manageData.FindByID).Methods("GET")
	vehicle.HandleFunc("/excel/upload/create_by/{create_by}", manageData.UploadExcel).Methods("POST")
	vehicle.HandleFunc("/excel/download", manageData.DownloadExcel).Methods("GET")
	vehicle.HandleFunc("/search_or_filter", manageData.Search).Methods("GET")

	//VehicleType
	vehicle_type := r.PathPrefix("/api/v1/vehicle_type").Subrouter()
	vehicle_type.HandleFunc("/add", manageData.Add).Methods("POST")
	vehicle_type.HandleFunc("/update/id/{id}", manageData.Update).Methods("POST")
	vehicle_type.HandleFunc("/delete/id/{id}", manageData.Delete).Methods("DELETE")
	vehicle_type.HandleFunc("", manageData.FindAll).Methods("GET")
	vehicle_type.HandleFunc("/id/{id}", manageData.FindByID).Methods("GET")
	vehicle_type.HandleFunc("/excel/upload/create_by/{create_by}", manageData.UploadExcel).Methods("POST")
	vehicle_type.HandleFunc("/excel/download", manageData.DownloadExcel).Methods("GET")
	vehicle_type.HandleFunc("/search_or_filter", manageData.Search).Methods("GET")

	//Driver
	driver := r.PathPrefix("/api/v1/driver").Subrouter()
	driver.HandleFunc("/add", manageData.Add).Methods("POST")
	driver.HandleFunc("/update/id/{id}", manageData.Update).Methods("POST")
	driver.HandleFunc("/delete/id/{id}", manageData.Delete).Methods("DELETE")
	driver.HandleFunc("", manageData.FindAll).Methods("GET")
	driver.HandleFunc("/id/{id}", manageData.FindByID).Methods("GET")
	driver.HandleFunc("/excel/upload/create_by/{create_by}", manageData.UploadExcel).Methods("POST")
	driver.HandleFunc("/excel/download", manageData.DownloadExcel).Methods("GET")
	driver.HandleFunc("/search_or_filter", manageData.Search).Methods("GET")

	//Helper
	helper := r.PathPrefix("/api/v1/helper").Subrouter()
	helper.HandleFunc("/add", manageData.Add).Methods("POST")
	helper.HandleFunc("/update/id/{id}", manageData.Update).Methods("POST")
	helper.HandleFunc("/delete/id/{id}", manageData.Delete).Methods("DELETE")
	helper.HandleFunc("", manageData.FindAll).Methods("GET")
	helper.HandleFunc("/id/{id}", manageData.FindByID).Methods("GET")
	helper.HandleFunc("/excel/upload/create_by/{create_by}", manageData.UploadExcel).Methods("POST")
	helper.HandleFunc("/excel/download", manageData.DownloadExcel).Methods("GET")
	helper.HandleFunc("/search_or_filter", manageData.Search).Methods("GET")

	//Matching Driver and Vehicle
	matching := r.PathPrefix("/api/v1/matching").Subrouter()
	matching.HandleFunc("/add/vehicle_id/{vehicle_id}", manageData.MatchDriver).Methods("POST")
	matching.HandleFunc("/remove/vehicle_id/{vehicle_id}", manageData.RemoveDriver).Methods("POST")

	//Address
	address := r.PathPrefix("/api/v1/address").Subrouter()
	address.HandleFunc("/add", manageData.Add).Methods("POST")
	address.HandleFunc("/update/id/{id}", manageData.Update).Methods("POST")
	address.HandleFunc("/delete/id/{id}", manageData.Delete).Methods("DELETE")
	address.HandleFunc("", manageData.FindAll).Methods("GET")
	address.HandleFunc("/id/{id}", manageData.FindByID).Methods("GET")
	address.HandleFunc("/excel/upload/create_by/{create_by}", manageData.UploadExcel).Methods("POST")
	address.HandleFunc("/excel/download", manageData.DownloadExcel).Methods("GET")
	address.HandleFunc("/search_or_filter", manageData.Search).Methods("GET")

	//GoodsIssue
	goods_issue := r.PathPrefix("/api/v1/goods_issue").Subrouter()
	goods_issue.HandleFunc("/add", manageData.Add).Methods("POST")
	goods_issue.HandleFunc("/update/id/{id}", manageData.Update).Methods("POST")
	goods_issue.HandleFunc("/delete/id/{id}", manageData.Delete).Methods("DELETE")
	goods_issue.HandleFunc("", manageData.FindAll).Methods("GET")
	goods_issue.HandleFunc("/id/{id}", manageData.FindByID).Methods("GET")
	goods_issue.HandleFunc("/excel/upload/create_by/{create_by}", manageData.UploadExcel).Methods("POST")
	goods_issue.HandleFunc("/excel/download", manageData.DownloadExcel).Methods("GET")
	goods_issue.HandleFunc("/search_or_filter", manageData.Search).Methods("GET")

	//DeliveryIssue
	delivery_issue := r.PathPrefix("/api/v1/delivery_issue").Subrouter()
	delivery_issue.HandleFunc("/add", manageData.Add).Methods("POST")
	delivery_issue.HandleFunc("/update/id/{id}", manageData.Update).Methods("POST")
	delivery_issue.HandleFunc("/delete/id/{id}", manageData.Delete).Methods("DELETE")
	delivery_issue.HandleFunc("", manageData.FindAll).Methods("GET")
	delivery_issue.HandleFunc("/id/{id}", manageData.FindByID).Methods("GET")
	delivery_issue.HandleFunc("/excel/upload/create_by/{create_by}", manageData.UploadExcel).Methods("POST")
	delivery_issue.HandleFunc("/excel/download", manageData.DownloadExcel).Methods("GET")
	delivery_issue.HandleFunc("/search_or_filter", manageData.Search).Methods("GET")

	//Shipment
	shipment := r.PathPrefix("/api/v1/shipment").Subrouter()
	shipment.HandleFunc("/add", manageData.Add).Methods("POST")
	shipment.HandleFunc("/update/id/{id}", manageData.Update).Methods("POST")
	shipment.HandleFunc("/delete/id/{id}", manageData.Delete).Methods("DELETE")
	shipment.HandleFunc("", manageData.FindAll).Methods("GET")
	shipment.HandleFunc("/id/{id}", manageData.FindByID).Methods("GET")
	shipment.HandleFunc("/excel/upload/create_by/{create_by}", manageData.UploadExcel).Methods("POST")
	shipment.HandleFunc("/excel/download", manageData.DownloadExcel).Methods("GET")
	shipment.HandleFunc("/search_or_filter", manageData.Search).Methods("GET")

	//ShipmentItem
	shipment_item := r.PathPrefix("/api/v1/shipment_item").Subrouter()
	shipment_item.HandleFunc("/add", manageData.Add).Methods("POST")
	shipment_item.HandleFunc("/update/id/{id}", manageData.Update).Methods("POST")
	shipment_item.HandleFunc("/delete/id/{id}", manageData.Delete).Methods("DELETE")
	shipment_item.HandleFunc("", manageData.FindAll).Methods("GET")
	shipment_item.HandleFunc("/id/{id}", manageData.FindByID).Methods("GET")
	shipment_item.HandleFunc("/excel/upload/create_by/{create_by}", manageData.UploadExcel).Methods("POST")
	shipment_item.HandleFunc("/excel/download", manageData.DownloadExcel).Methods("GET")
	shipment_item.HandleFunc("/search_or_filter", manageData.Search).Methods("GET")

	//Invoice
	invoice := r.PathPrefix("/api/v1/invoice").Subrouter()
	invoice.HandleFunc("/add", manageData.Add).Methods("POST")
	invoice.HandleFunc("/update/id/{id}", manageData.Update).Methods("POST")
	invoice.HandleFunc("/delete/id/{id}", manageData.Delete).Methods("DELETE")
	invoice.HandleFunc("", manageData.FindAll).Methods("GET")
	invoice.HandleFunc("/id/{id}", manageData.FindByID).Methods("GET")
	invoice.HandleFunc("/excel/upload/create_by/{create_by}", manageData.UploadExcel).Methods("POST")
	invoice.HandleFunc("/excel/download", manageData.DownloadExcel).Methods("GET")
	invoice.HandleFunc("/search_or_filter", manageData.Search).Methods("GET")

	//InvoiceItem
	invoice_item := r.PathPrefix("/api/v1/invoice_item").Subrouter()
	invoice_item.HandleFunc("/add", manageData.Add).Methods("POST")
	invoice_item.HandleFunc("/update/id/{id}", manageData.Update).Methods("POST")
	invoice_item.HandleFunc("/delete/id/{id}", manageData.Delete).Methods("DELETE")
	invoice_item.HandleFunc("", manageData.FindAll).Methods("GET")
	invoice_item.HandleFunc("/id/{id}", manageData.FindByID).Methods("GET")
	invoice_item.HandleFunc("/excel/upload/create_by/{create_by}", manageData.UploadExcel).Methods("POST")
	invoice_item.HandleFunc("/excel/download", manageData.DownloadExcel).Methods("GET")
	invoice_item.HandleFunc("/search_or_filter", manageData.Search).Methods("GET")

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowCredentials: true,
		AllowedMethods:   []string{"GET", "POST", "PATCH", "PUT", "DELETE"},
		Debug:            false,
	})

	handler := c.Handler(r)

	fmt.Println("Server Starting...")
	log.Fatal(http.ListenAndServe(":8866", handler))
	fmt.Printf("Server Started")
}
