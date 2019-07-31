package routes

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/wwleak/marvel/src/apis"
	"github.com/wwleak/marvel/src/config"
	"github.com/wwleak/marvel/src/model"
	"github.com/wwleak/marvel/src/util"
	"io/ioutil"
	"log"
	"net/http"
)

// GetCharacters godoc
// @Summary List characters
// @Description get characters
// @Tags characters
// @Accept  json
// @Produce  json
// @Param publicKey path string true "Your Public Key"
// @Param privateKey path string true "Your Private Key"
// @Success 200 {object} model.Response
// @Failure 400 {object} model.ResponseError
// @Failure 404 {object} model.ResponseError
// @Failure 500 {object} model.ResponseError
// @Router /character?publicKey={publicKey}&privateKey={privateKey} [get]
func GetCharacters(w http.ResponseWriter, r *http.Request) {

	(w).Header().Set("Access-Control-Allow-Origin", "*")
	(w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	(w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

	if (*r).Method == "OPTIONS" {
		return
	}

	conf := config.New()

	client := apis.New()

	url := conf.Marvel.Host + "characters" + util.GetAuthParams()

	resp, err := client.HttpClient.Get(url)

	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	if resp.StatusCode != 200 {
		var response model.ResponseError
		json.Unmarshal(body, &response)
		json.NewEncoder(w).Encode(response)
	} else {
		var response model.Response
		json.Unmarshal(body, &response)
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	}
}


// GetCharacterById godoc
// @Summary Shows a character
// @Description get string by ID
// @Tags characters
// @Accept  json
// @Produce  json
// @Param id path int true "Character ID"
// @Param publicKey path string true "Your Public Key"
// @Param privateKey path string true "Your Private Key"
// @Success 200 {object} model.Response
// @Failure 400 {object} model.ResponseError
// @Failure 404 {object} model.ResponseError
// @Failure 500 {object} model.ResponseError
// @Router /characters/{id}?publicKey={publicKey}&privateKey={privateKey} [get]
func GetCharacterById(w http.ResponseWriter, r *http.Request){

	(w).Header().Set("Access-Control-Allow-Origin", "*")
	(w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	(w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

	if (*r).Method == "OPTIONS" {
		return
	}

	conf := config.New()

	client := apis.New()

	vars := mux.Vars(r)
	id := vars["id"]

	url := conf.Marvel.Host + "characters/" + id + util.GetAuthParams()

	resp, err := client.HttpClient.Get(url)

	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)

	if resp.StatusCode != 200 {
		var response model.ResponseError
		json.Unmarshal(body, &response)
		json.NewEncoder(w).Encode(response)
	} else {
		var response model.Response
		json.Unmarshal(body, &response)
		json.NewEncoder(w).Encode(response)
	}

}
