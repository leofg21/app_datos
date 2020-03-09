package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/devc94/app_datos/config"
	"github.com/devc94/app_datos/models"
	"github.com/gorilla/mux"
)

func getCiudades(query string) (ciudades []models.Ciudades, err error) {
	db := config.GetConnection()
	defer db.Close()

	rows, err := db.Query(query)
	if err != nil {
		return
	}
	defer rows.Close()

	for rows.Next() {
		c := models.Ciudades{}
		err = rows.Scan(&c.IDCiudades, &c.CodigoPais, &c.Ciudad)
		if err != nil {
			return
		}

		ciudades = append(ciudades, c)
	}

	return ciudades, nil
}

// MostrarCiudades ...
func MostrarCiudades(w http.ResponseWriter, r *http.Request) {
	c, err := getCiudades("SELECT * FROM Ciudades;")
	if err != nil {
		fmt.Println(err)
	}

	err = json.NewEncoder(w).Encode(c)
}

// BuscarCiudad ...
func BuscarCiudad(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	ciudad := params["ciudad"]
	query := fmt.Sprintf("SELECT * FROM datos.ciudades WHERE ciudad = '%v';", ciudad)
	c, err := getCiudades(query)
	if err != nil {
		fmt.Println(err)
	}

	err = json.NewEncoder(w).Encode(c)
}

// NuevaCiudad ...
func NuevaCiudad(w http.ResponseWriter, r *http.Request) {
	db := config.GetConnection()
	if r.Method == "POST" {
		decoder := json.NewDecoder(r.Body)
		c := models.Ciudades{}
		err := decoder.Decode(&c)
		if err != nil {
			panic(err)
		}
		st, err := db.Prepare("INSERT INTO Ciudades(codigo_pais, ciudad) VALUES(?, ?)")
		if err != nil {
			panic(err)
		}
		st.Exec(c.CodigoPais, c.Ciudad)
	}
	defer db.Close()
}

// ModificarCiudad ...
func ModificarCiudad(w http.ResponseWriter, r *http.Request) {
	db := config.GetConnection()
	params := mux.Vars(r)
	idCiudad := params["id"]
	decoder := json.NewDecoder(r.Body)
	c := models.Ciudades{}
	err := decoder.Decode(&c)
	if err != nil {
		panic(err)
	}
	st, err := db.Prepare("UPDATE Ciudades SET codigo_pais = ?, ciudad = ? WHERE id_ciudades = ? ")
	if err != nil {
		panic(err)
	}
	st.Exec(c.CodigoPais, c.Ciudad, idCiudad)

	defer db.Close()
}

// EliminarCiudad ...
func EliminarCiudad(w http.ResponseWriter, r *http.Request) {
	db := config.GetConnection()
	params := mux.Vars(r)
	idCiudad := params["id"]
	st, err := db.Prepare("DELETE FROM Ciudades WHERE id_ciudades = ?")
	if err != nil {
		panic(err)
	}
	st.Exec(idCiudad)

	defer db.Close()
}
