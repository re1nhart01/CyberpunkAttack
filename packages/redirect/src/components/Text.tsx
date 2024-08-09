import { FC } from "react";

type textProps = {
    text: string;
}


export const Text: FC<textProps> = ({ text }) => {
    
    
    return (
        <div className="">
            {text}
        </div>
    )
}