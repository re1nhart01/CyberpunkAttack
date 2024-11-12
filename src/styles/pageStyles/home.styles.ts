import styled from 'styled-components';
import { ICONS } from '../../constant/icons';
import RowView from '../../components/layout/RowView';
import { HEADER_HEIGHT } from '../../constant/constants';

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
    background-image: url(${ICONS.mainBackground});
    background-repeat: repeat;
    background-size: auto;
    background-position: top left;
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
  PageSectionInner: styled.div`
    max-width: 57vw;
    min-width: 57vw;
    display: flex;
    flex-direction: column;
    justify-content: stretch;
    align-items: center;
    @media only screen and (max-width: 1024px) {
      max-width: 100%;
      min-width: 100%;
    }
  `,
  FormWrapper: styled.div`
    padding-top: 24px;
    display: flex;
    flex-direction: row;
    align-items: center;
    gap: 8px;
    width: 100%;
    flex-wrap: wrap;
    @media only screen and (max-width: 1024px) {
      width: 100%;
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
