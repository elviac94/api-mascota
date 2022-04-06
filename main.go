package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

// Pet model info
// @Description  Datos de la mascota registrada

type Pet struct {
	Nombre          string `json:"nombre"`           // Nombre de la mascota
	Especie         string `json:"especie"`          // Especie de la mascota
	Genero          string `json:"genero"`           // Genero de la mascota
	Edad            int    `json:"edad"`             // Edad de la mascota
	FechaNacimiento string `json:"fecha nacimiento"` // Fecha de nacimiento de la mascota
}

// KpiMascota model info
// @Description  Información general

type KpiMascota struct {
	EspecieMasNumerosa string  `json:"especie mas numerosa"`           // Especie más numerosa
	EdadPromedio       string  `json:"edad promedio perro"`            // Edad promedio de mis mascotas
	Desviacion         float32 `json:"desviación estándar edad perro"` // Desviación estándar de las edades de las mascotas
}

var arrPet []Pet

// createPet godoc
// @Summary  Crea el registro de una mascota con los datos introducidos
// @ID       create-pet
// @Produce  json
// @Param    pet  body Pet  true  "nombre,especie,genero,edad,fecha de nacimiento"
// @success 200 {object} Pet
// @Failure  405  not  allowed
// @Router   /crearmascota [post]
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

// getKpiPet godoc
// @Summary  Obtiene los datos de kpi mascota
// @ID       get-kpi-pet
// @Produce  json
// @success 200 {object}  KpiMascota
// @Failure  405  not  allowed
// @Router   /kpidemascotas [get]
func getKpiPet(w http.ResponseWriter, r *http.Request) {
	mascota := &KpiMascota{
		EspecieMasNumerosa: "Krill Antártico",
		EdadPromedio:       "8,5",
		Desviacion:         12.25,
	}
	json.NewEncoder(w).Encode(mascota)
}

// getListPets godoc
// @Summary  Obtiene el listado de mascotas
// @ID       get-list-pet
// @Produce  json
// @Success   200 {object} []Pet
// @Failure  405  not  allowed
// @Router   /lismascotas [get]
func getListPets(w http.ResponseWriter, r *http.Request) {

	json.NewEncoder(w).Encode(arrPet)
}

// @title        API - Mascota"
// @version      1.0
// @description  Microservicio con 3 endpoints:

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  MIT
// @license.url   https://opensource.org/licenses/MIT

// @host                     localhost:8022
// @BasePath                 /
// @query.collection.format  multi

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
	router.PathPrefix("/swagger/").Handler(http.StripPrefix("/swagger", http.FileServer(http.Dir("swagger"))))
	router.PathPrefix("/docs/").Handler(http.StripPrefix("/docs", http.FileServer(http.Dir("docs"))))

	log.Fatal(server.ListenAndServe())
}
