package tests

import (
	// "context"
	"bytes"
	"net/http"
	// "net/http/httptest"
	"testing"
	"encoding/json"
	// "github.com/APouzi/kv-store-project/services"
)

type RequestBody struct {
    Key string `json:"Key"`
    Value  string `json:"Value"`
}

type TestResponseAfterDeletion struct{
	Value string `json:"value"`
}

func TestDeletion(t *testing.T) {
	t.Log("Starting Delete test")
	postBody := RequestBody{Key: "Alice", Value: "Value of Alice"}

	jsonBody, err := json.Marshal(postBody)
	t.Log("jsonBody",string(jsonBody))
    if err != nil{
        t.Fatalf("issue with creating json body")
    }
	post, err := http.Post("http://localhost:8000/store","application/json", bytes.NewBuffer(jsonBody))
	if err != nil{
		t.Fatalf("Issue with posting")
	}
	var postResponse RequestBody
	json.NewDecoder(post.Body).Decode(&postResponse)
	
	
	if post.StatusCode != http.StatusOK{
		t.Fatalf("Status code is not 200")
	}
	t.Log("creation of key-value completed",postResponse.Key)
	// Get the value for the key from the store and verify it matches the added value


	// Deletion
	req, err := http.NewRequest("DELETE", "http://localhost:8000/store/"+postResponse.Key, nil)
	if err != nil {
		t.Log("Error")
	}

	deleteResp, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Log("Error")
	}
	defer deleteResp.Body.Close()

	


	res, err := http.Get("http://localhost:8000/store/"+postResponse.Key)
	if err != nil {
		t.Fatalf("failed to make request: %v", err)
	}
	// check response status code
	if res.StatusCode == http.StatusBadRequest {
		t.Logf("Successfully deleted")
	} else {
		t.Fatalf("Did not delete the key as expected")
	}
	t.Log(res)
}

func TestOverWrite(t *testing.T) {
	t.Log("Starting Overwrite test")
	postBody := RequestBody{Key: "Alice", Value: "Value of Alice"}

	jsonBody, err := json.Marshal(postBody)
	t.Log("Before OverWrite",string(jsonBody))
    if err != nil{
        t.Fatalf("issue with creating json body")
    }
	post, err := http.Post("http://localhost:8000/store","application/json", bytes.NewBuffer(jsonBody))
	if err != nil{
		t.Fatalf("Issue with posting")
	}
	var postResponse RequestBody
	json.NewDecoder(post.Body).Decode(&postResponse)
	
	
	if post.StatusCode != http.StatusOK{
		t.Fatalf("Status code is not 200")
	}
	t.Log("creation of key-value completed",postResponse.Key)

	// Overwrite
	t.Log("Creation completed, now time to overwrite said key")
	postBodyOvewrite := RequestBody{Key: "Alice", Value: "Changed Value"}

	jsonBodyOverWrite, err := json.Marshal(postBodyOvewrite)
	t.Log("OverWrite",string(jsonBody))
    if err != nil{
        t.Fatalf("issue with creating json body")
    }
	postOverWrite, err := http.Post("http://localhost:8000/store","application/json", bytes.NewBuffer(jsonBodyOverWrite))
	if err != nil{
		t.Fatalf("Issue with posting")
	}
	var postResponseOverWrite RequestBody
	json.NewDecoder(postOverWrite.Body).Decode(&postResponseOverWrite)
	
	
	if postOverWrite.StatusCode != http.StatusOK{
		t.Fatalf("Status code is not 200")
	}
	t.Log("creation of key-value completed",postResponseOverWrite.Key)

	if postResponseOverWrite == postResponse{
		t.Fatalf("Test failed with overwriting")
	}else{
		t.Log("Successfully overridden the keys")
	}
	// Delete the data created above
	req, err := http.NewRequest("DELETE", "http://localhost:8000/store/"+postResponseOverWrite.Key, nil)
	if err != nil {
		t.Log("Error")
	}

	deleteResp, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Log("Error")
	}
	defer deleteResp.Body.Close()
	t.Logf("OverWrite test is complete, deleting test data from DB")

}
