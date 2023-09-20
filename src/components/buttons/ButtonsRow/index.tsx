import React from 'react';
import {ButtonRowStyles} from "./styles";
import ButtonView from "../Button";


const { Wrapper, HeaderButton } = ButtonRowStyles;

const ButtonRowView = () => {
    return (
        <Wrapper>
                <HeaderButton onPress={() => console.log('zxc')} text="Rectag" />
        </Wrapper>
    );
};

export default ButtonRowView;
