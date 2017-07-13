package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Data ...
type Data struct {
	SnugStatus string `json:"snugStatus"`
	Name       string `json:"name"`
	Age        int    `json:"age"`
	Height     string `json:"height,omitempty"`
}

// DataHandler ...
func DataHandler(rw http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	var data Data

	json.NewDecoder(r.Body).Decode(&data)

	fmt.Printf("Data! %+v\n", data)

	json.NewEncoder(rw).Encode(data)
}
