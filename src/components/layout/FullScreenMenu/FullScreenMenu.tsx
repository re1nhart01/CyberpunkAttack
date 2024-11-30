import React from 'react';
import styled from 'styled-components';
import { useTranslation } from 'react-i18next';
import { ICONS } from '../../../constant/icons';
import { HEADER_HEIGHT, contactUs, discordLink, donateLink, instagramLink } from '../../../constant/constants';
import ButtonView from '../../buttons/Button';
import { TypographyComponents } from '../../typography/typography.styles';
import { svgs } from '../../../constant/svgs';
import { ButtonComponents } from '../../buttons/Button/components';

const { Text18Zekton700, Text14Zekton700 } = TypographyComponents;

const { FullScreenMenu, BurgerButtonItem, ButtonRowContainer, HorizontalRowContainer, ButtonViewDef } = {
  FullScreenMenu: styled.div<{ $isOpen: boolean }>`
    position: fixed;
    top: 0;
    left: 0;
    width: 100%;
    height: 100%;
    display: flex;
    flex-direction: column;
    justify-content: space-between;
    padding-bottom: 32px;
    padding-top: ${HEADER_HEIGHT + 70}px;
    align-items: center;
    background-image: url(${ICONS.mainBackground});
    background-repeat: repeat;
    background-size: auto;
    background-position: top left;
    transform: ${({ $isOpen }) => ($isOpen ? 'translateY(0)' : 'translateY(-100%)')};
    transition: transform 0.5s ease;
    z-index: 5;`,
  ButtonRowContainer: styled.div`
        display: flex;
        flex-direction: column;
        align-items: stretch;
        width: 100%;
        padding-left: 30px;
        padding-right: 30px;
        gap: 24px;
    `,
  ButtonViewDef: styled(ButtonView)``,
  HorizontalRowContainer: styled.div`
    display: flex;
    flex-direction: row;
    align-items: center;
    justify-content: center;
    gap: 24px;
    & > div > button > svg > path {
      fill: #01FFFF;
    }
  `,
  BurgerButtonItem: styled(ButtonView)`
    width: 100%;
    background-color: transparent;
    padding-top: 16px;
    padding-bottom: 16px;
    display: flex;
    flex-direction: row;
    justify-content: center;
    & > * {
      text-align: center;
    }
  `,
};

const { SubmitFormButton } = ButtonComponents;

type IProps = {
  onScrollIntoView: (v: 'subscribe' | 'about' | 'trailer' | 'start') => void;
  setIsOpen: React.Dispatch<React.SetStateAction<boolean>>;
  isOpen: boolean;
}

export const FullScreenMenuComponent = ({ isOpen, onScrollIntoView, setIsOpen }: IProps) => {
  const { t } = useTranslation();

  const overrideScroll = (v: 'subscribe' | 'about' | 'trailer' | 'start') => {
    onScrollIntoView(v);
    setIsOpen(false);
  };

  const onItemPress = (url: string) => {
    window.location.href = url;
  };

  const onDonatePress = () => {
    window.location.href = donateLink;
  };

  return (
    <FullScreenMenu $isOpen={isOpen}>
      <ButtonRowContainer>
        <BurgerButtonItem onPress={() => overrideScroll('subscribe')}>
          <Text18Zekton700 color="white">
            {t('header.subscribe')}
          </Text18Zekton700>
        </BurgerButtonItem>
        <BurgerButtonItem onPress={() => overrideScroll('about')}>
          <Text18Zekton700 color="white">
            {t('header.about')}
          </Text18Zekton700>
        </BurgerButtonItem>
        <BurgerButtonItem onPress={() => overrideScroll('trailer')}>
          <Text18Zekton700 color="white">
            {t('header.trailer')}
          </Text18Zekton700>
        </BurgerButtonItem>
        <SubmitFormButton onPress={onDonatePress}>
          <Text18Zekton700 color="black">
            {t('header.donate')}
          </Text18Zekton700>
        </SubmitFormButton>
        <HorizontalRowContainer>
          <ButtonViewDef onPress={() => onItemPress(instagramLink)}>
            <svgs.instagram />
          </ButtonViewDef>
          <ButtonViewDef onPress={() => onItemPress(discordLink)}>
            <svgs.discord />
          </ButtonViewDef>
          <ButtonViewDef onPress={() => onItemPress(contactUs)}>
            <svgs.contactUs />
          </ButtonViewDef>
        </HorizontalRowContainer>
      </ButtonRowContainer>
      <Text14Zekton700 color="rgbaw05">
        All rights reserved ©
        {' '}
        {new Date().getUTCFullYear()}
      </Text14Zekton700>
    </FullScreenMenu>
  );
};
