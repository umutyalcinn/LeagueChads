package main

import (
	"log"
	"net/http"
	"os"
	"strconv"

	riotClient "github.com/umutyalcinn/leaguechadsapi/internal/riot-client"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/gin-contrib/cors"
)

var client riotClient.ApiClient

func main() {

	err := godotenv.Load(".env")

	if(err != nil){
		log.Fatal("error loading .env")
	}

	api_key := os.Getenv("API_KEY")

	client = riotClient.New(api_key)

	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"0"},
		AllowHeaders:     []string{"Origin"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return origin == "https://github.com"
		},
	}))

	router.GET("/freechamps", getFreeChampionRotation)
	router.GET("/champions", getAllChampions)
	router.GET("/champions/:key", getChampionByKey)
	router.GET("/summoner", getSummonerByName)
	router.GET("/leagues/:summonerid", getLeaguesBySummonerId)

	router.Run("localhost:8080")
}

func getLeaguesBySummonerId(c *gin.Context) {

	summonerId := c.Param("summonerid")

	leagues, err := client.GetLeaguesBySummonerId(summonerId)

	if(err != nil){
		log.Fatal("Error getting leagues")
	}

	c.IndentedJSON(http.StatusOK, leagues)
}

func getSummonerByName(c *gin.Context) {

	summonerName, ok := c.GetQuery("summonername")

	if(!ok){
		c.String(http.StatusBadRequest, "Please provide summonername query param")
		return
	}

	summoner, err := client.GetSummonerByName(summonerName)

	if(err != nil){
		log.Fatal("Error getting summoner")
	}

	c.IndentedJSON(http.StatusOK, summoner)
}

func getFreeChampionRotation(c *gin.Context) {

	freeChampions, err := client.GetFreeChampionRotation()

	if(err != nil){
		log.Fatal("Error getting free champion rotation")
	}

	c.String(http.StatusOK, freeChampions)
}


func getAllChampions (c *gin.Context) {

	champions, err := client.GetAllChampions()

	if(err != nil){
		log.Fatal("Error getting all champs")
	}

	c.IndentedJSON(http.StatusOK, champions)
}

func getChampionByKey (c *gin.Context) {
	id, err := strconv.Atoi(c.Param("key"))

	if (err != nil) {
		c.String(http.StatusBadRequest, "Invalid champ key")
		return
	}

	champions, err := client.GetAllChampions()

	if (err != nil) {
		log.Fatal("Error getting all champs")
		return
	}

	for _, v := range champions {
		if v.Key == strconv.Itoa(id) {
			c.IndentedJSON(http.StatusOK, v)
			return
		}
	}

	c.String(http.StatusNotFound, "Champion not found")
}

