package routes

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/user/marvel/src/apis"
	"github.com/user/marvel/src/config"
	"github.com/user/marvel/src/model"
	"github.com/user/marvel/src/util"
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
// @Success 200 {object} model.Response
// @Failure 400 {string} string
// @Failure 404 {string} string
// @Failure 500 {string} string
// @Router /character [get]
func GetCharacters(w http.ResponseWriter, r *http.Request) {

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

	var response model.Response
	json.Unmarshal(body, &response)


	w.Header().Add("Content-Type", "application/json")

	json.NewEncoder(w).Encode(response)
}


// GetCharacterById godoc
// @Summary Show a character
// @Description get string by ID
// @Tags characters
// @Accept  json
// @Produce  json
// @Param id path int true "Character ID"
// @Success 200 {object} model.Response
// @Failure 400 {string} string
// @Failure 404 {string} string
// @Failure 500 {string} string
// @Router /characters/{id} [get]
func GetCharacterById(w http.ResponseWriter, r *http.Request){
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

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	var response model.Response
	json.Unmarshal(body, &response)

	json.NewEncoder(w).Encode(response)
}
