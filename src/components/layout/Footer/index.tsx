import React, { FC } from 'react';

import { useTranslation } from 'react-i18next';
import { FooterStyles } from './styles';
import { ImageViewComponents } from '../../images/ImageView/styles';
import { TypographyComponents } from '../../typography/typography.styles';

type footerViewProps = {};

const { Wrapper, ShrinkContainer, LogoContainer, LeftAlignedCorners, RightAlignedCorners } = FooterStyles;
const { Text14Zekton400, Text12Space400 } = TypographyComponents;

const { LogoRed, CALogo, InstaLogo } = ImageViewComponents;

const FooterView: FC<footerViewProps> = ({ }) => {
  const { t } = useTranslation();
  return (
    <Wrapper>
      <ShrinkContainer>
        <LogoContainer>
          <LogoRed />
          <Text12Space400 color="white">
            LAIDLON
          </Text12Space400>
        </LogoContainer>
        <Text14Zekton400 color="white">
          {t('footerText')}
          {' '}
          <Text14Zekton400 color="white" href="/privacy-policy" selector="a">
            {t('privacyPolicy')}
          </Text14Zekton400>
        </Text14Zekton400>
        <Text14Zekton400 color="white">
          All rights reserved ©
          {' '}
          {new Date().getUTCFullYear()}
        </Text14Zekton400>
        <LeftAlignedCorners />
        <RightAlignedCorners />
      </ShrinkContainer>
    </Wrapper>
  );
};

export default FooterView;
