package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/users",getUsers) //serve pra falar a rota e qual função executar quando bater nessa rota //o users é um endpoint
	fmt.Println("api is on 8080")

	log.Fatal(http.ListenAndServe(":8080", nil)) //se a api nao subir, ele quebra o serviço
}

type User struct {
	Id int
	Name string
}

func getUsers(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Contente=Type", "application/json")
	json.NewEncoder(w).Encode([]User{{
		Id: 1,
		Name: "Lucas",
	}})
}