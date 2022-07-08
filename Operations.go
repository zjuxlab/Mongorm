package Mongorm

import (
	. "go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func InsertOne[T any](val T, opts ...*options.InsertOneOptions) bool {
	col := GetCollection[T]()
	if col == nil {
		return false
	}
	_, err := col.InsertOne(ctx, val, opts...)
	if err != nil {
		logError[T]("Add", err)
	}
	return err == nil
}

func FindOne[T any](filter D, opts ...*options.FindOneOptions) (val T, ok bool) {
	col := GetCollection[T]()
	if col == nil {
		return val, false
	}
	res := col.FindOne(ctx, filter, opts...)
	if res.Err() != nil {
		logError[T]("FindOne", res.Err())
		return val, false
	}

	if err := res.Decode(&val); err != nil {
		logError[T]("FindOne", err)
		return val, false
	}

	return val, true
}

func FindMany[T any](filter D, opts ...*options.FindOptions) (vals []T, ok bool) {
	col := GetCollection[T]()
	if col == nil {
		return nil, false
	}
	res, err := col.Find(ctx, filter, opts...)
	if err != nil {
		logError[T]("FindMany", err)
		return nil, false
	}

	if err := res.All(ctx, &vals); err != nil {
		logError[T]("FindMany", err)
		return nil, false
	}

	return vals, true
}

func Count[T any](filter D, opts ...*options.CountOptions) (cnt int64, ok bool) {
	col := GetCollection[T]()
	if col == nil {
		return -1, false
	}
	cnt, err := col.CountDocuments(ctx, filter, opts...)
	if err != nil {
		logError[T]("Count", err)
		return -1, false
	}
	return cnt, true
}

func DeleteOne[T any](filter D, opts ...*options.DeleteOptions) bool {
	col := GetCollection[T]()
	if col == nil {
		return false
	}
	_, err := col.DeleteOne(ctx, filter, opts...)
	if err != nil {
		logError[T]("DeleteOne", err)
		return false
	}
	return true
}

func DeleteMany[T any](filter D, opts ...*options.DeleteOptions) bool {
	col := GetCollection[T]()
	if col == nil {
		return false
	}
	_, err := col.DeleteMany(ctx, filter, opts...)
	if err != nil {
		logError[T]("DeleteMany", err)
		return false
	}
	return true
}

func ReplaceOne[T any](filter D, new_val T, opts ...*options.FindOneAndReplaceOptions) (old_val T, ok bool) {
	col := GetCollection[T]()
	if col == nil {
		return old_val, false
	}
	res := col.FindOneAndReplace(ctx, filter, new_val, opts...)
	if res.Err() != nil {
		logError[T]("ReplaceOne", res.Err())
		return old_val, false
	}
	if err := res.Decode(&old_val); err != nil {
		logError[T]("ReplaceOne", err)
		return old_val, false
	}
	return old_val, true
}

func Aggregate[T any, Ret any](pipeline A, opts ...*options.AggregateOptions) (ret []Ret, ok bool) {
	col := GetCollection[T]()
	if col == nil {
		return nil, false
	}
	res, err := col.Aggregate(ctx, pipeline, opts...)
	if err != nil {
		logError[T]("Aggregate", err)
		return nil, false
	}
	if err := res.All(ctx, &ret); err != nil {
		logError[T]("Aggregate", err)
		return nil, false
	}
	return ret, true
}
