package parse

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/umutyalcinn/leaguechadsapi/internal/models"
)

func AllChampions(data string) []models.Champion {

	var data_json map[string]interface{}

	err := json.Unmarshal([]byte(data), &data_json)

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

	return champions
}

func MatchSummary(summoner *models.Summoner, matchData map[string]interface{}) (*models.MatchSummary, error) {

    info, _ := matchData["info"]

    infoMap := info.(map[string]interface{})

    participants, _ := infoMap["participants"]

    participantList := participants.([]interface{})

    for _, v := range participantList{
        participantMap := v.(map[string]interface{})

        puuid, _ := participantMap["puuid"]

        if(puuid.(string) == summoner.Puuid){

            championKey, _ := participantMap["championId"]
            championName, _ := participantMap["championName"]
            kills, _ := participantMap["kills"]
            deaths, _ := participantMap["deaths"]
            assists, _ := participantMap["assists"]
            item0, _ := participantMap["item0"]
            item1, _ := participantMap["item1"]
            item2, _ := participantMap["item2"]
            item3, _ := participantMap["item3"]
            item4, _ := participantMap["item4"]
            item5, _ := participantMap["item5"]

            // TODO: get asset urls from utility

            return &models.MatchSummary{
                ChampionKey: uint32(championKey.(float64)),
                ChampionIconUrl: fmt.Sprintf("https://ddragon.leagueoflegends.com/cdn/14.4.1/img/champion/%s.png", championName.(string)),
                Kills: uint32(kills.(float64)),
                Deaths: uint32(deaths.(float64)),
                Assists: uint32(assists.(float64)),
                Item0Id: uint32(item0.(float64)),
                Item1Id: uint32(item1.(float64)),
                Item2Id: uint32(item2.(float64)),
                Item3Id: uint32(item3.(float64)),
                Item4Id: uint32(item4.(float64)),
                Item5Id: uint32(item5.(float64)),
                Item0IconUrl: fmt.Sprintf("https://ddragon.leagueoflegends.com/cdn/14.4.1/img/item/%d.png", uint32(item0.(float64))),
                Item1IconUrl: fmt.Sprintf("https://ddragon.leagueoflegends.com/cdn/14.4.1/img/item/%d.png", uint32(item1.(float64))),
                Item2IconUrl: fmt.Sprintf("https://ddragon.leagueoflegends.com/cdn/14.4.1/img/item/%d.png", uint32(item2.(float64))),
                Item3IconUrl: fmt.Sprintf("https://ddragon.leagueoflegends.com/cdn/14.4.1/img/item/%d.png", uint32(item3.(float64))),
                Item4IconUrl: fmt.Sprintf("https://ddragon.leagueoflegends.com/cdn/14.4.1/img/item/%d.png", uint32(item4.(float64))),
                Item5IconUrl: fmt.Sprintf("https://ddragon.leagueoflegends.com/cdn/14.4.1/img/item/%d.png", uint32(item5.(float64))),
            }, nil
        }
    }

    return nil, nil
}
