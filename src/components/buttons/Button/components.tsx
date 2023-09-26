import React from 'react';
import styled from 'styled-components';
import ButtonView from './index';
import { ICONS } from '../../../services/constant/icons';
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
};
