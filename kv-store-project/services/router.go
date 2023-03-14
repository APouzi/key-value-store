package services

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/go-redis/redis/v8"
	"github.com/gorilla/mux"
)

type KeyValue struct{
	Key string `json:"Key"`
	Value string `json:"Value"`
}

var DB *redis.Client
func Router(){
	
	DB = RedisDB()
	router := mux.NewRouter()
	router.HandleFunc("/store/{Key}", KeyGETHandler).Methods("GET")
	router.HandleFunc("/store", KeyValuePOSTHandler).Methods("POST")
	router.HandleFunc("/store/{Key}", KeyDELETEHandler).Methods("DELETE")
	http.ListenAndServe(":8000", router)
}

func KeyGETHandler(w http.ResponseWriter, r *http.Request){
	key := mux.Vars(r)
	keyInput := key["Key"]
	ctx := context.Background()
	val, err := DB.Get(ctx, keyInput).Result()
    if err != nil {
		if errors.Is(err, redis.Nil){
			errorString := fmt.Sprintf("Value of this key doesn't exist: %s", keyInput)
			fmt.Println("Error retrieving key value:", err)
			http.Error(w, errorString, http.StatusBadRequest)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(val)
}

func KeyValuePOSTHandler( w http.ResponseWriter, r *http.Request){
	var keyvalue KeyValue
	ctx := context.Background()
	err := json.NewDecoder(r.Body).Decode(&keyvalue)
	if err != nil{
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	errDB := DB.Set(ctx, keyvalue.Key, keyvalue.Value, 0).Err()
	if errDB != nil{
		panic(err)
		json.NewEncoder(w).Encode("Issue with Setting the key or Redis")
	}
	json.NewEncoder(w).Encode(keyvalue)
	
}

func KeyDELETEHandler(w http.ResponseWriter, r *http.Request){
	key := mux.Vars(r)
	keyInput := key["Key"]
	ctx := context.Background()
	val, err := DB.Del(ctx, keyInput).Result()
    if err != nil {
		fmt.Println("Error retrieving key value:", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(val)
}