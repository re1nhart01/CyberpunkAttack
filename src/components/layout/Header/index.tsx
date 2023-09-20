import React, {FC} from 'react';
import {ShrinkContainer} from "../MainLayout/styles";

import {HeaderStyles} from "./styles";
import {TypographyComponents} from "../../typography/typography.styles";
import ButtonRowView from "../../buttons/ButtonsRow";


type headerViewProps = {};

const { Wrapper, LogoWrapper, ButtonsWrapper } = HeaderStyles;

const { LogoText, SubLogoText } = TypographyComponents;

const HeaderView: FC<headerViewProps> = ( {} ) => {
    return (
        <Wrapper>
            <LogoWrapper>
                <LogoText>Cyberpunk</LogoText>
                <SubLogoText>Attack</SubLogoText>
            </LogoWrapper>
            <ButtonsWrapper>
                <ButtonRowView></ButtonRowView>
            </ButtonsWrapper>
        </Wrapper>
    );
};

export default HeaderView;
