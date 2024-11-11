import React from 'react';
import styled from 'styled-components';
import ButtonView from './index';
import { ICONS } from '../../../constant/icons';
import { ImageViewComponents } from '../../images/ImageView/styles';

export const ButtonComponents = {
  ShipNovaPoshtaButton: styled(ButtonView).attrs({
    rightIcon: <ImageViewComponents.NovaPoshtaIcon source="npLogo" />,
  })`
      background-size: contain;
      width: 200px;
      height: 51px;
      background: url(${ICONS.npButtonBG}) no-repeat;
      background-size: contain;
      display: flex;
      flex-direction: row;
      align-items: center;
      padding-bottom: 8px;
      justify-content: center;
    `,
  BuyInOneClickButton: styled(ButtonView)`
    background: url(${ICONS.aboutSectionButtonBG});
    background-repeat: no-repeat;
    background-size: 100% 100%;
    padding: 15.5px 48px 15.5px 48px;
  `,
  SmallSocialButton: styled(ButtonView)`
    background: url(${ICONS.aboutSectionButtonBG});
    background-repeat: no-repeat;
    background-size: 100% 100%;
    padding: 12px 43px 12px 43px;
  `,
  SubmitFormButton: styled(ButtonView)`
    padding: 16px 42px 16px 42px;
    background-color: ${({ theme }) => theme.colors.main};
    border-radius: 40px;
    transition: 0.5s;
    cursor: pointer;
    &:hover {
      background-color: ${({ theme }) => theme.colors.white};
    }
    & > * {
      pointer-events: none;
    }
    @media only screen and (max-width: 1024px) {
      & > * {
        text-align: center;
        justify-self: center;
      }
    }
  `,
};
