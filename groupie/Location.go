package groupie

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type Vide struct {
	Index []DetailLocation
}

type DetailLocation struct {
	Id        int      `json:"id"`
	Locations []string `json:"locations"`
	Continent []int
}

func ApiLocation(Api int) *Vide { // recupere les locations des concerts

	var vide *Vide = new(Vide)

	url, err := http.Get("https://groupietrackers.herokuapp.com/api/locations")
	if err != nil {
		os.Exit(1)
	}
	Arstiste_json, err := ioutil.ReadAll(url.Body)
	if err != nil {
		log.Fatal(err)
	}
	json.Unmarshal(Arstiste_json, vide)

	fmt.Println(len(vide.Index))

	if Api == 2 {

		for i := 0; i < len(vide.Index); i++ {
			vide.Index[i].Continent = Cont(vide.Index[i].Locations)

			fmt.Println(i + 1)
		}
	}
	return vide

}
