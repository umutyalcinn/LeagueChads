import Summoner from "@/models/summoner";
import SummonerHeader from "./components/SummonerHeader";
import MatchSummaryRow from "./components/MatchSummaryRow";
import MatchSummary from "@/models/matchSummary";

async function getData(summonerName: string){
    const res = await fetch(`http://localhost:8080/getSummonerByName/${summonerName}`);

    if(!res.ok){
        throw new Error("Error fetching data");
    }

    return res.json()
}

export default async function Page ({params}:{params: {summonerName: string}}){

    const summonerData: Summoner = await getData(params.summonerName)

    console.log(summonerData)

    return (
        <>
            <SummonerHeader props={{
                summonerIcon: `https://ddragon.leagueoflegends.com/cdn/14.4.1/img/profileicon/${summonerData.profileIconId}.png`,
                summonerName: summonerData.name,
                summonerLevel: summonerData.summonerLevel
            }}/>

            {summonerData.matchHistory.map((v: MatchSummary, i) => {
                return <MatchSummaryRow key={`match-summary-${i}`} props={v}></MatchSummaryRow>
            })}
        </>
    )
}
