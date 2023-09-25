import React, {FC} from 'react';

import {FooterStyles} from "./styles";
import {ShrinkContainer} from "../MainLayout/styles";


type footerViewProps = {};

const { Wrapper } = FooterStyles;
const FooterView: FC<footerViewProps> = ({ }) => {
    return (
        <Wrapper>
            <ShrinkContainer>

            </ShrinkContainer>
        </Wrapper>
    );
};

export default FooterView;
