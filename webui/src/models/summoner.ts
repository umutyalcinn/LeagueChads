import MatchSummary from "./matchSummary";

export default interface Summoner {
    id: string,
    accountId: string,
    puuid: string,
    name: string,
    profileIconId: number,
    revisionDate: number,
    summonerLevel: number,
    matchHistory: MatchSummary[]
};
