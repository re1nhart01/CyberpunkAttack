import React, {FC} from 'react';
import {ShrinkContainer} from "../MainLayout/styles";

import {HeaderStyles} from "./styles";


type headerViewProps = {};

const { Wrapper } = HeaderStyles;


const HeaderView: FC<headerViewProps> = ( {} ) => {
    return (
        <Wrapper>
           <ShrinkContainer>
               aboba
           </ShrinkContainer>
        </Wrapper>
    );
};

export default HeaderView;
