package Mongorm

import (
	"fmt"
	"reflect"

	. "go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type CollectionNotFoundError[T any] struct{}

func (e CollectionNotFoundError[T]) Error() string {
	var v T
	typename := reflect.TypeOf(v).Name()

	return fmt.Sprintf("[Mongorm] Cannot Find Collection for type (%s)", typename)
}

func InsertOne[T any](val T, opts ...*options.InsertOneOptions) error {
	col := GetCollection[T]()
	if col == nil {
		return CollectionNotFoundError[T]{}
	}
	_, err := col.InsertOne(ctx, val, opts...)
	if err != nil {
		logError[T]("Add", err)
	}
	return err
}

func FindOne[T any](filter D, opts ...*options.FindOneOptions) (val T, err error) {
	col := GetCollection[T]()
	if col == nil {
		return val, CollectionNotFoundError[T]{}
	}
	res := col.FindOne(ctx, filter, opts...)
	if res.Err() != nil {
		logError[T]("FindOne", res.Err())
		return val, res.Err()
	}

	if err := res.Decode(&val); err != nil {
		logError[T]("FindOne", err)
		return val, err
	}

	return val, nil
}

func FindMany[T any](filter D, opts ...*options.FindOptions) (vals []T, err error) {
	col := GetCollection[T]()
	if col == nil {
		return nil, CollectionNotFoundError[T]{}
	}
	res, err := col.Find(ctx, filter, opts...)
	if err != nil {
		logError[T]("FindMany", err)
		return nil, err
	}

	if err := res.All(ctx, &vals); err != nil {
		logError[T]("FindMany", err)
		return nil, err
	}

	return vals, nil
}

func Count[T any](filter D, opts ...*options.CountOptions) (cnt int64, err error) {
	col := GetCollection[T]()
	if col == nil {
		return -1, CollectionNotFoundError[T]{}
	}
	cnt, err = col.CountDocuments(ctx, filter, opts...)
	if err != nil {
		logError[T]("Count", err)
		return -1, err
	}
	return cnt, nil
}

func DeleteOne[T any](filter D, opts ...*options.DeleteOptions) error {
	col := GetCollection[T]()
	if col == nil {
		return CollectionNotFoundError[T]{}
	}
	_, err := col.DeleteOne(ctx, filter, opts...)
	if err != nil {
		logError[T]("DeleteOne", err)
		return err
	}
	return nil
}

func DeleteMany[T any](filter D, opts ...*options.DeleteOptions) error {
	col := GetCollection[T]()
	if col == nil {
		return CollectionNotFoundError[T]{}
	}
	_, err := col.DeleteMany(ctx, filter, opts...)
	if err != nil {
		logError[T]("DeleteMany", err)
		return err
	}
	return nil
}

func ReplaceOne[T any](filter D, new_val T, opts ...*options.FindOneAndReplaceOptions) (old_val T, err error) {
	col := GetCollection[T]()
	if col == nil {
		return old_val, CollectionNotFoundError[T]{}
	}
	res := col.FindOneAndReplace(ctx, filter, new_val, opts...)
	if res.Err() != nil {
		logError[T]("ReplaceOne", res.Err())
		return old_val, res.Err()
	}
	if err := res.Decode(&old_val); err != nil {
		logError[T]("ReplaceOne", err)
		return old_val, err
	}
	return old_val, nil
}

func UpdateOne[T any](filter D, upd D, opts ...*options.UpdateOptions) error {
	col := GetCollection[T]()
	if col == nil {
		return CollectionNotFoundError[T]{}
	}
	_, err := col.UpdateOne(ctx, filter, upd, opts...)
	if err != nil {
		logError[T]("UpdateOne", err)
		return err
	}
	return nil
}

func UpdateMany[T any](filter D, upd D, opts ...*options.UpdateOptions) error {
	col := GetCollection[T]()
	if col == nil {
		return CollectionNotFoundError[T]{}
	}
	_, err := col.UpdateMany(ctx, filter, upd, opts...)
	if err != nil {
		logError[T]("UpdateMany", err)
		return err
	}
	return nil
}

// pipeline: bson.A or "Mongorm/PipelineBuilder".PB
func Aggregate[T any, Ret any](pipeline A, opts ...*options.AggregateOptions) (ret []Ret, err error) {
	col := GetCollection[T]()
	if col == nil {
		return nil, CollectionNotFoundError[T]{}
	}
	res, err := col.Aggregate(ctx, pipeline, opts...)
	if err != nil {
		logError[T]("Aggregate", err)
		return nil, err
	}
	if err := res.All(ctx, &ret); err != nil {
		logError[T]("Aggregate", err)
		return nil, err
	}
	return ret, nil
}
