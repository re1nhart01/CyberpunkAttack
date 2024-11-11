import React, { FC } from 'react';
import { useTranslation } from 'react-i18next';
import { ShrinkContainer } from '../MainLayout/styles';

import { HeaderStyles } from './styles';
import { TypographyComponents } from '../../typography/typography.styles';
import ButtonRowView from '../../buttons/ButtonsRow';
import { ImageViewComponents } from '../../images/ImageView/styles';
import { donateLink, instagramLink } from '../../../constant/constants';

type headerViewProps = {
  onScrollIntoView: (v: 'subscribe' | 'about' | 'trailer' | 'start') => void;
  hideButtons?: boolean;
};

const { Wrapper, LogoWrapper, ButtonsWrapper, HeaderButton, Separator } = HeaderStyles;

const { Text16Zekton400 } = TypographyComponents;
const { LaidlonLogo, CALogo, InstaLogo } = ImageViewComponents;

const HeaderView: FC<headerViewProps> = ({ onScrollIntoView, hideButtons }) => {
  const { t } = useTranslation();

  const onDonatePress = () => {
    window.location.href = donateLink;
  };

  const onInstagramPress = () => {
    window.location.href = instagramLink;
  };

  return (
    <Wrapper>
      <LogoWrapper onClick={() => onScrollIntoView('start')}>
        <LaidlonLogo />
        <CALogo />
      </LogoWrapper>
      <ButtonsWrapper>
        {!hideButtons && (
          <>
            <HeaderButton onPress={() => onScrollIntoView('subscribe')}>
              <Text16Zekton400>{t('header.subscribe')}</Text16Zekton400>
            </HeaderButton>
            <Separator />
            <HeaderButton onPress={() => onScrollIntoView('about')}>
              <Text16Zekton400>{t('header.about')}</Text16Zekton400>
            </HeaderButton>
            <Separator />
            <HeaderButton onPress={() => onScrollIntoView('trailer')}>
              <Text16Zekton400>{t('header.trailer')}</Text16Zekton400>
            </HeaderButton>
            <Separator />
          </>
        )}
        <HeaderButton onPress={onDonatePress}>
          <Text16Zekton400>{t('header.donate')}</Text16Zekton400>
        </HeaderButton>
        <Separator />
        <HeaderButton onPress={onInstagramPress}>
          <InstaLogo />
        </HeaderButton>
      </ButtonsWrapper>
    </Wrapper>
  );
};

export default HeaderView;
