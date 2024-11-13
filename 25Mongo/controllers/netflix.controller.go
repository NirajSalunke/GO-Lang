package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/NirajSalunke/modules/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const connectionString string = "mongodb+srv://itsme:nahnotme@niraj098.rjyug.mongodb.net/?retryWrites=true&w=majority&appName=Niraj098"
const dbName string = "Netflix"
const colName string = "WatchList"

// most imp

var collection *mongo.Collection

// connect with mongoDB

func init() {
	// Client
	clientOps := options.Client().ApplyURI(connectionString)
	// connection
	client, err := mongo.Connect(context.TODO(), clientOps)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to DB")
	collection = client.Database(dbName).Collection(colName)
	fmt.Println("Collection instance is ready")

}

// DB Helpers
// inserter
func insertOneMovie(movie models.Netflix) {
	result, err := collection.InsertOne(context.Background(), movie)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Inserted with insertID:- ", result.InsertedID)
}

func updateMovie(movieId string) {
	id, err := primitive.ObjectIDFromHex(movieId)
	if err != nil {
		log.Fatal(err)
	}
	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"watched": true}}
	// oldMovie, err := collection.FindOne(context.Background())
	result, err := collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Updated the movie:- ", result.ModifiedCount)
}

func deleteMovie(movieId string) {
	id, err := primitive.ObjectIDFromHex(movieId)
	if err != nil {
		log.Fatal(err)
	}
	filter := bson.M{"_id": id}
	result, err := collection.DeleteOne(context.Background(), filter, nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Deleted the movie:- ", result.DeletedCount)
}

func deleteAll() int64 {
	result, err := collection.DeleteMany(context.Background(), bson.M{}, nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Deleted the movie:- ", result.DeletedCount)
	return result.DeletedCount
}

func getAll() []bson.M {
	cur, err := collection.Find(context.Background(), bson.M{}, nil)
	if err != nil {
		log.Fatal(err)
	}
	var allDocs []bson.M

	// Here using cur.All
	err = cur.All(context.Background(), &allDocs)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("All Documents:-  ", allDocs)
	// using loops mannual cur.All
	// for cur.Next(context.Background()) {
	// 	var movie bson.M
	// 	err = cur.Decode(&movie)
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// 	allDocs = append(allDocs, movie)
	// }
	// defer cur.Close(context.Background())
	return allDocs
}

// real contollers

func GetAllMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	allMovies := getAll()
	json.NewEncoder(w).Encode(allMovies)

}
