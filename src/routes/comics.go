package routes

import (
	"encoding/json"
	"github.com/wwleak/marvel/src/apis"
	"github.com/wwleak/marvel/src/config"
	"github.com/wwleak/marvel/src/model"
	"github.com/wwleak/marvel/src/util"
	"io/ioutil"
	"log"
	"net/http"
)

// GetComics godoc
// @Summary List comics
// @Description get comics
// @Tags comics
// @Accept  json
// @Produce  json
// @Param publicKey path string true "Your Public Key"
// @Param privateKey path string true "Your Private Key"
// @Success 200 {object} model.Response
// @Failure 400 {object} model.ResponseError
// @Failure 404 {object} model.ResponseError
// @Failure 500 {object} model.ResponseError
// @Router /comic?publicKey={publicKey}&privateKey={privateKey} [get]
func GetComics(w http.ResponseWriter, r *http.Request) {

	(w).Header().Set("Access-Control-Allow-Origin", "*")
	(w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	(w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

	if (*r).Method == "OPTIONS" {
		return
	}

	conf := config.New()

	client := apis.New()

	url := conf.Marvel.Host + "comics" + util.GetAuthParams()

	resp, err := client.HttpClient.Get(url)

	if err != nil {
		log.Fatalln(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(" Internal Server Error"))
		return
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(" Internal Server Error"))
		return
	}

	if resp.StatusCode != 200 {
		var response model.ResponseError
		err := json.Unmarshal(body, &response)

		if err != nil{
			log.Fatalln(err)
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(" Internal Server Error"))
			return
		}

		w.Header().Add("Content-Type", "application/json")

		json.NewEncoder(w).Encode(response)
	} else {
		var response model.Response
		err := json.Unmarshal(body, &response)

		if err != nil{
			log.Fatalln(err)
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(" Internal Server Error"))
			return
		}

		w.Header().Add("Content-Type", "application/json")

		json.NewEncoder(w).Encode(response)
	}

}