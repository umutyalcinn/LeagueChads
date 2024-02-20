package riotclient

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
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

	log.Printf("%s", string(data))

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
