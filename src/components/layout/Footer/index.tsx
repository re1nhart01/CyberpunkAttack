import React, {FC} from 'react';

import {FooterStyles} from "./styles";


type footerViewProps = {};

const { Wrapper } = FooterStyles;
const FooterView: FC<footerViewProps> = ( {} ) => {
    return (
        <Wrapper>
           Footer
        </Wrapper>
    );
};

export default FooterView;
