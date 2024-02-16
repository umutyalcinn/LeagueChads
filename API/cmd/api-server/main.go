package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/umutyalcinn/leaguechadsapi/internal/models"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load(".env")

	if(err != nil){
		log.Fatal("fuck")
	}

	api_key := os.Getenv("API_KEY")

	log.Printf("%s", api_key)

	router := gin.Default()

	router.GET("/freechamps", getFreeChampionRotation)
	router.GET("/champions/all", getAllChampions)
	router.GET("/champions/all/:key", getChampionByKey)

	router.Run("localhost:8080")
}

func getFreeChampionRotation(c *gin.Context) {
	data := makeLeagueAPIRequest("GET", "https://tr1.api.riotgames.com/lol/platform/v3/champion-rotations", "RGAPI-8c0273f7-3510-446f-8e17-ed92e0aa4a1f")
	c.String(http.StatusOK, string(data))
}


func getAllChampions (c *gin.Context) {
	data := makeLeagueAPIRequest("GET", "https://ddragon.leagueoflegends.com/cdn/14.3.1/data/en_US/champion.json", "RGAPI-8c0273f7-3510-446f-8e17-ed92e0aa4a1f")
	var data_json map[string]interface{}

	err := json.Unmarshal(data, &data_json)

	if (err != nil) {
		log.Fatalf("error parsing all champ json")
	}

	champions_json := data_json["data"].(map[string]interface{})

	champions := make([]models.Champion, 0)

	for _, v := range champions_json {
		obj := v.(map[string]interface{})
		champions = append(champions, models.Champion{
			Id: obj["id"].(string),
			Key: obj["key"].(string),
			Name: obj["name"].(string),
			Title: obj["title"].(string),
		})
	}

	c.IndentedJSON(http.StatusOK, champions)
}

func getChampionByKey (c *gin.Context) {
	id, err := strconv.Atoi(c.Param("key"))

	if (err != nil) {
		c.String(http.StatusBadRequest, "Invalid champ key")
	}

	data := makeLeagueAPIRequest("GET", "https://ddragon.leagueoflegends.com/cdn/14.3.1/data/en_US/champion.json", "RGAPI-8c0273f7-3510-446f-8e17-ed92e0aa4a1f")
	var data_json map[string]interface{}

	err = json.Unmarshal(data, &data_json)

	if (err != nil) {
		log.Fatalf("error parsing all champ json")
	}

	champions_json := data_json["data"].(map[string]interface{})

	for _, v := range champions_json {
		obj := v.(map[string]interface{})
		champion := models.Champion{
			Id: obj["id"].(string),
			Key: obj["key"].(string),
			Name: obj["name"].(string),
			Title: obj["title"].(string),
		}

		if champion.Key == strconv.Itoa(id) {
			c.IndentedJSON(http.StatusOK, champion)
			return
		}
	}

	c.String(http.StatusNotFound, "Champion not found")
}

func makeLeagueAPIRequest(method string, url string, api_key string) []byte {

	req, err := http.NewRequest(method, url, nil)

	if(err != nil){
		log.Fatalf("error creating request")
	}

	req.Header = http.Header{
		"X-Riot-Token": {"RGAPI-8c0273f7-3510-446f-8e17-ed92e0aa4a1f"},
	}

	client := http.DefaultClient

	res, err := client.Do(req)

	if(err != nil) {
		log.Fatalf("error creating request")
	}

	data, err := io.ReadAll(res.Body)

	if(err != nil) {
		log.Fatalf("error reading body")
	}

	return data
}
