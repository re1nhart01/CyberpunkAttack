import React, {useCallback} from 'react';
import {ButtonRowStyles} from "./styles";
import ButtonView from "../Button";
import {TypographyComponents} from "../../typography/typography.styles";


const { Wrapper, HeaderButton, HeaderRightButton, Outer } = ButtonRowStyles;
const { Text16boxed600 } = TypographyComponents;
const ButtonRowView = () => {
    const handleButtonPress = useCallback(() => {

    }, [])

    return (
        <Outer>
        <Wrapper>
            <HeaderButton onPress={handleButtonPress} >
                <Text16boxed600 k="main" />
            </HeaderButton>
            <HeaderButton onPress={handleButtonPress} >
                <Text16boxed600 k="knowlegdeBase" />
            </HeaderButton>
            <HeaderButton onPress={handleButtonPress} >
                <Text16boxed600 k="aboutGame" />
            </HeaderButton>
            <HeaderButton onPress={handleButtonPress} >
                <Text16boxed600 k="DLC" />
            </HeaderButton>
            <HeaderButton onPress={handleButtonPress} >
                <Text16boxed600 k="LOCALE" />
            </HeaderButton>
        </Wrapper>
            <HeaderRightButton onPress={handleButtonPress}>
                <Text16boxed600 k="buyNow" />
            </HeaderRightButton>
        </Outer>
    );
};

export default ButtonRowView;
