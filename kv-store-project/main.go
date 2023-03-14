package main

import (
	"fmt"

	"github.com/APouzi/kv-store-project/services"
)

// One isntance of the DB client being passed around as a refrence to each request. This allows for better performance instead of constantly making a new on in each request.

func main(){
	fmt.Println("\nKey and Value store has launched.  \nPlease input /store with payload to input data or /store/someKey to retrieve some key \n")
	services.Router()
	
}

