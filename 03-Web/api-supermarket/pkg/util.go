package pkg

import (
	"api-supermarket/internal/domain"
	"encoding/json"
	"io"
	"log"
	"os"
)

// Un metodo util para cargar datos de json en bbdd
func FullfilDB(path string) []domain.Product {
	data, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	dataRead, err := io.ReadAll(data)
	if err != nil {
		log.Fatal(err)
	}
	slice := []domain.Product{}
	json.Unmarshal(dataRead, &slice)

	return slice
}

func FullfilClientsDB(path string) []domain.Client {
	data, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	dataRead, err := io.ReadAll(data)
	if err != nil {
		log.Fatal(err)
	}
	slice := []domain.Client{}
	json.Unmarshal(dataRead, &slice)

	return slice
}

// metodo que no se usa en esta app, pero puede ser util en algun momento.
// Gin ya implementa esto a traves de .JSON y .DATA
func MarshalingData(object any) (data []byte) {
	data, err := json.Marshal(object)
	if err != nil {
		log.Fatal(err)
	}

	return
}
