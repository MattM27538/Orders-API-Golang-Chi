package handler

import(
	"fmt"
	"net/http"
)

type Order struct{}

func(o* Order) Create(w http.ResponseWriter, r *http.Request){
	fmt.Println("Create an Order")
}

func(o* Order) List(w http.ResponseWriter, r *http.Request){
	fmt.Println("Called list")
}

func(o* Order) GetByID(w http.ResponseWriter, r *http.Request){
	fmt.Println("Get order by ID")
}

func(o* Order) UpdateByID(w http.ResponseWriter, r *http.Request){
	fmt.Println("Update order by ID")
}

func(o* Order) DeleteByID(w http.ResponseWriter, r *http.Request){
	fmt.Println("Delete order by ID")
}