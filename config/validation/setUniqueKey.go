package validation

import (
	"context"
	"os"
	"smartinno/config/db"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"reflect"
	"net/http"
	"fmt"
)

// set key to unique for check data in mongodb already exist.
func SetUniqueKey(col_name string, key string) {
	// connect mongodb
	collection, err := db.GetDBCollection(db.DB_NAME, col_name)
	// if connect mongo error return error
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	// Declare an index model object to pass to CreateOne() with options)
	mod := mongo.IndexModel{
	Keys: bson.M{
	key: 1, // index in descending order
	},
	// create UniqueIndex option
	Options: options.Index().SetUnique(true),
	}
	// Create an Index using the CreateOne() method
	ind, err := collection.Indexes().CreateOne(context.TODO(), mod)

	// Check if the CreateOne() method returned any errors
	if err != nil {
	fmt.Println("Indexes().CreateOne() ERROR:", err)
	os.Exit(1) // exit in case of error
	} else {
	// API call returns string of the index name
	fmt.Println("CreateOne() index:", ind)
	fmt.Println("CreateOne() type: ", reflect.TypeOf(ind))
	}
}

func PreConfigSetUniqueKey(w http.ResponseWriter, r *http.Request) {
	//create unique for zone collection with key = zone_name
	SetUniqueKey("zone", "zone_name")

	//create unique for hub collection with key = hub_name
	SetUniqueKey("hub", "hub_name")

	//create unique for carrier collection with key = carrier_name
	SetUniqueKey("carrier", "carrier_name")

	//create unique for delivery_issue collection with key = delivery_issue_name
	SetUniqueKey("delivery_issue", "delivery_issue_name")

	//create unique for driver collection with key = personal_id, driver_license_id, mobile, username
	SetUniqueKey("driver", "personal_id")
	SetUniqueKey("driver", "driver_license_id")
	SetUniqueKey("driver", "mobile")
	SetUniqueKey("driver", "username")

	//create unique for fleet collection with key = fleet_name
	SetUniqueKey("fleet", "fleet_name")

	//create unique for goods_issue collection with key = goods_issue_name
	SetUniqueKey("goods_issue", "goods_issue_name")

	//create unique for helper collection with key = personal_id, mobile
	SetUniqueKey("helper", "personal_id")
	SetUniqueKey("helper", "mobile")

	//create unique for address collection with key = address
	SetUniqueKey("address", "address_name")
	
	//create unique for vehicle collection with key = license
	SetUniqueKey("vehicle", "license")

	//create unique for vehicle_type collection with key = description
	SetUniqueKey("vehicle_type", "description")
	
}