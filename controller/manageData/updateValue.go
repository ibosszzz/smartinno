package manageData

import (
	"time"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func UpdateValue(collection_name string, myMap map[string]interface{}) *bson.M{
	t := time.Now()
	modify_by, _:= primitive.ObjectIDFromHex(myMap["modify_by"].(string))
	if collection_name == "zone" {
		update_value := bson.M{
			"zone_code": myMap["zone_code"],
			"zone_name": myMap["zone_name"],
			"active": myMap["active"],
			"modify_date": t.Format("2006-01-02 15:04:05"),
			"modify_by": modify_by,
		}
		return &update_value
	} else if collection_name == "hub"{
		zone_id, _:= primitive.ObjectIDFromHex(myMap["zone_id"].(string))
		update_value := bson.M{
			"hub_code": myMap["hub_code"],
			"hub_name": myMap["hub_name"],
			"active": myMap["active"],
			"address1": myMap["address1"],
			"address2": myMap["address2"],
			"zipcode": myMap["zipcode"],
			"zone_id": zone_id,
			"modify_date": t.Format("2006-01-02 15:04:05"),
			"modify_by": modify_by,
		}
		return &update_value
	} else if collection_name == "fleet"{
		hub_id, _:= primitive.ObjectIDFromHex(myMap["hub_id"].(string))
		update_value := bson.M{
			"fleet_code": myMap["fleet_code"],
			"fleet_name": myMap["fleet_name"],
			"active": myMap["active"],
			"hub_id": hub_id,
			"modify_date": t.Format("2006-01-02 15:04:05"),
			"modify_by": modify_by,
		}
		return &update_value
	} else if collection_name == "carrier"{
		fleet_id, _:= primitive.ObjectIDFromHex(myMap["fleet_id"].(string))
		update_value := bson.M{
			"carrier_code": myMap["carrier_code"],
			"carrier_name": myMap["carrier_name"],
			"contact_name": myMap["contact_name"],
			"email": myMap["email"],
			"mobile": myMap["mobile"],
			"active": myMap["active"],
			"fleet_id": fleet_id,
			"modify_date": t.Format("2006-01-02 15:04:05"),
			"modify_by": modify_by,
		}
		return &update_value
	} else if collection_name == "vehicle"{
		carrier_id, _:= primitive.ObjectIDFromHex(myMap["carrier_id"].(string))
		vehicle_type_id, _:= primitive.ObjectIDFromHex(myMap["vehicle_type_id"].(string))
		update_value := bson.M{
			"license": myMap["license"],
			"license_expired": myMap["license_expired"],
			"mileage": myMap["mileage"],
			"remark": myMap["remark"],
			"active": myMap["active"],
			"carrier_id": carrier_id,
			"vehicle_type_id": vehicle_type_id,
			"modify_date": t.Format("2006-01-02 15:04:05"),
			"modify_by": modify_by,
		}
		return &update_value
	} else if collection_name == "vehicle_type"{
		update_value := bson.M{
			"vehicle_type_code": myMap["vehicle_type_code"],
			"description": myMap["description"],
			"wheels": myMap["wheels"],
			"width": myMap["width"],
			"length": myMap["length"],
			"height": myMap["height"],
			"payload": myMap["payload"],
			"type": myMap["type"],
			"active": myMap["active"],
			"modify_date": t.Format("2006-01-02 15:04:05"),
			"modify_by": modify_by,
		}
		return &update_value
	} else if collection_name == "driver"{
		carrier_id, _:= primitive.ObjectIDFromHex(myMap["carrier_id"].(string))
		update_value := bson.M{
			"driver_name": myMap["driver_name"],
			"personal_id": myMap["personal_id"],
			"mobile": myMap["mobile"],
			"date_of_birth": myMap["date_of_birth"],
			"driver_license_id": myMap["driver_license_id"],
			"driver_license_expired": myMap["driver_license_expired"],
			"username": myMap["username"],
			"password": myMap["password"],
			"active": myMap["active"],
			"carrier_id": carrier_id,
			"modify_date": t.Format("2006-01-02 15:04:05"),
			"modify_by": modify_by,
		}
		return &update_value
	} else if collection_name == "helper"{
		carrier_id, _:= primitive.ObjectIDFromHex(myMap["carrier_id"].(string))
		update_value := bson.M{
			"helper_name": myMap["helper_name"],
			"helper_nickname": myMap["helper_nickname"],
			"personal_id": myMap["personal_id"],
			"mobile": myMap["mobile"],
			"date_of_birth": myMap["date_of_birth"],
			"active": myMap["active"],
			"carrier_id": carrier_id,
			"modify_date": t.Format("2006-01-02 15:04:05"),
			"modify_by": modify_by,
		}
		return &update_value
	} else if collection_name == "address"{
		update_value := bson.M{
			"address_code": myMap["address_code"],
			"address_name": myMap["address_name"],
			"pickup_or_delivery_point": myMap["pickup_or_delivery_point"],
			"building": myMap["building"],
			"address1": myMap["address1"],
			"address2": myMap["address2"],
			"zipcode": myMap["zipcode"],
			"city": myMap["city"],
			"country": myMap["country"],
			"contact_mobile": myMap["contact_mobile"],
			"contact_name": myMap["contact_name"],
			"latitude": myMap["latitude"],
			"longitude": myMap["longitude"],
			"radius": myMap["radius"],
			"suggested_latitude": myMap["suggested_latitude"],
			"suggested_longitude": myMap["suggested_longitude"],
			"remark": myMap["remark"],
			"active": myMap["active"],
			"modify_date": t.Format("2006-01-02 15:04:05"),
			"modify_by": modify_by,
		}
		return &update_value
	} else if collection_name == "goods_issue"{
		update_value := bson.M{
			"goods_issue_name": myMap["goods_issue_name"],
			"active": myMap["active"],
			"modify_date": t.Format("2006-01-02 15:04:05"),
			"modify_by": modify_by,
		}
		return &update_value
	} else {  // delivery_issue
		update_value := bson.M{
			"delivery_issue_name": myMap["delivery_issue_name"],
			"active": myMap["active"],
			"modify_date": t.Format("2006-01-02 15:04:05"),
			"modify_by": modify_by,
		}
		return &update_value
	}

}