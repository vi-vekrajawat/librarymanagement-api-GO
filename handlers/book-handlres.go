package handlers

import (
	"context"
	"encoding/json"
	"librarymanagement-system/config"
	"librarymanagement-system/models"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateBook(w http.ResponseWriter,r *http.Request){
	collection := config.DB.Database("books").Collection("library")

	ctx,cancel := context.WithTimeout(context.Background(),time.Second*10)
	defer cancel()

	var books models.Books
	json.NewDecoder(r.Body).Decode(&books)

	result,err := collection.InsertOne(ctx,books)

	if err!=nil{
		http.Error(w,err.Error(),500)
	}

	json.NewEncoder(w).Encode(result)


}

func GetAll (w http.ResponseWriter , r *http.Request){
	collection := config.DB.Database("books").Collection("library")

	ctx,cancel := context.WithTimeout(context.Background(),time.Second*10)

	defer cancel()
 
	var books []models.Books
    cursor,err := collection.Find(ctx,bson.M{})
	

	if err!=nil{
		http.Error(w,err.Error(),500)
	}

	if err =cursor.All(ctx,&books);err!=nil{
		http.Error(w,err.Error(),500)
	}

	json.NewEncoder(w).Encode(&books)

}

func UpdateBook(w http.ResponseWriter , r *http.Request){
	collection := config.DB.Database("books").Collection("library")

	ctx,cancel := context.WithTimeout(context.Background(),time.Second*10)

	defer cancel()

	vars := mux.Vars(r)
	id := vars["id"]
	objectId,err := primitive.ObjectIDFromHex(id)
	if err!=nil{
		http.Error(w,"Invalid Id format",400)
	}

	var Updatebook models.Books

	json.NewDecoder(r.Body).Decode(&Updatebook)

	

	update := bson.M{
		"$set":bson.M{
			"book_author" :     Updatebook.BookAuthor,
			"book_name" :       Updatebook.BookName,
			"book_price" :      Updatebook.BookPrice,
			"book_description" : Updatebook.BookDescription,
		},
	}

	result , err := collection.UpdateOne(ctx,bson.M{"_id":objectId},update)
	 if err!=nil{
		http.Error(w,err.Error(),500)
	 }

	 json.NewEncoder(w).Encode(result)
	 





}

func DeleteById(w http.ResponseWriter, r *http.Request){
	collection := config.DB.Database("books").Collection("library")

	ctx,cancel := context.WithTimeout(context.Background(),time.Second*10)
    defer cancel()

	vars := mux.Vars(r)
	id := vars["id"]
	objectId,err := primitive.ObjectIDFromHex(id)
	if err!=nil{
		http.Error(w,"Invalid Id format",400)
		return
	}
	result , err := collection.DeleteOne(ctx,bson.M{"_id":objectId})

	if err!=nil{
		http.Error(w,err.Error(),500)
	}
	json.NewEncoder(w).Encode(result)
}


type DeleteMultiple struct{
	IDs []string `json:ids`
}
func DeleteAll(w http.ResponseWriter, r *http.Request){
	collection := config.DB.Database("books").Collection("library")

	ctx,cancel := context.WithTimeout(context.Background(),time.Second*100)
	defer cancel()

	var bookData DeleteMultiple
	json.NewDecoder(r.Body).Decode(&bookData)

	var objectIds []primitive.ObjectID
	for _,id := range bookData.IDs{
		objectId,err:= primitive.ObjectIDFromHex(id)
		if err!=nil{
			http.Error(w , "Invalid id format",400)
			return
		}
		objectIds=append(objectIds,objectId)

	}

	filter:= bson.M{
		"_id":bson.M{
			"$in":objectIds,
		},
	}
	result , err := collection.DeleteMany(ctx,filter)

	if err!=nil{
		http.Error(w,err.Error(),500)
	}

	json.NewEncoder(w).Encode(result)
}