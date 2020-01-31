package manageData

import (
	"context"
	"encoding/json"
	"smartinno/config/db"
	master_data_model "smartinno/model/masterData"
	"io/ioutil"
	"net/http"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	res_format "smartinno/model/responseData"
	"github.com/gorilla/mux"
	"time"
)

func MatchDriver(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	// get data from param
	vars := mux.Vars(r)
	id := vars["vehicle_id"]
	objID, _ := primitive.ObjectIDFromHex(id)
	
	var vehicle master_data_model.Vehicle
	// read json body
	body, _ := ioutil.ReadAll(r.Body)
	// map json body to vehicle variable
	err := json.Unmarshal(body, &vehicle)

	var res res_format.ResponseResult
	// if map error return error
	if err != nil {
		res.Error = err.Error()
		res.Status = http.StatusInternalServerError
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(res)
		return
	}
	// connect mongodb to collection vehicle
	collection, err := db.GetDBCollection(db.DB_NAME, "vehicle")
	// if connect mongo error return error
	if err != nil {
		res.Error = err.Error()
		res.Status = http.StatusInternalServerError
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(res)
		return
	}
	// filter vehicle to update
	filter := bson.D{primitive.E{Key: "_id", Value: objID}}
	// update value
	t := time.Now()
	// var update 
	update := bson.D{primitive.E{Key: "$inc", Value: bson.D{primitive.E{Key:"version", Value: 1},
	}},
		{Key: "$set", Value: bson.M{
			"driver_id": vehicle.DriverID,
			"modify_date": t.Format("2006-01-02 15:04:05"),
			"modify_by": vehicle.ModifyBy,
		}},
	}
	
	// update data in mongodb
	_, err = collection.UpdateOne(context.TODO(), filter, update)

	// if error while update return error
	if err != nil {
		res.Error = err.Error()
		res.Status = http.StatusInternalServerError
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(res)
		return
	}

	// connect mongodb to collection driver
	collection, err = db.GetDBCollection(db.DB_NAME, "driver")
	// if connect mongo error return error
	if err != nil {
		res.Error = err.Error()
		res.Status = http.StatusInternalServerError
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(res)
		return
	}
	// filter driver to update
	filter = bson.D{primitive.E{Key: "_id", Value: vehicle.DriverID}}
	// var update 
	update = bson.D{primitive.E{Key: "$set", Value: bson.M{
		"ready_to_match": "false",
		"modify_date": t.Format("2006-01-02 15:04:05"),
		"modify_by": vehicle.ModifyBy,
	}}}
	
	// update data in mongodb
	_, err = collection.UpdateOne(context.TODO(), filter, update)

	// if error while update return error
	if err != nil {
		res.Error = err.Error()
		res.Status = http.StatusInternalServerError
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(res)
		return
	}

	// update success
	res.Data = "Update Driver Successful"
	res.Status = http.StatusOK
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)
	return

}

func RemoveDriver(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	// get data from param
	vars := mux.Vars(r)
	id := vars["vehicle_id"]
	objID, _ := primitive.ObjectIDFromHex(id)
	
	var vehicle master_data_model.Vehicle
	// read json body
	body, _ := ioutil.ReadAll(r.Body)
	// map json body to vehicle variable
	err := json.Unmarshal(body, &vehicle)

	var res res_format.ResponseResult
	// if map error return error
	if err != nil {
		res.Error = err.Error()
		res.Status = http.StatusInternalServerError
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(res)
		return
	}
	// connect mongodb
	collection, err := db.GetDBCollection(db.DB_NAME, "vehicle")
	// if connect mongo error return error
	if err != nil {
		res.Error = err.Error()
		res.Status = http.StatusInternalServerError
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(res)
		return
	}
	var result master_data_model.Vehicle
	// find vehicle_name in mongodb (validate by vehicle_name)
	err = collection.FindOne(context.TODO(), bson.D{primitive.E{Key: "_id", Value: objID}}).Decode(&result)
	if err != nil {
		res.Error = err.Error()
		res.Status = http.StatusNotFound
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(res)
		return
	}
	driverID := result.DriverID
	// filter vehicle to update
	filter := bson.D{primitive.E{Key: "_id", Value: objID}}
	// update value
	t := time.Now()
	// var update 
	update := bson.D{primitive.E{Key: "$inc", Value: bson.D{primitive.E{Key:"version", Value: 1},
	}},
		{Key: "$set", Value: bson.M{
			"driver_id": "000000000000000000000000",
			"modify_date": t.Format("2006-01-02 15:04:05"),
			"modify_by": vehicle.ModifyBy,
		}},
	}
	
	// update data in mongodb
	_, err = collection.UpdateOne(context.TODO(), filter, update)

	// if error while update return error
	if err != nil {
		res.Error = err.Error()
		res.Status = http.StatusInternalServerError
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(res)
		return
	}

	// connect mongodb to collection driver
	collection, err = db.GetDBCollection(db.DB_NAME, "driver")
	// if connect mongo error return error
	if err != nil {
		res.Error = err.Error()
		res.Status = http.StatusInternalServerError
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(res)
		return
	}
	// filter driver to update
	filter = bson.D{primitive.E{Key: "_id", Value: driverID}}
	// var update 
	update = bson.D{primitive.E{Key: "$set", Value: bson.M{
		"ready_to_match": "true",
		"modify_date": t.Format("2006-01-02 15:04:05"),
		"modify_by": vehicle.ModifyBy,
	}}}
	
	// update data in mongodb
	_, err = collection.UpdateOne(context.TODO(), filter, update)

	// if error while update return error
	if err != nil {
		res.Error = err.Error()
		res.Status = http.StatusInternalServerError
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(res)
		return
	}

	// update success
	res.Data = "Remove Driver Successful"
	res.Status = http.StatusOK
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)
	return

}

