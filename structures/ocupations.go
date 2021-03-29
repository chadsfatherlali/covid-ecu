package structures

import (
	"encoding/csv"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

type ocupations struct {
	Provincia                                 string
	Canton                                    string
	Disponibilidad_hospitalizacion            string
	Disponibilidad_uci                        string
	Disponibilidad_hospitalizacion_covid      string
	Disponibilidad_uci_adulto_covid           string
	Porcentaje_ocupadas_hospitalizacion       string
	Porcentaje_ocupadas_uci                   string
	Porcentaje_ocupadas_hospitalizacion_covid string
	Porcentaje_ocupadas_uci_adulto_covid      string
	Created_at                                string
}

func Ocupations() []ocupations {
	response, err := http.Get("https://raw.githubusercontent.com/andrab/ecuacovid/master/datos_crudos/camas/cantones.csv")

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

	var jsonData []ocupations

	for index, record := range csvData {
		if index > 0 {
			jsonData = append(jsonData, ocupations{
				Provincia:                                 record[0],
				Canton:                                    record[1],
				Disponibilidad_hospitalizacion:            record[2],
				Disponibilidad_uci:                        record[3],
				Disponibilidad_hospitalizacion_covid:      record[4],
				Disponibilidad_uci_adulto_covid:           record[5],
				Porcentaje_ocupadas_hospitalizacion:       record[6],
				Porcentaje_ocupadas_uci:                   record[7],
				Porcentaje_ocupadas_hospitalizacion_covid: record[8],
				Porcentaje_ocupadas_uci_adulto_covid:      record[9],
				Created_at:                                record[10],
			})
		}
	}

	return jsonData
}
