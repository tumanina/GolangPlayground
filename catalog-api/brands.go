package main

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Brand struct {
	Id   int32
	Name string
}

func getBrands() []Brand {
	brands := GetBrandsCollection()

	findOptions := options.Find()
	findOptions.SetLimit(50)

	cursor, err := brands.Find(context.TODO(), bson.D{})
	if err != nil {
		log.Fatal(err)
	}

	var results []Brand
	if err = cursor.All(context.Background(), &results); err != nil {
		log.Fatal(err)
	}

	return results
}

func addBrand(brand Brand) {
	brands := GetBrandsCollection()

	_, err := brands.InsertOne(context.TODO(), bson.D{{"id", brand.Id}, {"name", brand.Name}})
	if err != nil {
		log.Fatal(err)
	}
}

func deleteBrand(id int32) {
	brands := GetBrandsCollection()

	filter := bson.D{
		{"$and",
			bson.A{
				bson.D{
					{"id", bson.D{{"$gt", id}}},
				},
			},
		},
	}

	_, err := brands.DeleteOne(context.TODO(), filter)
	if err != nil {
		log.Fatal(err)
	}
}

func GetBrandsCollection() *mongo.Collection {
	databaseConfiguration := GetDatabaseConfiguration()
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(databaseConfiguration.ConnectionString))

	if err != nil {
		log.Fatal(err)
	}

	/*todo
	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			log.Fatal(err)
		}
	}()*/

	return client.Database(databaseConfiguration.DatabaseName).Collection("brands")
}
