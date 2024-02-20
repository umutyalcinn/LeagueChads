import Summoner from "@/models/summoner";

async function getData(summonerName: string){
    const res = await fetch(`http://localhost:8080/summoner?summonername=${summonerName}`);

    if(!res.ok){
        throw new Error("Error fetching data");
    }

    return res.json()
}

export default async function Page ({params}:{params: {summonerName: string}}){

    const summonerData: Summoner = await getData(params.summonerName)

    return (
        <>
            <div>{summonerData.id}</div>
            <div>{summonerData.accountId}</div>
            <div>{summonerData.name}</div>
            <div>{summonerData.puuid}</div>
            <div>{summonerData.revisionDate}</div>
            <img 
                src={`https://ddragon.leagueoflegends.com/cdn/14.3.1/img/profileicon/${summonerData.profileIconId}.png`}
                width={31}
                height={31}
            />
            <div>{summonerData.profileIconId}</div>
            <div>{summonerData.summonerLevel}</div>
        </>
    )
}
