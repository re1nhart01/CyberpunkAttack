import styled from 'styled-components';
import ButtonView from './index';
import { ICONS } from '../../../services/constant/icons';

export const ButtonStyles = {
  Wrapper: styled.div``,
  Button: styled.button<{ defaultcol?: boolean; gapInside?: number | string }>`
      display: flex;
      flex-direction: row;
      gap: ${({ gapInside }) => gapInside}px;
      background-color: ${({ defaultcol, theme }) => (!defaultcol ? 'transparent' : theme.colors.white)};
      border: ${({ defaultcol, theme }) => (!defaultcol ? 'none' : theme.colors.white)};
      color: ${({ defaultcol, theme }) => (!defaultcol ? theme.colors.white : theme.colors.main)};
    `,
};
