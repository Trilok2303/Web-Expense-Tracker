package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type User struct {
	Username string `bson:"username"`
	Email    string `bson:"email"`
	Password string `bson:"password"`
}

const (
	mongoURI       = "mongodb://localhost:27017"
	dbName         = "expense-tracker"
	collectionName = "users"
)

var client *mongo.Client

func initMongoDB() {
	clientOptions := options.Client().ApplyURI(mongoURI)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal("Failed to connect to MongoDB server")
	}

	fmt.Println("Connected to MongoDB")
}

func registerHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Failed to parse form data", http.StatusBadRequest)
		return
	}

	username := r.Form.Get("username")
	email := r.Form.Get("email")
	password := r.Form.Get("password")

	user := User{
		Username: username,
		Email:    email,
		Password: password,
	}

	collection := client.Database(dbName).Collection(collectionName)
	_, err = collection.InsertOne(context.Background(), user)
	if err != nil {
		http.Error(w, "Failed to register user", http.StatusInternalServerError)
		return
	}

	fmt.Println("User registered:", user)

	http.Redirect(w, r, "/static/login.html", http.StatusSeeOther)
}

func main() {
	initMongoDB()

	router := mux.NewRouter()

	// Register routes
	router.HandleFunc("/register", registerHandler).Methods(http.MethodPost)

	// Serve static files
	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("."))))

	// Start the server
	fmt.Println("Server started on http://localhost:8000")
	log.Fatal(http.ListenAndServe(":8000", router))

}
