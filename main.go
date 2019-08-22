package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Person struct{
	ID primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Firstname string `json:"firstname,omitempty" bson:"firstname,omitempty"`
	Lastname string `json:"lastname,omitempty" bson:"lastname,omitempty"`
}

var client *mongo.Client

func CreatePersonEndpoint(response http.ResponseWriter, request *http.Request){
	response.Header().Add("content-type", "application/json")
	var person Person
	json.NewDecoder(request.Body).Decode(&person)
	fmt.Println(person)
	collection := client.Database("thepolyglotdeveloper").Collection("people")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	result, _ := collection.InsertOne(ctx, person)
	json.NewEncoder(response).Encode(result)
}

func GetPersonEndpointByID(response http.ResponseWriter, request *http.Request){
	response.Header().Add("content-type", "application/json")
	var person Person
	params := mux.Vars(request)
	id, _ := primitive.ObjectIDFromHex(params["id"])
	collection := client.Database("thepolyglotdeveloper").Collection("people")
	ctx, _ := context.WithTimeout(context.Background(), 10 * time.Second)
	err := collection.FindOne(ctx, Person{ID: id}).Decode(&person)
	if err != nil{
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{"message" : "`+err.Error() +`"}`))
		return
	}
	json.NewEncoder(response).Encode(person)
}

func GetPersonEndpointByLastname(response http.ResponseWriter, request *http.Request){
	response.Header().Add("content-type", "application/json")
	var person Person
	params := mux.Vars(request)
	collection := client.Database("thepolyglotdeveloper").Collection("people")
	ctx, _ := context.WithTimeout(context.Background(), 10 * time.Second)
	err := collection.FindOne(ctx, Person{Lastname: params["lastname"]}).Decode(&person)
	if err != nil{
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{"message" : "`+err.Error() +`"}`))
		return
	}
	json.NewEncoder(response).Encode(person)
}

func GetPersonEndpoint(response http.ResponseWriter, request *http.Request){
	response.Header().Add("content-type", "application/json")
	var people []Person
	collection := client.Database("thepolyglotdeveloper").Collection("people")
	ctx, _ := context.WithTimeout(context.Background(), 10 * time.Second)
	cursor, err:= collection.Find(ctx, bson.M{})
	if err != nil{
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{"message" : "`+err.Error()+`"}`))
		return
	}
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var person Person
		cursor.Decode(&person)
		people = append(people, person)
	}
	if err := cursor.Err() ; err != nil{
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{"message" : "`+err.Error()+`"}`))
		return
	}
	json.NewEncoder(response).Encode(people)
}

func main(){
	fmt.Println("start running go main")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, _ = mongo.Connect(ctx, clientOptions)
	router := mux.NewRouter()
	_ = client
	router.HandleFunc("/person", CreatePersonEndpoint).Methods("POST")
	router.HandleFunc("/person", GetPersonEndpoint).Methods("GET")
	router.HandleFunc("/person/id/{id}", GetPersonEndpointByID).Methods("GET")
	router.HandleFunc("/person/lastname/{lastname}", GetPersonEndpointByLastname).Methods("GET")
	http.ListenAndServe(":12345", router)
}
