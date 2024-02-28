import MatchSummary from "@/models/matchSummary";

export default function MatchSummary ({props}: {props: MatchSummary}){

    return (
        <div className="flex items-center gap-8">
            <img  src={props.championIconUrl}/>
            <div>
                <p>{props.gameMode}</p>
                <span className="text-green-700">{props.kills}</span>/ 
                <span className="text-red-700">{props.deaths}</span>/ 
                <span className="text-gray-700">{props.assists}</span>
            </div>
            <div className="flex gap-2">
                <img src={props.item0Id == 0 ? "" : props.item0IconUrl} />
                <img src={props.item1Id == 0 ? "" : props.item1IconUrl} />
                <img src={props.item2Id == 0 ? "" : props.item2IconUrl} />
                <img src={props.item3Id == 0 ? "" : props.item3IconUrl} />
                <img src={props.item4Id == 0 ? "" : props.item4IconUrl} />
                <img src={props.item5Id == 0 ? "" : props.item5IconUrl} />
            </div>
        </div>
    )
}
