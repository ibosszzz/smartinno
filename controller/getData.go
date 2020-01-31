package controller

import (
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/bson"
	"smartinno/config/db"
	"net/http"
	"context"
	"fmt"
	"errors"
)

func Test(w http.ResponseWriter, r *http.Request) {
	oid, err := GetOID("zone", "zone_name", "name04")
	fmt.Println(oid)
	fmt.Println(err)
	oid, err = GetOID("zone", "zone_name", "name06")
	fmt.Println(oid)
	fmt.Println(err)
}

func GetOID(col_name string, key string, value string) (primitive.ObjectID, error){
	collection, err := db.GetDBCollection(db.DB_NAME, col_name)
	if err != nil{
		fmt.Println(err.Error())
	}
	var result map[string]primitive.ObjectID
	opts := options.FindOne().SetProjection(bson.M{"_id": 1})
	err = collection.FindOne(context.TODO(), bson.M{key: value}, opts).Decode(&result)
	if err != nil{
		return primitive.NilObjectID, errors.New("Not Found "+key+" = "+value+" in database")
	}
	return result["_id"], nil
}