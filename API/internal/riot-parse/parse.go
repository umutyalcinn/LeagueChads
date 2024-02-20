package parse

import (
	"encoding/json"
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
