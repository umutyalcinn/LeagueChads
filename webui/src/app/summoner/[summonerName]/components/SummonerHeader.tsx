export type SummonerHeaderProps = {
    summonerIcon: string
    summonerName: string
    summonerLevel: number
}
export default function SummonerHeader({props}: {props: SummonerHeaderProps}){
    console.log(props.summonerIcon)
    return(
        <div className="flex items-center">
            <img 
                src={props.summonerIcon}
                alt={props.summonerName}
                width={100}
                height={100}
            />
            <div>
                <p>{props.summonerName}</p>
                <p>Seviye: {props.summonerLevel}</p>
            </div>
        </div>
    )
}
