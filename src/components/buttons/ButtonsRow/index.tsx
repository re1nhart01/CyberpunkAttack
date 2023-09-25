import React, {useCallback} from 'react';
import {ButtonRowStyles} from "./styles";
import ButtonView from "../Button";


const { Wrapper, HeaderButton, HeaderRightButton, Outer } = ButtonRowStyles;

const ButtonRowView = () => {

    const handleButtonPress = useCallback(() => {

    }, [])

    return (
        <Outer>
        <Wrapper>
            <HeaderButton onPress={handleButtonPress} text="Rectag" />
            <HeaderButton onPress={handleButtonPress} text="Rectag" />
            <HeaderButton onPress={handleButtonPress} text="Rectag" />
            <HeaderButton onPress={handleButtonPress} text="Rectag" />
            <HeaderButton onPress={handleButtonPress} text="Rectag" />
        </Wrapper>
            <HeaderRightButton onPress={handleButtonPress} text="Rectag" />
        </Outer>
    );
};

export default ButtonRowView;
