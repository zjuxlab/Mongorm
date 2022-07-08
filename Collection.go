package Mongorm

import (
	"fmt"
	"reflect"

	"go.mongodb.org/mongo-driver/mongo"
)

var col_map map[reflect.Type](*mongo.Collection) = make(map[reflect.Type]*mongo.Collection)

func SetCollection[T any](col *mongo.Collection) {
	var v T
	col_map[reflect.TypeOf(v)] = col
}

func SetCollectionByName[T any](col_name string) {
	var v T
	if DB == nil {
		logError[T]("SetCollectionByName", fmt.Errorf("MongoDB not set"))
	}
	col_map[reflect.TypeOf(v)] = DB.Collection(col_name)
}

// Firstly, find Collection in col_map which is set by human;
// Then, try to open a Collection by TypeName in MongoDB.
func GetCollection[T any]() *mongo.Collection {
	var v T
	tp := reflect.TypeOf(v)
	col, ok := col_map[tp]
	if !ok {
		if DB != nil {
			return DB.Collection(tp.Name())
		}
		logError[T]("GetCollection", fmt.Errorf("Cannot Find Collection %s", tp.Name()))
		return nil
	}
	return col
}
