package structures

import (
	"encoding/csv"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

type death struct {
	Provincia  string
	Poblacion  string
	Total      string
	Nuevas     string
	Probables  string
	Lat        string
	Lng        string
	Created_at string
}

func Deaths() []death {
	response, err := http.Get("https://raw.githubusercontent.com/andrab/ecuacovid/master/datos_crudos/muertes/provincias.csv")

	if err != nil {
		log.Fatalln(err)
	}

	body, err := ioutil.ReadAll(response.Body)

	if err != nil {
		log.Fatalln(err)
	}

	reader := csv.NewReader(strings.NewReader(string(body)))

	csvData, err := reader.ReadAll()

	if err != nil {
		log.Fatalln(err)
	}

	var jsonData []death

	for index, record := range csvData {
		if index > 0 {
			jsonData = append(jsonData, death{
				Provincia:  record[0],
				Poblacion:  record[1],
				Total:      record[2],
				Nuevas:     record[3],
				Probables:  record[4],
				Lat:        record[5],
				Lng:        record[6],
				Created_at: record[7],
			})
		}
	}

	return jsonData
}
