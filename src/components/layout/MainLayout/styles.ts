import styled from 'styled-components';
import { HEADER_HEIGHT } from '../../../services/constant/icons';

export const MainLayoutStyles = {
  MainWrapper: styled.main`
      display: flex;
      flex-direction: column;
      min-height: 100%;
      background-color: ${({ theme }) => theme.components.backgrounds.primary};
    `,
  ContentWrapper: styled.div`
      display: flex;
      flex: 1 1 auto;
    `,
  ShrinkContainer: styled.div`
      max-width: 1440px;
      display: flex;
      flex: 1;
      margin-left: auto;
      margin-right: auto;
    `,
  FullContainer: styled.div`
      display: flex;
      flex: 1;
      margin-left: auto;
      margin-right: auto;
    `,
};

export const ShrinkContainer = MainLayoutStyles.ShrinkContainer;
export const FullContainer = MainLayoutStyles.FullContainer;
