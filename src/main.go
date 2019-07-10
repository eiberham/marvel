package main

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
)


type Response struct {
	Code      		int             	`json:"code"`
	Status    		string          	`json:"status"`
	Copyright 		string 				`json:"copyright"`
	AttributionText string 				`json:"attributionText"`
	AttributionHTML string 				`json:"attributionHTML"`
	Data 			json.RawMessage 	`json:"data"`
	Etag 			string  			`json:"etag"`
}

func init() {
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}

func main() {
	var router = mux.NewRouter()
	router.HandleFunc("/characters", getCharacters).Methods("GET")

	fmt.Println("Running server ...")
	log.Fatal(http.ListenAndServe(":3000", router))
}

func getCharacters(w http.ResponseWriter, r *http.Request) {
	now := time.Now()
	timestamp := strconv.FormatInt(now.UTC().UnixNano(), 10)

	API_URL, _     := os.LookupEnv("MARVEL_API_URL")
	API_PVT_KEY, _ := os.LookupEnv("MARVEL_API_PVT_KEY")
	API_PUB_KEY, _ := os.LookupEnv("MARVEL_API_PUB_KEY")

	str := string(timestamp) + API_PVT_KEY + API_PUB_KEY

	hasher := md5.New()
	hasher.Write([]byte(str))
	hash := hex.EncodeToString(hasher.Sum(nil))

	url := API_URL + "characters?apikey="+ API_PUB_KEY + "&hash=" + hash + "&ts=" + timestamp

	resp, err := http.Get(url)

	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	var response Response
	json.Unmarshal(body, &response)

	w.Header().Add("Content-Type", "application/json")

	json.NewEncoder(w).Encode(response)
}
