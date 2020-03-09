package models

// Ciudades estructura
type Ciudades struct {
	IDCiudades uint32 `json:"id_ciudad"`
	CodigoPais string `json:"codigo_pais"`
	Ciudad     string `json:"ciudad"`
}
