package groupie

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type Detail struct {
	Id           int      `json:"id"`
	Image        string   `json:"image"`
	Name         string   `json:"name"`
	Members      []string `json:"members"`
	CreationDate int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
}

type DetailSimple struct {
	Id            int
	Image         string
	Name          string
	Members       []string
	CreationDate  int
	FirstAlbum    string
	Locations     []string
	Dates         []string
	MembersFormat string
}

func Setup() *DetailSimple { // setup les variable pour la page artiste

	var DetailSimple *DetailSimple = new(DetailSimple)

	var zString []string
	zString = append(zString, "0")

	DetailSimple.Id = 0
	DetailSimple.Image = "0"
	DetailSimple.Name = "0"
	DetailSimple.Members = zString
	DetailSimple.CreationDate = 0
	DetailSimple.FirstAlbum = "0"
	DetailSimple.Locations = zString
	DetailSimple.Dates = zString

	return DetailSimple
}

func AllARtiste() []Detail { // recupere tout les artistes

	var m []Detail

	url, err := http.Get("https://groupietrackers.herokuapp.com/api/artists")
	if err != nil {
		os.Exit(1)
	}
	Arstiste_json, err := ioutil.ReadAll(url.Body)
	if err != nil {
		log.Fatal(err)
	}

	if err := json.Unmarshal([]byte(Arstiste_json), &m); err != nil {
		panic(err)
	}

	return m
}

func Format(DetailSimple *DetailSimple) *DetailSimple { // format les donnée pour la page artiste

	var zString []string
	zString = append(zString, "0")

	DetailSimple.MembersFormat = ""

	for i := 0; i < len(DetailSimple.Members); i++ {

		if i == len(DetailSimple.Members)-1 {
			DetailSimple.MembersFormat = DetailSimple.MembersFormat + DetailSimple.Members[i]
		} else {
			DetailSimple.MembersFormat = DetailSimple.MembersFormat + DetailSimple.Members[i] + ", "
		}
	}

	if DetailSimple.MembersFormat == "" {
		DetailSimple.MembersFormat = "0"
	}

	return DetailSimple

}
