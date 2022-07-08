# Mongorm
An ORM Framework for go-mongo

## How to Use

```go
// 1. init database
ConnectMongo(...)

// 2. init collection binding
// store all `Person` in `db.COL_People`
SetCollectionByName[Person]("COL_People")

// 3. enjoy CRUD
val, ok := Mongorm.FindOne[Person](bson.M{"Name": "xxx"}) 
// or Replace, Delete, ...

```

## To Begin With
Threr are 3 types to bind Mongo-Collections with Go-Type

1. `ConnectMongo` then `SetCollectionByName(col_name)` 
2. `ConnectMongo` and defaults to use `reflect.Type.Name()` as collection name
3. Init Mongo elsewhere and `SetCollection(col)` 

## Dependency
- Golang >= 1.18
- gomongo - [go.mongodb.org/mongo-driver/mongo](https://pkg.go.dev/go.mongodb.org/mongo-driver/mongo)
- logrus - [github.com/sirupsen/logrus](https://pkg.go.dev/github.com/sirupsen/logrus)