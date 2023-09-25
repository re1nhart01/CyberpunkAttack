import React, {FC} from 'react';
import {ShrinkContainer} from "../MainLayout/styles";

import {HeaderStyles} from "./styles";
import {TypographyComponents} from "../../typography/typography.styles";
import ButtonRowView from "../../buttons/ButtonsRow";


type headerViewProps = {};

const { Wrapper, LogoWrapper, ButtonsWrapper } = HeaderStyles;

const { Text22space400, Text11aquantix400 } = TypographyComponents;

const HeaderView: FC<headerViewProps> = ( {} ) => {
    return (
        <Wrapper>
            <LogoWrapper>
                <Text22space400>Cyberpunk</Text22space400>
                <Text11aquantix400>Attack</Text11aquantix400>
            </LogoWrapper>
            <ButtonsWrapper>
                <ButtonRowView></ButtonRowView>
            </ButtonsWrapper>
        </Wrapper>
    );
};

export default HeaderView;
