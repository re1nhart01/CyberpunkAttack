import React, {FC, PropsWithChildren} from 'react';
import { socialButtonStyles } from './styles';

const { SocialButtonInner, SocialButtonContainer, Wrapper, Icon } = socialButtonStyles;

type socialButtonProps = PropsWithChildren<{
    goTo: () => void;
}>;

export const SocialButton: FC<socialButtonProps> = ({ goTo, children }) => {
  return (
    <Wrapper>
      <SocialButtonContainer onClick={goTo}>
        <SocialButtonInner>
          {children}
        </SocialButtonInner>
      </SocialButtonContainer>
    </Wrapper>
  );
};
