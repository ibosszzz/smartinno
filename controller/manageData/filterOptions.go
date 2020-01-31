package manageData

import (
	"go.mongodb.org/mongo-driver/bson"
)

func FilterAndSearch(vars map[string]string) bson.M{
	findOptions := bson.M{}
	for key := range vars {
		if key == "search_value" && vars["search_value"] != "" {
			findOptions[vars["search_by"]] = bson.M{"$regex":vars["search_value"]}
		} else if key == "search_by" || key == "search_value" {
			continue
		} else if vars[key] != "" {
			findOptions[key] = vars[key]
		}
	}
	return findOptions
}