import styled from 'styled-components';
import { ICONS } from '../../../services/constant/icons';

export const blurRowStyles = {
  Wrapper: styled.div`
      display: flex;
      flex-direction: row;
      align-items: center;
      flex-wrap: wrap;
      background: url(${ICONS.blurRowImg}) no-repeat;
      background-size: contain;
      padding: 0px 24px 0px 12px;
      height: 100px;
      width: 100%;
    `,
};
