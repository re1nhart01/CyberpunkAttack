import styled from 'styled-components';
import { ICONS } from '../../../constant/icons';

export const socialButtonStyles = {
  Wrapper: styled.div`
  display: flex;
  flex-direction: row;
  justify-content: center;
  align-items: center;
`,
  SocialButtonContainer: styled.button`
    background: url("${ICONS.socialButtonBG}") no-repeat center;
    background-size: 100% 100%;
    max-width: 660px;
    width: 100%;
    height: 78px;
    box-sizing: border-box;
    transition: 0.5s;
    border: none;
    cursor: pointer;
    
    & * {
      color: #0a0a0a !important;
    }

    @media only screen and (min-width: 1024px) {
      &:hover {
        background: url("${ICONS.socialButtonBGActive}") no-repeat;
        background-size: 100% 100%;
        transition: 0.5s;
      };

      &:hover * {
        transition: 0.5s;
        color: white !important;
        fill: white !important;
      }
    }
  `,
  SocialButtonInner: styled.div`
  width: 100%;
  height: 100%;
  display: flex;
  flex-direction: row;
  justify-content: center;
  align-items: center;
  gap: 10px;
`,
  Icon: styled.img`
  width: 24px;
  height: 24px;
`,
};
