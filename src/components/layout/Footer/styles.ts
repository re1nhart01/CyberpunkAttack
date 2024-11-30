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
        padding-left: 20px;
        padding-right: 20px;
    }
    `,
  ShrinkContainer: styled.div`
      padding-left: 5.5%;
      padding-right: 5.5%;
      background-size: 40% 40%;
      background-image: url("${ICONS.footerBG}");
      background-repeat: repeat;
      min-width: 57%;
      gap: 44px;
      display: flex;
      position: relative;
      flex-direction: row;
      align-items: center;
      border: 1px solid rgba(255, 255, 255, 0.1);
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
    @media only screen and (max-width: 1024px) {
      flex-direction: column;
      gap: 8px;
    }
  `,
  LeftAlignedCorners: styled.div`
    position: absolute;
    left: -1px;
    top: 0;
    width: 12px;
    height: 12px;
    border-left: 2px solid #01E1FF;
    border-top: 2px solid #01E1FF;
  `,
  RightAlignedCorners: styled.div`
    position: absolute;
    right: -1px;
    top: 0;
    width: 12px;
    height: 12px;
    border-right: 2px solid #01E1FF;
    border-top: 2px solid #01E1FF;
  `,
};
