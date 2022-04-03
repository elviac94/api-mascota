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

type KpiMascota struct {
	EspecieMasNumerosa string  `json:"especie mas numerosa"`
	EdadPromedio       string  `json:"edad promedio perro"`
	Desviacion         float32 `json:"desviación estándar edad perro"`
}

var arrPet []Pet

func createPet(w http.ResponseWriter, r *http.Request) {
	var newPet Pet
	err := json.NewDecoder(r.Body).Decode(&newPet)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	arrPet = append(arrPet, newPet)

	json.NewEncoder(w).Encode(newPet)

}

func getKpiPet(w http.ResponseWriter, r *http.Request) {
	mascota := &KpiMascota{
		EspecieMasNumerosa: "Krill Antártico",
		EdadPromedio:       "8,5",
		Desviacion:         12.25,
	}
	json.NewEncoder(w).Encode(mascota)
}

func getListPets(w http.ResponseWriter, r *http.Request) {

	json.NewEncoder(w).Encode(arrPet)
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
	router.HandleFunc("/kpidemascotas", getKpiPet).Methods("GET")
	router.HandleFunc("/lismascotas", getListPets).Methods("GET")

	log.Fatal(server.ListenAndServe())
}
