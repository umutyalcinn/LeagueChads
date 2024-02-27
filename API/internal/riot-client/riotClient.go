// TODO: General fetch function by given url returns data

package riotclient

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/umutyalcinn/leaguechadsapi/internal/models"
	parse "github.com/umutyalcinn/leaguechadsapi/internal/riot-parse"
)

type ApiClient struct {
	apiKey string
}

func New(apiKey string) ApiClient {
	return ApiClient {
		apiKey,
	}
}

func (client *ApiClient) GetMatchByMatchId(matchId string) (map[string]interface{}, error) {
	url := fmt.Sprintf("https://europe.api.riotgames.com/lol/match/v5/matches/%s", matchId)

	data, err := makeLeagueAPIRequest("GET", url, client.apiKey)

	if(err != nil){
		return nil, err
	}

	var match map[string]interface{}

	err = json.Unmarshal([]byte(data), &match)

	if(err != nil){
		return nil, err
	}

    return match, nil

    /*

    var metadata, _ = match["metadata"] 

    metadataMap := metadata.(map[string]interface{})

    participants, _ := metadataMap["participants"]

    participantList := participants.([]interface{})

    var summonerList = make([]models.Summoner, 0, 10)

    for _, v := range participantList {
        summoner, err := client.GetSummonerByPuuid(v.(string))

        if(err != nil){
            return nil, err
        }

        summonerList = append(summonerList, *summoner)
    } 

	return summonerList, nil

    */
}

func (client *ApiClient) GetMatchHistoryByPuuid(puuid string, count uint8) ([]string, error) {
	url := fmt.Sprintf("https://europe.api.riotgames.com/lol/match/v5/matches/by-puuid/%s/ids?start=0&count=%d", puuid, count)

	data, err := makeLeagueAPIRequest("GET", url, client.apiKey)
    
	if(err != nil){
		return nil, err
	}

	var matchHistory []string

	err = json.Unmarshal([]byte(data), &matchHistory)

	if(err != nil){
		return nil, err
	}

	return matchHistory, nil
}

func (client *ApiClient) GetLeaguesBySummonerId(summonerId string) ([]models.League, error) {

	url := fmt.Sprintf("https://tr1.api.riotgames.com/lol/league/v4/entries/by-summoner/%s", summonerId)

	data, err := makeLeagueAPIRequest("GET", url, client.apiKey)

	if(err != nil){
		return nil, err
	}

	var leagues []models.League

	err = json.Unmarshal([]byte(data), &leagues)

	if(err != nil){
		return nil, err
	}

	return leagues, nil
}

func (client *ApiClient) GetAllChampions() ([]models.Champion, error) {

	data, err := makeLeagueAPIRequest("GET", "https://ddragon.leagueoflegends.com/cdn/14.3.1/data/en_US/champion.json", client.apiKey)

	if(err != nil){
		return nil, err
	}

	champions := parse.AllChampions(data)

	return champions, nil
}

func (client *ApiClient) GetFreeChampionRotation() (string, error) {

	data, err := makeLeagueAPIRequest("GET", "https://tr1.api.riotgames.com/lol/platform/v3/champion-rotations", client.apiKey)

	if(err != nil){
		return "", err
	}

	return string(data), nil
}

func (client *ApiClient) GetSummonerByName(summonerName string) (*models.Summoner, error) {

	queryString := strings.ReplaceAll(summonerName, " ", "%20")

	url := fmt.Sprintf("https://tr1.api.riotgames.com/lol/summoner/v4/summoners/by-name/%s", queryString)

	data, err := makeLeagueAPIRequest("GET", url, client.apiKey)

	if(err != nil){
		return nil, err
	}

	var summoner models.Summoner

	err = json.Unmarshal([]byte(data), &summoner)

	if(err != nil){
		return nil, err
	}

	return &summoner, nil
}

func (client *ApiClient) GetSummonerByPuuid(puuid string) (*models.Summoner, error) {

	queryString := strings.ReplaceAll(puuid, " ", "%20")

	url := fmt.Sprintf("https://tr1.api.riotgames.com/lol/summoner/v4/summoners/by-puuid/%s", queryString)

	data, err := makeLeagueAPIRequest("GET", url, client.apiKey)

	if(err != nil){
		return nil, err
	}

	var summoner models.Summoner

	err = json.Unmarshal([]byte(data), &summoner)

	if(err != nil){
		return nil, err
	}

	return &summoner, nil
}

func makeLeagueAPIRequest(method string, url string, api_key string) (string, error) {

	req, err := http.NewRequest(method, url, nil)

	if(err != nil){
		return "", err
	}

	req.Header = http.Header{
		"X-Riot-Token": { api_key },
	}

	client := http.DefaultClient

	res, err := client.Do(req)

	if(err != nil) {
		return "", err
	}

	data, err := io.ReadAll(res.Body)

	if(err != nil) {
		return "", err
	}

	return string(data), err
}
