package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

type Pet struct {
	Nombre          string `json:"nombre"`
	Especie         string `json:"especie"`
	Genero          string `json:"genero"`
	Edad            int    `json:"edad"`
	FechaNacimiento string `json:"fecha nacimiento"`
}

func createPet(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Crear mascota")
	var newPet Pet
	err := json.NewDecoder(r.Body).Decode(&newPet)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	json.NewEncoder(w).Encode(newPet)

}

func main() {
	router := mux.NewRouter()
	port := ":8022"
	server := &http.Server{
		Handler:      router,
		Addr:         port,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	router.HandleFunc("/creamascota", createPet).Methods("POST")

	log.Fatal(server.ListenAndServe())
}
