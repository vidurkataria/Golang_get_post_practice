package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	// "encoding/json"
)

type userInfo struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

func main() {
	fmt.Println("starting server")
	http.HandleFunc("/add_data/", AddUserInfo)
	http.ListenAndServe("127.0.0.1:9342", nil)
}

func AddUserInfo(w http.ResponseWriter, r *http.Request) {
	data := []userInfo{}
	info := userInfo{}

	err := json.NewDecoder(r.Body).Decode(&info)

	if err != nil {
		fmt.Println(err)
		// fmt.Errorf("Oh Common man!! %v", err)
	}
	fmt.Println("Data received: ", info)
	data = append(data, info)
	fmt.Println("New data added: ", data)

	err = json.NewEncoder(w).Encode(data)

	if err != nil {
		fmt.Println(err)
		// fmt.Errorf("Oh Common man!! %v", err)
	}
}
