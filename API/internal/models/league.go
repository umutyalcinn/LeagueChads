package models

type League struct {
	SummonerId string `json:"summonerId"`
	SummonerName string `json:"summonerName"`
	LeaguePoints uint32 `json:"leaguePoints"`
	Rank string `json:"rank"`
	Wins uint32 `json:"wins"`
	Losses uint32 `json:"losses"`
	Veteran bool `json:"veteran"`
	Inactive bool `json:"inactive"`
	FreshBlood bool `json:"freshBlood"`
	Hotstreak bool `json:"hotStreak"`
}
