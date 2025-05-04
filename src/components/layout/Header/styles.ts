import styled from 'styled-components';
import { HEADER_HEIGHT } from '../../../constant/constants';
import Button from '../../buttons/Button';

export const HeaderStyles = {
  Wrapper: styled.header`
      display: flex;
      flex-direction: row;
      justify-content: space-between;
      align-items: center;
      max-width: 100vw;
      height: ${HEADER_HEIGHT}px;
      background-color: rgba(0,0,0,0.1);
      padding-left: 32px;
      padding-right: 32px;
      position: sticky;
      width: 100%;
      box-sizing: border-box;
      flex-wrap: wrap;
      top: 0;
      z-index: 999;
      backdrop-filter: blur(10px);
      -webkit-backdrop-filter: blur(10px);
      border-bottom: 1px solid #FFFFFF33;
      @media only screen and (max-width: 1024px) {
        padding-left: 20px;
        padding-right: 20px;
      }
  `,
  LogoWrapper: styled.button`
      display: flex;
      flex-direction: row;
      gap: 12px;
      justify-content: space-between;
      align-items: center;
      cursor: pointer;
      border: none;
      background: none;
    `,
  ButtonsWrapper: styled.div`
    display: flex;
    flex-direction: row;
    align-items: center;
  `,
  HeaderButton: styled(Button)`
    padding-left: 24px;
    padding-right: 24px;
    background-color: transparent;
    z-index: 999;
    &, & > * {
      cursor: pointer;
    }
  `,
  HeaderButtonRound: styled(Button)`
    padding-left: 24px;
    padding-right: 24px;
    background-color: #01FFFF;
    display: flex;
    flex-direction: row;
    jusity-content: center;
    align-items: center;
    color: #000000 !important;
    height: 32px;
    border-radius: 70px;
    padding-top: 2px;
    padding-right: 24px;
    padding-left: 24px;
    margin-left: 24px;
    transition: 0.5s;
    gap: 10px;
    z-index: 999;
    &, & > * {
      cursor: pointer;
    }

    & > * {
    color: #000000;
    transition: color 0.5s;
  }


     &:hover {
      background-color: ${({ theme }) => theme.colors.black};
      & > button > span {
        color: #ffffff;
      }

    }
  `,
  Separator: styled.div`
    height: 12px;
    background-color: ${({ theme }) => theme.colors.main};
    width: 2px;
  `,
};
