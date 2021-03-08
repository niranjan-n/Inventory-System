package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Item struct {
	ID    string  `json:"ID"`
	Name  string  `json:"Name"`
	Desc  string  `json:"Description"`
	Price float64 `json:"Price"`
}

var inventory []Item

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello Go!!!")
}

func getInventory(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(inventory)
}
func addToInventory(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var item Item
	_ = json.NewDecoder(r.Body).Decode(&item)
	inventory = append(inventory, item)

	json.NewEncoder(w).Encode(item)
}

func deleteItem(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	fmt.Println(params)
	deleteItemAtID(params["ID"])
	json.NewEncoder(w).Encode(inventory)

}
func deleteItemAtID(i string) {
	for index, item := range inventory {
		if item.ID == i {
			//delete the item
			inventory = append(inventory[:index], inventory[index+1:]...)
			break
		}

	}
}

func updateItem(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var item Item
	params := mux.Vars(r)
	deleteItemAtID(params["ID"])
	_ = json.NewDecoder(r.Body).Decode(&item)
	inventory = append(inventory, item)
	json.NewEncoder(w).Encode(inventory)

}

func handleRequest() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", homePage).Methods("GET")
	router.HandleFunc("/inventory", getInventory).Methods("GET")
	router.HandleFunc("/inventory", addToInventory).Methods("POST")
	router.HandleFunc("/inventory/{ID}", deleteItem).Methods("DELETE")
	router.HandleFunc("/update-inventory", updateItem).Methods("PUT")

	log.Fatal(http.ListenAndServe(":8080", router))
}

func main() {

	inventory = append(inventory, Item{
		"1", "Headphone", "Buy this if you don't wanna get bored!!!", 1200,
	}, Item{
		"2", "Cookies", "Yum  Yum Yummmm!!!!", 30})
	handleRequest()
}
