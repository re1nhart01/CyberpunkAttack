import React, {FC} from 'react';
import {ButtonStyles} from "./styles";


const { Wrapper, Button } = ButtonStyles;

type buttonViewProps = {
    className?: string;
    text: string;
    leftIcon?: JSX.Element;
    rightIcon?: JSX.Element;
    onPress: () => void;
    normalized?: boolean;
}

const ButtonView: FC<buttonViewProps> = ({ leftIcon,
   rightIcon,
   text,
   onPress,
   normalized,
   className,
 }) => {
    return (
        <Wrapper className={className}>
            <Button defaultcol={normalized} onClick={onPress}>
                {leftIcon || null}
                {text}
                {rightIcon || null}
            </Button>
        </Wrapper>
    );
};

export default ButtonView;
