import React, {FC, PropsWithChildren} from 'react';
import {RowViewStyleComponent} from "./styles";


export type rowViewProps = PropsWithChildren<{
    pos: "row" | "column";
    className?: string;
    gap: number | string;
    leftM?: number | string;
    rightM?: number | string;
}>;

const RowView: FC<rowViewProps> = ({ leftM, rightM, gap, pos, className, children }) => {
    return (
        <RowViewStyleComponent
            leftM={leftM}
            rightM={rightM}
            gap={gap}
            className={className}
            pos={pos}
        >
            { children }
        </RowViewStyleComponent>
    );
};




export default RowView;