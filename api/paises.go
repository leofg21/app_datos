package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/devc94/app_datos/config"
	"github.com/devc94/app_datos/models"
	"github.com/gorilla/mux"
)

func getPaises(query string) (paises []models.Paises, err error) {
	db := config.GetConnection()
	defer db.Close()

	rows, err := db.Query(query)
	if err != nil {
		return
	}
	defer rows.Close()

	for rows.Next() {
		p := models.Paises{}
		err = rows.Scan(&p.Codigo, &p.Pais)
		if err != nil {
			return
		}

		paises = append(paises, p)
	}

	return paises, nil
}

// MostrarPaises ...
func MostrarPaises(w http.ResponseWriter, r *http.Request) {
	p, err := getPaises("SELECT * FROM Paises;")
	if err != nil {
		fmt.Println(err)
	}

	err = json.NewEncoder(w).Encode(p)
}

// BuscarPais ...
func BuscarPais(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	pais := params["pais"]
	query := fmt.Sprintf("SELECT * FROM Paises WHERE pais = '%v';", pais)
	p, err := getPaises(query)
	if err != nil {
		fmt.Println(err)
	}

	err = json.NewEncoder(w).Encode(p)
}

// NuevoPais ...
func NuevoPais(w http.ResponseWriter, r *http.Request) {
	db := config.GetConnection()
	if r.Method == "POST" {
		decoder := json.NewDecoder(r.Body)
		p := models.Paises{}
		err := decoder.Decode(&p)
		if err != nil {
			panic(err)
		}
		st, err := db.Prepare("INSERT INTO Paises(codigo, pais) VALUES(?, ?)")
		if err != nil {
			panic(err)
		}
		st.Exec(p.Codigo, p.Pais)
	}
	defer db.Close()
}

// ModificarPais ...
func ModificarPais(w http.ResponseWriter, r *http.Request) {
	db := config.GetConnection()
	params := mux.Vars(r)
	codigo := params["codigo"]
	decoder := json.NewDecoder(r.Body)
	p := models.Paises{}
	err := decoder.Decode(&p)
	if err != nil {
		panic(err)
	}
	st, err := db.Prepare("UPDATE Paises SET pais = ? WHERE codigo = ? ")
	if err != nil {
		panic(err)
	}
	st.Exec(p.Pais, codigo)

	defer db.Close()
}

// EliminarPais ...
func EliminarPais(w http.ResponseWriter, r *http.Request) {
	db := config.GetConnection()
	params := mux.Vars(r)
	codigo := params["codigo"]
	st, err := db.Prepare("DELETE FROM Paises WHERE codigo = ?")
	if err != nil {
		panic(err)
	}
	st.Exec(codigo)

	defer db.Close()
}
