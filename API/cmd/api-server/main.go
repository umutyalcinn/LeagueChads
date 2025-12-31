package main

import (
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/umutyalcinn/leaguechadsapi/internal/models"
	riotClient "github.com/umutyalcinn/leaguechadsapi/internal/riot-client"
	parse "github.com/umutyalcinn/leaguechadsapi/internal/riot-parse"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
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
    router.GET("/getSummonerByName/:summonerName", getSummonerByName)
    router.GET("/getSummonerByPuuid/:puuid", getSummonerByPuuid)
	router.GET("/leagues/:summonerId", getLeaguesBySummonerId)

	router.GET("/matchHistory/:puuid", getMatchHistoryByPuuid)
	router.GET("/match/:matchId", getMatchByMatchId)

	router.Run("192.168.1.101:8080")
}

func getMatchHistoryByPuuid(c *gin.Context){
	puuid := c.Param("puuid")

    countQuery, ok := c.GetQuery("count")

	if(!ok){
		c.String(http.StatusBadRequest, "Please provide count query param")
		return
	}

    count, err := strconv.Atoi(countQuery)

	if(err != nil){
        c.String(http.StatusBadRequest, "Invalid count query param")
        return
	}

    matchHistory, err := client.GetMatchHistoryByPuuid(puuid, uint8(count))

    c.IndentedJSON(http.StatusOK, matchHistory)
}

func getMatchByMatchId(c *gin.Context) {

	matchId := c.Param("matchId")

	match, err := client.GetMatchByMatchId(matchId)

	if(err != nil){
		log.Fatal("Error getting match")
	}

	c.IndentedJSON(http.StatusOK, match)
}

func getLeaguesBySummonerId(c *gin.Context) {

	summonerId := c.Param("summonerId")

	leagues, err := client.GetLeaguesBySummonerId(summonerId)

	if(err != nil){
		log.Printf("Error getting leagues")
	}

	c.IndentedJSON(http.StatusOK, leagues)
}

func getSummonerByName(c *gin.Context) {

	summonerName := c.Param("summonerName")

	summoner, err := client.GetSummonerByName(summonerName)

	if(err != nil){
		log.Printf("Error getting summoner by name")
	}

    matchHistory, err := client.GetMatchHistoryByPuuid(summoner.Puuid, 20)

    if(err != nil){
		log.Printf("Error getting match history by name")
        c.String(http.StatusInternalServerError, "Error getting history by name")
        return
    }

    matchSummaries := make([]models.MatchSummary, 0, 20)

    for _, v := range matchHistory{
        matchData, err := client.GetMatchByMatchId(v)

        if(err != nil){
            log.Printf("Error getting match")
            c.String(http.StatusInternalServerError, "Error getting")
            return
        }

        summary, err := parse.MatchSummary(summoner, matchData)

        if(err != nil){
            log.Printf("Error parsing match summary")
            c.String(http.StatusInternalServerError, "Error parsing match summary")
            return
        }

        matchSummaries = append(matchSummaries, *summary)

    }

    summoner.MatchHistory = matchSummaries

	c.IndentedJSON(http.StatusOK, summoner)
}

func getSummonerByPuuid(c *gin.Context) {

    puuid := c.Param("puuid")

	summoner, err := client.GetSummonerByPuuid(puuid)

	if(err != nil){
		log.Printf("Error getting summoner by puuid")
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

