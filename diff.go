package bsonrecord

import (
	"reflect"

	"github.com/globalsign/mgo/bson"
)

func diff(old bson.M, changed bson.M) bson.M {
	result := bson.M{}

	for key, value := range changed {
		if nestedValue, isNested := value.(bson.M); isNested {
			nestedDiff := diff(old[key].(bson.M), nestedValue)
			merge(withPrefix(nestedDiff, key+"."), result)
			continue
		}

		if !reflect.DeepEqual(value, old[key]) {
			result[key] = value
		}
	}

	return result
}

func merge(from, to bson.M) {
	for k, v := range from {
		to[k] = v
	}
}

func withPrefix(data bson.M, prefix string) bson.M {
	result := bson.M{}

	for k, v := range data {
		result[prefix+k] = v
	}

	return result
}
