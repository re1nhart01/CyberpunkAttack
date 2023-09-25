import React, {FC, PropsWithChildren} from 'react';
import {blurRowStyles} from "./styles";

type blurRowViewProps = PropsWithChildren<{}>

const { Wrapper } = blurRowStyles;

const BlurRowView: FC<blurRowViewProps> = ({ children }) => {
    return (
        <Wrapper>
            { children }
        </Wrapper>
    );
};

export default BlurRowView;