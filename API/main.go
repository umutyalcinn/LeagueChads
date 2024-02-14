package main

import (
	"io"
	"log"
	"net/http"
	"os"

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

	router.Run("localhost:8080")
}

func getFreeChampionRotation(c *gin.Context) {
	data := makeLeagueAPIRequest("GET", "https://tr1.api.riotgames.com/lol/platform/v3/champion-rotations", "RGAPI-8c0273f7-3510-446f-8e17-ed92e0aa4a1f")
	c.String(http.StatusOK, data)
}


func getAllChampions (c *gin.Context) {
	data := makeLeagueAPIRequest("GET", "https://ddragon.leagueoflegends.com/cdn/14.3.1/data/en_US/champion.json", "RGAPI-8c0273f7-3510-446f-8e17-ed92e0aa4a1f")
	c.String(http.StatusOK, data)
}

func makeLeagueAPIRequest(method string, url string, api_key string) string {

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

	return string(data)
}
