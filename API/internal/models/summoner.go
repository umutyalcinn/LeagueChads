package models

type Summoner struct {
	Id string `json:"id"`
	AccountId string `json:"accountId"`
	Puuid string `json:"puuid"`
	Name string `json:"name"`
	ProfileIconId uint32 `json:"profileIconId"`
	RevisionDate uint64 `json:"revisionDate"`
	SummonerLevel uint32 `json:"summonerLevel"`
    MatchHistory []MatchSummary `json:"matchHistory"`
}
