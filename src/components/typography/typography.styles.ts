import styled from 'styled-components';
import Typography, { typographyProps } from './index';

export const TypographyComponents = {
  Text22space400: styled(Typography).attrs<typographyProps>({
    color: 'white',
    fw: 'fw400',
    fsz: 'fz22',
    ff: 'space',
    selector: 'span',
  })`
      line-height: 120%;
      letter-spacing: 3.242px;
    `,
  Text11aquantix400: styled(Typography).attrs<typographyProps>({
    color: 'purple',
    fw: 'fw400',
    fsz: 'fz11',
    ff: 'aquantix',
    selector: 'span',
  })`
      line-height: 100%;
      letter-spacing: 18.526px;
    `,
  // FIRST BLOCK INDEX.TSX
  Text76orbitron700: styled(Typography).attrs<typographyProps>({
    color: 'white',
    fw: 'fw700',
    fsz: 'fz76',
    ff: 'orbitron',
    selector: 'span',
    fzType: 'rem',
  })`
      text-shadow: -2px -2px 0px ${({ theme }) => theme.colors.main}, 2px 2px 0px ${({ theme }) => theme.colors.purple};
      line-height: 100%;
      text-transform: uppercase;
    `,
  Text18boxed500: styled(Typography).attrs<typographyProps>({
    color: 'white',
    fw: 'fw500',
    fsz: 'fz18',
    ff: 'boxedRound',
    selector: 'span',
    fzType: 'px',
  })`
      line-height: 150%;
      opacity: 0.8;
    `,
  Text24boxed600: styled(Typography).attrs<typographyProps>({
    color: 'white',
    fw: 'fw600',
    fsz: 'fz24',
    ff: 'boxedRound',
    selector: 'span',
    fzType: 'px',
  })`
      line-height: 100%;
    `,
  Text16boxed500: styled(Typography).attrs<typographyProps>({
    color: 'main',
    fw: 'fw500',
    fsz: 'fz16',
    ff: 'boxedRound',
    selector: 'span',
    fzType: 'px',
  })`
      line-height: 110%;
    `,
  Text16boxed600: styled(Typography).attrs<typographyProps>({
    color: 'white',
    fw: 'fw600',
    fsz: 'fz16',
    ff: 'boxedRound',
    selector: 'span',
    fzType: 'px',
  })``,
  Text40orbitron700: styled(Typography).attrs<typographyProps>({
    color: 'white',
    fw: 'fw700',
    fsz: 'fz40',
    ff: 'orbitron',
    selector: 'span',
    fzType: 'px',
  })``,
  Text16boxed700: styled(Typography).attrs<typographyProps>({
    color: 'white',
    fw: 'fw700',
    fsz: 'fz16',
    ff: 'boxedRound',
    selector: 'span',
    fzType: 'px',
  })`
    line-height: 150%;
    opacity: 0.5;
    `,
  Text60orbitron700: styled(Typography).attrs<typographyProps>({
    color: 'white',
    fw: 'fw700',
    fsz: 'fz60',
    ff: 'orbitron',
    selector: 'span',
    fzType: 'px',
  })`
      text-shadow: -2px -1px 0px #01D1FF, 2px 1px 0px #FC2B87;
      text-transform: uppercase;
      line-height: 110%;
    `,
  Text20boxed400: styled(Typography).attrs<typographyProps>({
    color: 'white',
    fw: 'fw400',
    fsz: 'fz20',
    ff: 'boxedRound',
    selector: 'span',
    fzType: 'px',
  })``,
  Text170space700: styled(Typography).attrs<typographyProps>({
    color: 'dark',
    fw: 'fw700',
    fsz: 'fz170',
    ff: 'space',
    selector: 'span',
    fzType: 'rem',
  })`
    position: absolute;
    top: 0;
    left: 4.6%;
    right: 0;
    opacity: 0.05;
  `,
  Text60orbitron800: styled(Typography).attrs<typographyProps>({
    color: 'dark',
    fw: 'fw800',
    fsz: 'fz60',
    ff: 'orbitron',
    selector: 'span',
    fzType: 'px',
  })`
    line-height: 110%;
    text-transform: uppercase;
  `,
  Text24bebas400: styled(Typography).attrs<typographyProps>({
    color: 'main',
    fw: 'fw400',
    fsz: 'fz24',
    ff: 'bebas',
    selector: 'span',
    fzType: 'px',
  })`
    line-height: 80%; /* 19.2px */
    text-transform: uppercase;
  `,
  Text18boxed400: styled(Typography).attrs<typographyProps>({
    color: 'black',
    fw: 'fw400',
    fsz: 'fz14',
    ff: 'boxedRound',
    selector: 'span',
    fzType: 'px',
  })`
    font-weight: 400;
    line-height: 150%; /* 27px */
  `,
  Text16Zekton400: styled(Typography).attrs<typographyProps>({
    color: 'white',
    fw: 'fw400',
    fsz: 'fz16',
    ff: 'zek',
    selector: 'span',
    fzType: 'px',
  })``,
  Text18Zekton400: styled(Typography).attrs<typographyProps>({
    fw: 'fw400',
    fsz: 'fz18',
    ff: 'zek',
    selector: 'span',
    fzType: 'px',
  })``,
  Text16Zekton700: styled(Typography).attrs<typographyProps>({
    color: 'white',
    fw: 'fw700',
    fsz: 'fz16',
    ff: 'zek',
    selector: 'span',
    fzType: 'px',
  })``,
  Text14Zekton700: styled(Typography).attrs<typographyProps>({
    fw: 'fw700',
    fsz: 'fz16',
    ff: 'zek',
    selector: 'span',
    fzType: 'px',
  })``,
  Text18Zekton700: styled(Typography).attrs<typographyProps>({
    fw: 'fw700',
    fsz: 'fz18',
    ff: 'zek',
    selector: 'span',
    fzType: 'px',
  })``,
  Text48Orbitron700: styled(Typography).attrs<typographyProps>({
    color: 'white',
    fw: 'fw700',
    fsz: 'fz48',
    ff: 'orbitron',
    selector: 'div',
    fzType: 'px',
  })`
    text-shadow: #01D1FF -1px -1px 0px 0px, #FC2B87 1px 1px 0px 0px;
    text-align: center;
    @media only screen and (max-width: 1024px) {
      font-size: 28px !important;
    }
  `,
  Text24Zekton400: styled(Typography).attrs<typographyProps>({
    fw: 'fw400',
    fsz: 'fz24',
    ff: 'zek',
    selector: 'div',
    fzType: 'px',
  })`
    text-align: center;
    @media only screen and (max-width: 1024px) {
      font-size: 18px !important;
    }
  `,
  Text24Zekton700: styled(Typography).attrs<typographyProps>({
    color: 'white',
    fw: 'fw700',
    fsz: 'fz24',
    ff: 'zek',
    selector: 'div',
    fzType: 'px',
  })`
    text-align: center;
    @media only screen and (max-width: 1024px) {
      font-size: 18px !important;
    }
  `,
  Text56Bangers400: styled(Typography).attrs<typographyProps>({
    fw: 'fw400',
    fsz: 'fz56',
    ff: 'bangers',
    selector: 'span',
    fzType: 'px',
  })`
    align-self: flex-start;
  `,
  Text26Space400: styled(Typography).attrs<typographyProps>({
    fw: 'fw400',
    fsz: 'fz26',
    ff: 'space',
    selector: 'span',
  })`
    text-shadow: 0px 0px 4.44px 0px #01E1FF;
    @media only screen and (max-width: 1024px) {
      font-size: 14px !important;
    }
  `,
  Text12Space400: styled(Typography).attrs<typographyProps>({
    fw: 'fw400',
    fsz: 'fz12',
    ff: 'space',
    selector: 'span',
  })`
    text-shadow: 0px 0px 4.44px 0px #01E1FF;
    @media only screen and (max-width: 1024px) {
      font-size: 14px !important;
    }
    letter-spacing: 1.63px;
  `,
  Text14Zekton400: styled(Typography).attrs<typographyProps>({
    fw: 'fw400',
    fsz: 'fz14',
    ff: 'zek',
  })`
    opacity: 0.5;
  `,
};

export const OverrideTypographyComponents = {
  Text40orbitron700After: styled(TypographyComponents.Text40orbitron700)``,
  Text16Zekton400Black: styled(TypographyComponents.Text16Zekton400)`
    font-weight: 700 !important;
    color: ${({ theme }) => theme.colors.black} !important;
    text-align: center;
  `,
};
