import React, { useCallback } from 'react';
import { navigate } from 'gatsby';
import { ButtonRowStyles } from './styles';

import { TypographyComponents } from '../../typography/typography.styles';

const { Wrapper, HeaderButton, HeaderRightButton, Outer } = ButtonRowStyles;
const { Text16boxed600 } = TypographyComponents;
function ButtonRowView() {
  const handleAboutUs = useCallback(() => {
    navigate('about_us', { replace: false });
  }, []);

  const handleGoMain = useCallback(() => {
    navigate('/', { replace: false });
  }, []);

  const handleGoToYoutube = useCallback(() => {
    navigate('https://www.youtube.com/watch?v=W4agXL6aVow', { replace: false });
  }, []);

  return (
    <Outer>
      <Wrapper>
        <HeaderButton onPress={handleGoMain}>
          <Text16boxed600 k="main" />
        </HeaderButton>
        <HeaderButton onPress={handleAboutUs}>
          <Text16boxed600 k="aboutUs" />
        </HeaderButton>
        {/* <HeaderButton onPress={handleButtonPress}>
          <Text16boxed600 k="aboutGame" />
        </HeaderButton>
        <HeaderButton onPress={handleButtonPress}>
          <Text16boxed600 k="DLC" />
        </HeaderButton>
        <HeaderButton onPress={handleButtonPress}>
          <Text16boxed600 k="LOCALE" />
        </HeaderButton> */}
      </Wrapper>
      <HeaderRightButton onPress={handleGoToYoutube}>
        <Text16boxed600 k="promo" />
      </HeaderRightButton>
    </Outer>
  );
}

export default ButtonRowView;
