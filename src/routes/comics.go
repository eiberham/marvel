package routes

import (
	"encoding/json"
	"github.com/user/marvel/src/apis"
	"github.com/user/marvel/src/config"
	"github.com/user/marvel/src/model"
	"github.com/user/marvel/src/util"
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
	conf := config.New()

	client := apis.New()

	url := conf.Marvel.Host + "comics" + util.GetAuthParams()

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