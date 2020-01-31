package manageData

import (
	"context"
	"encoding/json"
	"smartinno/config/db"
	res_format "smartinno/model/responseData"
	"io/ioutil"
	"net/http"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"github.com/gorilla/mux"
	"strings"
	"fmt"
)

func Add(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	//get model to data_model variable
	prefix := strings.Split(r.URL.Path, "/")[3]
	var data_model = GetModelAndSetDefaultValue(prefix)
	// read json body
	body, _ := ioutil.ReadAll(r.Body)
	// map json body to zone variable
	err := json.Unmarshal(body, &data_model)

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
	collection, err := db.GetDBCollection(db.DB_NAME, prefix)
	// if connect mongo error return error
	if err != nil {
		res.Error = err.Error()
		res.Status = http.StatusInternalServerError
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(res)
		return
	}
	// insert data to mongodb
	_, err = collection.InsertOne(context.TODO(), data_model)
	if err != nil {
		res.Error = err.Error()
		res.Status = http.StatusConflict
		w.WriteHeader(http.StatusConflict)
		json.NewEncoder(w).Encode(res)
		return
	}
	// create success
	res.Data = "Create "+prefix+" Successful"
	res.Status = http.StatusCreated
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(res)
	return

}


func Update(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	// get data from param
	vars := mux.Vars(r)
	id := vars["id"]
	objID, _ := primitive.ObjectIDFromHex(id)
	// fmt.Println(objID)

	prefix := strings.Split(r.URL.Path, "/")[3]
	var data_model interface{} //GetModel(prefix)

	// read json body
	body, _ := ioutil.ReadAll(r.Body)
	// map json body to zone variable

	err := json.Unmarshal(body, &data_model)
	if err != nil{
		fmt.Println(err.Error())
	}

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
	collection, err := db.GetDBCollection(db.DB_NAME, prefix)
	// if connect mongo error return error
	if err != nil {
		res.Error = err.Error()
		res.Status = http.StatusInternalServerError
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(res)
		return
	}

	myMap := data_model.(map[string]interface{})
	// filter zone_name to update
	filter := bson.D{primitive.E{Key: "_id", Value: objID}}
	update := bson.D{primitive.E{Key: "$inc", Value: bson.D{primitive.E{Key:"version", Value: 1},
		}},
		{Key: "$set", Value: UpdateValue(prefix, myMap)},
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
	// update success
	res.Data = "Update "+prefix+" Successful"
	res.Status = http.StatusOK
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)
	return
}

func Delete(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	// get data from param
	vars := mux.Vars(r)
	id := vars["id"]
	objID, _ := primitive.ObjectIDFromHex(id)

	prefix := strings.Split(r.URL.Path, "/")[3]
	// connect mongodb
	collection, err := db.GetDBCollection(db.DB_NAME, prefix)
	// if connect mongo error return error
	var res res_format.ResponseResult
	if err != nil {
		res.Error = err.Error()
		res.Status = http.StatusInternalServerError
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(res)
		return
	}
	filter := bson.D{primitive.E{Key: "_id", Value: objID}}
	update := bson.D{primitive.E{Key: "$inc", Value: bson.D{primitive.E{Key:"version", Value: 1},
		}},
		{Key: "$set", Value: bson.M{"active":"false"}},
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
	// delete success
	res.Data = "Delete "+prefix+" Successful"
	res.Status = http.StatusOK
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)
	return
	
}

func FindAll(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	prefix := strings.Split(r.URL.Path, "/")[3]
	// connect mongodb
	collection, err := db.GetDBCollection(db.DB_NAME, prefix)
	// if connect mongo error return error
	var res res_format.ResponseResult
	if err != nil {
		res.Error = err.Error()
		res.Status = http.StatusInternalServerError
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(res)
		return
	}

	var results []*map[string]interface{}
	cur, err := collection.Find(context.TODO(), bson.D{})
	if err != nil {
		res.Error = err.Error()
		res.Status = http.StatusNotFound
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(res)
		return
	}

	// append data map struct to results variable
	for cur.Next(context.TODO()) {
		// create a value into which the single document can be decoded
		var data_model map[string]interface{}
		err := cur.Decode(&data_model)
		if err != nil {
			res.Error = err.Error()
			res.Status = http.StatusInternalServerError
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(res)
			return
		}
		results = append(results, &data_model)
	}

	// check non-result
	if len(results) == 0{
		res.Data = ""
		res.Status = http.StatusOK
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(res)
		return
	}

	// Close the cursor once finished
	cur.Close(context.TODO())

	// return results
	res.Data = results
	res.Status = http.StatusOK
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)
	return
}

func Search(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")

	// findOption := FilterAndSearch(mux.Vars(r))

	// get url query
	query := r.URL.Query()
	// convert type map[string][]string to map[string]string
	new_query := make(map[string]string)
	for key, value := range query {
        new_query[key] = value[0]
	}
	findOption := FilterAndSearch(new_query)

	// v := r.URL.Query()
    // username := v.Get("username")
    // email := v.Get("email")

	prefix := strings.Split(r.URL.Path, "/")[3]
	
	// connect mongodb
	collection, err := db.GetDBCollection(db.DB_NAME, prefix)
	// if connect mongo error return error
	var res res_format.ResponseResult
	if err != nil {
		res.Error = err.Error()
		res.Status = http.StatusInternalServerError
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(res)
		return
	}

	var results []*map[string]interface{}
	cur, err := collection.Find(context.TODO(), findOption)
	if err != nil {
		res.Error = err.Error()
		res.Status = http.StatusNotFound
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(res)
		return
	}

	// append data map struct to results variable
	for cur.Next(context.TODO()) {
		// create a value into which the single document can be decoded
		var data_model map[string]interface{}
		err := cur.Decode(&data_model)
		if err != nil {
			res.Error = err.Error()
			res.Status = http.StatusInternalServerError
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(res)
			return
		}
	
		results = append(results, &data_model)
	}

	// check non-result
	if len(results) == 0{
		res.Data = ""
		res.Status = http.StatusOK
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(res)
		return
	}

	// Close the cursor once finished
	cur.Close(context.TODO())

	// return results
	res.Data = results
	res.Status = http.StatusOK
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)
	return
}

func FindByID(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")

	prefix := strings.Split(r.URL.Path, "/")[3]
	// get data from param
	vars := mux.Vars(r)
	id := vars["id"]
	objID, _ := primitive.ObjectIDFromHex(id)
	// connect mongodb
	collection, err := db.GetDBCollection(db.DB_NAME, prefix)
	// if connect mongo error return error
	var res res_format.ResponseResult
	if err != nil {
		res.Error = err.Error()
		res.Status = http.StatusInternalServerError
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(res)
		return
	}
	var result map[string]interface{}
	// find by oid in mongodb
	err = collection.FindOne(context.TODO(), bson.D{primitive.E{Key: "_id", Value: objID}}).Decode(&result)
	if err != nil {
		res.Error = err.Error()
		res.Status = http.StatusNotFound
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(res)
		return
	}

	// return result find from id
	res.Data = result
	res.Status = http.StatusOK
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)
}