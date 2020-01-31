package manageData

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"smartinno/config/db"
	res_format "smartinno/model/responseData"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
	"github.com/tealeg/xlsx"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func DownloadExcel(w http.ResponseWriter, r *http.Request) {
	prefix := strings.Split(r.URL.Path, "/")[3]
	w.Header().Set("Content-Disposition", "attachment; filename="+prefix+".xlsx")
	w.Header().Set("Content-Type", r.Header.Get("Content-Type"))
	w.Header().Set("Content-Length", r.Header.Get("Content-Length"))
	fmt.Println("prefix: ", prefix)
	var x = "smartinno/templates/"
	var filePath = x + prefix + ".xlsx"
	fmt.Println("filePath: ", filePath)
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	_, err = io.Copy(w, file)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	return
}

func UploadExcel(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var res res_format.ResponseResultExcel
	vars := mux.Vars(r)
	create_by := vars["create_by"]
	//get prefix path from url. ex. zone, hub, fleet
	prefix := strings.Split(r.URL.Path, "/")[3]
	r.ParseMultipartForm(32 << 20)
	file, _, err := r.FormFile("file")
	if err != nil {
		res.Error = []error{err}
		res.Status = http.StatusInternalServerError
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(res)
		return
	}
	defer file.Close()

	f, err := os.OpenFile("upload_data.xlsx", os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		res.Error = []error{err}
		res.Status = http.StatusInternalServerError
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(res)
		return
	}
	defer f.Close()
	_, err = io.Copy(f, file)
	if err != nil {
		res.Error = []error{err}
		res.Status = http.StatusInternalServerError
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(res)
		return
	}
	//var errList []error
	// var a []error
	allData, errList := ReadExcel("upload_data.xlsx", prefix, create_by)
	if errList != nil {
		var res res_format.ResponseResultExcelWithError
		os.Remove("upload_data.xlsx")
		strErrors := make([]string, len(errList))

		for i, err := range errList {
			strErrors[i] = err.Error()
		}
		res.Error = strErrors
		res.Status = http.StatusInternalServerError
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(res)
		return
	}

	// connect mongodb
	collection, err := db.GetDBCollection(db.DB_NAME, prefix)
	if err != nil {
		res.Error = []error{err}
		res.Status = http.StatusInternalServerError
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(res)
		return
	}
	// count before add
	before_count, _ := collection.CountDocuments(context.TODO(), bson.M{})
	opts := options.InsertMany().SetOrdered(false)
	_, err = collection.InsertMany(context.TODO(), allData, opts)
	if err != nil {
		res.Error = []error{err}
	}
	os.Remove("upload_data.xlsx")
	after_count, _ := collection.CountDocuments(context.TODO(), bson.M{})
	if after_count-before_count == 0 {
		res.Data = "0 row Inserted"
		res.Status = http.StatusOK
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(res)
		return
	}

	res.Data = strconv.FormatInt((after_count-before_count), 10) + " rows Inserted"
	res.Status = http.StatusCreated
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(res)
	return

}

func ReadExcel(excelFileName string, prefix string, create_by string) ([]interface{}, []error) {
	create_by_oid, _ := primitive.ObjectIDFromHex(create_by)
	xlFile, err := xlsx.OpenFile(excelFileName)
	if err != nil {
		return nil, []error{err}
	}
	var allData []interface{}
	var errList []error
	for _, sheet := range xlFile.Sheets {
		for i := 1; i < len(sheet.Rows); i++ {
			row := sheet.Rows[i]
			data, err := GetModelAndSetDefaultValueExcel(prefix, create_by_oid, row.Cells)
			if err != nil {
				errList = append(errList, err)
			}
			allData = append(allData, &data)
		}
	}
	// fmt.Println(errList)
	if len(errList) != 0 {
		return nil, errList
	}
	return allData, nil

}
