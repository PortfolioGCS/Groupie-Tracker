package groupie

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type Index struct {
	Index []Date
}

type Date struct {
	Id    int      `json:"id"`
	Dates []string `json:"dates"`
}

func ApiDate() *Index { // recupere les dates des concerts

	var vide *Index = new(Index)

	url, err := http.Get("https://groupietrackers.herokuapp.com/api/dates")
	if err != nil {
		os.Exit(1)
	}
	Arstiste_json, err := ioutil.ReadAll(url.Body)
	if err != nil {
		log.Fatal(err)
	}
	json.Unmarshal(Arstiste_json, vide)

	return vide

}
