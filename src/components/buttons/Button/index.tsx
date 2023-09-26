import React, { FC, PropsWithChildren } from 'react';
import { ButtonStyles } from './styles';
import Typography from '../../typography';

const { Wrapper, Button } = ButtonStyles;

type buttonViewProps = PropsWithChildren<{
    className?: string;
    text?: string;
    leftIcon?: JSX.Element;
    rightIcon?: JSX.Element;
    onPress: () => void;
    normalized?: boolean;
    gapInside?: number | string;
}>;

const ButtonView: FC<buttonViewProps> = ({ leftIcon,
  rightIcon,
  text,
  onPress,
  normalized,
  className,
  children,
  gapInside,
}) => {
  return (
    <Wrapper className={className}>
      <Button gapInside={gapInside} defaultcol={normalized} onClick={onPress}>
        {leftIcon || null}
        { children || text }
        {rightIcon || null}
      </Button>
    </Wrapper>
  );
};

export default ButtonView;
