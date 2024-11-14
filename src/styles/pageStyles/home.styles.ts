import styled from 'styled-components';
import { ICONS } from '../../constant/icons';

export const HomePageStyles = {
  FullScreenMenu: styled.div<{ isOpen: boolean }>`
    position: fixed;
    top: 0;
    left: 0;
    width: 100%;
    height: 100%;
    display: flex;
    flex-direction: column;
    justify-content: center;
    align-items: center;
    background-image: url(${ICONS.mainBackground});
    background-repeat: repeat;
    background-size: auto;
    background-position: top left;
    transform: ${({ isOpen }) => (isOpen ? 'translateY(0)' : 'translateY(-100%)')};
    transition: transform 0.5s ease;
    z-index: 5;
`,
  PageContainer: styled.div`
    background: url(${ICONS.mainBackground}) repeat;
    background-size: 60% 15%;
    height: 100%;
    max-width: 100%;
  `,
  PageSection: styled.section`
    width: 100%;
    display: flex;
    flex-direction: row;
    justify-content: center;
    align-items: center;
    @media only screen and (max-width: 1024px) {
      padding-right: 20px;
      padding-left: 20px;
    }
  `,
  Separator: styled.div`
    height: 48px;
    background-color: ${({ theme }) => theme.colors.main};
    width: 2px;
    @media only screen and (max-width: 1024px) {
      width: 1px;
      height: 24px;
    }
  `,
  KickstarterContainer: styled.div`
    width: 453px;
    height: 90px;
    background-color: rgba(0,0,0,0.5);
    clip-path: polygon(100% 0, 100% 61%, 94% 100%, 0 100%, 0 0);
    position: absolute;
    bottom: 15%;
    left: 50%;
    display: flex;
    flex-direction: row;
    justify-content: space-between;
    align-items: center;
    padding-left: 24px;
    padding-right: 24px;

    @media only screen and (max-width: 1024px) {
      position: initial;
      background: transparent;
      flex-direction: column;
      min-width: 100%;
      height: initial;
      box-sizing: border-box;
      
      & > div {
        width: 100%;
        margin-top: 12px;
        margin-bottom: 24px;
      }
      
      #kickstarter-block > div:nth-child(1) {
        font-size: 32px !important;
      }

      #kickstarter-block > div:nth-child(2) {
        font-size: 18px !important;
      }
      
    }
    
    & > div {
      text-align: left;
    }
    
    & > div:nth-child(1) > div {
      font-size: 20px;
    }

    @media only screen and (min-width: 1024px) {
      & > div > div > button {
        width: 180px;
        height: 48px;
        background: url("${ICONS.smallBg}") no-repeat;
        background-size: 100% 100%;
      }

      & > div > div > button:hover {
        width: 180px;
        height: 48px;
        background: url("${ICONS.smallBgBlack}") no-repeat;
        background-size: 100% 100%;
      }
    }
    
   
    
  `,
  PageSectionInner: styled.div`
    max-width: 840px;
    min-width: 840px;
    display: flex;
    flex-direction: column;
    justify-content: stretch;
    align-items: center;
    @media only screen and (max-width: 1024px) {
      max-width: 100%;
      min-width: 100%;
    }
  `,
  FormWrapper: styled.div<{ isStretch?: boolean }>`
    padding-top: 24px;
    display: flex;
    flex-direction: row;
    align-items: center;
    gap: 8px;
    width: ${({ isStretch }) => (isStretch ? 'initial' : '100%')};
    flex-wrap: wrap;
    @media only screen and (max-width: 1024px) {
      width: ${({ isStretch }) => (isStretch ? 'initial' : '100%')};
      flex-direction: column;
      & > * {
        box-sizing: border-box;
        width: 100%;
      }
    }
  `,
  InnerWrapper: styled.div`
    padding-top: 16px;
  `,
  SocialButtonInner: styled.div`
    min-width: 100%;
    box-sizing: border-box;
    display: flex;
    flex-direction: column;
    justify-content: center;
    padding-top: 35px;
    gap: 35px;
    @media only screen and (max-width: 1024px) {
      gap: 0px;    
      padding-top: 0px;
      
      & svg {
        width: 24px;
        height: 24px;
      }
    }
  `,
  Del: styled.div`
      display: flex;
      flex-direction: row;
      align-items: center;
      margin-right: 5%;
    `,
  FBlockWrapper: styled.div`
      display: flex;
      flex-direction: column;
      justify-content: space-between;
      margin-bottom: 24px;
      position: relative;
      @media only screen and (max-width: 1024px) {
      }
    `,
  BorderView: styled.span`
    border: 8px solid ${({ theme }) => theme.colors.black}
  `,
  Spacer: styled.div<{ height: number | string }>`
    width: 1%;
    height: ${({ height }) => height}px;
    background-color: transparent;
  `,
  WhiteColorText: styled.div`
    & > * {
      color: ${({ theme }) => theme.colors.white};
      font-family: Zekton;
    }
  `,
};
