import MatchSummary from "@/models/matchSummary";

export default function MatchSummary ({props}: {props: MatchSummary}){

    return (
        <div className="flex">
            <img src={props.championIconUrl}/>
            <p>
                <span className="text-green-700">{props.kills}</span>/ 
                <span className="text-red-700">{props.deaths}</span>/ 
                <span className="text-gray-700">{props.assists}</span>
            </p>
        </div>
    )
}
