import styled from 'styled-components';
import ImageView from './index';

export const ImageViewComponents = {
  FbBlurRowImg: styled(ImageView).attrs({
    source: 'fbBoxGame',
    alterText: 'fbBoxGame',
  })`
      margin-right: 16px;
      max-width:100px;
      width:100%;
    `,
  NovaPoshtaIcon: styled(ImageView).attrs({
    source: 'npLogo',
    alterText: 'NP_LOGO_IMAGE',
  })`
      width: 24px;
      height: 24px;
    `,
  CardImageView: styled(ImageView).attrs({
    source: 'cardView',
    alterText: 'NP_LOGO_IMAGE',
  })`
      width: 95%;
      height: 95%;
    `,
  BackgroundAboutView: styled(ImageView).attrs({
    source: 'backgroundAbout',
    alterText: 'BACKGROUND_ABOUT',
  })`
      width: 100%;
      height: 100%;
      position: relative;
      top: 0;
      left: 0;
      right: 0;
      z-index: 2;
    `,
  SeparatorBlack: styled(ImageView).attrs({
    source: '',
  }),
};
