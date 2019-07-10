package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/user/marvel/src/config"
	"github.com/user/marvel/src/model"
	"github.com/user/marvel/src/util"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	var router = mux.NewRouter()
	router.HandleFunc("/characters", getCharacters).Methods("GET")
	router.HandleFunc("/characters/{id}", getCharacterById).Methods("GET")

	fmt.Println("Running server ...")
	log.Fatal(http.ListenAndServe(":3000", router))
}

func getCharacters(w http.ResponseWriter, r *http.Request) {
	config := config.New()

	url := config.Marvel.Host + "characters" + util.GetAuthParams()

	resp, err := http.Get(url)

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

func getCharacterById(w http.ResponseWriter, r *http.Request){
	//config := config.New()

	vars := mux.Vars(r)
	id := vars["id"]

	json.NewEncoder(w).Encode(map[string]string{"id": id})
}
