import styled from 'styled-components';
import { FOOTER_HEIGHT } from '../../../constant/constants';
import { ICONS } from '../../../constant/icons';

export const FooterStyles = {
  Wrapper: styled.footer`
    background: transparent;
    height: ${FOOTER_HEIGHT}px;    
    display: flex;
    flex-direction: row;
    justify-content: center;
    @media only screen and (max-width: 1024px) {
        height: initial;
        flex-direction: column;
    }
    `,
  ShrinkContainer: styled.div`
      padding-left: 5.5%;
      padding-right: 5.5%;
      background-size: 100% 100%;
      background-image: url("${ICONS.footerBG}");
      min-width: 57%;
      gap: 44px;
      display: flex;
      flex-direction: row;
      align-items: center;
    @media only screen and (max-width: 1024px) {
      flex-direction: column;
      padding-top: 30px;
      padding-bottom: 23px;
      text-align: center;
    }
    `,
  LogoContainer: styled.div`
    display: flex;
    flex-direction: row;
    align-items: center;
    gap: 12px;
  `,
};
