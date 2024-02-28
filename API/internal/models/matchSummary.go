package models

type MatchSummary struct{
    ChampionKey uint32 `json:"championKey"`
    ChampionIconUrl string `json:"championIconUrl"`
    Kills uint32 `json:"kills"`
    Deaths uint32 `json:"deaths"`
    Assists uint32 `json:"assists"`
    Item0Id uint32 `json:"item0Id"`
    Item1Id uint32 `json:"item1Id"`
    Item2Id uint32 `json:"item2Id"`
    Item3Id uint32 `json:"item3Id"`
    Item4Id uint32 `json:"item4Id"`
    Item5Id uint32 `json:"item5Id"`
    Item0IconUrl string `json:"item0IconUrl"`
    Item1IconUrl string `json:"item1IconUrl"`
    Item2IconUrl string `json:"item2IconUrl"`
    Item3IconUrl string `json:"item3IconUrl"`
    Item4IconUrl string `json:"item4IconUrl"`
    Item5IconUrl string `json:"item5IconUrl"`
    GameMode string `json:"gameMode"`
}
