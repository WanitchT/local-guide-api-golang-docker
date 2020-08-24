package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/app/local-guide-api-golang-docker/helper"
	"github.com/app/local-guide-api-golang-docker/models"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
)

type ResMsg struct {
	Response_status  string      `json:"response_status"`
	Response_message string      `json:"response_message"`
	Response_data    interface{} `json:"response_data"`
}

func getCommentHeader(w http.ResponseWriter) http.ResponseWriter {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	return w
}
func getComment(w http.ResponseWriter, r *http.Request) {

	w = getCommentHeader(w)
	//var comments models.Comments
	// we created Book array
	var comment_array []models.Comments

	// we get params with mux.
	var params = mux.Vars(r)

	// string to primitive.ObjectID
	guideid, _ := params["guideid"]
	fmt.Println(guideid)

	//Connection mongoDB with helper class
	collection := helper.ConnectDB()

	// bson.M{},  we passed empty filter. So we want to get all data.
	//options := options.Find()
	//sort := bson.D{}
	//sort = append(sort, bson.E{"createdAt", -1})

	//options.SetSort(sort)

	filter := bson.M{"guideid": guideid}
	//cur, err := collection.Find(context.TODO(), filter, options)
	cur, err := collection.Find(context.TODO(), filter)

	if err != nil {
		helper.GetError(err, w)
		return
	}

	// Close the cursor once finished
	/*A defer statement defers the execution of a function until the surrounding function returns.
	simply, run cur.Close() process but after cur.Next() finished.*/
	defer cur.Close(context.TODO())

	for cur.Next(context.TODO()) {

		// create a value into which the single document can be decoded
		var comments models.Comments
		// & character returns the memory address of the following variable.
		err := cur.Decode(&comments) // decode similar to deserialize process.
		if err != nil {
			log.Fatal(err)
		}

		// add item our array	//line 20
		comment_array = append(comment_array, comments)
	}

	var Response ResMsg
	Response.Response_status = "success"
	Response.Response_message = "Get Comment Success!"
	Response.Response_data = comment_array

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

	json.NewEncoder(w).Encode(Response) // encode similar to serialize process.
}

func main() {

	r := mux.NewRouter()
	r.HandleFunc("/go/getComment/{guideid}", getComment).Methods("POST")
	//r.HandleFunc("/go/getCommentById/{id}", getNewByID).Methods("GET")
	//r.HandleFunc("/admin/addNews", createComments).Methods("POST")

	log.Fatal(http.ListenAndServe(":8000", r))

}
