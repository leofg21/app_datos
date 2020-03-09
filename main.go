package main

import (
	"log"
	"net/http"

	"github.com/devc94/app_datos/api"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/api/paises", api.MostrarPaises)
	r.HandleFunc("/api/paises/{pais}", api.BuscarPais).Methods("GET")
	r.HandleFunc("/api/pais/nuevo", api.NuevoPais).Methods("POST")
	r.HandleFunc("/api/pais/modificar/{codigo}", api.ModificarPais).Methods("PUT")
	r.HandleFunc("/api/pais/eliminar/{codigo}", api.EliminarPais).Methods("DELETE")

	r.HandleFunc("/api/ciudades", api.MostrarCiudades)
	r.HandleFunc("/api/ciudades/{ciudad}", api.BuscarCiudad).Methods("GET")
	r.HandleFunc("/api/ciudad/nueva", api.NuevaCiudad).Methods("POST")
	r.HandleFunc("/api/ciudad/modificar/{id}", api.ModificarCiudad).Methods("PUT")
	r.HandleFunc("/api/ciudad/eliminar/{id}", api.EliminarCiudad).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":5000", r))
}
