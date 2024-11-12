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
    source: 'blackSeparator',
    alterText: 'BLACK_SEPARATOR_1',
  })`
    width: 100%;
  `,
  IRLGameView: styled(ImageView).attrs({
    source: 'irlGame',
    alterText: 'IRL_GAME_TOP_VIEW',
  })`
    max-width: 517px;
    height: 420px;
  `,
  RolesBannerView: styled(ImageView).attrs({
    source: 'rolesBanner',
    alterText: 'ROLES_BANNER_IMAGE',
  })`
    max-width: 770px;
  `,
  SocialImage: styled(ImageView).attrs(({ source }) => ({
    source,
    alterText: 'SOCIAL_MEDIA_BUTTON_IMAGE',
  }))`
    width: 24px;
    height: 24px;
  `,
  AboutUsImage: styled(ImageView).attrs(({ source }) => ({
    source,
    alterText: 'ABOUT_US',
  }))`
    width: 50vw;
    object-fit: contain;
  `,
  LaidlonLogo: styled(ImageView).attrs({
    source: 'laidlonLogo',
    alterText: 'LAIDLON_LAIDLON_IMAGE',
  })`
    max-width: 33px;
    max-height: 33px;
  `,
  CALogo: styled(ImageView).attrs({
    source: 'caLogo',
    alterText: 'CYBERPUNK_ATTACK_LOGO',
  })`
    max-width: 187px;
    @media only screen and (max-width: 1024px) {
      max-width: 140px;
    }
  `,
  InstaLogo: styled(ImageView).attrs({
    source: 'instagramIcon',
    alterText: 'CYBERPUNK_INSTAGRAM_LOGO',
  })``,
  HeaderIllustration: styled(ImageView).attrs<{ isMobile: boolean; }>(({ isMobile }) => ({
    source: isMobile ? 'headerIllustrationMobile' : 'headerIllustration',
    alterText: 'HEADER_ILLUSTRATION_IMAGE',
  }))`
    width: 100%;
    overflow: hidden;
  `,
  AlertImage: styled(ImageView).attrs({
    source: 'alert',
    alterText: 'SECTION_ALERT_IMAGE',
  })`
  `,
  SeparatorBlue: styled(ImageView).attrs({
    source: 'separator',
    alterText: 'SECTION_ALERT_IMAGE',
  })`
    width: 100%;
  `,
  VersusIllustration: styled(ImageView).attrs({
    source: 'versusIllustration',
    alterText: 'VERSUS_ILLUSTRATION_IMAGE',
  })`
    width: 100%;
    height: 99% !important;
    object-fit: fill;
  `,
  CyberpunkText: styled(ImageView).attrs({
    source: 'cyberpunktext',
    alterText: 'CYBERPUNK_ATTACK_LOGO_TEXT',
  })`
    max-width: 100%;
  `,
  CyberbodyImage: styled(ImageView).attrs({
    source: 'cyberbodyImagePNG',
    alterText: 'CYBERBODY_ATTACK_LOGO_TEXT',
  })`
    width: 100%;
  `,
  CardsList: styled(ImageView).attrs({
    source: 'cardslist',
    alterText: 'CYBERBODY_ATTACK_LOGO_TEXT',
  })`
    width: 100%;
  `,
  LogoRed: styled(ImageView).attrs({
    source: 'logoRed',
    alterText: 'LOGO_LAIDLON_RED',
  })`
    max-width: 37.5px;
    @media only screen and (max-width: 1024px) {
      max-width: 56px;
    }
  `,
};
